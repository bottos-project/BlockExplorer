<template>
  <div class="main">
    <div class="border">
      <div class="title">
        <h3><span>{{$t('transfers.title')}}</span></h3>
        <!-- <div class="selectd">
          <select v-model="currency">
            <option value="BTO">BTO</option>
            <option value="BTC">BTC</option>
            <option value="ETH">ETH</option>
          </select>
        </div> -->
      </div>
      <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
        <el-table-column prop="" width="40">
        </el-table-column>
        <el-table-column prop="transactionId" :label="$t('transfers.table1')"  :show-overflow-tooltip="true" min-width="350">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/transfersDetail/'+scope.row.transactionId">
              {{scope.row.transactionId}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="block" :label="$t('transfers.table2')" min-width="80">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/blocksDetail/'+scope.row.blockNum">
              {{scope.row.blockNum}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="tradeDate" :label="$t('transfers.table3')" min-width="100">
          <template slot-scope="scope">
            {{scope.row.tradeDate | dateformat}}
          </template>
        </el-table-column>
        <el-table-column :label="$t('transfers.table4')" min-width="180">
          <template slot-scope="scope">
            <span class="vh">
              <router-link :to="'/blockchain/accountsDetail/'+scope.row.sender">
              {{scope.row.sender}}
              </router-link>
            </span> -> 
            <span class="vh">
              <router-link :to="'/blockchain/accountsDetail/'+scope.row.receiver">
              {{scope.row.receiver}}
              </router-link>
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="tradeAmount" :label="$t('transfers.table5')" min-width="130">
          <template slot-scope="scope">
            {{scope.row.tradeAmount / Math.pow(10,8)}} {{scope.row.currency}}
          </template>
        </el-table-column>
      </el-table>

      <div v-show="!listLoading" class="pagination-container">
        <el-pagination background
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="1"
            :page-sizes="[20, 50, 100]"
            :page-size="listQuery.length"
            layout="total, sizes, prev, pager, next, jumper"
            :total="listQuery.iTotalDisplayRecords">
        </el-pagination>
      </div>
    </div>
  </div>
    
</template>
<script>
import {queryTransferlList} from '@/api/blockchain_transfers'
import {querParamListAuto} from '@/api/common'
    export default {
      data() {
        return {
          tableData: [],
          listLoading: true,
          listQuery: {
            iTotalDisplayRecords:null,
            start: 0,
            length: 20
          },
          currency:'BTO'
        }
      },
      created(){
        this.fetchData()
      },
      watch: {
        currency: function (newQuestion, oldQuestion) {
          this.fetchData()
        }
      },
      methods:{
        fetchData() {
          let pe = {
            type:"suppCurrency"
          }
          querParamListAuto(pe).then(response => {
            console.log(response.data.paramList)
          }).catch(function(error){
            console.log('加载出错……');
          })

          let para = {
            currency:this.currency,
            start:this.listQuery.start,
            length:this.listQuery.length
          }
          queryTransferlList(para).then(response => {
            this.tableData = response.data
            this.listQuery.iTotalDisplayRecords = response.iTotalDisplayRecords
            this.listLoading = false
          }).catch(function(error){
            console.log('加载出错……');
          })
        },

        //操作分页
        handleSizeChange(val) {
          this.listQuery.length = val;
          this.fetchData();
        },
        //操作分页
        handleCurrentChange(val) {
          console.log(val)
          this.listQuery.start = (val-1) * this.listQuery.length;
          this.fetchData();
        },
        tableRowClassName({row, rowIndex}) {
          if (rowIndex%2 === 0) {
            return 'row';
          }
          return '';
        }
      }
    }
</script>
<style lang="less" scoped>

</style>
