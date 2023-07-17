package merkleTree

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
	Hash  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Hash = hash[:]
	} else {
		prevHashes := append(left.Hash, right.Hash...)
		hash := sha256.Sum256(prevHashes)
		node.Hash = hash[:]
	}
	node.Left = left
	node.Right = right
	node.Data = data
	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, dat := range data {
		node := NewMerkleNode(nil, nil, dat)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode
		if len(nodes) == 1 {
			break
		}
		for j := 0; j < len(nodes); j += 2 {
			if j+1 < len(nodes) {
				node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
				newLevel = append(newLevel, *node)
			} else {
				node := NewMerkleNode(&nodes[j], &nodes[j], nil)
				newLevel = append(newLevel, *node)
			}
		}
		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}

	return &mTree
}

func (mTree *MerkleTree) PrintTree() {
	fmt.Println("Merkle Tree:")
	printNode(mTree.RootNode, 0)
}

func printNode(node *MerkleNode, level int) {
	format := ""
	for i := 0; i < level; i++ {
		format += "    "
	}
	format += "|--- "
	hash := hex.EncodeToString(node.Hash)
	fmt.Printf(format+"Hash: %s\n", hash)
	if node.Left != nil {
		printNode(node.Left, level+1)
	}
	if node.Right != nil {
		printNode(node.Right, level+1)
	}
}

func main() {
	data := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
		[]byte("5"),
		[]byte("6"),
	}
	tree := NewMerkleTree(data)
	tree.PrintTree()
}
