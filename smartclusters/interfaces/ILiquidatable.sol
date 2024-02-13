pragma solidity ^0.4.0;

interface ILiquidatable {

    function isLiquidatable(address owner, uint256 id) external view returns (bool);
}
