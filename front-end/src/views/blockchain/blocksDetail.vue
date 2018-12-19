<template>
  <div class="main">
    <div class="border">
        <div class="title"><h3><span>{{$t('blocksDetail.title')}}</span></h3></div>
        <div class="detail">
            <ul>
                <li>
                    <div class="tit">{{$t('blocksDetail.state')}}</div>
                    <div class="con">
                        <span class="confirm" v-show="Detail.blockStatus === '已确认'">{{$t('blocksDetail.state_yes')}}</span>
                        <span class="Unconfirm" v-show="Detail.blockStatus === '未确认'">{{$t('blocksDetail.state_no')}}</span>
                    </div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('blocksDetail.BlockID')}}</div>
                    <div class="con">{{Detail.blockHash}}</div>
                </li>
                <li>
                    <div class="tit">{{$t('blocksDetail.height')}}</div>
                    <div class="con">{{Detail.blockNum}}</div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('blocksDetail.time')}}</div>
                    <div class="con">
                        {{Detail.tradeDate | dateformat}} 
                        （{{Detail.tradeDate | dateformatNow}}）
                    </div>
                </li>
                <li>
                    <div class="tit">{{$t('blocksDetail.TransactionNumber')}}</div>
                    <div class="con">{{Detail.transCount}}</div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('blocksDetail.ThePreviousBlockID')}}</div>
                    <div class="con">
                        <router-link :to="'/blockchain/blocksDetail/'+Detail.prevBlockNum">
                            {{Detail.prevBlockHash}}
                        </router-link>                    
                    </div>
                </li>
                <li>
                    <div class="tit">{{$t('blocksDetail.Chunk')}}</div>
                    <div class="con">
                        <router-link :to="'/blockchain/accountsDetail/'+Detail.delegate">
                            {{Detail.delegate}}
                        </router-link>
                    </div>
                </li>
            </ul>
        </div>
    </div>
    <div class="border">
        <div class="tabs">
            <el-tabs>
                <el-tab-pane :label="$t('blocksDetail.tabs1')">
                    <keep-alive>
                        <transactions></transactions>
                    </keep-alive>
                </el-tab-pane>
                <el-tab-pane :label="$t('blocksDetail.tabs2')">
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
    import transactions from './blocks/transactions'
    import transfers from './blocks/transfers'
    import { queryBlockDetail } from '@/api/blockchain_blocks'
    export default {
      components: { transactions,transfers},
      data() {
        return {
          custId:this.$route.params && this.$route.params.id,
          Detail:{}
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
            blockNum:this.custId
          }
          queryBlockDetail(para).then(response => {
              console.log({response})
            this.Detail = response
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
