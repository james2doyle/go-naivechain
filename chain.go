package blockchain

import (
  "crypto/sha256"
  "errors"
  "fmt"
  "log"
)

type (
  // Chain is an array of Blocks
  Chain struct {
    Links []Block
  }
)

// GetBlock returns a block by the index
func (c Chain) GetBlock(index int) Block {
  return c.Links[index]
}

// GetLatestBlock returns the newest block
func (c Chain) GetLatestBlock() Block {
  if len(c.Links) > 0 {
    return c.Links[len(c.Links)-1]
  }

  return c.Links[0]
}

// GetBlockLength returns the block length
func (c Chain) GetBlockLength() int {
  return len(c.Links)
}

// CreateHash returns a hash for a given string
func (c Chain) CreateHash(hashString string) string {
  h := sha256.New()
  h.Write([]byte(hashString))

  return fmt.Sprintf("%x", h.Sum(nil))
}

// CalculateHashForBlock returns a hash for a given block
func (c Chain) CalculateHashForBlock(b Block) string {
  return c.CreateHash(b.GetHashableString())
}

// IsValidBlock compares 2 blocks
func (c Chain) IsValidBlock(b1 Block, b2 Block) (bool, error) {
  if b2.Index == 0 {
    return true, nil
  }

  if b1.PreviousHash != b2.Hash {
    log.Printf("block 1: %s and block 2: %s", b1.PreviousHash, b2.Hash)
    return false, errors.New("invalid previoushash")
  } else if c.CalculateHashForBlock(b1) != b1.Hash {
    return false, fmt.Errorf("invalid hash: %s %s", c.CalculateHashForBlock(b1), b1.Hash)
  }

  return true, nil
}

// IsValidNewBlock checks that the block being added is valid
func (c Chain) IsValidNewBlock(b1 Block, b2 Block) (bool, error) {
  if (b1.Index <= b2.Index) || (b1.Index != b2.Index+1) {
    return false, errors.New("invalid index")
  }

  return c.IsValidBlock(b1, b2)
}

// CheckValidity returns the status for the chain
func (c Chain) CheckValidity() (bool, error) {
  for _, b := range c.Links {
    // we made it to the end
    if b.Index-1 <= 0 {
      return true, nil
    }

    pass, err := c.IsValidBlock(b, c.GetBlock(b.Index-1))
    if err != nil {
      return false, err
    }
    if !pass {
      return false, fmt.Errorf("invalid block at %s with hash %s", b.Index, b.Hash)
    }
  }

  return true, nil
}

// AddBlock adds a new block to the chain
func (c *Chain) AddBlock(b Block) (Block, error) {
  pass, err := c.IsValidNewBlock(b, c.GetLatestBlock())
  if err != nil {
    return b, err
  }

  if pass {
    // c.Links[b.Index] = b
    c.Links = append(c.Links, b)
    return b, nil
  }

  return b, errors.New("failed to add new block")
}

// CreateNewBlock adds a block to the chain given some data
func (c *Chain) CreateNewBlock(data string) (Block, error) {
  newBlock := c.GetLatestBlock().GenerateChild(data)
  return c.AddBlock(newBlock)
}
