<template>
<div class="inner cf">
  <div class="main">
    <div class="border">
      <div class="title">
        <h3><span>{{$t('contracts.title')}}</span></h3>
        <div class="input-group">
          <input type="text" v-model="searchVal" class="input" clearable :placeholder="$t('contracts.searchs')" value="">
          <a class="filterBtn" @click="fetchData">{{$t('contracts.searchsbtn')}}</a>
        </div>
      </div>
      <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
        <el-table-column prop="" width="40">
        </el-table-column>
        <el-table-column :label="$t('contracts.table1')" min-width="100" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            {{scope.row.table1}}
          </template>
        </el-table-column>
        <el-table-column :label="$t('contracts.table2')" min-width="320">
          <template slot-scope="scope">
            <router-link :to="'/contracts/detail/'+scope.row.table2">
              {{scope.row.table2}}
            </router-link>
          </template>
        </el-table-column>
        <el-table-column :label="$t('contracts.table3')" min-width="100">
          <template slot-scope="scope">
            {{scope.row.table3}}
          </template>
        </el-table-column>
        <el-table-column :label="$t('contracts.table4')" min-width="100">
          <template slot-scope="scope">
            {{scope.row.table4}}
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
</div>    
</template>
<script>
//import {queryTradeList} from '@/api/blockchain_transactions'
    export default {
      data() {
        return {
          tableData: [
            {table1:'eosfuneos111',table2:'TVkNuE1BYxECWq85d8UR9zsv6WppBns9iH',table3:'28,701,093',table4:'28,701,093'}
          ],
          listLoading: false,
          listQuery: {
            iTotalDisplayRecords:null,
            start: 0,
            length: 20
          },
          searchVal:''
        }
      },
      created(){
        //this.fetchData()
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
.input-group {
    display: flex;
    position: absolute;
    width: 320px;
    right: 30px;
    top: 24px;
    .input {
      line-height: 1.5;
      border-radius: 40px;
      background-color: #fff;
      border: solid 1px #a3a3a3;
      color: #333;
      font-size: 14px;
      padding: 10px 15px;
      width: 100%;
    }
    .filterBtn {
      position: absolute;
      top: 6px;
      right: 10px;
      height: 30px;
      padding: 0 10px;
      line-height: 30px;
      text-align: center;
    }
  }
@media only screen and (max-width:690px){
    .title h3 {
      margin-bottom: 10px;
    }
    .input-group {
      margin: 0 auto;
      position: relative;    
      top: 0; right: 0;  
    }
	}
</style>
