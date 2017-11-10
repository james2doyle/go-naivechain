# go-naivechain

> An implementation of a block chain in Go inspired by the naivechain npm package

Note:

_I am not a crypto expert (or a blockchain expert) so this could be a terrible implementation of a blockchain, or have bugs, so please be kind. I am also an amateur Go programmer, so if there is room to improve, please let me know._

Usage:

To install go-naivechain:

```
$ go get github.com/james2doyle/go-naivechain
```

Example:

```go
package main

import (
  "log"

  "github.com/james2doyle/blockchain"
)

func main() {
  // create a new chain
  myChain := blockchain.NewChain()

  // add some data to the chain
  _, err := myChain.CreateNewBlock("my new data")
  if err != nil {
    log.Printf("%v", err)
  }

  // check that the chain is valid
  valid, err := myChain.CheckValidity()
  if err != nil {
    log.Printf("%v", err)
  }

  log.Printf("Chain is valid: %t", valid)
}
```
