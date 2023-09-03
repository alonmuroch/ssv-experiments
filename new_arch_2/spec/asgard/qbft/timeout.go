package qbft

import "time"

var (
	quickTimeoutThreshold = 8               //nolint
	quickTimeout          = 2 * time.Second //nolint
	slowTimeout           = 2 * time.Minute //nolint
	// CutoffRound which round the instance should stop its timer and progress no further
	CutoffRound = 15 // stop processing instances after 8*2+120*6 = 14.2 min (~ 2 epochs)
)
