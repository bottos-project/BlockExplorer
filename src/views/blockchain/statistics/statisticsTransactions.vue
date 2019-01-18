<template>
  <div class="card-body">
      <div id="myChart"></div>
  </div>
</template>
<script>
import { queryLastFTTradeSummary } from '@/api/blockchain_statistics'
import moment from 'moment';
export default {
    data(){
      return {
        listQuery:{
          startDate:moment().subtract(1, 'days').format('YYYY-MM-DD'),
          endDate:moment().format('YYYY-MM-DD')
        },
        chartColumn: null,
        chartData:{
          tradeDate:[],
          dailyTransCount:[]
        }
      }
    },
    created(){
      this.fetchData()
    },        
    methods: {
      fetchData() {
        let para = {
          startDate:this.listQuery.startDate,
          endDate:this.listQuery.endDate
        }
        queryLastFTTradeSummary(para).then(response => {
          this.chartData.tradeDate = response.map(function (item) {
            return item.tradeDate
          })
          this.chartData.dailyTransCount = response.map(function (item) {
            return item.dailyTransCount
          })
          this.drawColumnChart()
        }).catch(error => {
          this.$message({
            message: this.$i18n.t('tips.error'),
            type: 'error'
          });
        })
      },
      drawColumnChart() {
      this.chartColumn = this.$echarts.init(document.getElementById("myChart"));
        this.chartColumn.setOption({
          color: ['#3398DB', '#00517a', '#ff085f', '#008489'],
          tooltip: {
              trigger: 'axis',
              axisPointer: {
                  type: 'cross',
                  crossStyle: {
                      color: '#000'
                  }
              }
          },
          grid: {  
              left: '60',  
              right: '20',  
              bottom: '80',  
              top: '20'
          },
          dataZoom: [
              {
                show: true,
                realtime: true,
                start: 60,
                end: 100
              },
              {
                type: 'inside',
                realtime: true,
                start: 65,
                end: 85
              }
          ],
          xAxis: {
            data: this.chartData.tradeDate,
            axisPointer: {
              type: 'shadow'
            }
          },
          yAxis: {
            axisPointer: {
                label: {
                    formatter: function (params) {
                        return params.value.toFixed(0);
                    }
                }
            }
          },
          series: [
              {
                name:this.$i18n.t('home.card1_series_name'),
                type:'line',
                data:this.chartData.dailyTransCount
              }
          ]
        });
      } 
    },
    mounted: function() {
    this.drawColumnChart();
    },
    updated: function() {
      this.drawColumnChart();
    }
}
</script>
<style scoped lang="less">

.card-body {
  padding: 15px;
  min-height: 250px;
}
#myChart { 
  width: 100%;
  height: 350px;
}
</style>
