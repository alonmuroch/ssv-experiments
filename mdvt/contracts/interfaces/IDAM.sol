// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IDAM {
    /**
    * @dev Emitted when a operators are added to a cluster.
     * @param id is the unique DAM ID
     * @param operatorIds The operator IDs managing the cluster.
     * @param shares snappy compressed shares(a set of encrypted and public shares)
     */
    event ClusterAddedOperators(uint indexed id, uint64[] operatorIds, bytes shares, bytes extraData);
    /**
    * @dev Emitted when a operators are added to a cluster.
     * @param id is the unique DAM ID
     * @param operatorIds The operator IDs managing the cluster.
     * @param shares snappy compressed shares(a set of encrypted and public shares)
     */
    event ClusterRemovedOperators(uint indexed id, uint64[] operatorIds, bytes shares, bytes extraData);
}
