package blockstructure

type Block struct {
	Version      string
	Block_height uint64
	MinerReward  float64
	TimeStamp    string
	CurrentHash  string
	PreviousHash string
	Transaction  *TransactionDetails
	Status       bool
}
type TransactionDetails struct {
	FromSend string
	ToSend   string
	Amount   float64
	Fees     float64
}
