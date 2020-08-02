package p2p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"simpleblockchain/blocks"

	"github.com/davecgh/go-spew/spew"
)

// StartClient starts the p2p client
func StartClient() {
	// Starting P2P client
	service := fmt.Sprintf(":%d", PORT)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}

	// Write the blockheight that we want to the p2p server
	_, err = conn.Write(big.NewInt(4).Bytes())
	if err != nil {
		panic(err)
	}

	// Get the serialized block
	blockSerialized, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	block := &blocks.Block{}
	json.Unmarshal(blockSerialized, block)
	spew.Dump(block)
}
