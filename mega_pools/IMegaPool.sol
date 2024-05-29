// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IMegaPool {
    /// @notice Gets the mega-pool capacity
    /// @return overall (per owner) validator capacity and per admin node capacity
    function getCapacity() external view returns (uint,uint);

    /// @notice Gets the number of registered validators for admin
    /// @param owner of registered validators
    /// @return number of registered validators
    function getRegisteredValidators(address owner) external view returns (uint);

    /// @notice Returns true if address is owner
    /// @param owner as defined in the contract
    /// @return true if owner
    function isOwner(address owner) external view returns (bool);

    /// @notice Registers operators for an owner (per capacity)
    /// @param publicKeys is an array of public keys (per the number of registering operators)
    /// @return new operator ids
    function registerOperators(bytes[] calldata publicKeys) external returns (uint64[]);
}
