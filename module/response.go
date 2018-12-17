package module

type ResDataStruct struct {
	Data []interface{} `json:"data"`
}

// 首页统计
type ResTotalCount struct {
	TransactionCount string `json:"lastTradeCount"`
	NodeCount        string `json:"nodeCount"`
	BlockNumber      uint64 `json:"blockNum"`
	Accounts         int64  `json:"rtCustCount"`
}

type ResPageList struct {
	Data          interface{} `json:"data"`
	TotalRecordes int         `json:"iTotalDisplayRecords"`
}
