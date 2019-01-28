<template>
  <div class="card-body">
      <div id="myChart2"></div>
  </div>
</template>
<script>
import {queryRigiterSummary} from '@/api/blockchain_statistics'
import moment from 'moment';
export default {
    data(){
      return {
        listQuery:{
          startDate:moment().subtract(90, 'days').format('YYYY-MM-DD'),
          endDate:moment().format('YYYY-MM-DD')
        },
        chartColumn: null,
        chartData:{
          regDate:[],
          totalRegCount:[]
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
        queryRigiterSummary(para).then(response => {
          this.chartData.regDate = response.map(function (item) {
            var date = new Date(item.regDate)
            return date.toLocaleDateString()
          })
          this.chartData.totalRegCount = response.map(function (item) {
            return item.totalRegCount
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
      this.chartColumn = this.$echarts.init(document.getElementById("myChart2"));
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
            data: this.chartData.regDate,
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
                name:this.$i18n.t('home.card2_series_name'),
                type:'line',
                data:this.chartData.totalRegCount
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
#myChart2 { 
  width: 100%;
  height: 350px;
}
</style>
