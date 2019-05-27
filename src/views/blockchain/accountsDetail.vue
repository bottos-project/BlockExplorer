<template>
  <div class="main">
    <div class="border">
        <div class="title"><h3><span>{{$t('accountsDetail.title')}}</span></h3></div>
        <div class="detail-left">
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
                  <div class="tit">{{$t('accountsDetail.TransactionNumber')}}</div>
                  <div class="con">{{Detail.tradeCount}}</div>
              </li>
              <li class="row">
                  <div class="tit">{{"delegate"}}</div>
                  <div class="con">{{Detail.vote && Detail.vote.delegate || ""}}</div>
              </li>
              <li>
                  <div class="tit">{{"vote"}}</div>
                  <div class="con">{{Detail.vote && Detail.vote.votes / Math.pow(10,8) || 0}}</div>
              </li>
              <li class="row">
                  <div class="tit">{{$t('accountsDetail.Transfer')}}</div>
                  <div class="con">
                    <img src="../../assets/arrow_r.png" width="16" /> {{Detail.receiveCount}}
                    &nbsp;&nbsp;
                    <img src="../../assets/arrow_s.png" width="16" /> {{Detail.sendCount}}
                  </div>
              </li>
              <!-- <li>
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
              </li> -->
          </ul>
        </div>
        <div class="detail-right">
          <ul>
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
                    {{Detail.resource.stakeUsedTime || ""}}
                  </div>
              </li>
              <li>
                <div class="tit">{{"多签账户授权列表"}}</div>
                <div class="con">{{msignAccounts.param && msignAccounts.param.authority || ""}}</div>
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
    import { queryCustDetail,queryMsignAccount,queryMsignProposal } from '@/api/blockchain_accounts'
    export default {
      components: { transactions,transfers},
      data() {
        return {
          custId:this.$route.params && this.$route.params.id,
          Detail:{},
          currencyAmount: '',
          currencyAmountOptions: [],
          msignAccounts:{},
          msignProposal:{}
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
        this.getMsignAccount()
        this.getMsignProposal()
      },
      methods:{
        fetchData() {
          let para = {
            accountName:this.custId
          }
          queryCustDetail(para).then(response => {
            console.log({queryCustDetail:response})
            this.Detail = response.data
            console.log({Detail:this.Detail})
            this.currencyAmountOptions = response.data.currencyList
            this.currencyAmount = response.balance
          }).catch(error => {
            this.$message({
                message: this.$i18n.t('tips.error'),
                type: 'error'
            });
          })
        },
        getMsignAccount(){
          queryMsignAccount({author_account:this.custId}).then(response=>{
            console.log({queryMsignAccount:response})
            if(response){
              this.msignAccounts = response[0]
            }
          }).catch(error=>{
            console.log({error})
          })
        },
        getMsignProposal(){
          queryMsignProposal({author_account:this.custId}).then(response=>{
            console.log({queryMsignProposal:response})
            if(response){
              this.msignProposal = response && response[0]
            }
          }).catch(error=>{
            console.log({error})
          })
        }
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

/*详情*/
.detail-left {
  float: left;
  width: 49%;
  border:2px ;
  border-radius:5px;
	ul {
		li {
			overflow: hidden;
			line-height: 28px;
			padding: 10px 30px;
			.tit {float: left; width: 40%;}
			.con {
        float: left;
				word-break:break-all;
				padding-left: 30%;
				.confirm {
					color: #27ce00;
					padding: 3px 10px;
					border: solid 1px #27ce00;
					border-radius: 15px;
				}
				.Unconfirm {
					color: #eb0000;
					padding: 3px 10px;
					border: solid 1px #eb0000;
					border-radius: 15px;
				}
				pre {outline: 1px solid #f6f6f6; background:#f6f6f6; padding: 15px 20px; margin-bottom:20px;
				  .string { color: green;word-break:normal; white-space:pre-wrap;word-wrap : break-word ;overflow: hidden ; }
				  .number { color: darkorange; }
				  .boolean { color: blue; }
				  .null { color: magenta; }
				  .key { color: red; }
				}
			}
		}
		.row {
			background-color: #f6f6f6;
		}
		li:last-child {
			border-bottom: 0;
		}
		a {
			color: #0099CC;
    }
	}
}


/*详情*/
.detail-right {
  float: right;
  width: 49%;
  border:2px ;
  border-radius:5px;
	ul {
		li {
			overflow: hidden;
			line-height: 28px;
			padding: 10px 30px;
			.tit {float: left; width: 40%;}
			.con {
        float: left;
				word-break:break-all;
				padding-left: 40%;
				.confirm {
					color: #27ce00;
					padding: 3px 10px;
					border: solid 1px #27ce00;
					border-radius: 15px;
				}
				.Unconfirm {
					color: #eb0000;
					padding: 3px 10px;
					border: solid 1px #eb0000;
					border-radius: 15px;
				}
				pre {outline: 1px solid #f6f6f6; background:#f6f6f6; padding: 15px 20px; margin-bottom:20px;
				  .string { color: green;word-break:normal; white-space:pre-wrap;word-wrap : break-word ;overflow: hidden ; }
				  .number { color: darkorange; }
				  .boolean { color: blue; }
				  .null { color: magenta; }
				  .key { color: red; }
				}
			}
		}
		.row {
			background-color: #f6f6f6;
		}
		li:last-child {
			border-bottom: 0;
		}
		a {
			color: #0099CC;
    }
	}
}
</style>