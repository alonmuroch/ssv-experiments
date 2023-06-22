package process

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch/p2p"
	qbft2 "ssv-experiments/new_arch/pipeline/qbft"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch/spec_test"
	"testing"
)

type SpecTest struct {
	Pre      *qbft.Instance
	Post     *qbft.Instance
	Messages []*qbft.SignedMessage `ssz-max:"256"`
}

// Test will run the test, fail if errors during test and will return a post run test object to be compared with
func (test *SpecTest) Test(t *testing.T) *spec_test.TestResult {
	p := qbft2.NewQBFTPipeline(test.Pre)
	for _, msg := range test.Messages {
		byts, err := msg.MarshalSSZ()
		require.NoError(t, err)

		p2pMessage := &p2p.Message{
			MsgType: p2p.SSVConsensusMsgType,
			MsgID:   msg.Message.Identifier,
			Data:    byts,
		}

		err, _ = p.ProcessMessage(p2pMessage)
		require.NoError(t, err)
	}

	return &spec_test.TestResult{
		ExpectedResult: []spec_test.TestObject{test.Post},
		Actual:         []spec_test.TestObject{test.Pre},
	}
}
