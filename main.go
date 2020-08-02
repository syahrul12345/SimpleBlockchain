package main

import (
	"encoding/hex"
	"simpleblockchain/database"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	database.InitDB()
	// block := blocks.CreateBlock("this is fifth block", "")
	// block.Save(database.GetUtilsDB(), database.GetHeightDB(), database.GetBlocksDB())

	block, err := database.GetBlockByHeight(database.GetHeightDB(), database.GetBlocksDB(), "01")
	if err != nil {
		panic(err)
	}
	spew.Dump(hex.EncodeToString(block.Hash))
	block2, err := database.GetBlockByHash(database.GetBlocksDB(), hex.EncodeToString(block.Hash))
	if err != nil {
		panic(err)
	}
	spew.Dump(hex.EncodeToString(block2.Hash))
}
