package day09

import "fmt"

type Block struct {
	fileId int
	from   int
	length int
}

func NewBlock(fileId int, from int, length int) *Block {
	if length == 0 {
		panic("block length is zero")
	}
	return &Block{fileId: fileId, from: from, length: length}
}

func (b *Block) String() string {
	return fmt.Sprintf("Block[%d, from:%d, len:%d]", b.fileId, b.from, b.length)
}

func (b *Block) To() int {
	return b.from + b.length - 1
}

func (b *Block) Overlaps(other *Block) bool {
	return !(b.To() < other.from || b.from > other.To())
}
