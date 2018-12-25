<template>
  <div class="main">
    <div class="border">
      <div class="title">
        <h3><span>{{$t('blocks.title')}}</span></h3>
      </div>
      <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
        <el-table-column prop="" width="40">
        </el-table-column>
        <el-table-column prop="height" :label="$t('blocks.table1')" min-width="110">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/blocksDetail/'+scope.row.blockNum">
              {{scope.row.blockNum}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="tradeDate" :label="$t('blocks.table2')" min-width="150">
          <template slot-scope="scope">
            {{scope.row.tradeDate | dateformat}}
          </template>
        </el-table-column>
        <el-table-column prop="transCount" :label="$t('blocks.table3')" min-width="100">
        </el-table-column>
        <el-table-column prop="delegate" :label="$t('blocks.table4')" min-width="140">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/accountsDetail/'+scope.row.delegate">
              {{scope.row.delegate}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="" :label="$t('blocks.table6')" min-width="160">
          <template slot-scope="scope">
            2.28310501 BTO
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
import {queryBlockList} from '@/api/blockchain_blocks'
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
          };
          queryBlockList(para).then(response => {
            this.tableData = response.data
            this.listQuery.iTotalDisplayRecords = response.iTotalDisplayRecords
            this.listLoading = false
          }).catch(function(error){
            this.$message({
              message: this.$i18n.t('tips.error'),
              type: 'error'
            });
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
