<template>
  <div class="main">
    <div class="border">
      <div class="title">
        <h3><span>{{$t('transactions.title')}}</span></h3>
      </div>
      <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
        <el-table-column prop="" width="40">
        </el-table-column>
        <el-table-column prop="transactionId" :label="$t('transactions.table1')" min-width="320" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/transactionsDetail/'+scope.row.transactionId">
              {{scope.row.transactionId}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="blockNum" :label="$t('transactions.table2')" min-width="80">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/blocksDetail/'+scope.row.blockNum">
              {{scope.row.blockNum}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="tradeDate" :label="$t('transactions.table3')" min-width="100">
          <template slot-scope="scope">
            {{scope.row.tradeDate | dateformat}}
          </template>
        </el-table-column>
        <el-table-column prop="contract" :label="$t('transactions.table4')" min-width="100">
        </el-table-column>
        <!-- <el-table-column prop="method" :label="$t('transactions.method_name')" min-width="100">
        </el-table-column> -->
        <el-table-column :label="$t('transactions.computing_resource')" min-width="100">
          <template slot-scope="scope">
            1 ms 760 µs
          </template>
        </el-table-column>
        <el-table-column :label="$t('transactions.network_resource')" min-width="100">
          <template slot-scope="scope">
            18 B
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
import {queryTradeList} from '@/api/blockchain_transactions'
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
        }
      },
      created(){
        this.fetchData()
      },
      methods:{
        fetchData() {
          let para = {
            start:this.listQuery.start,
            length:this.listQuery.length
          }
          queryTradeList(para).then(response => {
            this.tableData = response.data.data
            this.listQuery.iTotalDisplayRecords = response.data.iTotalDisplayRecords
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
