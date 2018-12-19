<template>
  <div>
      <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
        <el-table-column prop="" width="40">
        </el-table-column>
        <el-table-column :label="$t('contractsDetail.table1')" min-width="180" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/transactionsDetail/'+scope.row.transactionId">
              {{scope.row.table1}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column :label="$t('contractsDetail.table2')" min-width="90" >
          <template slot-scope="scope">
            <router-link :to="'/blockchain/blocksDetail/'+scope.row.blockNum">
              {{scope.row.table2}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column :label="$t('contractsDetail.table3')" min-width="120" >
          <template slot-scope="scope">
            {{scope.row.table3}}
          </template>
        </el-table-column>
        <el-table-column :label="$t('contractsDetail.table4')" min-width="110" >
          <template slot-scope="scope">
            {{scope.row.table4}}
          </template>
        </el-table-column>
        <el-table-column :label="$t('contractsDetail.table5')" min-width="290" >
          <template slot-scope="scope">
            {{scope.row.table5}}
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
</template>
<script>
//import {queryTradeList} from '@/api/blockchain_transactions'

export default {
    data(){
      return {
        custId:this.$route.params && this.$route.params.id,
        tableData: [
          {table1:"0abeee7d4db7e41976ff8a1c61fdc4e00ef959e4769f07935f927f0af5e9151e",table2:"56999",table3:"bottostasd216",table4:"1天2小时前",table5:{"memo":"frompowertesting","from":"bottostqwe216","to":"bottostasd216","value":"800000000"}}
        ],
        listLoading: false,
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
      //this.fetchData()
    },        
    methods: {
      fetchData() {
        let para = {
          accountName:this.custId,
          start:this.listQuery.start,
          length:this.listQuery.length
        };
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
<style scoped lang="less">

</style>
