package module

import (
	"gopkg.in/mgo.v2/bson"
)

type DBAccounts struct {
	AccountName        string `bson:"account_name"`
	PubKey             string `bson:"pubkey"`
	Balance            string `bson:"balance"`
	StakedBalance      string `bson:"staked_balance"`
	StakedSpaceBalance string `bson:"staked_space_balance"`
	StakedTimeBalance  string `bson:"staked_time_balance"`
	UnStakingBalance   string `bson:"unStaking_balance"`
	UnStakdingTime     uint64 `bson:"unStaking_time"`
	Vote               DBVote `bson:"vote"`
}

type DBVote struct {
	Delegate              string `bson:"delegate"`
	Votes                 string `bson:"votes"`
	AvailableSpaceBalance uint64 `bson:"available_space_balance"`
	UsedSpaceBalance      uint64 `bson:"used_space_balance"`
	AvailableTimeBalance  uint64 `bson:"available_time_balance"`
	UsedTimeBalance       uint64 `bson:"used_time_balance"`
}

type DBBlocks struct {
	BlockHash     string          `bson:"block_hash"`
	PrevBlockHash string          `bson:"prev_block_hash"`
	BlockNumber   uint64          `bson:"block_number"`
	TimeStamp     int64           `bson:"timestamp"`
	Delegate      string          `bson:"delegate"`
	Transactions  []bson.ObjectId `bson:"transactions"`
	CreateTime    int64           `bson:"create_time"`
}

type DBTransaction struct {
	ID            bson.ObjectId `bson:"_id"`
	BlockNumber   uint64        `bson:"block_number"`
	TrxHash       string        `bson:"trx_hash"`
	TransactionID string        `bson:"transaction_id"`
	SequenceNum   int32         `bson:"sequence_num"`
	BlockHash     string        `bson:"block_hash"`
	CursorNum     uint32        `bson:"cursor_num"`
	CursorLabel   uint32        `bson:"cursor_label"`
	LifeTime      uint64        `bson:"lifetime"`
	Sender        string        `bson:"sender"`
	Contract      string        `bson:"contract"`
	Method        string        `bson:"method"`
	Param         interface{}   `bson:"param"`
	SigAlg        uint32        `bson:"sig_alg"`
	Signature     string        `bson:"signature"`
	CreateTime    int64         `bson:"create_time"`
}
