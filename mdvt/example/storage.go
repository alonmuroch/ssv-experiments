package example

import (
	"github.com/attestantio/go-eth2-client/spec/deneb"
	ssz "github.com/ferranbt/fastssz"
	"github.com/herumi/bls-eth-go-binary/bls"
	"ssv-experiments/mdvt/module"
)

type ExtraData struct {
	ValidatorPK [48]byte `ssz-size:"48"`
}

type ClusterData struct {
	SK        *bls.SecretKey
	Share     *module.Share
	ExtraData *ExtraData
}

type Storage struct {
	Clusters map[uint64]*ClusterData
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) AddedToCluster(id uint64, share *module.Share, encryptedShare []byte) error {
	ed := &ExtraData{}
	if err := ed.UnmarshalSSZ(share.ExtraData); err != nil {
		return err
	}

	sk := &bls.SecretKey{}
	if err := sk.Deserialize(encryptedShare); err != nil {
		return err
	}

	s.Clusters[id] = &ClusterData{
		SK:        sk,
		Share:     share,
		ExtraData: ed,
	}
	return nil
}

func (s *Storage) IsSlashableBlock(block *deneb.BeaconBlock) error {
	panic("implement")
}

func (s *Storage) SignBeaconObject(obj ssz.HashRoot, slot uint64, domain [4]byte) ([]byte, error) {
	panic("implement")
}
