<template>

    <div class="main">
      <div class="border">
        <div class="Stats">
          <el-row>
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
              <div class="item">
                <div class="con">{{statistics.nodeCount}}</div>
                <div class="tit">{{$t('nodes.stats1')}}</div>
              </div>
            </el-col>
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
              <div class="item">
                <div class="con">{{statistics.country}}</div>
                <div class="tit">{{$t('nodes.stats2')}}</div>
              </div>
            </el-col>
          </el-row>
        </div>
      </div>
      

    <!-- <div class="maps">
      <Maps></Maps>
    </div> -->


      <div class="border">
        <div class="title">
          <h3><span>{{$t('nodes.title')}}</span></h3>
          <div class="input-group">
            <input type="text" v-model="searchVal" class="input" clearable :placeholder="$t('nodes.searchs')" value="">
            <a class="filterBtn" @click="fetchData">{{$t('nodes.searchsbtn')}}</a>
          </div>
        </div>
        <el-table :data="tableData" size="medium" v-loading="listLoading" :row-class-name="tableRowClassName" class="tables">
          <el-table-column prop="" width="40">
          </el-table-column>
          <el-table-column :label="$t('nodes.table1')" min-width="250" :show-overflow-tooltip="true">
            <template slot-scope="scope">
              {{scope.row.URL}}
              <!-- <router-link :to="'/nodes/detail/'+scope.row.table1">
              {{scope.row.table1}}
              </router-link>
              <br />
              <a :href="scope.row.url" target="_blank">
                {{scope.row.url}}
              </a> -->
            </template>
          </el-table-column>
         
          <el-table-column :label="$t('nodes.table2')" min-width="150">
            <template slot-scope="scope">
              {{scope.row.node_type}}
            </template>
          </el-table-column>
          <el-table-column :label="$t('nodes.table3')" min-width="150">
            <template slot-scope="scope">
              {{scope.row.node_country}}
            </template>
          </el-table-column>
          <!-- <el-table-column :label="$t('nodes.table4')" min-width="100">
            <template slot-scope="scope">
              {{scope.row.table4}}
            </template>
          </el-table-column>
          <el-table-column :label="$t('nodes.table5')" min-width="100">
            <template slot-scope="scope">
              {{scope.row.table5}}
            </template>
          </el-table-column>
           <el-table-column :label="$t('nodes.table6')" min-width="250">
            <template slot-scope="scope">
              <a :href="scope.row.table6" target="_blank">
                {{scope.row.table6}}
              </a>
            </template>
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

    </div>

    

</template>
<script>
  import Maps from './maps' //地图
  import {queryTopNodeSummary,queryNodeRank} from '@/api/nodes'
    export default {
      components:{
        Maps
      },
      data() {
        return {
          statistics:'',
          tableData: [
            {table1:"Bitfinex",url:"https://www.genesis-mining.com/eos",table2:"存储服务",table3:"United States",table4:"10,344",table5:"5,88",table6:"https://mainnet.eosauthority.com:9393",}
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
        this.statisticsData()
        this.fetchData()
      },
      methods:{
        statisticsData(){
          let para = {}
          queryTopNodeSummary(para).then(response => {
            this.statistics = response.data
          }).catch(function(error){
            this.$message({
              message: this.$i18n.t('tips.error'),
              type: 'error'
            });
          })
        },
        fetchData() {
          let para = {}
          queryNodeRank(para).then(response => {
            console.log({response})
            this.tableData = response
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
  

.maps {
  margin-bottom: 20px;
}

</style>
