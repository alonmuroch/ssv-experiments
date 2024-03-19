package optimizedbft

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"ssv-experiments/optimizedbft/utils"
	"testing"
)

func TestHappyFlow(t *testing.T) {
	states := []*State{
		NewState(),
		NewState(),
		NewState(),
	}

	share := &Share{
		Quorum: 3,
	}

	shares, _ := utils.GenerateShares(3, 4)
	data := []byte("hello world")

	msgs := []*Message{
		{Type: Propose, OperatorID: 1, Round: 1},

		{Type: Prepare, OperatorID: 1, Round: 1},
		{Type: Prepare, OperatorID: 2, Round: 1},
		{Type: Prepare, OperatorID: 3, Round: 1},

		{Type: Commit, OperatorID: 1, Round: 1, PartialSignature: shares[1].SignByte(data)},
		{Type: Commit, OperatorID: 2, Round: 1, PartialSignature: shares[2].SignByte(data)},
		{Type: Commit, OperatorID: 3, Round: 1, PartialSignature: shares[3].SignByte(data)},
	}

	for _, msg := range msgs {
		for _, s := range states {
			switch msg.Type {
			case Propose:
				require.NoError(t, s.ProcessPropose(share, msg))
			case Prepare:
				require.NoError(t, s.ProcessPrepare(share, msg))
			case Commit:
				require.NoError(t, s.ProcessCommit(msg))
			}
		}
	}

	for _, s := range states {
		require.True(t, s.MessageQuorumForRound(share, Commit, 1))
	}
}

func TestSingleRoundChangeBeforePrepare(t *testing.T) {
	states := []*State{
		NewState(),
		NewState(),
		NewState(),
	}

	share := &Share{
		Quorum: 3,
	}

	shares, _ := utils.GenerateShares(3, 4)
	data := []byte("hello world")

	msgs := []*Message{
		{Type: Propose, OperatorID: 1, Round: 1},

		{Type: Prepare, OperatorID: 1, Round: 1},
		{Type: Prepare, OperatorID: 2, Round: 1},

		{Type: Propose, OperatorID: 2, Round: 2},

		{Type: Prepare, OperatorID: 1, Round: 2},
		{Type: Prepare, OperatorID: 2, Round: 2},
		{Type: Prepare, OperatorID: 3, Round: 2},

		{Type: Commit, OperatorID: 1, Round: 2, PartialSignature: shares[1].SignByte(data)},
		{Type: Commit, OperatorID: 2, Round: 2, PartialSignature: shares[2].SignByte(data)},
		{Type: Commit, OperatorID: 3, Round: 2, PartialSignature: shares[3].SignByte(data)},
	}

	bumpRoundAfterMsgIndex := 2

	for i, msg := range msgs {
		for _, s := range states {
			switch msg.Type {
			case Propose:
				require.NoError(t, s.ProcessPropose(share, msg))
			case Prepare:
				require.NoError(t, s.ProcessPrepare(share, msg))
			case Commit:
				require.NoError(t, s.ProcessCommit(msg))
			}

			if i == bumpRoundAfterMsgIndex {
				s.BumpRound()
			}
		}
	}

	for _, s := range states {
		require.True(t, s.MessageQuorumForRound(share, Commit, 2))
	}
}

func TestSingleRoundChangeAfterPrepare(t *testing.T) {
	states := []*State{
		NewState(),
		NewState(),
		NewState(),
	}

	share := &Share{
		Quorum: 3,
	}

	shares, _ := utils.GenerateShares(3, 4)
	data := []byte("hello world")

	msgs := []*Message{
		{Type: Propose, OperatorID: 1, Round: 1, Data: data},

		{Type: Prepare, OperatorID: 1, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 2, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 3, Round: 1, Data: data},

		{Type: Commit, OperatorID: 1, Round: 1, Data: data, PartialSignature: shares[1].SignByte(data)},

		// change round
		{Type: Propose, OperatorID: 2, Round: 2, Data: data},

		{Type: Prepare, OperatorID: 1, Round: 2, Data: data},
		{Type: Prepare, OperatorID: 2, Round: 2, Data: data},
		{Type: Prepare, OperatorID: 3, Round: 2, Data: data},

		{Type: Commit, OperatorID: 1, Round: 2, Data: data, PartialSignature: shares[1].SignByte(data)},
		{Type: Commit, OperatorID: 2, Round: 2, Data: data, PartialSignature: shares[2].SignByte(data)},
		{Type: Commit, OperatorID: 3, Round: 2, Data: data, PartialSignature: shares[3].SignByte(data)},
	}

	bumpRoundAfterMsgIndex := 4

	for i, msg := range msgs {
		for _, s := range states {
			switch msg.Type {
			case Propose:
				require.NoError(t, s.ProcessPropose(share, msg))
			case Prepare:
				require.NoError(t, s.ProcessPrepare(share, msg))
			case Commit:
				require.NoError(t, s.ProcessCommit(msg))
			}

			if i == bumpRoundAfterMsgIndex {
				s.BumpRound()
			}
		}
	}

	for _, s := range states {
		require.True(t, s.MessageQuorumForRound(share, Commit, 2))
	}
}

