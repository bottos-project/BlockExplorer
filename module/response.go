package module

import "time"

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
	NodeCount        int    `json:"nodeCount"`
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
	BlockNumber    uint64 `bson:"block_number" json:"blockNum"`
	TransactionID  string `bson:"transaction_id" json:"transactionId"`
	Sender         string `bson:"sender" json:"sender"`
	Contract       string `bson:"contract" json:"contract"`
	Method         string `bson:"method" json:"method"`
	TimeStamp      int64  `bson:"timestamp" json:"tradeDate"`
	SpaceTokenCost int32  `bson:"space_token_cost" json:"spaceTokenCost"`
	TimeTokenCost  int32  `bson:"time_token_cost" json:"timeTokenCost"`
}

// transfers
type ResTransfers struct {
	BlockNumber    uint64  `bson:"block_number" json:"blockNum"`
	From           string  `json:"sender"`
	To             string  `json:"receiver"`
	Value          float64 `json:"tradeAmount"`
	CoinType       string  `json:"currency"`
	TransactionID  string  `bson:"transaction_id" json:"transactionId"`
	TimeStamp      int64   `bson:"timestamp" json:"tradeDate"`
	SpaceTokenCost int32   `bson:"space_token_cost" json:"spaceTokenCost"`
	TimeTokenCost  int32   `bson:"time_token_cost" json:"timeTokenCost"`
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
	AccountName           string `bson:"account_name" json:"accountName"`
	ReceiveCount          int    `json:"receiveCount"`
	SendCount             int    `json:"sendCount"`
	TradeCount            int    `json:"tradeCount"`
	Balance               string `bson:"balance" json:"balance"`
	StakedBalance         string `bson:"staked_balance" json:"stakedBalance"`
	UnStakedBalance       string `bson:"unStaking_balance" json:"unStakingBalance"`
	ResultType            string `json:"resultType"`
	UnClaimedReward       string `bson:"un_claimed_reward" json:"unClaimedReward"`
	AvailableSpaceBalance string `bson:"available_space_balance" json:"availableSpaceBalance"`
	AvailableTimeBalance  string `bson:"available_time_balance" json:"availableTimeBalance"`

	Resource DBResource `bson:"resource" json:"resource"`
}
type DBResource struct {
	FreeAvailableSpace  uint64 `bson:"free_available_space" json:"freeAvailableSpace"`
	FreeUsedSpace       uint64 `bson:"free_used_space" json:"freeUsedSpace"`
	StakeAvailableSpace uint64 `bson:"stake_available_space" json:"stakeAvailableSpace"`
	StakeUsedSpace      uint64 `bson:"stake_used_space" json:"stakeUsedSpace"`
	FreeAvailableTime   uint64 `bson:"free_available_time" json:"freeAvailableTime"`
	FreeUsedTime        uint64 `bson:"free_used_time" json:"freeUsedTime"`
	StakeAvailableTime  uint64 `bson:"stake_available_time" json:"stakeAvailableTime"`
	StakeUsedTime       uint64 `bson:"stake_used_time" json:"stakeUsedTime"`
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

// transaction statistic
type ResStatisticTransaction struct {
	DailyTransCount int       `bson:"trx_count" json:"dailyTransCount"`
	TradeDate       time.Time `bson:"day" json:"tradeDate"`
	TotalTransCount int       `bson:"total_trx_count" json:"totalTrxCount"`
}

// account statistic
type ResStatisticAccount struct {
	DailyRegCount int       `bson:"account_count" json:"dailyRegCount"`
	RegDate       time.Time `bson:"day" json:"regDate"`
	TotalRegCount int       `bson:"total_account_count" json:"totalRegCount"`
}

type ResPersonalTransactionByMethod struct {
	BlockNumber   uint64                 `bson:"block_number" json:"blockNum"`
	TransactionID string                 `bson:"transaction_id" json:"transactionId"`
	SequenceNum   int32                  `bson:"sequence_num" json:"sequenceNum"`
	BlockHash     string                 `bson:"block_hash" json:"blockHash"`
	CursorNum     uint32                 `bson:"cursor_num" json:"cursonNum"`
	CursorLabel   uint32                 `bson:"cursor_label" json:"cursorLabel"`
	LifeTime      uint64                 `bson:"lifetime" json:"lifeTime"`
	Sender        string                 `bson:"sender" json:"sender"`
	Contract      string                 `bson:"contract" json:"contract"`
	Method        string                 `bson:"method" json:"method"`
	Param         map[string]interface{} `bson:"param" json:"param"`
	SigAlg        uint32                 `bson:"sig_alg" json:"sigAlg"`
	Signature     string                 `bson:"signature" json:"signature"`
	TimeStamp     int64                  `bson:"timestamp" json:"timestamp"`
}

type ResNodeSuperSummary struct {
	DelegateCount   int32 `json:"delegateCount"`
	VoteCustCount   int32 `json:"voteCustCount"`
	AllTransitVotes int64 `json:"allTransitVotes"`
}

type ResServiceNodeSummary struct {
	DelegateCount int `json:"delegateCount"`
	ServiceCount  int `json:"serviceNodeCount"`
}
