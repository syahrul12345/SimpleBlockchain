package main

import (
	"simpleblockchain/database"
	"simpleblockchain/p2p"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	// Create block
	// block := blocks.CreateBlock("this is sixth block", "")
	// block.Save(database.GetUtilsDB(), database.GetHeightDB(), database.GetBlocksDB())

	// blockRes, err := database.GetBlockByHeight(database.GetHeightDB(), database.GetBlocksDB(), big.NewInt(6).Bytes())
	// if err != nil {
	// 	panic(err)
	// }
	// spew.Dump(blockRes)
	// block2, err := database.GetBlockByHash(database.GetBlocksDB(), block.Hash)
	// if err != nil {
	// 	panic(err)
	// }
	// spew.Dump(block2)
	go p2p.StartServer()
	go p2p.StartClient()
	StartWebServer()
}

//StartWebServer : Starts Start the webserver
func StartWebServer() {
	app := gin.New()
	app.Run(":800")
}
