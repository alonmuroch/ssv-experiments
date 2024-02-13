pragma solidity ^0.4.0;

interface IOperator {
    /// @notice Registers a new operator, should be called by the ssv contract
    function registerOperator(uint64 id);

    /// @notice Removes an existing operator, should be called by the ssv contract
    function removeOperator(uint64 id);

    /// @notice Returns true if all operators are registered
    function areOperatorRegistered(uint64[] ids) external view returns (bool);
}
