// SPDX-License-Identifier: MIT
pragma solidity ^0.8.5;

contract trading {
    //dataIds
    uint256 public dataId;
    address public seller;
    address public arbitrator;
    address public buyer;
    bytes32 public delta_1;
    bytes32 public delta_2;
    uint256 public depth_2;
    bytes32 public mk;
    uint256 public p;
    uint256 public p1;
    bytes32 public kc;
    uint256 public opreationtime;
    uint  public timeLimit = 600 ;
    State public state;
    enum State { start, commitData, challenge, response, payleft, ReleaseSampleKey, Payleft, releaseMasterkey, dispute, done}
    //order info structure
    event commitDataEvent(uint256 id, bytes32 delta_1);
    event challengeEvent(bytes32 kc, uint256 p, uint256 p1);
    event respondeEvent(bytes32 _delta_2, uint256 _length2);
    event buyerRefundEvent(address userAddr);
    event payleftEvent(address userAddr, uint256 value);
    event releaseMasterkeyEvent(bytes32 masterkey);
    event initiationArbitrationEvent(address userAddr);
    event verdictEvent(address userAddr, address winner);
    event onchianArbitrationEvent(uint status);
    constructor (){
        arbitrator = xxxxx;
        seller = xxxxx;
        buyer = xxxxx;
        state = State.start;
    }
    
    function commitData (uint256 id, bytes32 _delta_1) public onlySeller {
       require(state == State.start);
        dataId = id;
        delta_1 = _delta_1;
        state = State.commitData;
        emit commitDataEvent(id, delta_1);
    }

    function challenge (bytes32 _kc, uint256 _p, uint256 _p1) payable public onlyBuyer{
        require(state == State.commitData);
        //check payment  
        require(msg.value == _p1);
        kc = _kc;
        p = _p;
        p1 = _p1;
        opreationtime = block.timestamp;
        state = State.challenge;
        emit challengeEvent(_kc, _p, _p1);
    }


    function response (bytes32 _delta_2, uint256 _depth_2) payable public {
       require(state == State.challenge);
       require(opreationtime + timeLimit > block.timestamp, "Responde time is over ");
        delta_2 = _delta_2;
        depth_2 = _depth_2;
        state = State.response;
        emit respondeEvent(_delta_2, _depth_2);
    }

    // timeout refund
    function buyerRefund() public {
        //check order state
        require(state == State.challenge || state == State.payleft);
        //check timeout
        require(opreationtime + timeLimit > block.timestamp, "Operation time is not over yet");
        payable(msg.sender).transfer(p1);
        state = State.done;
        emit buyerRefundEvent(msg.sender);
    }    

  
    function payleft() public payable onlyBuyer{
        //check order state
        require(state == State.response);
        require(msg.value == (p - p1));
        //change order state
        state = State.payleft;
        emit payleftEvent(msg.sender, msg.value);
    }
    
    function releaseMasterkey(bytes32 _mk) public onlySeller {
        // check order state
        require(state == State.payleft);
        mk = _mk;
        state = State.releaseMasterkey;
        emit releaseMasterkeyEvent(mk);
    }
     //on-chain arttribute
    function RaiseOnChainArbitration(bytes32 index, bytes32 _committedSubKey, bytes32[] memory _merkleTreePath) public  returns (string memory)   {
        bytes32 computedSubkey = hmac(mk, index);
        bytes32 committedNode;
        if (computedSubkey != _committedSubKey){
            committedNode = keccak256(abi.encode(_committedSubKey));
            for (uint i = 0; i < depth_2; i++)
                committedNode = keccak256(abi.encode(committedNode, _merkleTreePath[i]));
            if (committedNode == delta_2)
            {
                payable(msg.sender).transfer(10);
                state = State.done;
                return "1"; 
            }else{
                payable(seller).transfer(10);
                state = State.done;
                return "2";
            }
       }
       return "3";
    }

    function RaiseOffChainArbitration() public {
        // check order state
        require(state != State.ReleaseSampleKey);
        state = State.dispute; 
        emit initiationArbitrationEvent(msg.sender);
    }

    function verdict(address userAddr,address winner) public onlyArbitrator{
        // check order state
        require(state == State.dispute);
        if(userAddr == winner){
            //buyer win
            payable(userAddr).transfer(p1);
        }else {
            //seller win
            payable(seller).transfer(p1);
        }
        state = State.done;
        emit verdictEvent(userAddr, winner);
    }

    function hmac(bytes32 key, bytes32 message) public pure returns (bytes32) {
        bytes32  opad = 0x5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c;
        bytes32  ipad = 0x3636363636363636363636363636363636363636363636363636363636363636;
        bytes32 inner = sha2562(key ^ ipad, message);
        return sha2562(key ^ opad, inner);
    }
    
    
    function test256(bytes32 key)public pure returns (bytes32) {
        return sha256(abi.encodePacked(key));

    }
    function sha2562(bytes32 key, bytes32 message) public pure returns (bytes32) {
        return sha256(abi.encodePacked(key, message));
    }
    modifier onlySeller {
        require(msg.sender == seller);
        _;
    }
    modifier onlyBuyer {
        require(msg.sender == buyer);
        _;
    }
    modifier onlyArbitrator {
        require(msg.sender == arbitrator);
        _;
    }  
}
