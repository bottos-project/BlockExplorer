<template>
  <div class="main">
    <div class="border">
      <div class="detail">
        <div class="information">
          <el-row :gutter="100">
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
              <div class="box bg">
                <dl>
                  <dd>{{Detail.name}}</dd>
                  <dd><a target="_blank" :href="Detail.nodeUrl">{{Detail.nodeUrl}}</a></dd>
                  <dd><span class="icon_position">{{Detail.city}},{{Detail.country}}</span></dd>
                </dl>
              </div>
            </el-col>
          </el-row>
          <div class="des">
            <div :class="{showall:true,active:showall}">
              {{Detail.nodeDesc}}
            </div>
            <a class="open" @click="openDes" v-if="showall == false">{{$t('supernodeDetail.expand')}}</a>
          </div>
        </div>
        <ul>
            <li class="row">
                <div class="tit">{{$t('supernodeDetail.account')}}</div>
                <div class="con">
                  <router-link :to="'/blockchain/accountsDetail/'+Detail.accountName">
                    {{Detail.accountName}}
                  </router-link>
                </div>
            </li>
            <li>
                <div class="tit">{{$t('supernodeDetail.serviceApi')}}</div>
                <div class="con">
                  <a target="_blank" :href="Detail.serviceApi">{{Detail.serviceApi}}</a>
                </div>
            </li>
        </ul>
        <el-row :gutter="40" class="computer">
          <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
            <div class="box">
              <div class="tit">内存</div>
              <div class="circ">
                <progress-circle :list='20' :jindu='20'></progress-circle>
              </div>
              <div class="info">5.8 kB / 8.1 kB</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
            <div class="box">
              <div class="tit">计算资源</div>
              <div class="circ">
                <progress-circle :list='60' :jindu='60'></progress-circle>
              </div>
              <div class="info">1 ms 464 µs / 18 ms 413 µs</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
            <div class="box">
              <div class="tit">网络资源</div>
              <div class="circ">
                <progress-circle :list='49' :jindu='49'></progress-circle>
              </div>
              <div class="info">105 B / 388.1 kB</div>
            </div>
          </el-col>
        </el-row>
        
      </div>
    </div>
  </div>
</template>
<script>
    import progressCircle from '@/components/progeress-circle'
    import { queryNodeDetail } from '@/api/supernode'
    export default {
      components:{
        progressCircle
      },
      data() {
        return {
          custId:this.$route.params && this.$route.params.id,
          Detail:{
            name:'EOS-Huobipool',
            nodeUrl:'https://www.genesis-mining.com/eos',
            city:'Beijing',
            country:'China',
            nodeDesc:'EOS-Huobipool旨在为EOS社区的每位成员提供一站式金融科技服务和解决方案，并在区块链领域构建紧密的合作伙伴关系。EOS-Huobipool隶属于火币集团，火币品牌在全球区块链领域拥有优异的知名度和较好的口碑。EOS-Huobipool旨在为EOS社区的每位成员提供一站式金融科技服务和解决方案，并在区块链领域构建紧密的合作伙伴关系。EOS-Huobipool隶属于火币集团，火币品牌在全球区块链领域拥有优异的知名度和较好的口碑。EOS-Huobipool旨在为EOS社区的每位成员提供一站式金融科技服务和解决方案，并在区块链领域构建紧密的合作伙伴关系。EOS-Huobipool隶属于火币集团，火币品牌在全球区块链领域拥有优异的知名度和较好的口碑。EOS-Huobipool旨在为EOS社区的每位成员提供一站式金融科技服务和解决方案，并在区块链领域构建紧密的合作伙伴关系。EOS-Huobipool隶属于火币集团，火币品牌在全球区块链领域拥有优异的知名度和较好的口碑。',

            accountName:'TMdDoWAAMdMooANGn97KA1Jp5gN1BmfxJv',
            serviceApi:'http://node869-mainnet.eosauthority.com:9393',
          },
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
.computer {
  background-color: #f6f6f6;
  padding: 30px 30px;
  .box {
    text-align: center;
    .tit {
      margin-bottom: 10px;
      font-size: 16px; font-weight: bold;
    }
    .circ {
      margin-bottom: 10px;
    }
    .info {
      margin-bottom: 10px;
    }
  }
}
</style>