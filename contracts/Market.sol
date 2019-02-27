pragma solidity ^0.4.0;

import "./SafeMath.sol";
import "./Ownable.sol";

contract Market is Ownable {
    using SafeMath for uint;
    uint256 public lastDealId = 1;
    uint256 depositAmount = 3 ether;

    struct Deal {
        address senderAddress;
        address receiverAddress;
        uint256 deposit;
    }

    mapping(uint256 => Deal) public deals;

    //Getters
    function getDealSender(uint256 id) public view returns (address) {
        return deals[id].senderAddress;
    }

    function getDealReceiver(uint256 id) public view returns (address) {
        return deals[id].receiverAddress;
    }

    function getDealDeposit(uint256 id) public view returns (uint256) {
        return deals[id].deposit;
    }


    function createDeal(address sender, address receiver, uint256 deposit) internal returns (uint256) {
        Deal storage p = deals[lastDealId];
        uint256 id = lastDealId;
        p.senderAddress = sender;
        p.receiverAddress = receiver;
        p.deposit = deposit;

        deals[lastDealId] = p;
        lastDealId = lastDealId.add(1);
        return id;
    }


    function setDealSender(address sender, uint256 id) public onlyOwner returns (bool) {
        Deal storage p = deals[id];
        p.senderAddress = sender;
        deals[id] = p;
        return true;
    }

    function setDealReceiver(address receiver, uint256 id) public onlyOwner returns (bool) {
        Deal storage p = deals[id];
        p.receiverAddress = receiver;
        deals[id] = p;
        return true;
    }

    function returnDeposit(uint256 id) public onlyOwner returns (bool) {
        Deal storage p = deals[id];
        uint256 amount = p.deposit;
        p.deposit = 0;
        p.receiverAddress.transfer(amount);
        deals[id] = p;
        return true;
    }

    // ----------------------------------------------------------------------------
    // Deposit management mechanism
    //
    // ----------------------------------------------------------------------------
    //Two mappings to hold addresses of senders and receivers
    mapping(address => address) sender_receiver;
    mapping(address => address) receiver_sender;

    /**
        @dev record sender/receiver relationship
        @param _sender      purchase ETH sender address
        @param _receiver    purchase Token receiver address
    */
    function setSenderReceiverPair(address _sender, address _receiver) public onlyOwner {
        require(_sender != 0 && _receiver != 0);

        bool senderUniqueForReceiver = sender_receiver[_sender] == 0;
        bool senderUniqueForSender = receiver_sender[_sender] == 0;

        bool receiverNotInTokenEth = receiver_sender[_receiver] == 0;
        bool receiverNotInEthToken = sender_receiver[_receiver] == 0;

        if (senderUniqueForReceiver && senderUniqueForSender &&
        receiverNotInEthToken && receiverNotInTokenEth) {
            sender_receiver[_sender] = _receiver;
            receiver_sender[_receiver] = _sender;
        } else {
            revert();
        }
    }


    //Check who receives goods for the indicated sender
    function getReceiver(address _sender) public view returns (address){
        return sender_receiver[_sender];
    }

    //Check who sends goods for the indicated receiver
    function getTokenSender(address _receiver) public view returns (address){
        return receiver_sender[_receiver];
    }


    /**
       @dev record new deposit
       @param _sender      ETH sender address
       @param _receiver    receiver address
    */

    mapping(address => uint256) public deposits;

    function getDeposit(address _investor) public view returns (uint256) {
        return deposits[_investor];
    }

    function() payable public {
        if (msg.value > 0) {
            require(msg.value == depositAmount);
            address _receiver = getReceiver(msg.sender);
            require(_receiver != 0);
            uint256 id = createDeal(msg.sender, _receiver, msg.value);
            deposits[_receiver] = deposits[_receiver].add(msg.value);
        } else {
            revert();
        }
    }
}