<template>
  <div>
      <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
        <el-table-column prop="" width="40">
        </el-table-column>
        <el-table-column prop="accountName" :label="$t('supernodeDetail.table5')" min-width="120">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/accountsDetail/'+scope.row.accountName">
              {{scope.row.param.voter}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="blockHeight" :label="$t('supernodeDetail.table1')" min-width="120">
          <template slot-scope="scope">
              {{scope.row.block_number}}
          </template>
        </el-table-column>
        <el-table-column prop="voteTime" :label="$t('supernodeDetail.table9')" min-width="120">
          <template slot-scope="scope">
              {{scope.row.timestamp | dateformatNow}}
          </template>
        </el-table-column>
        <!-- <el-table-column prop="voteCount" :label="$t('supernodeDetail.table6')" min-width="120">
        </el-table-column>
        <el-table-column prop="currRate" :label="$t('supernodeDetail.table7')" min-width="120">
        </el-table-column>
        <el-table-column prop="totalRate" :label="$t('supernodeDetail.table8')" min-width="120">
        </el-table-column> -->
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
import {queryVoteList} from '@/api/supernode'

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
    created(){
      this.fetchData()
    },        
    methods: {
      fetchData() {
        let para = {
          nodeName:this.custId,
          start:this.listQuery.start,
          length:this.listQuery.length
        };
        queryVoteList(para).then(response => {
          console.log({queryVoteList:response})
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
<style scoped lang="less">

</style>
