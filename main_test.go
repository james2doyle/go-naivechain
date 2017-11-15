package blockchain

import (
  "fmt"
  "testing"
  "time"

  . "github.com/smartystreets/goconvey/convey"
)

func Test_Create_New_Chain(t *testing.T) {
  Convey("Create new chain", t, func() {
    defer func() {
      So(recover(), ShouldBeNil)
    }()

    myChain := NewChain()

    So(myChain.GetBlockLength(), ShouldEqual, 1)
    So(myChain.GetLatestBlock().Hash, ShouldEqual, myChain.GetBlock(0).Hash)
    So(myChain.GetLatestBlock().Index, ShouldEqual, myChain.GetBlock(0).Index)
    So(myChain.GetBlock(0).Index, ShouldEqual, 0)
  })
}

func Test_Add_Block_To_Chain(t *testing.T) {
  Convey("Add block to chain", t, func() {
    defer func() {
      So(recover(), ShouldBeNil)
    }()

    myChain := NewChain()

    block, err := myChain.CreateNewBlock("my new data")

    _, err2 := myChain.CreateNewBlock("my new data 2")
    So(err2, ShouldBeNil)

    _, err3 := myChain.CreateNewBlock("my new data 3")
    So(err3, ShouldBeNil)

    So(err, ShouldBeNil)
    So(block.Data, ShouldEqual, "my new data")
    So(myChain.GetBlockLength(), ShouldEqual, 4)
  })
}

func Test_Check_Validity_Of_The_Chain(t *testing.T) {
  Convey("Check validity of the chain", t, func() {
    defer func() {
      So(recover(), ShouldBeNil)
    }()

    myChain := NewChain()

    valid, err := myChain.CheckValidity()

    So(err, ShouldBeNil)
    So(valid, ShouldEqual, true)
  })
}

func Test_Check_Validity_Of_A_Long_Chain(t *testing.T) {
  Convey("Check validity of a long chain", t, func() {
    defer func() {
      So(recover(), ShouldBeNil)
    }()

    myChain := NewChain()

    myChain.CreateNewBlock("my new data 1")
    myChain.CreateNewBlock("my new data 2")
    myChain.CreateNewBlock("my new data 3")

    valid, err2 := myChain.CheckValidity()

    So(err2, ShouldBeNil)
    So(valid, ShouldEqual, true)
  })
}

func Test_A_Block_Can_Be_Added_To_A_Chain(t *testing.T) {
  Convey("Check a block can be added to a chain", t, func() {
    defer func() {
      So(recover(), ShouldBeNil)
    }()

    myChain := NewChain()

    nextIndex := 1
    nextTimestamp := time.Now()

    previousHash := myChain.GetLatestBlock().Hash

    newHash := fmt.Sprintf("%d%s%s%s", 1, previousHash, nextTimestamp, "new data here!")

    newBlock := Block{
      Index:        nextIndex,
      PreviousHash: previousHash,
      Timestamp:    nextTimestamp,
      Data:         "new data here!",
      Hash:         myChain.CreateHash(newHash),
    }

    block, err := myChain.AddBlock(newBlock)

    So(err, ShouldBeNil)
    So(block.Hash, ShouldEqual, newBlock.Hash)
    So(block.Hash, ShouldEqual, myChain.CreateHash(newHash))

    pass, err2 := myChain.IsValidBlock(block, myChain.GetBlock(block.Index-1))
    So(err2, ShouldBeNil)
    So(pass, ShouldEqual, true)

    valid, err3 := myChain.CheckValidity()
    So(err3, ShouldBeNil)

    So(valid, ShouldEqual, true)
  })
}

func Test_An_Invalid_Block_Cannot_Be_Added_To_A_Chain(t *testing.T) {
  Convey("Check an invalid block cannot be added to a chain", t, func() {
    defer func() {
      So(recover(), ShouldBeNil)
    }()

    myChain := NewChain()

    nextIndex := 4
    nextTimestamp := time.Now()

    previousHash := myChain.GetLatestBlock().Hash

    newHash := fmt.Sprintf("%d%s%s%s", 4, previousHash, nextTimestamp, "new data here!")

    newBlock := Block{
      Index:        nextIndex,
      PreviousHash: previousHash,
      Timestamp:    nextTimestamp,
      Data:         "new data here!",
      Hash:         myChain.CreateHash(newHash),
    }

    _, err := myChain.AddBlock(newBlock)

    So(err, ShouldNotBeNil)
  })
}
