<template>
  <div class="main">
    <div class="border">
      <div class="title">
        <h3><span>{{$t('statistics.title')}}</span></h3>
      </div>
    </div>
    <div class="statistics">
      <el-row :gutter="20">
        <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
          <div class="card">
            <div class="card-header"><div class="icon"><img src="../../assets/home_card_1.png" height="42" /></div><h3>{{$t('statistics.card1')}}</h3></div>
            <transactions></transactions>
          </div>
        </el-col>
        <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
          <div class="card">
            <div class="card-header"><div class="icon"><img src="../../assets/home_card_2.png" height="42" /></div><h3>{{$t('statistics.card2')}}</h3></div>
            <accounts></accounts>
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
      
    
</template>
<script>
import transactions from './statistics/statisticsTransactions' //24小时交易数
import accounts from './statistics/statisticsAccounts'//账户增长
    export default {
      components:{
        transactions,accounts
      },
      data() {
        return {
        }
      }
    }
</script>
<style lang="less" scoped>
.statistics {
  .el-row {
  }
  .card {
    margin-bottom: 20px;
    border: solid 1px #d8d8d8;
    border-radius: 5px;
    background-color: #fff;
    border-bottom:solid 5px #1d2327; 
  }
  .card-header {
    padding: 15px 15px 5px 15px;
    position: relative;
    border-bottom: solid 1px #f0f0f0;
    text-align: center;
    h3 {
      font-size: 20px;height: 40px; line-height:40px;
      color: #5bc1e8;
    }
  }
}

</style>
