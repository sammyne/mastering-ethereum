// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract Private_Bank {
    mapping(address => uint256) public balances;
    uint256 public MinDeposit = 1 ether;
    Log TransferLog;

    //function Private_Bank(address memory _log) {
    constructor(address _log) {
        TransferLog = Log(_log);
    }

    function Deposit() public payable {
        if (msg.value >= MinDeposit) {
            balances[msg.sender] += msg.value;
            TransferLog.AddMessage(msg.sender, msg.value, "Deposit");
        }
    }

    function CashOut(uint256 _am) public {
        if (_am <= balances[msg.sender]) {
            (bool ok,) = msg.sender.call{value: _am}("");
            if (ok) {
                balances[msg.sender] -= _am;
                TransferLog.AddMessage(msg.sender, _am, "CashOut");
            }
        }
    }

    receive() external payable {}
}

contract Log {
    struct Message {
        address Sender;
        string Data;
        uint256 Val;
        uint256 Time;
    }

    Message[] public History;
    Message LastMsg;

    function AddMessage(
        address _adr,
        uint256 _val,
        string memory _data
    ) public {
        LastMsg.Sender = _adr;
        LastMsg.Time = block.timestamp;
        LastMsg.Val = _val;
        LastMsg.Data = _data;
        History.push(LastMsg);
    }
}
