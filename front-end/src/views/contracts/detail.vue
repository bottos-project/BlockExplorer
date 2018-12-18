<template>
  <div class="main">
    <div class="border">
        <div class="title"><h3><span>{{$t('contractsDetail.title')}}</span></h3></div>
        <div class="detail">
            <ul>
                <li>
                    <div class="tit">{{$t('contractsDetail.name')}}</div>
                    <div class="con">betdiceadmin</div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('contractsDetail.address')}}</div>
                    <div class="con">TVkNuE1BYxECWq85d8UR9zsv6WppBns9iH</div>
                </li>
                <li>
                    <div class="tit">{{$t('contractsDetail.TransactionNumber')}}</div>
                    <div class="con">77,585,810</div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('contractsDetail.AccountSize')}}</div>
                    <div class="con">17,400</div>
                </li>
            </ul>
        </div>
    </div>
    <div class="border">
      <div class="tabs">
        <el-tabs>
          <el-tab-pane :label="$t('contractsDetail.Record')">
              <keep-alive>
                  <records></records>
              </keep-alive>
          </el-tab-pane>
          <el-tab-pane :label="$t('contractsDetail.code')">
              <keep-alive>
                  <codes></codes>
              </keep-alive>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </div>
</template>
<script>
    import records from './component/record'
    import codes from './component/code'
    //import { queryNodeDetail } from '@/api/supernode'
    export default {
      components: { records,codes },
      data() {
        return {
          custId:this.$route.params && this.$route.params.id,
          Detail:{},
          currencyAmount: '',
          currencyAmountOptions: [],
          showall:false
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
      methods:{
        fetchData() {
          let para = {
            nodeName:this.custId
          }
          queryNodeDetail(para).then(response => {
            this.Detail = response.data
          }).catch(error => {
            this.$message({
                message: this.$i18n.t('tips.error'),
                type: 'error'
            });
          })
        },
        openDes(){
          this.showall = true
        }
      }
    }
</script>
<style lang="less" scoped>
.information {
  padding: 30px 50px;
  overflow: hidden;
  position: relative;
  .horn {
    position: absolute;
    top: 0;
    right: 0;
  }
  .box {
    padding: 10px 20px;
    border-radius: 10px;
    label {
      display: inline-block;
      width: 100px;
    }
  }
  .box.bg {
    background-color: #5bc1e8;
    color: #fff;
    a {
      color: #fff;
    }
  }
  .des {
    padding-top:20px; 
    position: relative;
    line-height: 24px;
    padding-bottom: 24px;
    .open {
      position: absolute;
      bottom:0;
      right: 0; 
      color: #5bc1e8;
    }
  }
  .showall {
    height: 72px;
    overflow: hidden;
    clear: both;
  }
  .showall.active {  //点击了查看更多，就高度就不管了。随实际的大小变化
    height: auto;
  }
}
</style>