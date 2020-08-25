// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract EtherGame {
    uint256 public payoutMileStone1 = 3 ether;
    uint256 public mileStone1Reward = 2 ether;
    uint256 public payoutMileStone2 = 5 ether;
    uint256 public mileStone2Reward = 3 ether;
    uint256 public finalMileStone = 10 ether;
    uint256 public finalReward = 5 ether;
    uint256 public depositedWei;

    mapping(address => uint256) redeemableEther;

    function play() public payable {
        require(msg.value == 0.5 ether);
        uint256 currentBalance = depositedWei + msg.value;
        // ensure no players after the game has finished
        require(currentBalance <= finalMileStone);
        if (currentBalance == payoutMileStone1) {
            redeemableEther[msg.sender] += mileStone1Reward;
        } else if (currentBalance == payoutMileStone2) {
            redeemableEther[msg.sender] += mileStone2Reward;
        } else if (currentBalance == finalMileStone) {
            redeemableEther[msg.sender] += finalReward;
        }
        depositedWei += msg.value;
        return;
    }

    function claimReward() public {
        // ensure the game is complete
        require(depositedWei == finalMileStone);
        // ensure there is a reward to give
        require(redeemableEther[msg.sender] > 0);

        uint256 transferValue = redeemableEther[msg.sender];
        redeemableEther[msg.sender] = 0;
        msg.sender.transfer(transferValue);
    }
}
