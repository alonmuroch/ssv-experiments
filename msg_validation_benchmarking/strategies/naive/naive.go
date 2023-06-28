package naive

import (
	"fmt"
	"ssv-experiments/msg_validation_benchmarking/strategies"
	"sync"
	"time"
)

type Naive struct {
	network       strategies.INetwork
	msgValidation func(msgValidationF []byte)

	queue                  [][]byte
	poppedCnt              int
	accumulatedProcessTime time.Duration
	queueL                 sync.RWMutex
}

func NewNaiveStrategy() *Naive {
	return &Naive{}
}

func (n *Naive) Init(
	network strategies.INetwork,
	msgValidation func(msgValidationF []byte),
) {
	n.network = network
	n.msgValidation = msgValidation

	n.queue = make([][]byte, 0)
	n.queueL = sync.RWMutex{}
}

func (n *Naive) Start() {
	n.network.Start()
	go n.processMessages()

	// add msgs to queue
loop:
	for {
		select {
		case msg := <-n.network.GetMessagesChannel():
			n.AppendToQueue(msg)
		case <-n.network.GetStopChannel():
			fmt.Printf("naive: stopping\n")
			break loop
		}
	}
}

func (n *Naive) processMessages() {
	for {
		if msg := n.PopQueue(); msg != nil {
			n.poppedCnt++
			start := time.Now()
			n.msgValidation(msg)
			n.accumulatedProcessTime += time.Now().Sub(start)
		} else {
			time.Sleep(time.Millisecond * 5)
		}

	}
}

func (n *Naive) AppendToQueue(msg []byte) {
	n.queueL.Lock()
	defer n.queueL.Unlock()
	n.queue = append(n.queue, msg)
}

func (n *Naive) PopQueue() []byte {
	n.queueL.Lock()
	defer n.queueL.Unlock()

	if len(n.queue) == 0 {
		return nil
	}
	ret := n.queue[0]
	n.queue = n.queue[1:]
	return ret
}

func (n *Naive) GetQueue() [][]byte {
	n.queueL.RLock()
	defer n.queueL.RUnlock()

	return n.queue
}

func (n *Naive) GetQueueStats() (pending, popped int, avgProcessTimeSec float64) {
	avgProcessTimeSec = n.accumulatedProcessTime.Seconds() / float64(n.poppedCnt-1)
	return len(n.GetQueue()), n.poppedCnt, avgProcessTimeSec
}
