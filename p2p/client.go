package p2p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"simpleblockchain/blocks"

	"github.com/davecgh/go-spew/spew"
)

var client *P2PClient

// P2PClient is the client which makes p2p calls
type P2PClient struct {
	Connection net.Conn
}

// StartClient starts the p2p client
func StartClient() {
	// Starting P2P client
	// Get the seed nodes
	service := fmt.Sprintf("%s:%d", SEEDS[0], PORT)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)

	log.Printf("Connecting to node: %s", tcpAddr)

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
