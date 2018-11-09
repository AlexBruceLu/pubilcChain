package blc

import (
	"fmt"
	"log"

	"github.com/bolt"
)

const BlockDB = "blockChain.db"
const BlockBucket = "blocks"

type BlockChain struct {
	Tip []byte // 最后一个区块的哈希
	// Blocks []*Block
	DB *bolt.DB
}

// 创建带有创世区块的BlockChain
func NewBlockChain() *BlockChain {
	//return &BlockChain{[]*Block{NewGenesisBlock()}}
	// 1. 创建或打开数据库
	// 2. db.update
	var tip []byte
	db, err := bolt.Open(BlockDB, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockBucket))
		if bucket == nil {
			fmt.Println("No found BlockBucket, Creating a new bucket")
			genesis := NewGenesisBlock()
			bucket, err = tx.CreateBucket([]byte(BlockBucket))
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put([]byte("1"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash

		} else {
			tip = bucket.Get([]byte("1"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &BlockChain{tip, db}

}

// 新增区块
func (blockChain *BlockChain) AddBlock(data string) {
	//	preBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	//	newBlock := NewBlock(data, preBlock.Hash)
	//	blockChain.Blocks = append(blockChain.Blocks, newBlock)
	newBlock := NewBlock(data, blockChain.Tip)
	// upodate 数据
	err := blockChain.DB.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(BlockBucket))
		if bucket == nil {

			fmt.Println("NO found blockchain. Pls. check!")
			return nil

		}

		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = bucket.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		blockChain.Tip = newBlock.Hash
		return nil

	})
	if err != nil {
		log.Panic(err)
	}

}
