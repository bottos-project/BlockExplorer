<template>
  <div class="main">
    <div class="border">
      <div class="title">
        <h3><span>{{$t('accounts.title')}}</span>
          <small>{{$t('accounts.title_statistics')}} <i>{{CustSummary}}</i></small>
        </h3>
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
        <el-table-column prop="accountName" :label="$t('accounts.table1')"  :show-overflow-tooltip="true" min-width="150">
          <template slot-scope="scope">
            <router-link :to="'/blockchain/accountsDetail/'+scope.row.accountName">
              {{scope.row.accountName}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column prop="availableAmount" :label="$t('accounts.table2')" min-width="180">
          <template slot-scope="scope">
            {{scope.row.availableAmount / Math.pow(10,8)}} {{scope.row.currency}} 
          </template>
        </el-table-column>
        <el-table-column prop="holdRate" :label="$t('accounts.table3')" min-width="150">
          <template slot-scope="scope">
            {{parseFloat(scope.row.availableAmount / Math.pow(10,17) * 100).toFixed(8)}} %
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
import {queryCustList,queryCustSummary} from '@/api/blockchain_accounts'
    export default {
      data() {
        return {
          CustSummary:'',
          tableData: [],
          totalAccounts:null,
          listLoading: true,
          listQuery: {
            iTotalDisplayRecords:null,
            start: 0,
            length: 20
          },
          currency:'BTO',
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
          let para = {
            currency:this.currency,
            start:this.listQuery.start,
            length:this.listQuery.length
          };
          queryCustSummary(para).then(response => {
            console.log({responsesdfssdfds:response})
            this.CustSummary = response.rtCustCount
          }).catch(function(error){
            this.$message({
              message: this.$i18n.t('tips.error'),
              type: 'error'
            });
          })
          queryCustList(para).then(response => {
            console.log({response})
            this.tableData = window.sortData(response.data)
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
