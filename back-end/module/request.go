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
	BlockNum string `json:"blockNum"`
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
	AccountName string `json:"accountName"`
}

// GetTransactionList
type ReqTransactionList struct {
	AccountName string `json:"accountName"`
	BlockNum    string `json:"blockNum"`
	Start       int    `json:"start"`
	Length      int    `json:"length"`
}
