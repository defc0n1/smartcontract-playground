pragma solidity ^0.4.10;

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


contract VoteCollector{
    function countVote(bytes32 option, uint votes);
}

contract ModumToken is ERC20Interface {
    bool mintingAllowed = true;
    
    uint avaiableSupply = 0;
    uint lockedSupply = 0;
    uint lockCorrectedAirdroppedWei = 0;

    address owner;
    Proposal proposal;
    
    mapping(address => TokenHolder) balances;
    mapping(address => mapping (address => uint)) allowed;
    
    string public constant name = "Modum Token";
    string public constant symbol = "MOD";
    
    event Voted(address voter, VoteCollector voteCollector, bytes32 option, uint256 votes);
    event WeiDropped(uint ethAmount, uint elegibleTokens);

    
    function ModumToken(){
        owner = msg.sender;
    }
    
    struct Proposal{
        VoteCollector voteCollector;
        uint blockStart;
        uint blockEnd;
    }
    
    struct TokenHolder{
        uint lastVoteReceive;
        uint lastAirDropReceived;
        uint avaiableTokensForVote;
        uint shareTokensAmount;
        uint weiAirdropAvaiable;
    }
    
     function proposalIsPresent() returns(bool){
        return proposal.blockEnd != 0;
    }
    
    function proposalIsActive() returns(bool){
        return proposalIsPresent()  && block.number >= proposal.blockStart && block.number < proposal.blockEnd;
    }
    
    function mint(address receiver, uint value){
        require(msg.sender == owner);
        require(!proposalIsActive());
        require(mintingAllowed);
        TokenHolder storage from = touchAndGet(receiver);
        from.shareTokensAmount += value;
        avaiableSupply+= value;
    }
    
    function mintLocked(uint value){
        require(msg.sender == owner);
        require(!proposalIsActive());
        require(mintingAllowed);
        lockedSupply += value;
    }
    
    function stopMinting(){
        require(msg.sender == owner);
        require(mintingAllowed);
        mintingAllowed = false;
    }
    
    function airDrop() payable{
        require(!mintingAllowed);
        uint totSup = avaiableSupply+lockedSupply;
        require(avaiableSupply != 0);
        uint lockCorrected = (msg.value*totSup)/avaiableSupply;
        lockCorrectedAirdroppedWei += lockCorrected;
        WeiDropped(msg.value, avaiableSupply);
    }
    
    function registerProposal(VoteCollector _proposal, uint blockStart, uint blockEnd){
        require(msg.sender == owner);
        require(!proposalIsPresent());
        require(blockStart < blockEnd && blockStart > block.number);
        proposal = Proposal(_proposal, blockStart, blockEnd);
    }
    
    //Gets the holder and calculate current votes, always use this as getter, never the map directly
    function touchAndGet(address account) internal returns(TokenHolder storage){
        require(!mintingAllowed);
        TokenHolder storage holder =  balances[account];
        if(proposalIsActive()){
            if(holder.lastVoteReceive < proposal.blockStart){
                holder.avaiableTokensForVote = holder.shareTokensAmount;
            }   
            holder.lastVoteReceive = block.number;
        } else {
            holder.avaiableTokensForVote = 0;
        }
        
        uint remindingDrop = lockCorrectedAirdroppedWei - holder.lastAirDropReceived;
        if(remindingDrop != 0){
            holder.weiAirdropAvaiable += (remindingDrop*holder.shareTokensAmount)/(lockedSupply+avaiableSupply);
            holder.lastAirDropReceived = lockCorrectedAirdroppedWei;
        }
        return holder;
    }
    
    function unlockSupply(address receiver, uint amount){
        require(!proposalIsActive());
        require(proposalIsPresent() && msg.sender == address(proposal.voteCollector));
        require(lockedSupply >= amount);
        lockedSupply -= amount;
        avaiableSupply += amount;
        touchAndGet(receiver).shareTokensAmount += amount;
    }
    
    function deleteProposal(){
       require(!proposalIsActive());
       require(proposalIsPresent() && msg.sender == address(proposal.voteCollector));
       delete proposal;
    }
    
    function vote(bytes32 option){
        require(proposalIsActive());
        TokenHolder storage voter = touchAndGet(msg.sender);
        uint votes = voter.avaiableTokensForVote;
        require(votes > 0);
        voter.avaiableTokensForVote = 0;
        proposal.voteCollector.countVote(option, votes);
        Voted(msg.sender, proposal.voteCollector, option, votes);
    }
    
    function withdrawWei(){
        TokenHolder storage account = touchAndGet(msg.sender);
        uint eth = account.weiAirdropAvaiable;
        if(eth == 0) return;
        account.weiAirdropAvaiable = 0;
        msg.sender.transfer(eth);
    }
    
    function totalSupply() constant returns (uint256 totalSupply){
        return lockedSupply+avaiableSupply;
    }

    // Get the account balance of another account with address _owner
    function balanceOf(address _owner) constant returns (uint256 balance){
        return balances[_owner].shareTokensAmount;
    }
    
    function votingPower(address _owner) returns (uint256 balance){
        return touchAndGet(_owner).avaiableTokensForVote;
    }
    
    function weiAvaiable(address _owner) returns (uint256 balance){
        return touchAndGet(_owner).weiAirdropAvaiable;
    }

    // Send _value amount of tokens to address _to
    function transfer(address _to, uint256 _value) returns (bool success){
        TokenHolder storage from = touchAndGet(msg.sender);
        TokenHolder storage to = touchAndGet(_to);
        if(from.shareTokensAmount >= _value) return false;
        if(_value == 0) return false;
        from.shareTokensAmount -= _value;
        to.shareTokensAmount += _value;
        Transfer(msg.sender, _to, _value);
        return true;
    }
    
    function transferFrom(address _from, address _to, uint _value) returns (bool success) {
        TokenHolder storage from = touchAndGet(_from);
        TokenHolder storage to = touchAndGet(_to);
        
        if(from.shareTokensAmount >= _value) return false;
        if(allowed[_from][msg.sender] >= _value) return false;
        if(_value == 0) return false;
        
        from.shareTokensAmount -= _value;
        allowed[_from][msg.sender] -= _value;
        to.shareTokensAmount += _value;
        
        Transfer(_from, _to, _value);
        return true;
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