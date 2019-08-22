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
type LastBlockNum struct {
	BlockNumber int64 `bson:"block_number" json:"block_number"`
}

type DBVote struct {
	Delegate string `bson:"delegate"`
	Votes    string `bson:"votes"`
}

type SuperNodeInfo struct {
	Description string `bson:"description" json:"description"`
	Avatar string `bson:"avatar" json:"avatar"`
	AvatarThumb string `bson:"avatar_thumb" json:"avatar_thumb"`
}

type VoteDelegateParam struct {
	Delegate string `bson:"delegate" json:"delegate"`
	Voter    string  `bson:"voter" json:"voter"`
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

type DBTransactions struct {
	ID            bson.ObjectId          `bson:"_id" json:"_id"`
	BlockNumber   uint64                 `bson:"block_number" json:"block_number"`
	TransactionID string                 `bson:"transaction_id" json:"transaction_id"`
	//SequenceNum   int32                  `bson:"sequence_num" json:"sequence_num"`
	BlockHash     string                 `bson:"block_hash" json:"block_hash"`
	CursorNum     uint32                 `bson:"cursor_num" json:"cursor_num"`
	CursorLabel   uint32                 `bson:"cursor_label" json:"cursor_label"`
	LifeTime      uint64                 `bson:"lifetime" json:"life_time"`
	Sender        string                 `bson:"sender" json:"sender"`
	Contract      string                 `bson:"contract" json:"contract"`
	Method        string                 `bson:"method" json:"method"`
	Param         struct{
		Account string `bson:"account" json:"account"`
		Proposal string `bson:"proposal" json:"proposal"`
		Proposer string `bson:"proposer" json:"proposer"`
		Transfer  string `bson:"transfer" json:"transfer"`
	}
	Space_token_cost int32				`bson:"space_token_cost" json:"space_token_cost"`
	Time_token_cost int32				`bson:"time_token_cost" json:"time_token_cost"`
	SigAlg        uint32                 `bson:"sig_alg" json:"sig_alg"`
	Signature     string                 `bson:"signature" json:"signature"`
	TimeStamp     int64                  `bson:"timestamp" json:"time_stamp"`
}
type ProposalParam struct {
	Account string `bson:"account" json:"account"`
	Proposal string `bson:"proposal" json:"proposal"`
	Proposer string `bson:"proposer" json:"proposer"`
	Transfer string `bson:"transfer" json:"transfer"`
}

type DBNodeService struct {
	NodeType    string `bson:"node_type" json:"node_type"`
	NodeCountry string `bson:"node_country" json:"node_country"`
	IP          string `bson:"service_api" json:"ip"`
	URL         string `bson:"service_url" json:"serviceUrl"`
}

type DBNodeSuper struct {
	// NodeName            string `bson:"node_name" json:"nodeName"`
	// VoteCount           int64  `bson:"transit_votes" json:"voteCount"`
	// VoteCustCount       int32  `bson:"voted_account_num" json:"voteCustCount"`
	// NodeDesc            string `bson:"node_desc" json:"nodeDesc"`
	// HomePage            string `bson:"homepage" json:"homepage"`
	Delegate string `bson:"delegate" json:"delegate"`
	// StakedBalance       string `bson:"staked_balance" json:"stakedBalance"`
	// ProducedBlocksCount string `bson:"produced_blocks" json:"producedBlocksCount"`
	// StateOfDelegate     bool   `bson:"state_of_delegate" json:"stateOfDelegate"`
	AllTransitVotes int64  `bson:"all_transit_votes" json:"allTransitVotes"`
	VoteRank        int    `bson:"vote_rank" json:"voteRank"`
	Votes           string `bson:"votes" json:"votes"`
	Description     string `bson:"description" json:"description"`
	Active          bool   `bson:"active" json:"active"`
	TransitVotes    string `bson:"transit_votes" json:"transit_votes"`
	Location        string `bson:"location" json:"location"`
	PubKey          string `bson:"public_key" json:"public_key"`
}

type NodeSuper struct {
	Avatar string `bson:"avatar" json:"avatar"`
	AvatarThumb string `bson:"avatar_thumb" json:"avatar_thumb"`
	Delegate string `bson:"delegate" json:"delegate"`
	Description string `bson:"description" json:"description"`
}

type PubKey struct {
	Pkey string `bson:"pubkey" json:"pubkey"`
}

type PushProposal struct {
	Sender string `bson:"sender" json:"sender"`
	Res    string `bson:"res" json:"res"`
	Param struct{
		Account string `bson:"account" json:"account"`
		Proposal string `bson:"proposal" json:"proposal"`
		Proposer string `bson:"proposer" json:"proposer"`
		Transfer string `bson:"transfer" json:"transfer"`
		From     string `bson:"from" json:"from"`
		To 		 string `bson:"to" json:"to"`
		Value    float64 `bson:"value" json:"value"`
		Memo     string `bson:"memo" json:"memo"`
	}
}

type ExecProposal struct {
	Sender string `bson:"sender" json:"sender"`
	Method string `bson:"method" json:"method"`
	Param struct{
		Proposal string `bson:"proposal" json:"proposal"`
		Proposer string `bson:"proposer" json:"proposer"`
		From     string `bson:"from" json:"from"`
		To 		 string `bson:"to" json:"to"`
		Value    float64 `bson:"value" json:"value"`
		Memo     string `bson:"memo" json:"memo"`
	}
}





type DBMsignAccountModule struct {
	AuthorAccount string `bson:"author_account" json:"author_account"`
}

type AuthorAccountModule struct {
	Contract string `bson:"contract" json:"contract"`
	Method   string `bson:"method" json:"method"`
	Param    struct {
		Account   string `bson:"account" json:"account"`
		Authority []struct {
			AuthorAccount string `bson:"author_account" json:"author_account"`
			Weight        int32  `bson:"weight" json:"weight"`
			Transfer      TransferDetail
		} `bson:"authority" json:"authority"`
		Threshold int32 `bson:"threshold" json:"threshold"`
	} `bson:"param" json:"param"`
	Sender string `bson:"sender" json:"sender"`
}

type Param struct {
	Account   string      `bson:"account" json:"account"`
	Authority []Authority `bson:"authority" json:"authority"`
	Threshold int32       `bson:"threshold" json:"threshold"`
}

type Authority struct {
	AuthorAccount string `bson:"author_account" json:"author_account"`
	weight        int32  `bson:"weight" json:"weight"`
}

type TransferDetail struct{
	From string `bson:"from" json:"from"`
	To string	`json:"to"`
	Amount string `json:"amount"`
	Memo string	`json:"memo"`
}