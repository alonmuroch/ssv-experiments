pragma solidity ^0.4.0;

import "./interfaces/IBalance.sol";
import "./interfaces/IOperator.sol";
import "./interfaces/IValidator.sol";
import "./interfaces/IFee.sol";
import "./interfaces/ILiquidatable.sol";

contract SmartClusterFactory is IBalance {
    struct Instance {
        IOperator operators;
        IValidator validators;
        IFee fee;
    }

    mapping (address => Instance) instances;
    mapping (address => mapping (address => uint256)) balances;

    function initialize() external {

    }

    function newInstance(
        IOperator operators,
        IValidator validators,
        IFee fee
    ) external {
        instances[msg.sender] = Instance(balances, operators, validators, fee);
    }

    /// @notice Registers a new validator, called directly by the ssv contract
    function registerValidators(IValidator.RegisterData calldata validators_data) public {
        // operators registered are already verified by ssv
        // validator unique already verified by ssv
        // operators size 4,7,10,13 verified by ssv
        // pub key verified by ssv

        // verify instance exists
        if (instances[msg.sender] == address(0)) {
            revert("instance not found");
        }

        // call register
        instances[msg.sender].validators.registerValidators(validators_data);

        // update fee
        uint256 id = instances[msg.sender].fee.startFeeCharge(validators_data.operatorIDs);

        // verify not liquidatable
        if (instances[msg.sender].fee.isLiquidatable(msg.sender, id) == true) {
            revert ("not enough balance to register validator");
        }

        // validator added event fired on ssv side
    }

    function deposit(IERC20 token, uint256 amount) public {
        // move tokens
        token.transferFrom(msg.sender, address(this), amount);

        // update
        balances[msg.sender][address(token)] += amount;
    }

    function withdraw(IERC20 token, uint256 amount) public {
        // verify instance exists
        if (instances[msg.sender] == address(0)) {
            revert("instance not found");
        }



    }

    function getBalance(address owner, IERC20 token) external view returns (uint256) {
        return balances[owner][address(token)];
    }
}
