<template>
  <div>
      <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
        <el-table-column prop="" width="40">
        </el-table-column>
        <el-table-column prop="transactionId" :label="$t('transactions.table1')" min-width="200" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/transactionsDetail/'+scope.row.transactionId">
              {{scope.row.transactionId}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="contract" :label="$t('transactions.table4')">
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
</template>
<script>
import {queryTradeList} from '@/api/blockchain_blocks'

export default {
    data(){
      return {
        custId:this.$route.params && this.$route.params.id,
        tableData: [],
        listLoading: true,
        listQuery: {
          iTotalDisplayRecords:null,
          start: 0,
          length: 20
        },
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
    methods: {
      fetchData() {
        let para = {
          blockNum:this.custId,
          start:this.listQuery.start,
          length:this.listQuery.length
        };
        queryTradeList(para).then(response => {
          console.log({transactionsresponse:response})
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
<style scoped lang="less">

</style>
