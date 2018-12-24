<template>
    <div class="card-body scrollbar">
       <ul class="list">
            <li v-for="(item,index) in tableList" :class="{on:index%2==0,off:index%2!=0}">
                <router-link :to="'/blockchain/transfersDetail/'+item.transactionId">
                <div class="item">
                    <span class="ids oh" style="width: 70%;">{{$t('home.card4_transactions')}}ID {{item.transactionId}}</span>
                    <div class="time">{{item.tradeDate | dateformat}}</div>
                </div>
                <div class="item">
                    <div class="fl" style="width:60%">
                        <span class="oh fl" style="max-width: 40%;">{{item.sender}}</span>
                        <span class="fl"><img src="../../assets/icon_arrow.png" width="20" height="20" /></span>
                        <span class="oh fl" style="max-width: 40%;">{{item.receiver}}</span>
                    </div>
                    <div class="fr">
                        <span class="icon"><img src="../../assets/icon_transfers.png" width="20" height="20" /></span>
                        {{item.tradeAmount}}
                    </div>
                </div>
                </router-link>
            </li>
            <div class="noData" v-if="tableList.length === 0">{{$t('tips.nodata')}}</div>
        </ul>
    </div>
</template>
<script>
import { queryTransferlList } from '@/api/home'
export default {
    data(){
      return {
        tableList:{}
      }
    },
    created(){
      this.fetchData()
      //window.refsDate = setInterval(this.fetchData,2000)
    },        
    methods: {
          
      fetchData() {
        let para = {
          start:0,
          length:20
        }
        queryTransferlList(para).then(response => {
          this.tableList = response.data
          if(response.data.data.length === 0){
            
          }
        }).catch(error => {
        //   this.$message({
        //     message: this.$i18n.t('tips.error'),
        //     type: 'error'
        //   });
        })
      }, 
    },
    beforeDestroy(){ // 可以清除定时器 或者清除事件绑定
        //clearInterval(window.refsDate);
    }
}
</script>
<style scoped lang="less">
.card-body {
    overflow-y: scroll; height: 470px;
}
.list {
    li {
        border-bottom: solid 1px #eee;
    }
    li.on {
        background-color: #f6f6f6;
    }
    li:last-child {
        border-bottom: 0;
    }
    a {
        display: block;
        padding: 15px 20px 10px 17px;
        border-left: solid 3px transparent;
    }
    a:hover {
       color: #333;
        background-color: #f1f1f1;
        border-left: solid 3px #33cde5;
    }
    .item {
        overflow: hidden;
        padding: 3px 0;
        .icon {
            padding: 5px;
            border-radius: 50%;
            height: 20px;
            width: 20px;
            background-color: #fff;
            box-shadow: 0 0 5px rgba(0,0,0,0.1);
            margin: 0 5px;
            img {
                width:18px; height: 18px;
            }
        }
    }
    .ids {
        float: left;
    }
    .time {
        float: right;
        color: #666666;
    }
    .oh {
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        display: inline-block;
        vertical-align: middle;
    }
}
@media only screen and (max-width:690px){
    .card-body {
        overflow-y: auto; height: auto;
    }
    .list {font-size: 12px;}
}

.scrollbar::-webkit-scrollbar{    
    width: 6px;    
    height: 6px;    
    background-color: #f5f5f5;
}
.scrollbar::-webkit-scrollbar-track{    
    background-color: #f5f5f5;}
.scrollbar::-webkit-scrollbar-thumb{    
    height: 10px;    
    background-color: #555;
}
</style>
