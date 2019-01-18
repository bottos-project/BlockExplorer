<template>
  <div class="main">
    <div class="border">
      <div class="Stats">
        <el-row>
          <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
            <div class="item"> 
              <div class="tit">{{$t('supernode.stats1')}}</div>
              <div class="con">{{statistics.delegateCount}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
            <div class="item">             
              <div class="tit">{{$t('supernode.stats2')}}</div>
              <div class="con">{{statistics.voteCustCount}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
            <div class="item">
              <div class="tit">{{$t('supernode.stats3')}}</div>
              <div class="con">{{statistics.allTransitVotes}}</div>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
    <div class="border">
      <div class="title"><h3><span>{{$t('supernode.title')}}</span></h3></div>
      <div class="supernode">
        <!-- <div class="searchs">
          <div class="input-group">
            <input type="text" v-model="searchVal" class="input" clearable :placeholder="$t('supernode.searchs')" value="">
            <a class="clearInput" @click="clearInput">X</a>
          </div>
        </div> -->
        <div class="lists">
          <el-row :gutter="20" v-if="list.length > 0">
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" v-for="item in list" key="item.id">
              <div class="item">
                <div class="ranking">{{item.rank}}</div>
                <div class="box">
                  <div class="tit">
                    <!-- <router-link :to="'/supernode/detail/'+item.nodeName"> -->
                    {{item.nodeName}}
                    <!-- </router-link> -->
                  </div>
                  <div class="progress">
                    <span class="progress-bar" :style="'width:'+item.votesRate+'%'"></span>
                    <span class="text">{{item.votesRate}}%</span>
                  </div>
                </div>
                <div class="box">
                  <div class="url"><a :href="item.nodeUrl" target="_blank">{{item.nodeUrl}}</a></div>
                  <div class="fr">
                    <div class="ticket">
                      {{item.voteCount / Math.pow(10,8)}} {{$t('supernode.lists_ticket')}}<br />
                      {{item.voteCustCount}} {{$t('supernode.lists_throw')}}
                    </div>
                  </div>
                </div>
                <div class="block">
                  <span class="toBlock" v-if="item.nodeStatus === '正在出块'">{{$t('supernode.lists_block')}}</span>
                </div>
                <div class="des">
                  {{item.nodeDesc}}  
                  <router-link :to="'/supernode/detail/'+item.delegate" class="more">{{$t('supernode.lists_more')}}</router-link>
                </div>
              </div>
            </el-col>
          </el-row>
          <el-row v-else class="noData">
            {{$t('supernode.lists_nodata')}}
          </el-row>            
        </div>
      </div>      
    </div>

  </div>
</template>
<script>
import {queryNodeSummary,queryNodeList} from '@/api/supernode'
    export default {
      data() {
        return {
          statistics:{},
          tableData: [],
          searchVal:''
        }
      },
      created(){
        this.fetchData()
        this.statisticsData()
      },
      methods:{
        statisticsData(){
          let para = {}
          queryNodeSummary(para).then(response => {
            this.statistics = response.data
          }).catch(function(error){
            this.$message({
              message: this.$i18n.t('tips.error'),
              type: 'error'
            });
          })
        },
        fetchData() {
          let para = {}
          queryNodeList(para).then(response => {
            this.tableData = response
          }).catch(function(error){
            this.$message({
              message: this.$i18n.t('tips.error'),
              type: 'error'
            });
          })
        },
        clearInput(){
          this.searchVal = ""
        }
      },
      computed:{
        list: function(){
            var _this = this;
            var arrByZM = [];
            for (var i=0;i<this.tableData.length;i++){
                //for循环数据中的每一项（根据name值）
                if(this.tableData[i].nodeName.search(this.searchVal) != -1){
                    //判断输入框中的值是否可以匹配到数据，如果匹配成功
                    arrByZM.push(this.tableData[i]);
                }
            }
            return arrByZM;
        }
    }
  }
</script>
<style lang="less" scoped>
.searchs {
  width: 100%;
  overflow: hidden;
  .input-group {
    float: right; 
    display: flex;
    position: relative;
    width: 320px;
    .input {
      line-height: 1.5;
      border-radius: 40px;
      background-color: #fff;
      border: solid 1px #a3a3a3;
      color: #333;
      font-size: 14px;
      padding: 10px 15px;
      width: 100%;
    }
    .clearInput {
      position: absolute;
      top: 6px;
      right: 10px;
      width: 30px;
      height: 30px;
      line-height: 30px;
      text-align: center;
    }
  }
}
@media only screen and (max-width:690px){
	.searchs {
    .input-group {
      margin: 0 auto;
      float: none;
    }
	}
} 
.supernode {
  padding: 20px;
}  

.lists {
  width: 100%;
  .item {
    position: relative;
    padding: 20px 30px;
    border: solid 1px #333333;
    border-radius: 10px;
    margin-top: 20px;
    overflow: hidden;
    a {
      color: #5bc1e8;
    }
    .ranking {
      position: absolute;
      top: 0; 
      left: -10px;
      color: #fff;
      width: 40px;
      height: 60px;
      padding-left: 20px;
      line-height: 30px;
      background: url(../../assets/supernode_bg.png) no-repeat;
    }
    .box {
      overflow: hidden;
      padding: 5px 0;
    }
    .tit {
      width: 55%;
      float: left;
      font-size: 22px;
      line-height: 30px;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
      a {
        color: #333333;
      }
    }
    .progress {
      float: right;
      margin-top: 6px;
      position: relative;
      width: 40%;
      height: 20px;
      line-height: 20px;
      background-color: #eee;
      border-radius: 10px;
      .progress-bar {
        float: left;
        display: block;
        width: 0%;
        height: 20px;
        border-radius: 10px;
        background-color: #5bc1e8;
      }
      .text{
        position: absolute;
        right: 5px;
        line-height: 20px;
        font-size: 13px;
      }
    }
    .url {
      float: left;
      width: 60%;
      line-height: 20px;
    }
     .ticket {
      text-align: right;
      line-height: 20px;
    }
    .block {
      margin-top: -20px;
      height: 24px;
      margin-bottom: 10px;
      .toBlock {
        height: 24px;
        line-height: 24px;
        display: inline-block;
        padding: 0 15px;
        background-color: #5bc1e8;
        color: #fff;
        border-radius: 15px;
      }
    }
   
    .des {
      line-height: 22px;
      font-size: 12px;
      height: 44px;
      overflow: hidden;
      position: relative;
      .more {
        position: absolute;
        bottom: 0; right: -10px;
        padding: 0px 10px;
        background-color: #fff;
      }
    }
  }
}
</style>
