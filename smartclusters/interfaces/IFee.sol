pragma solidity ^0.4.0;

import "./ILiquidatable.sol";
import "./IBalance.sol";

interface IFee is ILiquidatable,IBalance {
    function setFee(uint64 operator) external;

    function resetFeeCharge(uint64[] operatorIDs) public;

    function startFeeCharge(uint64[] operatorIDs) public;

    function endFeeCharge(uint64[] operatorIDs) public ;

    /// @notice Returns the accumulated fee index for a specific operator ID and a payee
    function getAccumulatedIndex(address payee, uint64 operatorID) external view returns (uint256);
}
