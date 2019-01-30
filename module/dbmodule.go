package module

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DBAccounts struct {
	AccountName           string    `bson:"account_name"`
	PubKey                string    `bson:"pubkey"`
	Balance               string    `bson:"balance"`
	TimeStamp             int64     `bson:"timestamp"`
	CreateTime            time.Time `bson:"create_time"`
	StakedBalance         string    `bson:"staked_balance"`
	StakedSpaceBalance    string    `bson:"staked_space_balance"`
	StakedTimeBalance     string    `bson:"staked_time_balance"`
	UnStakingBalance      string    `bson:"unStaking_balance"`
	UnStakdingTimestamp   uint64    `bson:"unStaking_timestamp"`
	AvailableSpaceBalance uint64    `bson:"available_space_balance"`
	UsedSpaceBalance      uint64    `bson:"used_space_balance"`
	AvailableTimeBalance  uint64    `bson:"available_time_balance"`
	UsedTimeBalance       uint64    `bson:"used_time_balance"`
	UnClaimedReward       string    `bson:"un_claimed_reward"` // 出块奖励
	DeployContractList    string    `bson:"deploy_contract_list"`
	Vote                  DBVote    `bson:"vote"`
}

type DBVote struct {
	Delegate string `bson:"delegate"`
	Votes    string `bson:"votes"`
}

type DBBlocks struct {
	BlockHash        string          `bson:"block_hash"`
	PrevBlockHash    string          `bson:"prev_block_hash"`
	BlockNumber      uint64          `bson:"block_number"`
	TimeStamp        int64           `bson:"timestamp"`
	Delegate         string          `bson:"delegate"`
	Transactions     []bson.ObjectId `bson:"transactions"`
	CreateTime       int64           `bson:"create_time"`
	TransactionCount int             `bson:"transaction_count"`
}

type DBTransaction struct {
	ID            bson.ObjectId          `bson:"_id"`
	BlockNumber   uint64                 `bson:"block_number"`
	TransactionID string                 `bson:"transaction_id"`
	SequenceNum   int32                  `bson:"sequence_num"`
	BlockHash     string                 `bson:"block_hash"`
	CursorNum     uint32                 `bson:"cursor_num"`
	CursorLabel   uint32                 `bson:"cursor_label"`
	LifeTime      uint64                 `bson:"lifetime"`
	Sender        string                 `bson:"sender"`
	Contract      string                 `bson:"contract"`
	Method        string                 `bson:"method"`
	Param         map[string]interface{} `bson:"param"`
	SigAlg        uint32                 `bson:"sig_alg"`
	Signature     string                 `bson:"signature"`
	TimeStamp     int64                  `bson:"timestamp"`
}

type DBNodeService struct {
	NodeType    string `bson:"node_type" json:"node_type"`
	NodeCountry string `bson:"node_country" json:"node_country"`
	IP          string `bson:"service_api" json:"ip"`
	URL         string `bson:"service_url" json:"serviceUrl"`
}

// bson.M{"country": "Europe-Germany", "nodeCount": 1},
type DBNodeSuper struct {
	NodeName            string `bson:"node_name" json:"nodeName"`
	VoteCount           string `bson:"transit_votes" json:"voteCount"`
	VoteCustCount       int32  `bson:"voted_account_num" json:"voteCustCount"`
	NodeDesc            string `bson:"node_desc" json:"nodeDesc"`
	HomePage            string `bson:"homepage" json:"homepage"`
	Delegate            string `bson:"delegate" json:"delegate"`
	StakedBalance       string `bson:"staked_balance" json:"stakedBalance"`
	ProducedBlocksCount string `bson:"produced_blocks" json:"producedBlocksCount"`
	StateOfDelegate     bool   `bson:"state_of_delegate" json:"stateOfDelegate"`
	AllTransitVotes     string `bson:"all_transit_votes" json:"allTransitVotes"`
	VoteRank            int    `bson:"vote_rank" json:"voteRank"`
}
