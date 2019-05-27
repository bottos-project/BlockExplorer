import enLocale from 'element-ui/lib/locale/lang/en'
const messagesEn = {
    menus:{
        home: 'Home',
        blockchain: 'Blockchain',
        blocks: 'Blocks',
        transactions:'Transactions',
        transfers:'Transfers',
        accounts: 'Accounts',
        statistics: 'Statistics',
        contracts: 'Contracts',
        nodes: 'Service Nodes',
        supernode:'Production Nodes'
    },
    network:{
        network:"network",
        mainnet:"mainnet",
        testnet:"testnet"
    },
    footer:{
        title:'Contact Us'
    },
    tips:{
        error:"Data load error, please check the data interface!",
        nodata:"No Data!"
    },
    home:{
        banner_title:'Bottos Blocks Explorer',
        searchs:'Search Accounts,Blocks Height,Blocks Hash,Transactions Hash',
        searchs_tips:'Record not found!',
        stats1:'Nodes',
        stats2:'Block Height',
        stats3:'Transactions Last Day',
        stats4:'Total Accounts',
        card1:'Past 14 days of total transactions',
        card1_series_name:'Total Transactions',
        card2:'Accounts Growth',
        card2_series_name:'Total Accounts',
        card3:'Blocks',
        card3_block:'Block',
        card3_transactions:'Transactions',
        card3_delegate:'Miner',
        card3_reward:'Block Reward',
        card4:'Transfers',
        card4_transactions:'Transactions',
        card_btn:'View All',
        testnet:'Testnet',
    },
    blocks:{
        title:"Blocks",
        table1:"Height",
        table2:"Age",
        table3:"Txn",
        table4:"Miner",
        table5:"Size",
        table6:"Reward"
    },
    blocksDetail:{
        title:"Block Detail",
        state:"Status",
        state_yes:"Confirmed",
        state_no:"Unconfirmed",
        BlockID:"Block Hash",
        height:"Height",
        time:"Age",
        TransactionNumber:"Transactions",
        ThePreviousBlockID:"Parent Block Hash",
        Chunk:"Miner",
        Size:"Size",
        tabs1:"Transactions",
        tabs2:"Transfers"
    },
    transactions:{
        title:"Transactions",
        table1:"Transaction Hash",
        table2:"Block",
        table3:"Age",
        table4:"Sender",
        table5:"Contract Name",
        method_name:"Method Name",
        time_resource:"Time Res",
        space_resource:"Space Res",
    },
    transactionsDetail:{
        title:"Transaction Detail",
        state:"Status",
        state_yes:"Confirmed",
        state_no:"Unconfirmed",
        transactionsID:"Transaction Hash",
        block:"Block",
        time:"Age",
        time_resource:"Time Res",
        space_resource:"Space Res",
        title2:"Contract",
        name:"Contract Name",
        method_name:"Method Name",
        input:"Parameter"
    },
    transfers:{
        title:"Transfers",
        table1:"Transaction Hash",
        table2:"Block",
        table3:"Age",
        table4:"From",
        table5:"To",
        table6:"Value",
        table7:"->",
    },
    transfersDetail:{
        title:"Transfer Detail",
        state:"Status",
        state_yes:"Confirmed",
        state_no:"Unconfirmed",
        transactionsID:"Transaction Hash",
        block:"Block",
        time:"Age",
        time_resource:"Time Res",
        space_resource:"Space Res",
        title2:"Contract - Transfer",
        name:"Contract Name",
        method_name:"Method Name",
        input:"Parameter"
    },
    accounts:{
        title:"Accounts",
        title_statistics:"Total Accounts",
        table1:"Transfers",
        table2:"Balance",
        table3:"Possession"
    },
    accountsDetail:{
        title:"Account Detail",
        accounts:"Account",
        balance:"Balance", 
        stakedBalance: "Staked Balance",
        unStakingBalance: "UnStaking Balance",
        freeAvailableSpace: "Free Available Space",
        freeUsedSpace:"Free Used Space",
        stakeAvailableSpace: "Stake Available Space",
        stakeUsedSpace: "Stake Used Space",
        freeAvailableTime: "Free Available Time",
        freeUsedTime: "Free Used Time",
        stakeAvailableTime: "Stake Available Time",
        stakeUsedTime: "Stake Used Time",
        TransactionNumber:"Transactions",
        Transfer:"Transfers",
        tabs1:"Transactions",
        AllReward:"Undrawn rewards",
        AllBlockReward:"Undrawn mining rewards",
        AllVoteReward:"Undrawn voting rewards",
        tabs2:"Transfers"
    },
    statistics:{
        title:"Statistics",
        card1:"Transactions Per Day",
        card2:"Account Growth"
    },
    contracts:{
        title:"Contracts",
        searchs:"Search for contract name",
        searchsbtn:"Searchs",
        table1:"ContractName",
        table2:"Address",
        table3:"TxCount",
        table4:"Account size"
    },
    contractsDetail:{
        title:"Contract Details",
        name:"Contract Name",
        address:"Contract Address",
        TransactionNumber:"Transaction Number",
        AccountSize:"Account size",
        Record:"Invoke Record",
        code:"Code",
        table1:"Transaction Hash",
        table2:"Block",
        table3:"Invoke Account",
        table4:"Invoke Time",
        table5:"Detail",
        contractsAPI:"Contract API",
        contractsByte:"Contract Byte"
    },
    nodes:{
        stats1:"Production Nodes",
        stats2:"Service Nodes",
        title:'Nodes',
        searchs:"Nodes Name",
        searchsbtn:"Searchs",
        table1:'Node',
        table2:'Service Type',
        table3:'Region',
        table4:'Connection',
        table5:'Surplus',
        table6:'Service API',
    },
    supernode:{
        stats1:"Production Node",
        stats2:"Total Number of Votes",
        stats3:"Total Number Turnout",
        title:"Production Node",
        searchs:"Search For Production Node Name",
        lists_block:"Now Productivity",
        lists_ticket:"Ticket",
        lists_throw:"People Already Voted",
        lists_more:"More",
        lists_nodata:"No Related Nodes Were Searched."
    },
    supernodeDetail:{
        ranking:"Ranking",
        got_vote:"Got Vote",
        ticket:"Ticket",
        cast_vote:"Cast Vote",
        people:"People",
        expand:"Expand",
        account:"Account",
        quality_deposit:"Quality Deposit",
        number_of_blocks:"Blocks Produced",
        block_success_rate:"Block Success Rate",
        serviceApi:"Service API",
        tabs1:"Produced Blocks",
        tabs2:"Voter",
        table1:"Hight",
        table2:"Age",
        table3:"Transactions",
        table4:"Size",
        table5:"Voter People",
        table6:"Cast Vote Number",
        table7:"This Node Voting Ratio",
        table8:"Overall Voting Share",
        table9:"Vote Time",
        table10:"Un Claimed Reward",
        table11:"Public Key"
    },
    aboutUS:{
        weChat:"WeChat",
        technologyDevelopmentZone:"Dev.Zone"
    },
    ...enLocale
}

export default messagesEn