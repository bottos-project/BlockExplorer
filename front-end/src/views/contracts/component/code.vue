<template>
  <div>
    <div class="box">
      <div class="tit">{{$t('contractsDetail.contractsAPI')}}</div>
      <div class="con">
        {{inputdata1}}
      </div>
    </div>
    <div class="box">
      <div class="tit">{{$t('contractsDetail.contractsByte')}}</div>
      <div class="con">
        {{inputdata2}}
      </div>
    </div>
  </div>
</template>
<script>
//import {queryTradeList} from '@/api/blockchain_transactions'

export default {
    data(){
      return {
        custId:this.$route.params && this.$route.params.id,
        inputdata1:'{"entrys":[{"name":"setRandomUser","inputs":[{"name":"_addr","type":"address"}],"type":2,"stateMutability":3},{"name":"setImporter","inputs":[{"name":"_addr","type":"address"},{"name":"_useful","type":"bool"}],"type":2,"stateMutability":3},{"name":"doUnpause","type":2,"stateMutability":3},{"name":"doPause","type":2,"stateMutability":3},{"name":"setAdmin","inputs":[{"name":"_newAdmin","type":"address"}],"type":2,"stateMutability":3},{"name":"importSeedFromThird","inputs":[{"name":"_bettor","type":"address"},{"name":"_seed","type":"uint256"}],"type":2,"stateMutability":3},{"constant":true,"name":"isPaused","outputs":[{"type":"bool"}],"type":2,"stateMutability":2},{"type":1,"stateMutability":3}]}',
        inputdata2:'60806040526000805460ff1916905534801561001a57600080fd5b5060008054610100330261010060a860020a031991821681179091161790556103bf806100486000396000f3006080604052600436106100825763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663018e087381146100875780632b643cdc146100aa57806330efb8d3146100d057806367d0661d146100e5578063704b6c02146100fa57806376eb5cb21461011b578063b187bd261461013f575b600080fd5b34801561009357600080fd5b506100a8600160a060020a0360043516610168565b005b3480156100b657600080fd5b506100a8600160a060020a036004351660243515156101b3565b3480156100dc57600080fd5b506100a861020f565b3480156100f157600080fd5b506100a8610237565b34801561010657600080fd5b506100a8600160a060020a0360043516610262565b34801561012757600080fd5b506100a8600160a060020a03600435166024356102c8565b34801561014b57600080fd5b5061015461038a565b604080519115158252519081900360200190f35b6000546101009004600160a060020a0316331461018457600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000546101009004600160a060020a031633146101cf57600080fd5b600160a060020a03821615156101e457600080fd5b600160a060020a03919091166000908152600260205260409020805460ff1916911515919091179055565b6000546101009004600160a060020a0316331461022b57600080fd5b6000805460ff19169055565b6000546101009004600160a060020a0316331461025357600080fd5b6000805460ff19166001179055565b6000546101009004600160a060020a0316331461027e57600080fd5b600160a060020a038116151561029357600080fd5b60008054600160a060020a039092166101000274ffffffffffffffffffffffffffffffffffffffff0019909216919091179055565b3360009081526002602052604090205460ff16806102f557506000546101009004600160a060020a031633145b151561030057600080fd5b600154604080517f67c18aa1000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015260248201859052915191909216916367c18aa191604480830192600092919082900301818387803b15801561036e57600080fd5b505af1158015610382573d6000803e3d6000fd5b505050505050565b60005460ff16815600a165627a7a723058202aa146dbe4b2f04208a5e0fe7f55aa95e904e78fca15932ca6da935642d78f780029'
      }
    },
    watch:{
      $route(){ // 只要路径变化 重新获取数据
        this.custId = this.$route.params && this.$route.params.id
        this.fetchData()
      }
    },
    created(){
      //this.fetchData()
    },        
    methods: {
      fetchData() {
        let para = {
          accountName:this.custId,
          start:this.listQuery.start,
          length:this.listQuery.length
        };
        queryTradeList(para).then(response => {
          this.tableData = response.data.data
          this.listQuery.iTotalDisplayRecords = response.data.iTotalDisplayRecords
          this.listLoading = false
        }).catch(function(error){
          console.log('加载出错……');
        })
        
      },

      
    }
}
</script>
<style scoped lang="less">
  .box {
    padding: 10px 30px;
    .tit {
      font-size: 16px;
      margin-bottom: 10px;
    }
    .con {
      outline: 1px solid #f6f6f6;
      background: #f6f6f6;
      padding: 15px 20px;
      margin-bottom: 20px;
      word-wrap:break-word
    }
  }
</style>
