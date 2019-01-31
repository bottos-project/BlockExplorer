<template>
  <div class="main">
    <div class="border">
        <div class="title"><h3><span>{{$t('transactionsDetail.title')}}</span></h3></div>
        <div class="detail">
            <ul>
                <li>
                    <div class="tit">{{$t('transactionsDetail.state')}}</div>
                    <div class="con">
                        <span class="confirm" v-show="Detail.tradeStatus === '已确认'">{{$t('blocksDetail.state_yes')}}</span>
                        <span class="Unconfirm" v-show="Detail.tradeStatus === '未确认'">{{$t('blocksDetail.state_no')}}</span>
                    </div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('transactionsDetail.transactionsID')}}</div>
                    <div class="con">{{Detail.transactionId}}</div>
                </li>
                <li>
                    <div class="tit">{{$t('transactionsDetail.block')}}</div>
                    <div class="con">
                        <router-link :to="'/blockchain/blocksDetail/'+Detail.blockNum">
                            {{Detail.blockNum}}
                        </router-link>
                    </div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('transactionsDetail.time')}}</div>
                    <div class="con">
                        {{Detail.tradeDate | dateformat}} 
                        （{{Detail.tradeDate | dateformatNow}}）
                    </div>
                </li>
                <li>
                    <div class="tit">{{$t('transactionsDetail.time_resource')}}</div>
                    <div class="con">
                        {{Detail.timeTokenCost}} µs
                    </div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('transactionsDetail.space_resource')}}</div>
                    <div class="con">
                        {{Detail.spaceTokenCost}} B
                    </div>
                </li>
            </ul>
        </div>
    </div>
    <div class="border">
        <div class="title"><h3><span>{{$t('transactionsDetail.title2')}}</span></h3></div>
        <div class="detail">
            <ul>
                <li>
                    <div class="tit">{{$t('transactionsDetail.name')}}</div>
                    <div class="con">
                        {{Detail.contract}}
                    </div>
                </li>
                <li class="row">
                    <div class="tit">{{$t('transactionsDetail.method_name')}}</div>
                    <div class="con">
                        {{Detail.method}}
                    </div>
                </li>
                <li>
                    <div class="tit">{{$t('transactionsDetail.input')}}</div>
                    <div class="con">
                      <pre v-html="inputdata">
                        {{inputdata}}
                      </pre>
                    </div>
                </li>
            </ul>
        </div>
    </div>
  </div>
</template>
<script>
import {queryTradeDetail} from '@/api/blockchain_transactions'

    export default {
      data() {
        return {
          custId:this.$route.params && this.$route.params.id,
          Detail:{},
          inputdata:null
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
            transactionId:this.custId
          }
          queryTradeDetail(para).then(response => {
            //   console.log({response})
            this.Detail = response.data
            console.log({data:response.data})
            let param = response.data.param
            if (param && param.value ){
                param.value = (param.value / Math.pow(10,8)) + 'BTO'
            }
            this.inputdata = this.syntaxHighlight(response.data.param)
          }).catch(error => {
            this.$message({
                message: this.$i18n.t('tips.error'),
                type: 'error'
            });
          })
        },
        syntaxHighlight(json) {
          if (typeof json != 'string') {
              json = JSON.stringify(json, undefined, 2);
          }
          console.log(json)
          json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
          return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function(match) {
              var cls = 'number';
              if (/^"/.test(match)) {
                  if (/:$/.test(match)) {
                      cls = 'key';
                  } else {
                      cls = 'string';
                  }
              } else if (/true|false/.test(match)) {
                  cls = 'boolean';
              } else if (/null/.test(match)) {
                  cls = 'null';
              }
              return '<span class="' + cls + '">' + match + '</span>';
          });
        }
      }
    }
</script>