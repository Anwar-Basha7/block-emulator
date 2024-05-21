package pbft

import (
    "testing"

    "block-emulator/emulator" // Import the emulator package
)

func TestPBFTNode(t *testing.T) {
    node := NewPBFTNode(0, 4)

    // Use the emulator.Block type instead of defining a new Block type
    block, err := emulator.NewBlock(1, "Block 1 data")
    if err != nil {
        t.Errorf("Error creating block: %v", err)
    }
    prePrepareMsg := PrePrepareMessage{Block: *block, View: 0, Seq: 1}

    node.HandlePrePrepare(prePrepareMsg)

    if len(node.PrepareMsgs[1]) != 1 {
        t.Errorf("Expected 1 prepare message, got %d", len(node.PrepareMsgs[1]))
    }

    prepareMsg := PrepareMessage{NodeID: 1, View: 0, Seq: 1, Digest: "1:Block 1 data"}
    node.HandlePrepare(prepareMsg)

    if len(node.PrepareMsgs[1]) != 2 {
        t.Errorf("Expected 2 prepare messages, got %d", len(node.PrepareMsgs[1]))
    }

    commitMsg := CommitMessage{NodeID: 1, View: 0, Seq: 1, Digest: "1:Block 1 data"}
    node.HandleCommit(commitMsg)

    if len(node.CommitMsgs[1]) != 1 {
        t.Errorf("Expected 1 commit message, got %d", len(node.CommitMsgs[1]))
    }

    node.HandleCommit(commitMsg)

    if len(node.Blocks) != 1 {
        t.Errorf("Expected 1 block, got %d", len(node.Blocks))
    }

    if node.Blocks[0].ID != 1 || node.Blocks[0].Data != "1:Block 1 data" {
        t.Errorf("Unexpected block data: %+v", node.Blocks[0])
    }
}
