pragma solidity ^0.4.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBalance {
    function deposit(IERC20 token, uint256 amount) public;

    function withdraw(IERC20 token, uint256 amount) public;

    function getBalance(address owner, IERC20 token) external view returns (uint256);
}
