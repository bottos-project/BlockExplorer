<template>
    <div class="card-body scrollbar">
        <ul class="list">
            <li v-for="(item,index) in tableList" :class="{on:index%2==0,off:index%2!=0}">
                <router-link :to="'/blockchain/blocksDetail/'+item.blockNum">
                <div class="item">
                    <span class="block">{{$t('home.card3_block')}} {{item.blockNum}}</span>
                    <span class="icon"><img src="../../assets/icon_transfers.png" width="20" height="20" /></span>
                    <span>{{item.transCount}} {{$t('home.card3_transactions')}}</span>
                    <div class="time">{{ item.tradeDate | dateformat}}</div>
                </div>
                <div class="item">
                    <div class="fl oh">{{$t('home.card3_delegate')}}：{{item.delegate}}</div>
                    <!-- <div class="fr">{{$t('home.card3_reward')}}: 0.000584 DTO</div> -->
                </div>
                </router-link>
            </li>
            <div class="noData" v-if="tableList.length === 0">{{$t('tips.nodata')}}</div>
        </ul>
    </div>
</template>
<script>
import { queryBlockList } from '@/api/home'

export default {
    data(){
      return {
        tableList:{}
      }
    },
    created(){
      this.fetchData()
    }, 
    methods: {
      fetchData() {
        let para = {
          start:0,
          length:20
        }
        queryBlockList(para).then(response => {
          this.tableList = window.sortData(response.data)
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
        .block { 
            color: #5bc1e8; 
            padding: 5px 7px;
            border:dotted 1px #5bc1e8;
            border-radius: 5px;
        }
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
    .time {
        float: right;
        color: #666666;
    }
    .oh {
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        width: 50%;
    }
}
@media only screen and (max-width:690px){
    .card-body {
        overflow-y: auto; height: auto;
    }
	.list .oh { width: 35%;}
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
