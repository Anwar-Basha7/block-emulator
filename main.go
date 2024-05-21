package main

import (
    "flag"
    "log"
    "sync"
    "time"

    "block-emulator/emulator"
    "block-emulator/logger"
    "block-emulator/pbft"
)

func main() {
    var nodeID, totalNodes, start, end, mode int
    flag.IntVar(&nodeID, "n", 0, "Node ID")
    flag.IntVar(&totalNodes, "N", 4, "Total number of nodes")
    flag.IntVar(&start, "s", 0, "Start block ID")
    flag.IntVar(&end, "S", 10, "End block ID")
    flag.IntVar(&mode, "m", 0, "Mode")
    flag.Parse()

    logger.Init()

    var wg sync.WaitGroup
    nodes := make([]*pbft.PBFTNode, totalNodes)
    for i := 0; i < totalNodes; i++ {
        nodes[i] = pbft.NewPBFTNode(i, totalNodes)
    }

    // Emulate blocks and start PBFT protocol
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := start; i <= end; i++ {
            block, err := emulator.NewBlock(i, "Data for block "+string(i))
            if err != nil {
                log.Printf("Error creating block %d: %v\n", i, err)
                continue
            }
            prePrepareMsg := pbft.PrePrepareMessage{
                Block: *block,
                View:  0,
                Seq:   i,
            }
            for _, node := range nodes {
                node.HandlePrePrepare(prePrepareMsg)
            }
            time.Sleep(1 * time.Second) // Simulate network delay
        }
    }()

    wg.Wait()
    log.Println("Emulation and PBFT consensus completed")
}
