package blockchain

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/ror0/wiki/common"
)

type Block struct {
	version int64
	time int64
	prev_hdr_hash []byte
	hash []byte

	data []byte
}

func (b *Block) CalcHeaderHash() []byte {
	version := []byte(strconv.FormatInt(b.version, 10))
	time := []byte(strconv.FormatInt(b.time, 10))

	headerBytes := bytes.Join(
		[][]byte{version, time, b.prev_hdr_hash, b.hash},
		[]byte{},
	)

	return common.CalcHash(headerBytes)
}

func (b *Block) PrintBlockInfo(detailed bool) {
	fmt.Printf("Version: %d\n", b.version)
	fmt.Printf("Time: %d\n", b.time)
	fmt.Printf("Previous Header Hash:\n\t%x\n", b.prev_hdr_hash)
	fmt.Printf("Hash:\n\t%x\n", b.hash)

	if detailed {
		fmt.Printf("Data:\n\t%s\n", b.data)
	}
}

func NewBlock(data []byte, prev_hash []byte) *Block {
	return &Block{
		common.VERSION,
		time.Now().Unix(),
		prev_hash,
		common.CalcHash(data),
		data,
	}
}

func NewGenesisBlock() *Block {
	return NewBlock([]byte("Initial Block"), []byte{})
}


