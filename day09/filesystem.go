package day09

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
)

type Filesystem struct {
	blocks []*Block // keep the blocks ordered!
}

func NewFilesystem() *Filesystem {
	return &Filesystem{
		blocks: make([]*Block, 0),
	}
}

func (fs *Filesystem) String() string {
	return fmt.Sprintf("%v", fs.blocks)
}

func (fs *Filesystem) PrettyPrint() string {
	ret := ""

	for i, block := range fs.blocks {
		ret += strings.Repeat(strconv.Itoa(block.fileId), block.length)
		if i > 0 && i < len(fs.blocks)-1 {
			ret += strings.Repeat(".", fs.blocks[i+1].from-block.To()-1)
		}
	}

	return ret
}

func (fs *Filesystem) AddBlockUnsafe(block *Block) {
	if !fs.isFree(block.from, block.length) {
		panic("space not free for block")
	}
	fs.blocks = append(fs.blocks, block)
}

func (fs *Filesystem) insertBlockAtPosition(pos int, newBlock *Block) {
	if pos < 0 {
		fmt.Println("Invalid position")
		return
	}

	fs.blocks = append(fs.blocks[:pos], append([]*Block{newBlock}, fs.blocks[pos:]...)...)
}

func (fs *Filesystem) removeAt(pos int) {
	if pos < 0 || pos >= len(fs.blocks) {
		fmt.Println("Invalid position")
		return
	}
	fs.blocks = append(fs.blocks[:pos], fs.blocks[pos+1:]...)
}

func (fs *Filesystem) isFree(from int, length int) bool {
	freeBlock := NewBlock(-1, from, length)
	for _, block := range fs.blocks {
		if block.Overlaps(freeBlock) {
			return false
		}
	}
	return true
}

func (fs *Filesystem) Compress() {
	// while there are gaps
	firstGap, exists, indexBlockBeforeGap := fs.FirstGap()
	for exists {
		// take first gap and fill block-by-block from the end
		lastBlock := fs.blocks[len(fs.blocks)-1]

		// how many blocks to copy?
		lengthToCopy := common.IntMin(lastBlock.length, firstGap.length)

		if lengthToCopy < lastBlock.length {
			// reduce size of last block
			lastBlock.length -= lengthToCopy
		} else {
			// remove last block entirely
			fs.blocks = fs.blocks[:len(fs.blocks)-1]
		}

		fs.insertBlockAtPosition(indexBlockBeforeGap+1, NewBlock(lastBlock.fileId, firstGap.from, lengthToCopy))

		firstGap, exists, indexBlockBeforeGap = fs.FirstGap()
	}
}

func (fs *Filesystem) GetBlockByFileId(fileId int) (*Block, int) {
	for i, block := range fs.blocks {
		if block.fileId == fileId {
			return block, i
		}
	}
	panic("block not found")
}

func (fs *Filesystem) CompressFiles() {
	// go backwards through files (same as blocks at the start)
	// use fileId for search, to avoid moving the same block twice
	for fileId := len(fs.blocks) - 1; fileId >= 0; fileId-- {
		block, index := fs.GetBlockByFileId(fileId)

		// get first (left-most) gap of at least length block.length, left of the current block
		firstGap, exists, indexBlockBeforeGap := fs.FirstGapLargerThanUntilPos(block.length, block.from)

		if exists {
			// remove old block
			fs.removeAt(index)

			// insert new block in gap
			fs.insertBlockAtPosition(indexBlockBeforeGap+1, NewBlock(block.fileId, firstGap.from, block.length))
		}
	}
}

func (fs *Filesystem) FirstGap() (gap *Block, exists bool, indexBlockBeforeGap int) {
	if len(fs.blocks) <= 1 {
		return nil, false, -1
	}

	// we know there is no gap at the start

	for i := 0; i < len(fs.blocks)-1; i++ {
		distToNextBlock := fs.blocks[i+1].from - fs.blocks[i].To() - 1
		if distToNextBlock > 0 {
			return NewBlock(-1, fs.blocks[i].To()+1, distToNextBlock), true, i
		}
	}

	return nil, false, -1
}

func (fs *Filesystem) FirstGapLargerThanUntilPos(minLength int, pos int) (gap *Block, exists bool, indexBlockBeforeGap int) {
	if len(fs.blocks) <= 1 {
		return nil, false, -1
	}

	// we know there is no gap at the start

	for i := 0; i < len(fs.blocks)-1; i++ {
		if fs.blocks[i].To() > pos {
			break
		}
		distToNextBlock := fs.blocks[i+1].from - fs.blocks[i].To() - 1
		if distToNextBlock >= minLength {
			return NewBlock(-1, fs.blocks[i].To()+1, distToNextBlock), true, i
		}
	}

	return nil, false, -1
}

func (fs *Filesystem) checksum() int {
	sum := 0

	for _, block := range fs.blocks {
		// dumb loop inside the block - can be replaced with a formula if too slow
		for i := 0; i < block.length; i++ {
			sum += block.fileId * (block.from + i)
		}
	}

	return sum
}
