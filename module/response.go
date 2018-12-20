package module

// response success
type ResSuccess struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type ResErr struct {
	Data    error  `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// response data struct
type ResDataStruct struct {
	Data interface{} `json:"data"`
}

// home statistic
type ResTotalCount struct {
	TransactionCount int    `json:"lastTradeCount"`
	NodeCount        string `json:"nodeCount"`
	BlockNumber      uint64 `json:"blockNum"`
	Accounts         int    `json:"rtCustCount"`
}

type ResPageList struct {
	Data          interface{} `json:"data"`
	TotalRecordes int         `json:"iTotalDisplayRecords"`
}

// block list
type ResBlockList struct {
	BlockHash        string `bson:"block_hash" json:"blockHash"`
	BlockNumber      uint64 `bson:"block_number" json:"blockNum"`
	TimeStamp        int64  `bson:"timestamp" json:"tradeDate"`
	Delegate         string `bson:"delegate" json:"delegate"`
	TransactionCount int32  `bson:"transaction_count" json:"transCount"`
}

// transactions
type ResTransactions struct {
	BlockNumber   uint64 `bson:"block_number" json:"blockNum"`
	TransactionID string `bson:"transaction_id" json:"transactionId"`
	Sender        string `bson:"sender" json:"sender"`
	Contract      string `bson:"contract" json:"contract"`
	Method        string `bson:"method" json:"method"`
	TimeStamp     int64  `bson:"timestamp" json:"tradeDate"`
}

// transfers
type ResTransfers struct {
	BlockNumber   uint64  `bson:"block_number" json:"blockNum"`
	From          string  `json:"sender"`
	To            string  `json:"receiver"`
	Value         float64 `json:"tradeAmount"`
	CoinType      string  `json:"currency"`
	TransactionID string  `bson:"transaction_id" json:"transactionId"`
	TimeStamp     int64   `bson:"timestamp" json:"tradeDate"`
}

// account list
type ResAccountList struct {
	AccountName string `bson:"account_name" json:"accountName"`
	Balance     string `bson:"balance" json:"availableAmount"`
}

// block detail
type ResBlockDetail struct {
	BlockHash        string `bson:"block_hash" json:"blockHash"`
	PrevBlockHash    string `bson:"prev_block_hash" json:"prevBlockHash"`
	BlockNumber      uint64 `bson:"block_number" json:"blockNum"`
	PrevBlockNum     uint64 `json:"prevBlockNum"`
	BlockStatus      string `json:"blockStatus"`
	TimeStamp        int64  `bson:"timestamp" json:"tradeDate"`
	Delegate         string `bson:"delegate" json:"delegate"`
	TransactionCount int    `bson:"transaction_count" json:"transCount"`
	ResultType       string `json:"resultType"`
}

// account detail
type ResAccountDetail struct {
	AccountName  string `bson:"account_name" json:"accountName"`
	ReceiveCount int    `json:"receiveCount"`
	SendCount    int    `json:"sendCount"`
	TradeCount   int    `json:"tradeCount"`
	Balance      string `bson:"balance" json:"balance"`
	ResultType   string `json:"resultType"`
}

// transfer detail
type ResTransferDetail struct {
	BlockNumber   uint64                 `bson:"block_number" json:"blockNum"`
	TransactionID string                 `bson:"transaction_id" json:"transactionId"`
	Sender        string                 `bson:"sender" json:"sender"`
	Contract      string                 `bson:"contract" json:"contract"`
	Method        string                 `bson:"method" json:"method"`
	Param         map[string]interface{} `bson:"param" json:"param"`
	TimeStamp     int64                  `bson:"timestamp" json:"tradeDate"`
	TradeStatus   string                 `json:"tradeStatus"`
	ResultType    string                 `json:"resultType"`
}
