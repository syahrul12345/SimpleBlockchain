package p2p

import (
	"fmt"
	"log"
	"net"
	"simpleblockchain/database"
)

// Maps of seed nodes.
// Define here on initial launch
// Other nodes should use the same executable
const (
	PORT int = 7788
)

// StartServer starts the p2pserver
func StartServer() {
	log.Printf("Staring the p2p server at port: %d", PORT)

	// Set up the tcp client
	service := fmt.Sprintf(":%d", PORT)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		panic(err)
	}

	// Get the client to listen
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	// Blocks the thread and listens to connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			// Skip this loop
			continue
		}
		go handleClient(conn)
	}
}

// handleclient helper to hanlde all clients which are conencting
// returns the block which is requested
func handleClient(conn net.Conn) {
	log.Println("Connection to node detected")
	defer conn.Close()
	// Set up a buffer 512 bytes long
	var buf [512]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		panic(err)
	}
	// Check the height and simply returns the block
	height := buf[0:n]
	blockSerialized, err := database.GetBlockByHeight(database.GetHeightDB(), database.GetBlocksDB(), height)
	if err != nil {
		panic(err)
	}
	conn.Write(blockSerialized)
}
