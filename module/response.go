package module

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

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

type TransferHistory struct {
	Account string `bson:"_id" json:"account_name"`
	Block string `bson:"block_number" json:"block_number"`
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
	Method        string   `bson:"method" json:"method"`
}

// account list
type ResAccountList struct {
	AccountName string `bson:"account_name" json:"accountName"`
	StakeBalance int64`bson:"staked_balance" json:"staked_balance"`
	StakedTimeBalance int64`bson:"staked_time_balance" json:"staked_time_balance"`
	StakedSpaceBalance int64`bson:"staked_space_balance" json:"staked_space_balance"`
	Balance     int64 `bson:"balance" json:"balance"`
	UnClaimedReward int64 `bson:"un_claimed_reward" json:"un_claimed_reward"`
	TotalBalance int64  `bson:"total_balance" json:"total_balance"`

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
	Balance               interface{} `bson:"balance" json:"balance"`
	StakedBalance         interface{} `bson:"staked_balance" json:"stakedBalance"`
	StakedTimeBalance	  interface{}  `bson:"staked_time_balance" json:"staked_time_balance"`
	StakedSpaceBalance	  interface{}  `bson:"staked_space_balance" json:"staked_space_balance"`
	UnStakedBalance       interface{} `bson:"unStaking_balance" json:"unStakingBalance"`
	ResultType            string `json:"resultType"`
	UnClaimedReward       string `bson:"un_claimed_reward" json:"unClaimedReward"`
	UnClaimedVoteReward   interface{} `bson:"un_claimed_vote_reward" json:"unClaimedVoteReward"`
	UnClaimedBlockReward  interface{} `bson:"un_claimed_block_reward" json:"unClaimedBlockReward"`
	AvailableSpaceBalance interface{} `bson:"available_space_balance" json:"availableSpaceBalance"`
	AvailableTimeBalance  interface{} `bson:"available_time_balance" json:"availableTimeBalance"`
	Vote                  struct {
		Delegate string `bson:"delegate" json:"delegate"`
		Votes    interface{} `bson:"votes" json:"votes"`
	} `bson:"vote" json:"vote"`
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
	BlockNumber    uint64                 `bson:"block_number" json:"blockNum"`
	TransactionID  string                 `bson:"transaction_id" json:"transactionId"`
	Sender         string                 `bson:"sender" json:"sender"`
	Contract       string                 `bson:"contract" json:"contract"`
	Method         string                 `bson:"method" json:"method"`
	Param          map[string]interface{} `bson:"param" json:"param"`
	TimeStamp      int64                  `bson:"timestamp" json:"tradeDate"`
	TradeStatus    string                 `json:"tradeStatus"`
	ResultType     string                 `json:"resultType"`
	SpaceTokenCost int32                  `bson:"space_token_cost" json:"spaceTokenCost"`
	TimeTokenCost  int32                  `bson:"time_token_cost" json:"timeTokenCost"`
}

// transaction statistic
type ResStatisticTransaction struct {
	DailyTransCount int    `bson:"trx_count" json:"trx_count"`
	TradeDate       string `bson:"day" json:"day"`
	TotalTransCount int    `bson:"account_count" json:"account_count"`
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

type HomePageResponse struct {
	LastestBlockInfo int    `bson:"lastest_block_info" json:"lastest_block_info"`
	NodeCount        int    `bson:"node_count" json:"node_count"`
	TransactionCount int	`bson:"transaction_count" json:"transaction_count"`
	BlockNumber      int64 `bson:"block_number" json:"block_number"`
	AccountCount     int    `bson:"account_count" json:"account_count"`
	Stake            StakeSum
	Tps              TpsCount
	StatisticData    []StatisticData
	Time             time.Time
}

type StakeSum struct {
	StakeTimeAmount int64 `json:"stake_time_amount"`
	StakeSpaceAmount int64 `json:"stake_space_amount"`
	StakeVoteAmount int64 `json:"stake_vote_amount"`
}

type TpsCount struct {
	TpsPeak   int32 `bson:"tps_peak" json:"tps_peak"`
	Tps       int32 `json:"tps"`
}

type BlockTps struct {
	TransactionCount int32 `bson:"transaction_count" json:"transaction_count"`
}

type StatisticData struct {
	Day string `bson:"day" json:"day"`
	AccountCount int32 `bson:"account_count" json:"account_count"`
	TrxCount int32 `bson:"trx_count" json:"trx_count"`
}

type StakeData struct{
	Id  bson.ObjectId `json:"_id"`
	Total int64 	`json:"total"`
}