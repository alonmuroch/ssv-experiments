pragma solidity ^0.4.0;

interface IValidator {
    struct RegisterData {
        bytes[] publicKeys;
        uint64[] operatorIDs;
        bytes[] sharesData;
        bytes freeData; // used to forward implementation specific data, decoded however the implementation needs
    }

    event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares);

    /// @notice Registers a new validator, returns a unique group ID for which validators were registered.this
    /// should be called by the ssv contract
    function registerValidators(RegisterData calldata validators_data) public returns (uint256 id);

    /// @notice Removes existing validator, should be called by the ssv contract
    function removeValidators(bytes[] publicKeys, bytes freeData);
}