func TestSingleRoundChangeWithWrongDataAfterPrepare(t *testing.T) {
	states := []*State{
		NewState(),
		NewState(),
		NewState(),
	}

	share := &Share{
		Quorum: 3,
	}

	shares, _ := utils.GenerateShares(3, 4)
	data := []byte("hello world")
	wrongData := []byte("goodbye world")

	msgs := []*Message{
		{Type: Propose, OperatorID: 1, Round: 1, Data: data},

		{Type: Prepare, OperatorID: 1, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 2, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 3, Round: 1, Data: data},

		{Type: Commit, OperatorID: 1, Round: 1, Data: data, PartialSignature: shares[1].SignByte(data)},

		// change round
		{Type: Propose, OperatorID: 2, Round: 2, Data: wrongData},
	}

	bumpRoundAfterMsgIndex := 4

	for i, msg := range msgs {
		for _, s := range states {
			switch msg.Type {
			case Propose:
				if msg.Round == 2 {
					require.EqualError(t, s.ProcessPropose(share, msg), "propose data invalid")
				} else {
					require.NoError(t, s.ProcessPropose(share, msg))
				}

			case Prepare:
				require.NoError(t, s.ProcessPrepare(share, msg))
			case Commit:
				require.NoError(t, s.ProcessCommit(msg))
			}

			if i == bumpRoundAfterMsgIndex {
				s.BumpRound()
			}
		}
	}

	for _, s := range states {
		require.False(t, s.MessageQuorumForRound(share, Commit, 2))
	}
}

func TestSplitPrepareAndRecover(t *testing.T) {
	states := []*State{
		NewState(),
		NewState(),
		NewState(),
		NewState(),
	}

	share := &Share{
		Quorum: 3,
	}

	shares, _ := utils.GenerateShares(3, 4)
	data := []byte("hello world")
	differentData := []byte("goodbye world")

	msgs := []*Message{
		{Type: Propose, OperatorID: 1, Round: 1, Data: data},

		{Type: Prepare, OperatorID: 1, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 2, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 3, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 4, Round: 1, Data: data},

		{Type: Commit, OperatorID: 1, Round: 1, Data: data, PartialSignature: shares[1].SignByte(data)},
		{Type: Commit, OperatorID: 2, Round: 1, Data: data, PartialSignature: shares[1].SignByte(data)},

		// change round
		{Type: Propose, OperatorID: 2, Round: 2, Data: differentData},

		{Type: Prepare, OperatorID: 3, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 4, Round: 1, Data: data},

		// change round
		{Type: Propose, OperatorID: 3, Round: 3, Data: data},

		{Type: Prepare, OperatorID: 1, Round: 3, Data: data},
		{Type: Prepare, OperatorID: 2, Round: 3, Data: data},
		{Type: Prepare, OperatorID: 3, Round: 3, Data: data},
		{Type: Prepare, OperatorID: 4, Round: 3, Data: data},

		{Type: Commit, OperatorID: 1, Round: 3, Data: data, PartialSignature: shares[1].SignByte(data)},
		{Type: Commit, OperatorID: 2, Round: 3, Data: data, PartialSignature: shares[1].SignByte(data)},
		{Type: Commit, OperatorID: 3, Round: 3, Data: data, PartialSignature: shares[1].SignByte(data)},
		{Type: Commit, OperatorID: 4, Round: 3, Data: data, PartialSignature: shares[1].SignByte(data)},
	}

	msgAcceptanceMatrix := map[uint64]map[MessageType]map[uint64]bool{
		1: {
			Propose: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
			Prepare: map[uint64]bool{1: true, 2: true, 3: false, 4: false},
			Commit:  map[uint64]bool{1: true, 2: true, 3: false, 4: false},
		},
		2: {
			Propose: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
			Prepare: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
		},
		3: {
			Propose: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
			Prepare: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
			Commit:  map[uint64]bool{1: true, 2: true, 3: true, 4: true},
		},
	}

	bumpRoundAfterMsgIndex := map[uint64]bool{
		6: true,
		9: true,
	}

	for i, msg := range msgs {
		for OperatorIndex, s := range states {
			if msgAcceptanceMatrix[msg.Round][msg.Type][uint64(OperatorIndex+1)] {
				switch msg.Type {
				case Propose:
					s.ProcessPropose(share, msg)
				case Prepare:
					s.ProcessPrepare(share, msg)
				case Commit:
					s.ProcessCommit(msg)
				}
			}

			if bumpRoundAfterMsgIndex[uint64(i)] {
				s.BumpRound()
			}
		}
	}

	require.EqualValues(t, 3, states[0].PreparedRound)
	require.EqualValues(t, 3, states[1].PreparedRound)
	require.EqualValues(t, 3, states[2].PreparedRound)
	require.EqualValues(t, 3, states[3].PreparedRound)
}

