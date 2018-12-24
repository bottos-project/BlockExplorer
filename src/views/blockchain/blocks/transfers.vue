<template>
  <div>
    <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
      <el-table-column prop="" width="40">
        </el-table-column>
      <el-table-column prop="transactionId" :label="$t('transfers.table1')"  :show-overflow-tooltip="true" min-width="200">
        <template slot-scope="scope">
          <router-link :to="'/blockchain/transfersDetail/'+scope.row.transactionId">
            {{scope.row.transactionId}}
          </router-link>
        </template>
      </el-table-column>
      <el-table-column :label="$t('transfers.table4')" min-width="100">
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
      <el-table-column prop="tradeAmount" :label="$t('transfers.table5')" width="160">
        <template slot-scope="scope">
          {{scope.row.tradeAmount / Math.pow(10,8)}}
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
import {queryTransferlList} from '@/api/blockchain_blocks'
    export default {
      data() {
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
      methods:{
        fetchData() {
          let para = {
            blockNum:this.custId,
            start:this.listQuery.start,
            length:this.listQuery.length
          };
          queryTransferlList(para).then(response => {
            console.log({queryTransferlList:response})
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
.vh{
      overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        max-width: 42%;
        display: inline-block;
        vertical-align: middle;
    }
</style>
