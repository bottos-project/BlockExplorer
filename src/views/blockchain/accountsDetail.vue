<template>
  <div class="main">
    <div class="border">
        <div class="title"><h3><span>{{$t('accountsDetail.title')}}</span></h3></div>
        <div class="detail">
          <ul>
              <li>
                  <div class="tit">{{$t('accountsDetail.accounts')}}</div>
                  <div class="con">{{Detail.accountName}}</div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.balance')}}</div>
                  <div class="con">
                    {{Detail.balance / Math.pow(10,8)}} BTO
                    <!--{{currencyAmount}}-->
                    <!-- <select v-model="currencyAmount" class="select">
                      <option v-for="option in currencyAmountOptions" v-bind:value="option.availableAmount">
                        {{ option.currency }}
                      </option>
                    </select> -->
                  </div>
              </li>
              <li>
                  <div class="tit">{{$t('accountsDetail.stakedBalance')}}</div>
                  <div class="con">
                    {{Detail.stakedBalance / Math.pow(10,8)}} BTO
                  </div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.unStakingBalance')}}</div>
                  <div class="con">
                    {{Detail.unStakingBalance / Math.pow(10,8)}} BTO
                  </div>
              </li>
              <li>
                  <div class="tit">{{$t('accountsDetail.freeAvailableSpace')}}</div>
                  <div class="con">
                    {{Detail.resource.freeAvailableSpace}}
                  </div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.freeUsedSpace')}}</div>
                  <div class="con">
                    {{Detail.resource.freeUsedSpace}}
                  </div>
              </li>
              <li>
                  <div class="tit">{{$t('accountsDetail.stakeAvailableSpace')}}</div>
                  <div class="con">
                    {{Detail.resource.stakeAvailableSpace}}
                  </div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.stakeUsedSpace')}}</div>
                  <div class="con">
                    {{Detail.resource.stakeUsedSpace}}
                  </div>
              </li>
              <li>
                  <div class="tit">{{$t('accountsDetail.freeAvailableTime')}}</div>
                  <div class="con">
                    {{Detail.resource.freeAvailableTime}}
                  </div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.freeUsedTime')}}</div>
                  <div class="con">
                    {{Detail.resource.freeUsedTime}}
                  </div>
              </li>
                 <li>
                  <div class="tit">{{$t('accountsDetail.stakeAvailableTime')}}</div>
                  <div class="con">
                    {{Detail.resource.stakeAvailableTime}}
                  </div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.stakeUsedTime')}}</div>
                  <div class="con">
                    {{Detail.resource.stakeUsedTime}}
                  </div>
              </li>
              
              <li>
                  <div class="tit">{{$t('accountsDetail.TransactionNumber')}}</div>
                  <div class="con">{{Detail.tradeCount}}</div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.Transfer')}}</div>
                  <div class="con">
                    <img src="../../assets/arrow_r.png" width="16" /> {{Detail.receiveCount}}
                    &nbsp;&nbsp;
                    <img src="../../assets/arrow_s.png" width="16" /> {{Detail.sendCount}}
                  </div>
              </li>
              <li>
                  <div class="tit">{{$t('accountsDetail.AllBlockReward')}}</div>
                  <div class="con">{{Detail.unClaimedBlockReward / Math.pow(10,8)}} BTO</div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.AllVoteReward')}}</div>
                  <div class="con">{{Detail.unClaimedVoteReward / Math.pow(10,8)}} BTO</div>
              </li>
              <li>
                  <div class="tit">{{$t('accountsDetail.AllReward')}}</div>
                  <div class="con">{{Detail.unClaimedReward / Math.pow(10,8)}} BTO</div>
              </li>
          </ul>
        </div>
    </div>
    <div class="border">
      <div class="tabs">
        <el-tabs>
          <el-tab-pane :label="$t('accountsDetail.tabs1')" class="tabsMenu">
              <keep-alive>
                  <transactions></transactions>
              </keep-alive>
          </el-tab-pane>
          <el-tab-pane :label="$t('accountsDetail.tabs2')">
              <keep-alive>
                  <transfers></transfers>
              </keep-alive>
          </el-tab-pane>
      </el-tabs>
      </div>
    </div>
  </div>
</template>
<script>
    import transactions from './accounts/transactions'
    import transfers from './accounts/transfers'
    import { queryCustDetail } from '@/api/blockchain_accounts'
    export default {
      components: { transactions,transfers},
      data() {
        return {
          custId:this.$route.params && this.$route.params.id,
          Detail:{},
          currencyAmount: '',
          currencyAmountOptions: []
        }
      },
      watch:{
        $route(){ // 只要路径变化 重新获取数据
          this.custId = this.$route.params && this.$route.params.id
          this.fetchData()
        }
      },
      created(){
        this.fetchData()
      },
      methods:{
        fetchData() {
          let para = {
            accountName:this.custId
          }
          queryCustDetail(para).then(response => {
            this.Detail = response.data
            this.currencyAmountOptions = response.data.currencyList
            this.currencyAmount = response.balance
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
<style lang="less" scoped>
select{
  border: solid 1px #000;
  appearance:none;
  -moz-appearance:none;
  -webkit-appearance:none;
  background: url("../../assets/select_arrow.png") no-repeat scroll right center transparent;
  padding: 3px 15px 3px 8px;
  margin-left: 10px;
 
}
select::-ms-expand { display: none; }

</style>