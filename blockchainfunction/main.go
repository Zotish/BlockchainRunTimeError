package blockchainfunction

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	blockstructure "github.com/Zotish/blockStructure"
)

func CreateGenBlock() *blockstructure.Block {
	var GenBlock *blockstructure.Block
	GenBlock.Block_height = 0
	GenBlock.TimeStamp = time.Now().String()
	GenBlock.MinerReward = 0
	GenBlock.Transaction.FromSend = ""
	GenBlock.Transaction.ToSend = ""
	GenBlock.Transaction.Amount = 0
	GenBlock.Version = "v1.1.0"
	GenBlock.PreviousHash = ""
	GenBlock.Status = true
	GenBlock.CurrentHash = CalculateHash(GenBlock)
	return GenBlock
}
func CalculateHash(Blk *blockstructure.Block) string {
	Hash := strconv.FormatUint(uint64(Blk.Block_height), 32) + Blk.Transaction.FromSend + Blk.Transaction.ToSend + strconv.FormatFloat(Blk.Transaction.Amount, 'f', -1, 64) + strconv.FormatFloat(Blk.Transaction.Fees, 'f', -1, 64) + Blk.PreviousHash + Blk.TimeStamp
	ConvToSHA := sha256.New()
	ConvToSHA.Write([]byte(Hash))
	Hashed := ConvToSHA.Sum(nil)
	return hex.EncodeToString(Hashed)
}
func CreateNewBlock(PreviousHash *blockstructure.Block, Transaction *blockstructure.TransactionDetails) blockstructure.Block {
	var newBlock blockstructure.Block
	time := time.Now()
	newBlock.Block_height = PreviousHash.Block_height + 1
	newBlock.TimeStamp = time.String()
	newBlock.Transaction = Transaction
	newBlock.PreviousHash = PreviousHash.CurrentHash
	newBlock.CurrentHash = CalculateHash(&newBlock)
	return newBlock
}
func IsValidBlock(StoreNewBlock *blockstructure.Block, StorePreviousBlock *blockstructure.Block) bool {
	if StoreNewBlock.Block_height != StorePreviousBlock.Block_height {
		return false
	}
	if StoreNewBlock.PreviousHash != StorePreviousBlock.CurrentHash {
		return false
	} else {
		return true
	}
}
func IsValidChain(Blockcahain []blockstructure.Block) bool {
	for index := range Blockcahain {

		CurrentBlock := Blockcahain[index]
		preViousBlock := Blockcahain[index-1]

		if CurrentBlock.CurrentHash != preViousBlock.PreviousHash {
			return false
		}
		if CurrentBlock.PreviousHash != preViousBlock.CurrentHash {
			return false
		}
	}
	return true
}
