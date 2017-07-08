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
    mapping (address => uint) accounts;
    mapping(address => mapping (address => uint256)) allowed;
    
    uint unlockedTokens;
    uint lockedTokens = 10 * 1000 * 1000;
    uint maxTokens = 30 * 1000 * 1000;
    //https://www.epochconverter.com -> 1.9.2017 + 1 week
    uint mintPhaseEnd = 1504224000 + 1 weeks; //1 week after ICO, to be defined
    uint votingDuration = 2 weeks;
    
    string public constant name = "Modum Token";
    string public constant symbol = "MOD";
    
    struct Proposal {
        string addr;
        bytes32 hash;
        uint value; //proposal with 0 amount is invalid
        uint startTime;
        uint yay;
        uint nay;
        uint nrVotes;
    }
    
    Proposal currentProposal;

    function ModumToken() {
        owner = msg.sender;
    }
    
    function proposal(string _addr, bytes32 _hash, uint _value) returns (bool) {
        //only if no proposal is pending
        if(currentProposal.value > 0) {
            return false;
        }
        uint _yay = unlockedTokens; // default is yes
        uint _nay = 0;
        currentProposal = Proposal(_addr, _hash, _value, now, _yay, _nay, 0);
        return true;
    }
    
    function vote(bool _vote) returns (uint) {
        uint votingPower = accounts[msg.sender];
        if(! _vote) {
            currentProposal.yay -= votingPower;
            currentProposal.nay += votingPower;
        }
        currentProposal.nrVotes++;
        return votingPower;
    }
    
    function collect() returns (bool) {
        if(msg.sender == owner &&
                currentProposal.value > 0 &&
                now > currentProposal.startTime + votingDuration &&
                currentProposal.yay > currentProposal.nay &&
                lockedTokens - currentProposal.value >=0) {
            accounts[msg.sender] += currentProposal.value;
            unlockedTokens += currentProposal.value;
            lockedTokens -= currentProposal.value;
            delete currentProposal;
            return true;
        }
        return false;
    }
    
    function mint(address _recipient, uint _value) returns (bool) {
        if(msg.sender == owner && 
                (now < mintPhaseEnd) &&
                (lockedTokens + unlockedTokens + _value) <= maxTokens) {
            accounts[_recipient] += _value;
            unlockedTokens += _value;
            return true;
        }
        return false;
    }
    
    function totalSupply() constant returns (uint) {
        return unlockedTokens + lockedTokens;
    }
    function balanceOf(address _owner) constant returns (uint balance) {
        return accounts[_owner];
    }
    function transfer(address _to, uint _value) returns (bool success) {
        if (accounts[msg.sender] >= _value 
            && _value > 0
            && accounts[_to] + _value > accounts[_to]) {
            accounts[msg.sender] -= _value;
            accounts[_to] += _value;
            Transfer(msg.sender, _to, _value);
            return true;
        } else {
            return false;
        }
    }
    
    function transferFrom(address _from, address _to, uint _value) returns (bool success) {
        if (accounts[_from] >= _value
            && allowed[_from][msg.sender] >= _value
            && _value > 0
            && accounts[_to] + _value > accounts[_to]) {
            accounts[_from] -= _value;
            allowed[_from][msg.sender] -= _value;
            accounts[_to] += _value;
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