func TestLivenessFail(t *testing.T) {
	states := []*State{
		NewState(),
		NewState(),
		NewState(),
		NewState(),
	}

	share := &Share{
		Quorum: 3,
	}

	shares, _ := utils.GenerateShares(3, 4)
	data := []byte("hello world")
	differentData := []byte("goodbye world")
	differentData2 := []byte("goodbye2 world")

	msgs := []*Message{
		{Type: Propose, OperatorID: 1, Round: 1, Data: data},

		{Type: Prepare, OperatorID: 1, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 2, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 3, Round: 1, Data: data},
		{Type: Prepare, OperatorID: 4, Round: 1, Data: data},

		{Type: Commit, OperatorID: 1, Round: 1, Data: data, PartialSignature: shares[1].SignByte(data)},

		// change round
		{Type: Propose, OperatorID: 2, Round: 2, Data: differentData},

		{Type: Prepare, OperatorID: 2, Round: 2, Data: differentData},
		{Type: Prepare, OperatorID: 3, Round: 2, Data: differentData},
		{Type: Prepare, OperatorID: 4, Round: 2, Data: differentData},

		// change round
		{Type: Propose, OperatorID: 3, Round: 3, Data: differentData2},

		{Type: Prepare, OperatorID: 3, Round: 3, Data: differentData2},
		{Type: Prepare, OperatorID: 4, Round: 3, Data: differentData2},
	}

	msgAcceptanceMatrix := map[uint64]map[MessageType]map[uint64]bool{
		1: {
			Propose: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
			Prepare: map[uint64]bool{1: true, 2: false, 3: false, 4: false},
			Commit:  map[uint64]bool{1: true, 2: true, 3: true, 4: true},
		},
		2: {
			Propose: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
			Prepare: map[uint64]bool{1: false, 2: true, 3: false, 4: false},
			Commit:  map[uint64]bool{1: true, 2: true, 3: true, 4: true},
		},
		3: {
			Propose: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
			Prepare: map[uint64]bool{1: true, 2: true, 3: true, 4: true},
			Commit:  map[uint64]bool{1: true, 2: true, 3: true, 4: true},
		},
	}

	bumpRoundAfterMsgIndex := map[uint64]bool{
		6:  true,
		10: true,
	}

	for i, msg := range msgs {
		for OperatorIndex, s := range states {
			if msgAcceptanceMatrix[msg.Round][msg.Type][uint64(OperatorIndex+1)] {
				switch msg.Type {
				case Propose:
					s.ProcessPropose(share, msg)
				case Prepare:
					s.ProcessPrepare(share, msg)
				case Commit:
					s.ProcessCommit(msg)
				}
			}

			if bumpRoundAfterMsgIndex[uint64(i)] {
				s.BumpRound()
			}
		}
	}

	// f nodes prepared on round 1 value 1, f nodes prepare on round 2 value 2 and the rest can't prepare
	require.EqualValues(t, 1, states[0].PreparedRound)
	d, err := states[0].MessageQuorumDataForRound(Prepare, 1)
	require.NoError(t, err)
	require.EqualValues(t, data, d)

	require.EqualValues(t, 2, states[1].PreparedRound)
	d, err = states[1].MessageQuorumDataForRound(Prepare, 2)
	require.NoError(t, err)
	require.EqualValues(t, differentData, d)

	require.EqualValues(t, 0, states[2].PreparedRound)
	_, err = states[2].MessageQuorumDataForRound(Prepare, 2)
	require.EqualValues(t, err, fmt.Errorf("data not found"))
	require.EqualValues(t, differentData, d)

	require.EqualValues(t, 0, states[3].PreparedRound)
	_, err = states[3].MessageQuorumDataForRound(Prepare, 2)
	require.EqualValues(t, err, fmt.Errorf("data not found"))
	require.EqualValues(t, differentData, d)
}
