pragma solidity ^0.4.10;

//https://theethereum.wiki/w/index.php/ERC20_Token_Standard
contract ERC20Interface {

    // Get the total token supply
    function totalSupply() constant returns (uint256 totalSupply);

    // Get the account balance of another account with address _owner
    function balanceOf(address _owner) constant returns (uint256 balance);

    // Send _value amount of tokens to address _to
    function transfer(address _to, uint256 _value) returns (bool success);

    // Send _value amount of tokens from address _from to address _to
    function transferFrom(address _from, address _to, uint256 _value) returns (bool success);

    // Allow _spender to withdraw from your account, multiple times, up to the _value amount.
    // If this function is called again it overwrites the current allowance with _value.
    // this function is required for some DEX functionality
    function approve(address _spender, uint256 _value) returns (bool success);

    // Returns the amount which _spender is still allowed to withdraw from _owner
    function allowance(address _owner, address _spender) constant returns (uint256 remaining);
    // Triggered when tokens are transferred.
    event Transfer(address indexed _from, address indexed _to, uint256 _value);

    // Triggered whenever approve(address _spender, uint256 _value) is called.
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
}

contract ModumToken is ERC20Interface {
    
    address owner;
    mapping (address => Account) accounts;
    mapping(address => mapping (address => uint256)) allowed;
    struct Account {
		uint bonusWei;
		uint lastAirdropWei;
		uint valueMod;
		uint valueModVote;
        uint lastProposalStartTime;
    }
    
    uint bonusWei;
    
    uint unlockedTokens;
    uint lockedTokens = 10 * 1000 * 1000;
    uint maxTokens = 30 * 1000 * 1000;
    //https://www.epochconverter.com -> 1.9.2017 + 1 week
    bool mintDone = false;
    uint votingDuration = 2 minutes;
    uint8 quorumWeight = 200; //100 -> is 1, 150 is 1.5, with 200 you need 2/3 of the votes
    
    string public constant name = "Modum Token";
    string public constant symbol = "MOD";
    
    uint rounding = 0;
    
    struct Proposal {
        string addr;
        bytes32 hash;
        uint valueMod; //proposal with 0 amount is invalid
        uint startTime;
        uint yay;
        uint nay;
    }
    
    Proposal currentProposal;

    function ModumToken() {
        owner = msg.sender;
        //we are not a bonus member, so don't add msg.sender to bonusMembers
        accounts[msg.sender] = Account(0,0,0,0,0);
    }
    
    function proposal(string _addr, bytes32 _hash, uint _value) returns (bool) {
        require(mintDone); //minting phase needs to be over
        require(currentProposal.valueMod == 0); // no vote is ongoing
        require(msg.sender == owner); // proposal ony by onwer
        require(lockedTokens >= _value); //proposal cannot be larger than locked tokens
        require(_value > 0);
        uint _yay = unlockedTokens; // default is yes // discussion
        uint _nay = 0;
        currentProposal = Proposal(_addr, _hash, _value, now, _yay, _nay);
        return true;
    }
    
    function vote(bool _vote) returns (uint) {
        require(isVoteOngoing()); // vote is ongoing
        votingPower(msg.sender);
        require(accounts[msg.sender].valueModVote > 0);
        
        if(! _vote) {
            currentProposal.yay -= accounts[msg.sender].valueModVote;
            currentProposal.nay += accounts[msg.sender].valueModVote;
        }
        accounts[msg.sender].valueModVote = 0;
        return accounts[msg.sender].valueModVote;
    }
    
    function claimProposal() returns (bool) {
        require(msg.sender == owner);
        require(currentProposal.valueMod > 0 && now > currentProposal.startTime + votingDuration); // vote just ended
        if(currentProposal.yay > (currentProposal.nay * quorumWeight) / 100) {
            accounts[msg.sender].valueMod += currentProposal.valueMod;
            unlockedTokens += currentProposal.valueMod;
            lockedTokens -= currentProposal.valueMod;
        }
        delete currentProposal;
        return true;
    }
    
    function mint(address _recipient, uint _value) returns (bool) {
        require(msg.sender == owner);
        if(!mintDone && (totalSupply() + _value) <= maxTokens) {
            
            weiPower(_recipient);
            
            accounts[_recipient].valueMod += _value;
            unlockedTokens += _value;
            return true;
        }
        return false;
    }
    
    function mintFinished() returns (bool) {
        require(msg.sender == owner);
        mintDone = true;
        return true;
    }
    
    //default function to pay bonus, anybody that sends eth to this contract will distribute the wei
    //to their token holders
    function() payable {
        uint value = msg.value + rounding;
        rounding = value % unlockedTokens;
        bonusWei += ((value-rounding) * totalSupply()) / unlockedTokens;
    }
    
    function getUnlockedTokens() constant returns (uint) {
        return unlockedTokens;
    }
    
    function claimBonus() returns (bool) {
        weiPower(msg.sender);
        uint sendValue = accounts[msg.sender].bonusWei;
        accounts[msg.sender].bonusWei = 0;
        accounts[msg.sender].lastAirdropWei = bonusWei;
        msg.sender.transfer(sendValue);
        return true;
    }
    
    function totalSupply() constant returns (uint) {
        return unlockedTokens + lockedTokens;
    }
    function balanceOf(address _owner) constant returns (uint balance) {
        return accounts[_owner].valueMod;
    }
    
    function isVoteOngoing() internal returns (bool)  {
        return currentProposal.valueMod > 0 && now >= currentProposal.startTime && now < currentProposal.startTime + votingDuration;
    }
    
    function votingPower(address _addr) internal {
        if(accounts[_addr].lastProposalStartTime < currentProposal.startTime) { // the user did set his token power yet
            accounts[_addr].valueModVote = accounts[_addr].valueMod;
            accounts[_addr].lastProposalStartTime = currentProposal.startTime;
        }
    }
    
    function weiPower(address _addr) internal {
        uint bonus = (bonusWei -  accounts[_addr].lastAirdropWei);
        accounts[_addr].bonusWei += (bonus * accounts[_addr].valueMod) / totalSupply();
        accounts[_addr].lastAirdropWei = bonusWei;
    }
    
    function transfer(address _to, uint _value) returns (bool success) {
        if (accounts[msg.sender].valueMod >= _value 
            && _value > 0
            && accounts[_to].valueMod + _value > accounts[_to].valueMod) {
                
                //if we are in a voting phase, make sure we won't change the voting power
                if(isVoteOngoing()) {
                    votingPower(msg.sender);
                    votingPower(_to);
                }
                
                weiPower(msg.sender);
                weiPower(_to);
                
                accounts[msg.sender].valueMod -= _value;
                accounts[_to].valueMod += _value;
                Transfer(msg.sender, _to, _value);
                return true;
        } else {
            return false;
        }
    }
    
    function transferFrom(address _from, address _to, uint _value) returns (bool success) {
        if (accounts[_from].valueMod >= _value
            && allowed[_from][msg.sender] >= _value
            && _value > 0
            && accounts[_to].valueMod + _value > accounts[_to].valueMod) {
                
                //if we are in a voting phase, make sure we won't change the voting power
                if(isVoteOngoing()) {
                    votingPower(_from);
                    votingPower(_to);
                }
                
                weiPower(_from);
                weiPower(_to);
                
                accounts[_from].valueMod -= _value;
                allowed[_from][msg.sender] -= _value;
                accounts[_to].valueMod += _value;
                return true;
        } else {
            return false;
        }
    }
    function approve(address _spender, uint _value) returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        Approval(msg.sender, _spender, _value);
        return true;
    }
    function allowance(address _owner, address _spender) constant returns (uint remaining) {
        return allowed[_owner][_spender];
    }
}


