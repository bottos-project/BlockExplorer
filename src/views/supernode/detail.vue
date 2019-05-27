<template>
  <div class="main">
    <div class="border">
      <div class="detail">
        <!-- <div class="information">
          <div class="horn">
            <img src="../../assets/supernode_detail_horn.png" />
          </div>
          <el-row :gutter="100">
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
              <div class="box bg">
                <dl>
                  <dd>{{Detail.nodeName}}</dd>
                  <dd><a href="#">{{Detail.nodeUrl}}</a></dd>
                  <dd><span class="icon_position">{{Detail.city}},{{Detail.country}}</span></dd>
                </dl>
              </div>
            </el-col>
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
              <div class="box">
                <dl>
                  <dd><label>{{$t('supernodeDetail.ranking')}}</label> {{Detail.voteRank}}</dd>
                  <dd><label>{{$t('supernodeDetail.got_vote')}}</label> {{Detail.votes/ Math.pow(10,8)}} {{$t('supernodeDetail.ticket')}}（{{(Detail.voteCount/Detail.allTransitVotes*100).toFixed(8)}}%）</dd>
                  <dd><label>{{$t('supernodeDetail.cast_vote')}}</label> {{Detail.voteCustCount}} {{$t('supernodeDetail.people')}}</dd>
                </dl>
              </div>
            </el-col>
          </el-row>
          <div class="des">
            <div :class="{showall:true,active:showall}">
              {{Detail.nodeDesc}}
            </div>
            <a class="open" @click="openDes" v-if="showall == false">{{$t('supernodeDetail.expand')}}</a>
          </div>
        </div> -->
        <ul>
            <li class="row">
                <div class="tit">{{$t('supernodeDetail.account')}}</div>
                <div class="con">
                  <router-link :to="'/blockchain/accountsDetail/'+Detail.delegate">
                    {{Detail.delegate}}
                  </router-link>
                </div>
            </li>
            <li>
              <div class="tit">{{$t('supernodeDetail.table11')}}</div>
              <div class="con">{{Detail.public_key}}</div>
            </li>
            <li class="row">
              <div class="tit">{{$t('supernodeDetail.quality_deposit')}}</div>
              <div class="con">{{Detail.detail.staked_balance / Math.pow(10,8)}} BTO</div>
            </li>
            <li>
              <div class="tit">{{$t('supernodeDetail.quality_deposit')}}</div>
              <div class="con">{{Detail.votes / Math.pow(10,8)}}</div>
            </li>
            <li class="row">
              <div class="tit">{{$t('supernodeDetail.tabs2')}}</div>
              <div class="con">{{Detail.voters}}</div>
            </li>
            <li>
              <div class="tit">{{$t('supernodeDetail.number_of_blocks')}}</div>
              <div class="con">{{Detail.produce_block_count}}</div>
            </li>
            <li class="row">
              <div class="tit">{{$t('supernodeDetail.table10')}}</div>
              <div class="con">{{Detail.detail.un_claimed_reward / Math.pow(10,8)}} BTO</div>
            </li>
        </ul>
      </div>
    </div>
    <div class="border">
      <div class="tabs">
        <el-tabs>
          <el-tab-pane :label="$t('supernodeDetail.tabs1')">
              <keep-alive>
                  <BlockLists></BlockLists>
              </keep-alive>
          </el-tab-pane>
          <el-tab-pane :label="$t('supernodeDetail.tabs2')">
              <keep-alive>
                  <Voters></Voters>
              </keep-alive>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </div>
</template>
<script>
    import BlockLists from './BlockLists'
    import Voters from './Voters'
    import { queryNodeDetail } from '@/api/supernode'
    export default {
      components: { BlockLists,Voters},
      data() {
        return {
          custId:this.$route.params && this.$route.params.id,
          Detail:{},
          currencyAmount: '',
          currencyAmountOptions: [],
          showall:false
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
            nodeName:this.custId
          }
          queryNodeDetail(para).then(response => {
            this.Detail = response
          }).catch(error => {
            this.$message({
                message: this.$i18n.t('tips.error'),
                type: 'error'
            });
          })
        },
        openDes(){
          this.showall = true
        }
      }
    }
</script>
<style lang="less" scoped>
.information {
  padding: 30px 50px;
  overflow: hidden;
  position: relative;
  .horn {
    position: absolute;
    top: 0;
    right: 0;
  }
  .box {
    padding: 10px 20px;
    border-radius: 10px;
    label {
      display: inline-block;
      width: 100px;
    }
  }
  .box.bg {
    background-color: #5bc1e8;
    color: #fff;
    a {
      color: #fff;
    }
  }
  .des {
    padding-top:20px; 
    position: relative;
    line-height: 24px;
    padding-bottom: 24px;
    .open {
      position: absolute;
      bottom:0;
      right: 0; 
      color: #5bc1e8;
    }
  }
  .showall {
    height: 72px;
    overflow: hidden;
    clear: both;
  }
  .showall.active {  //点击了查看更多，就高度就不管了。随实际的大小变化
    height: auto;
  }
}
</style>