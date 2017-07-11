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
		uint valueMod;
		uint8 active;
    }
    address[] internal bonusMembers;
    
    uint unlockedTokens;
    uint lockedTokens = 10 * 1000 * 1000;
    uint maxTokens = 30 * 1000 * 1000;
    //https://www.epochconverter.com -> 1.9.2017 + 1 week
    uint mintPhaseEnd = 1504224000 + 1 weeks; //1 week after ICO, to be defined
    uint votingDuration = 2 weeks;
    uint8 quorumWeight = 200; //100 -> is 1, 150 is 1.5, with 200 you need 2/3 of the votes
    
    string public constant name = "Modum Token";
    string public constant symbol = "MOD";
    
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
        accounts[msg.sender] = Account(0,0,1);
    }
    
    function proposal(string _addr, bytes32 _hash, uint _value) returns (bool) {
        //only if no proposal is pending
        require(currentProposal.valueMod == 0);
        uint _yay = unlockedTokens; // default is yes
        uint _nay = 0;
        currentProposal = Proposal(_addr, _hash, _value, now, _yay, _nay);
        return true;
    }
    
    function vote(bool _vote) returns (uint) {
        uint votingPower = accounts[msg.sender].valueMod;
        if(! _vote) {
            currentProposal.yay -= votingPower;
            currentProposal.nay += votingPower;
        }
        return votingPower;
    }
    
    function claimProposal() returns (bool) {
        require(msg.sender == owner);
        if(currentProposal.valueMod > 0 &&
                now > currentProposal.startTime + votingDuration &&
                lockedTokens >= currentProposal.valueMod) {
            if(currentProposal.yay > (currentProposal.nay * quorumWeight) / 100) {
                accounts[msg.sender].valueMod += currentProposal.valueMod;
                unlockedTokens += currentProposal.valueMod;
                lockedTokens -= currentProposal.valueMod;
            }
            delete currentProposal;
            return true;
        }
        return false;
    }
    
    function mint(address _recipient, uint _value) returns (bool) {
        require(msg.sender == owner);
        if(now < mintPhaseEnd &&
                (lockedTokens + unlockedTokens + _value) <= maxTokens) {
            if(accounts[_recipient].active == 0) {
                accounts[_recipient].active = 1;
                bonusMembers.push(_recipient);
            }
            accounts[_recipient].valueMod += _value;
            unlockedTokens += _value;
            
            return true;
        }
        return false;
    }
    
    function bonus() payable returns (bool) {
        require(msg.sender == owner);
        //distribute the wei to the token holders
        uint totalBonus;
        //don't add our tokens, as we don't get any bonus
        uint tokenCount = unlockedTokens - accounts[msg.sender].valueMod;
        for (uint i = 0; i < bonusMembers.length; i++) {
            uint bonus = (msg.value * accounts[bonusMembers[i]].valueMod) / tokenCount;
            accounts[bonusMembers[i]].bonusWei += bonus;
            totalBonus += bonus;
        }
        //rounding goes to us
        accounts[msg.sender].bonusWei += msg.value - totalBonus;
    }
    
    function claimBonus() returns (bool) {
        msg.sender.transfer(accounts[msg.sender].bonusWei);
        accounts[msg.sender].bonusWei = 0;
        return true;
    }
    
    function totalSupply() constant returns (uint) {
        return unlockedTokens + lockedTokens;
    }
    function balanceOf(address _owner) constant returns (uint balance) {
        return accounts[_owner].valueMod;
    }
    function transfer(address _to, uint _value) returns (bool success) {
        if (accounts[msg.sender].valueMod >= _value 
            && _value > 0
            && accounts[_to].valueMod + _value > accounts[_to].valueMod) {
                
                if(accounts[_to].active == 0) {
                    accounts[_to].active = 1;
                    bonusMembers.push(_to);
                }
                
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
                
                if(accounts[_to].active == 0) {
                    accounts[_to].active = 1;
                    bonusMembers.push(_to);
                }
                
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
    
    event Transfer(address indexed _from, address indexed _to, uint _value);

    event Approval(address indexed _owner, address indexed _spender, uint _value);
}
