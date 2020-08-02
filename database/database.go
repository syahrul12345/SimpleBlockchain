package database

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var blocksdb *leveldb.DB
var heightdb *leveldb.DB
var utilsdb *leveldb.DB
var err error

// InitDB intializes the db
func InitDB() {
	blocksdb, err = leveldb.OpenFile("./storage/blocks/", nil)
	if err != nil {
		panic(err)
	}
	utilsdb, err = leveldb.OpenFile("./storage/utils/", nil)
	if err != nil {
		panic(err)
	}
	heightdb, err = leveldb.OpenFile("./storage/height/", nil)
	if err != nil {
		panic(err)
	}
}

// GetBlocksDB returns the db that stores the blocks
func GetBlocksDB() *leveldb.DB {
	return blocksdb
}

// GetUtilsDB returns the utilsDB
func GetUtilsDB() *leveldb.DB {
	return utilsdb
}

// GetHeightDB returns the heightDB
func GetHeightDB() *leveldb.DB {
	return heightdb
}

// GetBlockByHeight returns the block based on height. Height has to be a hexadecimal string
func GetBlockByHeight(heightdb *leveldb.DB, blocksdb *leveldb.DB, heightHex []byte) ([]byte, error) {
	// Get the hash by height
	hash, err := heightdb.Get(heightHex, nil)
	if err != nil {
		return nil, err
	}
	// Using the hash as key, we can get the block
	blockSerializedBytes, err := blocksdb.Get(hash, nil)
	if err != nil {
		return nil, err
	}
	return blockSerializedBytes, nil
}

// GetBlockByHash returns the block based on the hash.
func GetBlockByHash(blocksdb *leveldb.DB, hashBytes []byte) ([]byte, error) {
	// Using the hash as key, we can get the block
	blockSerializedBytes, err := blocksdb.Get(hashBytes, nil)
	if err != nil {
		return nil, err
	}
	return blockSerializedBytes, nil
}
