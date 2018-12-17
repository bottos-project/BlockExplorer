package module

// GetBlockList
type ReqBlockList struct {
	Start  int `json:"start"`
	Length int `json:"length"`
}

type ReqAccountList struct {
	Start  int `json:"start"`
	Length int `json:"length"`
}

// GetBlockDetail
type ReqBlockDetail struct {
	BlockNum uint64 `json:"block_num"`
}

// get block detail by blockhash
type ReqBlockDetailByHash struct {
	BlockHash string `json:"block_hash"`
}

// GetTransaction
type ReqGetTransaction struct {
	TrxHash string `json:"trx_hash"`
}

// GetAccountInfo
type ReqAccountInfo struct {
	AccountName string `json:"account_name"`
}

// GetTransactionList
type ReqTransactionList struct {
	BlockNum int `json:"block_number"`
	Start    int `json:"start"`
	Length   int `json:"length"`
}
