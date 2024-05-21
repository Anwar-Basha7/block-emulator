package pbft

import (
	"fmt"
	"sync"

	"block-emulator/logger"
	"block-emulator/emulator"
)

type PBFTNode struct {
	ID          int
	TotalNodes  int
	PrepareMsgs map[int][]PrepareMessage
	CommitMsgs  map[int][]CommitMessage
	mu          sync.Mutex
}

type PrePrepareMessage struct {
	Block emulator.Block // Use the Block type from the emulator package
	View  int
	Seq   int
}

type PrepareMessage struct {
	NodeID int
	View   int
	Seq    int
	Digest string
}

type CommitMessage struct {
	NodeID int
	View   int
	Seq    int
	Digest string
}

func NewPBFTNode(id int, totalNodes int) *PBFTNode {
	return &PBFTNode{
		ID:          id,
		TotalNodes:  totalNodes,
		PrepareMsgs: make(map[int][]PrepareMessage),
		CommitMsgs:  make(map[int][]CommitMessage),
	}
}

func (node *PBFTNode) HandlePrePrepare(msg PrePrepareMessage) {
	logger.InfoLogger.Printf("Node %d received PrePrepare message: %+v", node.ID, msg)

	prepareMsg := PrepareMessage{
		NodeID: node.ID,
		View:   msg.View,
		Seq:    msg.Seq,
		Digest: fmt.Sprintf("%d:%s", msg.Block.ID, msg.Block.Data),
	}

	node.mu.Lock()
	node.PrepareMsgs[msg.Seq] = append(node.PrepareMsgs[msg.Seq], prepareMsg)
	node.mu.Unlock()

	node.BroadcastPrepare(prepareMsg)
}

func (node *PBFTNode) BroadcastPrepare(msg PrepareMessage) {
	// Simulate sending prepare message to other nodes
	for i := 0; i < node.TotalNodes; i++ {
		if i != node.ID {
			// Here we would normally send the message over the network
			logger.InfoLogger.Printf("Node %d broadcasts Prepare message to Node %d: %+v", node.ID, i, msg)
		}
	}
}

func (node *PBFTNode) HandlePrepare(msg PrepareMessage) {
	logger.InfoLogger.Printf("Node %d received Prepare message: %+v", node.ID, msg)

	node.mu.Lock()
	node.PrepareMsgs[msg.Seq] = append(node.PrepareMsgs[msg.Seq], msg)
	prepareCount := len(node.PrepareMsgs[msg.Seq])
	node.mu.Unlock()

	if prepareCount >= (2 * node.TotalNodes / 3) {
		commitMsg := CommitMessage{
			NodeID: node.ID,
			View:   msg.View,
			Seq:    msg.Seq,
			Digest: msg.Digest,
		}

		node.mu.Lock()
		node.CommitMsgs[msg.Seq] = append(node.CommitMsgs[msg.Seq], commitMsg)
		node.mu.Unlock()

		node.BroadcastCommit(commitMsg)
	}
}

func (node *PBFTNode) BroadcastCommit(msg CommitMessage) {
	// Simulate sending commit message to other nodes
	for i := 0; i < node.TotalNodes; i++ {
		if i != node.ID {
			// Here we would normally send the message over the network
			logger.InfoLogger.Printf("Node %d broadcasts Commit message to Node %d: %+v", node.ID, i, msg)
		}
	}
}

// HandleCommit method will be added here
