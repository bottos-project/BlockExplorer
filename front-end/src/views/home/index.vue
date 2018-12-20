<template>
  <div class="home">
    <div class="banner">
      <div class="inner cf">
        <h1>{{$t('home.banner_title')}}</h1>
        <div class="input-group" :class="{shake:isShake}">
          <input type="text" v-model="searchInput" class="input" :placeholder="$t('home.searchs')" value="" @keyup.13="searchData">
          <div class="input-btn">
            <button @click="searchData"><img src="../../assets/icon_view.png" width="26" height="26" /></button>
          </div>
        </div>
        <!-- <div class="homeStats">
          <el-row>
            <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6">
             1
            </el-col>
            <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6">
              2
            </el-col>
            <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6">
              3
            </el-col>
            <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6">
              4
            </el-col>
          </el-row>
        </div> -->
        <div class="homeStats">
          <el-row>
            <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6">
              <router-link to="/nodes">
                <div class="con">
                  {{PrimarySummary.nodeCount}}
                </div>
                <div class="tit">{{$t('home.stats1')}}</div>
              </router-link>
            </el-col>
            <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6">
              <router-link to="/blockchain/blocks">
                <div class="con">
                  {{PrimarySummary.blockNum}}
                </div>
                <div class="tit">{{$t('home.stats2')}}</div>
              </router-link>
            </el-col>
            <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6">
              <router-link to="/blockchain/transactions">
                <div class="con">
                  {{PrimarySummary.lastTradeCount}}
                </div>
                <div class="tit">{{$t('home.stats3')}}</div>
              </router-link>
            </el-col>
            <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6">
              <router-link to="/blockchain/accounts">
                <div class="con">
                  {{PrimarySummary.rtCustCount}}
                </div>
                <div class="tit">{{$t('home.stats4')}}</div>
              </router-link>
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
    
    <div class="row">
      <div class="inner cf">
        <el-row :gutter="20">
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <div class="card">
              <div class="card-header"><div class="icon"><img src="../../assets/home_card_1.png" height="42" /></div><h3>{{$t('home.card1')}}</h3></div>
              <transactions></transactions>
            </div>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <div class="card">
              <div class="card-header"><div class="icon"><img src="../../assets/home_card_2.png" height="42" /></div><h3>{{$t('home.card2')}}</h3></div>
              <accounts></accounts>
            </div>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <div class="card">
              <div class="card-header"><div class="icon"><img src="../../assets/home_card_3.png" height="42" /></div><h3>{{$t('home.card3')}}</h3></div>
              <blocks></blocks>
              <router-link class="view" to="/blockchain/blocks">{{$t('home.card_btn')}}</router-link>
            </div>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
            <div class="card">
              <div class="card-header"><div class="icon"><img src="../../assets/home_card_4.png" height="42" /></div><h3>{{$t('home.card4')}}</h3></div>
              <transfers></transfers>
               <router-link class="view" to="/blockchain/transfers">{{$t('home.card_btn')}}</router-link>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>
