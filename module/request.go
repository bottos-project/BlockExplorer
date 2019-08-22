package module

// GetBlockList
type ReqBlockList struct {
	DelegateName string `json:"delegateName"`
	Start        int    `json:"start"`
	Length       int    `json:"length"`
}

type ReqAccountList struct {
	Start  int `json:"start"`
	Length int `json:"length"`
	Order  string `json:"order"`
	By     string `json:"by"`
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
	TransactionID string `json:"transactionId"`
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

// Home search by account、blockNum、transactionID、blockHash
type ReqHomeSearch struct {
	Condition string `json:"condition"`
}

// super node detail
type ReqNodeSuperDetail struct {
	NodeName string `json:"nodeName"`
	Start    int    `json:"start"`
	Length   int    `json:"length"`
}

type ReqQueryStaticInfo struct {
	StartTime string `json:"startDate"`
	EndTime   string `json:"EndDate"`
}

// person transaction search list
type ReqPersonTransactionByMethod struct {
	AccountName string `json:"accountName"`
	BlockNum    string `json:"blockNum"`
	Start       int    `json:"start"`
	Length      int    `json:"length"`
	Method      string `json:"method"`
}
