package blc

import (
	"fmt"
	"time"
)

type Block struct {
	TimeStamp int64  //时间戳
	Data      []byte // 交易数据
	PrevHash  []byte // 前区块哈希
	Hash      []byte // 当前区块的哈希
	Nonce     int    // 难度值
}

// func (block *Block) SetHash() {
// 	// 将时间戳转换为字节数组
// 	timeStr := strconv.FormatInt(block.TimeStamp, 2) // 第二个参数2-36，代表转换的数字进制
// 	timeStamp := []byte(timeStr)
// 	// 将除了哈希以外的数据，转化为字节数组拼接起来
// 	header := bytes.Join([][]byte{timeStamp, block.Data, block.PrevHash, block.Hash}, []byte{})
// 	// 将拼接起来的数据进行256哈希
// 	hash := sha256.Sum256(header)
// 	// 将哈希值赋值给哈希属性
// 	block.Hash = hash[:]
// }

// 创建区块的工厂方法，返回新区块
func NewBlock(data string, prevBlockHash []byte) *Block {

	block := &Block{time.Now().Unix(), []byte(prevBlockHash), []byte(data), []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	// block.SetHash()
	isValid := pow.ValidData()
	fmt.Println(isValid)
	return block

}

// 添加创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
