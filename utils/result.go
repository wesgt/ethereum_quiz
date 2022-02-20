package utils

type Block struct {
	Number     uint64 `json:"block_num"`
	Hash       string `json:"block_hash"`
	ParentHash string `json:"parent_hash"`
	Time       uint64 `json:"block_time"`
}

type BlockWithTx struct {
	// Number       uint64   `json:"block_num"`
	// Hash         string   `json:"block_hash"`
	// ParentHash   string   `json:"parent_hash"`
	// Time         uint64   `json:"block_time"`
	Block
	Transactions []string `json:"transactions"`
}

type Transaction struct {
	// no data filed
	Hash  string `json:"tx_hash"`
	From  string `json:"from"`
	To    string `json:"to"`
	Value uint64 `json:"value"`
	Nonce uint64 `json:"nonce"`
	Logs  []Log  `json:"logs"`
}

type Log struct {
	Index uint   `json:"index"`
	Data  string `json:"data"`
}
