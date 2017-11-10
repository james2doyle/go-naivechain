package blockchain

import (
  "crypto/sha256"
  "fmt"
  "time"
)

type (
  // Block is the unit of item inside the Chain array
  Block struct {
    Index        int
    PreviousHash string
    Timestamp    time.Time
    Data         string
    Hash         string
  }
)

// GetHashableString function
func (b Block) GetHashableString() string {
  return fmt.Sprintf("%s%s%s%s", b.Index, b.PreviousHash, b.Timestamp, b.Data)
}

// CreateHash function
func (b Block) CreateHash(index int, hash string, timestamp time.Time, data string) string {
  h := sha256.New()
  hashable := fmt.Sprintf("%s%s%s%s", index, hash, timestamp, data)
  h.Write([]byte(hashable))

  return fmt.Sprintf("%x", h.Sum(nil))
}

// GenerateChild creates a new block off the current block
func (b Block) GenerateChild(data string) Block {
  nextIndex := b.Index + 1
  nextTimestamp := time.Now()

  return Block{
    Index:        nextIndex,
    PreviousHash: b.Hash,
    Timestamp:    nextTimestamp,
    Data:         data,
    Hash:         b.CreateHash(nextIndex, b.Hash, nextTimestamp, data),
  }
}
