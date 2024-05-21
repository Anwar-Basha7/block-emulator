package emulator

import (
    "errors"
    "fmt"
)

type Block struct {
    ID   int
    Data string
}

func NewBlock(id int, data string) (*Block, error) {
    if id < 0 {
        return nil, errors.New("block ID cannot be negative")
    }
    if data == "" {
        return nil, errors.New("block data cannot be empty")
    }
    return &Block{
        ID:   id,
        Data: data,
    }, nil
}

func (b *Block) String() string {
    return fmt.Sprintf("Block ID: %d, Data: %s", b.ID, b.Data)
}
