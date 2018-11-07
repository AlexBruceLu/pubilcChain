package blc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 20

var maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block   // 当前区块
	target *big.Int // 挖矿的难度
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) // 把一个数向左移位uint位
	pow := &ProofOfWork{block, target}
	return pow
}

//拼接数据返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{IntToHex(pow.block.TimeStamp), pow.block.PrevHash, IntToHex(int64(targetBits)), pow.block.Data, IntToHex(int64(nonce))}, []byte{})
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hash [32]byte
	var hashInt big.Int
	nonce := 0
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash) // \r循环打印时覆盖前面的数据
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 { //hashInt < pow.target -1
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}

// 判断pow的有效性
func (pow *ProofOfWork) ValidData() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