<script>
import { PrimarySummary,queryFuzzyInfo } from '@/api/home'
import transactions from './cardTransactions' //过去14天交易数
import accounts from './cardAccounts'//账户增长
import blocks from './cardBlocks'//账户增长
import transfers from './cardTransfers'//账户增长
//import NumberGrow from '@/components/NumberGrow'//账户增长
export default {
    components:{
      transactions,accounts,blocks,transfers
    },
    data(){
      return {
        searchInput:'',
        isShake:false,
        PrimarySummary:{
          nodeCount:null,//线上节点数
          blockNum:null,//区块高度
          lastTradeCount:null,//过去一天交易数
          rtCustCount:null//实时总账户数
        }
      }
    },
    created(){
      this.PrimarySummaryData()
    },        
    methods: {
      searchData(){
        let para = {
          condition:this.searchInput !== ""? this.searchInput : "null"
        };
        queryFuzzyInfo(para).then(response => {
          let than = this
          if(response.message === "无结果"){
            this.isShake = true
            setTimeout(function(){
              than.isShake = false
            }, 3000 )
            this.$message({
              message: this.$i18n.t('home.searchs_tips'),
              type: 'warning'
            });
          }
          //区块
          if(response.message === "搜索成功"){
            if(response.data.resultType === "block"){
              console.log('区块高度跳转')
              console.log(response.data.blockNum)
              this.$router.push('/blockchain/blocksDetail/'+response.data.blockNum);
            }
            if(response.data.resultType === "trade"){
              console.log('交易ID跳转')
              this.$router.push('/blockchain/transactionsDetail/'+response.data.transactionId);
            }
            if(response.data.resultType === "cust"){
              this.$router.push('/blockchain/accountsDetail/'+response.data.accountName);
            }
          }
              
        }).catch(error => {
          this.$message({
            message: this.$i18n.t('tips.error'),
            type: 'error'
          });
        })
      },
      PrimarySummaryData() {
        let para = {}
        PrimarySummary(para).then(response => {
          console.log({response})
          this.PrimarySummary = response.data
        }).catch(error => {
          this.$message({
            message: this.$i18n.t('tips.error'),
            type: 'error'
          });
        })
      },
    }
}
</script>
<style scoped lang="less">
.banner {
   background: url(../../assets/homebg.jpg) no-repeat top center;
   background-size: cover;
   widows: 100%;
   padding: 0;
   position: relative;
   color: #fff;
   padding: 50px 0 85px 0;
   .inner {
    position: relative;
  }
   h1 {
    font-size: 50px;
    line-height: 50px;
    text-align: center;
  }
  .input-group {
      display: flex;
      margin: 50px auto 0 auto;
      position: relative;
      width: 50%;
      .input {
        line-height: 1.5;
        border-radius: 10px;
        border-top-right-radius: 0;
        border-bottom-right-radius: 0;
        background-color: transparent;
        border: solid 1px #fff;
        border-right: 0;
        font-size: 14px;
        padding: 15px 30px;
        width: 100%;
        color: #fff;
      }
      .input::-webkit-input-placeholder {color:#fff;}
      .input:-moz-placeholder {color:#fff;}
      .input::-moz-placeholder {color:#fff;}
      .input:-ms-input-placeholder {color:#fff;}
  }
  .input-btn {
    margin-left: -1px;
    display: flex;
    box-sizing: border-box;
    text-align: center;
    button {
      line-height: 1.5;
      border-radius: 10px;
      border-top-left-radius: 0;
      border-bottom-left-radius: 0;
      border: 1px solid transparent;
      border-color: #5bc1e8;
      background-color: #5bc1e8;
      color: #fff;
      cursor: pointer;
      transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out, border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out, -webkit-box-shadow 0.15s ease-in-out;
      padding: 0 20px;
    }
    button:hover {
      background-color: #45a5ca;
    }
  }
  .shake {
    -webkit-animation-name: shake;
    animation-name: shake;
    -webkit-animation-duration: 1s;
    animation-duration: 1s;
    -webkit-animation-fill-mode: both;
    animation-fill-mode: both;
  }
 }
@media only screen and (max-width:690px){
	.banner {
    padding: 50px 0;
	}
} 
@media only screen and (max-width:650px){
	.banner {
    h1 {
      font-size: 40px
    }
    .input-group {
      width: 100%;
    }
  }
}
@media only screen and (max-width:450px){
	.banner {
    h1 {
      font-size: 30px
    }
  }
}
@media only screen and (max-width:350px){
	.banner {
    h1 {
      font-size: 24px
    }
  }
}
 
.homeStats {
  margin-top: 50px;
  padding: 20px 0;
  a {
    display: block;
    padding: 15px;
    text-align: center;
    color: #fff;
  }
  a:hover {
    color: #ccc;
  }
  .tit {
    font-size: 16px;
    position: relative;
  }
  .tit:before {
        content: "";
        display: block;
        position: absolute;
        bottom: -6px;
        left: 50%;
        margin-left: -12px;
        width: 24px;
        height: 3px;
        background-color: #5bc1e8;
        border-radius: 3px;
    }
  .con {
    font-size: 40px;
    line-height: 60px;
    height: 60px;
  }
}
@media only screen and (max-width:690px){
	.homeStats {
    margin-top: 50px;
    text-align: center;
    padding: 15px 0;
    a {
      padding: 10px;
    }
  }
  .el-row>div:last-child a{
    border-bottom: 0;
  }
  .row {
    margin-top: 0;
  }
}

.row {
  padding: 20px 0 20px 0;
  background-color: #e8e8e8;
  margin-top: -100px;
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
  .view {
    padding: 0 10px;
    background-color: #5bc1e8;
    color: #fff;
    border-radius: 5px;
    display: block;
    width: 110px;
    text-align: center;
    margin: 15px auto;
    line-height: 40px;
  }
  .view:hover {
    background-color: #45a5ca;
  }
}
@media only screen and (max-width:690px){
  .row {
    margin-top: 0;
  }
}

@-webkit-keyframes shake {
  from,to {
    -webkit-transform: translate3d(0, 0, 0);
    transform: translate3d(0, 0, 0);
  }
  10%,30%,50%,70%,90% {
    -webkit-transform: translate3d(-10px, 0, 0);
    transform: translate3d(-10px, 0, 0);
  }
  20%,40%,60%,80% {
    -webkit-transform: translate3d(10px, 0, 0);
    transform: translate3d(10px, 0, 0);
  }
}

@keyframes shake {
  from,to {
    -webkit-transform: translate3d(0, 0, 0);
    transform: translate3d(0, 0, 0);
  }
  10%,30%,50%,70%,90% {
    -webkit-transform: translate3d(-10px, 0, 0);
    transform: translate3d(-10px, 0, 0);
  }
  20%,40%,60%,80% {
    -webkit-transform: translate3d(10px, 0, 0);
    transform: translate3d(10px, 0, 0);
  }
}

</style>
