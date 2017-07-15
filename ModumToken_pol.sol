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
    mapping(address => Account) accounts;
    mapping(address => mapping (address => uint256)) allowed;
    
    enum UpdateMode{None,Wei,Vote,Both}

	struct Account {
	    uint lastProposalStartTime; //For checking at which proposal valueModVote was last updated 
		uint lastAirdropWei; //For checking after which airDrop bonusWei was last updated 
		uint bonusWei;      //airDrop/Dividend payout avaiable for withdrawl
		uint valueModVote;  // votes avaiable for voting on active Proposal
		uint valueMod;      // the owned tokens
    }
    
    uint bonusLockCorrectedWei = 0;         //totally airdropped eth with a correction to accpount for locked tokens
    uint rounding = 0;                      //airdrops not accounted yet to make system rounding error proof

    uint unlockedTokens = 0;                //tokens that can vote, transfer, receive dividend
    uint lockedTokens = 10 * 1000 * 1000;   //token that need to be unlocked by voting
    uint maxTokens = 30 * 1000 * 1000;      //max distributable tokens
    //https://www.epochconverter.com -> 1.9.2017 + 1 week
    bool mintDone = false;                  //distinguisher for minting phase
    uint votingDuration = 2 minutes;        //time to vote
    
    string public constant name = "Modum Token";
    string public constant symbol = "MOD";
    
    struct Proposal {
        string addr;        //Uri for more info
        bytes32 hash;       //Hash of the uri content for checking 
        uint valueMod;      //token to unlock: proposal with 0 amount is invalid
        uint startTime;     
        uint16 quorumWeight; //100 -> is 1, 150 is 1.5, with 200 you need 2/3 of the votes
        uint yay;   
        uint nay;
    }
    
    Proposal currentProposal; 

    function ModumToken() {
        owner = msg.sender;
    }
    
    function proposal(string _addr, bytes32 _hash, uint _value, uint16 _quorumWeight) {
        require(mintDone); //minting phase needs to be over
        require(currentProposal.valueMod == 0); // no vote is ongoing
        require(msg.sender == owner); // proposal ony by onwer
        require(lockedTokens >= _value); //proposal cannot be larger than remaining locked tokens
        require(_value > 0);            //proposal with 0 unlock are invalid
        uint _yay = unlockedTokens; // default is yes // discussion
        uint _nay = 0;
        currentProposal = Proposal(_addr, _hash, _value, now, _quorumWeight, _yay, _nay);
    }
    
    function vote(bool _vote) returns (uint) {
        require(mintDone); //minting phase needs to be over
        require(isVoteOngoing()); // vote needs to be ongoing
        Account storage account = getAccount(msg.sender, UpdateMode.Vote);
        uint votes = account.valueModVote;  //avaiable votes
        require(votes > 0);  //voter must have a vote left
        
        if(! _vote) {   //yes votes are default
            currentProposal.yay -= votes;
            currentProposal.nay += votes;
        }
        
        account.valueModVote = 0;
        return votes;
    }
    
    function claimProposal(){
        require(mintDone); //minting phase needs to be over
        require(msg.sender == owner); //only owner can claim proposal
        require(currentProposal.valueMod > 0); // no proposal active
        require(now > currentProposal.startTime + votingDuration); // voting has already ended
        if(currentProposal.yay > (currentProposal.nay * currentProposal.quorumWeight) / 100) {
            //It was accepted
            Account storage account = getAccount(owner, UpdateMode.Both);
            uint valueMod = currentProposal.valueMod;
            account.valueMod += valueMod; //uadd to owner
            unlockedTokens += valueMod; //unlock
            lockedTokens -= valueMod; //unlock
        }
        delete currentProposal; //proposal ended
    }
    
    function mint(address _recipient, uint _value)  {
        require(msg.sender == owner); //only owner can claim proposal
        require(!mintDone); //only during minting
        require((totalSupply() + _value) <= maxTokens); //do not exceed max
        Account storage account = getAccount(_recipient, UpdateMode.None);
        account.valueMod += _value; //create the tokens and add to recipient
        unlockedTokens += _value; //create tokens

    }
    
    function mintFinished() {
        require(msg.sender == owner); //only owner
        require(!mintDone); //only in minting
        mintDone = true; //end the minting
    }
    
    //default function to pay bonus, anybody that sends eth to this contract will distribute the wei
    //to their token holders
    //Dividend payment / Airdrop
    function() payable {
        require(mintDone);      //only after minting
        uint value = msg.value + rounding; //add old runding
        rounding = value % unlockedTokens; //ensure no rounding error
        bonusLockCorrectedWei += ((value-rounding) * totalSupply()) / unlockedTokens; //account for locked tokens and add the drop
    }
    
    function getUnlockedTokens() constant returns (uint) {
        return unlockedTokens;
    }
    
    function claimBonus() {
        Account storage account = getAccount(msg.sender, UpdateMode.Wei);
        uint sendValue = account.bonusWei;
        if(sendValue != 0){
            account.bonusWei = 0;
            msg.sender.transfer(sendValue);
        }
    }
    
    function totalSupply() constant returns (uint) {
        return unlockedTokens + lockedTokens;
    }
    
    function balanceOf(address _owner) constant returns (uint balance) {
        return accounts[_owner].valueMod;
    }
    
    function isVoteOngoing() internal returns (bool)  {
        return currentProposal.valueMod != 0 && now >= currentProposal.startTime && now < currentProposal.startTime + votingDuration;
    }
    
	function getAccount(address _addr, UpdateMode mode) internal returns(Account storage){
        if(mode != UpdateMode.None) require(mintDone);
        
        Account storage account = accounts[_addr];
		if(mode == UpdateMode.Vote || mode == UpdateMode.Both){
		    if(isVoteOngoing() && account.lastProposalStartTime < currentProposal.startTime) { // the user did set his token power yet
                account.valueModVote = account.valueMod;
                account.lastProposalStartTime = currentProposal.startTime;
            }
		}
		
		if(mode == UpdateMode.Wei || mode == UpdateMode.Both){
            uint bonus = (bonusLockCorrectedWei -  account.lastAirdropWei);
            if(bonus != 0){
    			account.bonusWei += (bonus * account.valueMod) / totalSupply();
    			account.lastAirdropWei = bonusLockCorrectedWei;
    		}
		}
		
		return account;
    }
    
    function transfer(address _to, uint _value) returns (bool success) {
        require(mintDone);
        Account storage tmpFrom = getAccount(msg.sender, UpdateMode.None);
        if (tmpFrom.valueMod >= _value  && _value > 0){
                Account storage from = getAccount(msg.sender, UpdateMode.Both);
                Account storage to = getAccount(_to, UpdateMode.Both);
                from.valueMod -= _value;
                to.valueMod += _value;
                Transfer(msg.sender, _to, _value);
                return true;
        } 
        return false;
    }
    
    function transferFrom(address _from, address _to, uint _value) returns (bool success) {
        require(mintDone);
        Account storage tmpFrom = getAccount(msg.sender, UpdateMode.None);
        if (tmpFrom.valueMod >= _value  && _value > 0 && allowed[_from][msg.sender] >= _value){
                Account storage from = getAccount(msg.sender, UpdateMode.Both);
                Account storage to = getAccount(_to, UpdateMode.Both);
                from.valueMod -= _value;
                to.valueMod += _value;
                allowed[_from][msg.sender] -= _value;
                Transfer(msg.sender, _to, _value);
                return true;
        } 
        return false;
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


