package blockchain

import (
  "fmt"
  "time"
)

// NewChain create a new chain from scratch
func NewChain() *Chain {
  newChain := &Chain{}

  now := time.Now()
  genesisData := "my genesis block!!"

  newHashString := fmt.Sprintf("%s%s%s%s", 0, "0", now, genesisData)

  block := Block{
    Index:        0,
    PreviousHash: "0",
    Timestamp:    now,
    Data:         genesisData,
    Hash:         newChain.CreateHash(newHashString),
  }

  newChain.Links = append(newChain.Links, block)

  return newChain
}
