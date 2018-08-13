package blockchain

import (
	"github.com/internet-cash/prototype/transaction"
	"math"
	"github.com/internet-cash/prototype/common"
)

type Merkle struct {
}

func (self Merkle) BuildMerkleTreeStore(transactions []*transaction.Tx) ([]*common.Hash) {
	// Calculate how many entries are required to hold the binary merkle
	// tree as a linear array and create an array of that size.
	nextPoT := self.nextPowerOfTwo(len(transactions))
	arraySize := nextPoT*2 - 1
	merkles := make([]*common.Hash, arraySize)

	// Create the base transaction hashes and populate the array with them.
	for i, tx := range transactions {
		// If we're computing a witness merkle root, instead of the
		// regular txid, we use the modified wtxid which includes a
		// transaction's witness data within the digest. Additionally,
		// the coinbase's wtxid is all zeroes.
		witness := false
		switch {
		case witness && i == 0:
			var zeroHash common.Hash
			merkles[i] = &zeroHash
		case witness:
			//wSha := tx.MsgTx().WitnessHash()
			//merkles[i] = &wSha
			continue
		default:
			merkles[i] = tx.Hash()
		}

	}

	// Start the array offset after the last transaction and adjusted to the
	// next power of two.
	offset := nextPoT
	for i := 0; i < arraySize-1; i += 2 {
		switch {
		// When there is no left child node, the parent is nil too.
		case merkles[i] == nil:
			merkles[offset] = nil

			// When there is no right child, the parent is generated by
			// hashing the concatenation of the left child with itself.
		case merkles[i+1] == nil:
			newHash := self.HashMerkleBranches(merkles[i], merkles[i])
			merkles[offset] = newHash

			// The normal case sets the parent node to the double sha256
			// of the concatentation of the left and right children.
		default:
			newHash := self.HashMerkleBranches(merkles[i], merkles[i+1])
			merkles[offset] = newHash
		}
		offset++
	}

	return merkles
}

func (self Merkle) nextPowerOfTwo(n int) int {
	// Return the number if it's already a power of 2.
	if n&(n-1) == 0 {
		return n
	}

	// Figure out and return the next power of two.
	exponent := uint(math.Log2(float64(n))) + 1
	return 1 << exponent // 2^exponent
}

// HashMerkleBranches takes two hashes, treated as the left and right tree
// nodes, and returns the hash of their concatenation.  This is a helper
// function used to aid in the generation of a merkle tree.
func (self Merkle) HashMerkleBranches(left *common.Hash, right *common.Hash) *common.Hash {
	// Concatenate the left and right nodes.
	var hash [common.HashSize * 2]byte
	copy(hash[:common.HashSize], left[:])
	copy(hash[common.HashSize:], right[:])

	newHash := common.DoubleHashH(hash[:])
	return &newHash
}
