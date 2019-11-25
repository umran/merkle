package merkle

import (
	"math"

	"github.com/umran/crypto"
)

// Tree ...
type Tree struct {
	nodes     []crypto.Hash
	root      crypto.Hash
	lenLeaves int
}

// Hashable ...
type Hashable interface {
	Hash() crypto.Hash
}

// New ...
func New(leaves []Hashable) *Tree {
	lenLeaves := len(leaves)
	lenTreeUpperbound := determineLenTreeUpperbound(lenLeaves)

	tree := &Tree{
		nodes:     make([]crypto.Hash, lenLeaves, lenTreeUpperbound),
		lenLeaves: lenLeaves,
	}

	// populate tree with leaves
	for i, leaf := range leaves {
		tree.nodes[i] = leaf.Hash()
	}

	// generate all upper branches
	lenCurrentLevel := lenLeaves
	offset := 0

	for lenCurrentLevel > 1 {
		lenNextLevel := determineLenNextLevel(lenCurrentLevel)
		tree.nodes = tree.nodes[:offset+lenCurrentLevel+lenNextLevel]

		// populate the next level
		for i := 0; i < lenNextLevel; i++ {
			tree.nodes[offset+lenCurrentLevel+i] = generateNextNode(tree.nodes, lenCurrentLevel, offset, i)
		}

		offset += lenCurrentLevel
		lenCurrentLevel = lenNextLevel
	}

	tree.root = tree.nodes[offset]

	return tree
}

func generateNextNode(nodes []crypto.Hash, lenCurrentLevel, offset, currentIndex int) crypto.Hash {
	var (
		leftNode  crypto.Hash
		rightNode crypto.Hash
	)

	leftNode = nodes[offset+2*currentIndex]

	switch {
	case 2*currentIndex+1 > lenCurrentLevel-1:
		rightNode = nodes[offset+2*currentIndex]
	default:
		rightNode = nodes[offset+2*currentIndex+1]
	}

	return leftNode.Merge(rightNode)
}

func determineLenTreeUpperbound(lenLeaves int) int {
	return int(math.Ceil(float64(2*lenLeaves) + math.Log2(float64(lenLeaves-1))))
}

func determineLenNextLevel(lenCurrentLevel int) int {
	switch {
	case lenCurrentLevel == 2:
		return 1
	case isDivisibleByTwo(lenCurrentLevel) == true:
		return lenCurrentLevel / 2
	default:
		return (lenCurrentLevel + 1) / 2
	}
}

func isDivisibleByTwo(value int) bool {
	switch {
	case math.Mod(float64(value), 2) == 0:
		return true
	default:
		return false
	}
}

// Nodes ...
func (tree *Tree) Nodes() []crypto.Hash {
	return tree.nodes
}

// Root ...
func (tree *Tree) Root() crypto.Hash {
	return tree.root
}
