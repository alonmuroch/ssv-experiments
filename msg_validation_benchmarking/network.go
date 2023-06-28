package main

import (
	"ssv-experiments/msg_validation_benchmarking/strategies"
	"time"
)

type Network struct {
	Fuzzer               *FuzzerHeartBeat
	BroadcastingDuration time.Duration
	// Mps - messages per second
	Mps int

	messagesC chan []byte
	stopC     chan bool
}

func NewNetwork(fuzzer *FuzzerHeartBeat, broadcastingDuration time.Duration, mps int) strategies.INetwork {
	return &Network{
		Fuzzer:               fuzzer,
		BroadcastingDuration: broadcastingDuration,
		Mps:                  mps,
		messagesC:            make(chan []byte),
		stopC:                make(chan bool),
	}
}

func (n *Network) Start() {
	go func() {
		t := time.NewTicker(time.Second)
		timePassed := float64(0)

	loop:
		for {
			select {
			case <-t.C:
				if timePassed >= n.BroadcastingDuration.Seconds() {
					n.stopC <- true
					break loop
				}

				for i := 0; i < n.Mps; i++ {
					go func() {
						n.messagesC <- test
					}()
				}

				timePassed++
			}
		}
	}()
}

func (n *Network) GetMessagesChannel() chan []byte {
	return n.messagesC
}

func (n *Network) GetStopChannel() chan bool {
	return n.stopC
}
