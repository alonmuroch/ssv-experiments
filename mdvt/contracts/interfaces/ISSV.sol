// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ISSV {
    /// @notice Registers a DAM to the SSV network
    /// @param damAddress the DAM contract address
    function registerDAM(address damAddress) external;

    /**
    * @dev Emitted when a new DAM is added
     * @param damAddress the DAM contract address
     * @param id is the unique DAM ID
     */
    event DAMAdded(address indexed damAddress, uint indexed id);

    /// @notice Updates cluster fee rate
    /// @param id is the unique DAM ID
    /// @param new fee rate to be set
    /// @param amount to deposit with update, could be 0
    function updateClusterFeeRate(uint id, uint newFeeRate, uint amountToDeposit) external;

    /**
   * @dev Emitted when a new DAM is added
     * @param id is the unique DAM ID
     * @param new fee rate to be set
     * @param current cluster fee balance
     */
    event FeeRateUpdated(uint id, uint newFeeRate, uint balance);
}