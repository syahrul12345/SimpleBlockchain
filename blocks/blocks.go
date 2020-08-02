package blocks

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"log"
	"math/big"

	"github.com/syndtr/goleveldb/leveldb"
)

// Block : A block in the blockchain
type Block struct {
	Height   []byte
	Hash     []byte
	PrevHash []byte
	Data     []byte
}

// CreateBlock will create a block with the provided data
// Provide identifier if it is a genesis
func CreateBlock(data string, prevhash string) *Block {
	hashString := bytes.Join(append([][]byte{}, []byte(data), []byte(prevhash)), []byte{})
	hash := sha256.Sum256(hashString)
	return &Block{
		Hash:     hash[:],
		PrevHash: []byte(prevhash),
		Data:     []byte(data),
	}
}

// Save will save the block.
// Save the latest height and the blockhash
func (b *Block) Save(utilsdb, heightdb, blockdb *leveldb.DB) error {
	// Get the previous block height
	height, err := utilsdb.Get([]byte("chain_height"), nil)
	var newHeight []byte
	if err != nil {
		log.Println("Error getting height, might be genesis")
		newHeight = []byte{0}
	} else {
		// Append 1 to the new height
		newHeight = big.NewInt(0).Add(big.NewInt(0).SetBytes(height), big.NewInt(1)).Bytes()
	}
	b.Height = newHeight
	blockSerialized, _ := json.Marshal(b)
	// update chain height
	err = utilsdb.Put([]byte("chain_height"), b.Height, nil)
	if err != nil {
		log.Println("Failed to update height to 0")
		return err
	}
	// Put hash -> block data
	// Put height -> hash -> block data
	err = heightdb.Put(b.Height, b.Hash, nil)
	err = blockdb.Put(b.Hash, blockSerialized, nil)
	if err != nil {
		log.Println("Failed to create block")
		return err
	}
	return nil
}
