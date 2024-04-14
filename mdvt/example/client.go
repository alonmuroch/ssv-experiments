package example

import (
	"context"
	"fmt"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"ssv-experiments/mdvt"
)

// Client is an example DAM for a beacon chain validator client
type Client struct {
	storage   *Storage
	damAPI    mdvt.DAM
	nodeAPI   mdvt.API
	beaconAPI BeaconAPI
}

func New() *Client {
	s := NewStorage()

	return &Client{
		storage: s,
		damAPI:  NewAPIClient(s),
	}
}

func (c *Client) ExecuteProposalDuty(clusterID uint64, slot uint64) error {
	cluster := c.storage.Clusters[clusterID]
	if cluster == nil {
		return fmt.Errorf("cluster not found")
	}

	instance := &Instance{Slot: slot, ClusterID: clusterID}
	id, err := instance.HashTreeRoot()
	if err != nil {
		return err
	}

	// sign and collect randao signatures
	randaoSig, err := c.storage.SignBeaconObject(SSZUint64(SlotToEpoch(slot)), slot, DomainRandao)
	if err != nil {
		return err
	}
	randaoSignatures, err := c.nodeAPI.CollectSignatures(context.Background(), id, randaoSig)
	if err != nil {
		return err
	}
	reconstructedRandao, err := ReconstructPartialSignatures(randaoSignatures)
	if err != nil {
		return err
	}

	// decide on block
	block, err := c.beaconAPI.GetBlock(slot, reconstructedRandao.Serialize())
	if err != nil {
		return err
	}
	byts, err := block.MarshalSSZ()
	if err != nil {
		return err
	}
	validateF := func(in []byte) error {
		decoded := &deneb.BeaconBlock{}
		if err := decoded.UnmarshalSSZ(in); err != nil {
			return err
		}

		if err := c.storage.IsSlashableBlock(decoded); err != nil {
			return err
		}
		// un-marshal block, validate it, slashing check, etc
		return nil
	}
	decidedValue, _, err := c.nodeAPI.Decide(context.Background(), id, byts, validateF)
	if err != nil {
		return err
	}

	// sign and collect block
	decidedBlock := &deneb.BeaconBlock{}
	if err := decidedBlock.UnmarshalSSZ(decidedValue); err != nil {
		return err
	}
	blockSig, err := c.storage.SignBeaconObject(decidedBlock, slot, DomainProposer)
	if err != nil {
		return err
	}
	blockSignatures, err := c.nodeAPI.CollectSignatures(context.Background(), id, blockSig)
	if err != nil {
		return err
	}
	reconstructedBlock, err := ReconstructPartialSignatures(blockSignatures)
	if err != nil {
		return err
	}

	// broadcast
	return c.beaconAPI.SubmitBlock(decidedBlock, reconstructedBlock.Serialize())
}
