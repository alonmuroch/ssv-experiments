package example

import "ssv-experiments/mdvt/module"

type DAMAPIClient struct {
	storage *Storage
}

func NewAPIClient(s *Storage) *DAMAPIClient {
	return &DAMAPIClient{
		storage: s,
	}
}

func (api *DAMAPIClient) AddedToCluster(id uint64, share *module.Share, encryptedShare []byte) {
	if err := api.storage.AddedToCluster(id, share, encryptedShare); err != nil {
		// TODO
	}
}

func (api *DAMAPIClient) RemovedFromCluster(id uint64) {
	delete(api.storage.Clusters, id)
}
