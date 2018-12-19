<template>
  <div class="border">
      <div id="maps"></div>
  </div>
</template>
<script>
import '../../../node_modules/echarts/map/js/world.js'

import {queryNodeDistribute} from '@/api/nodes'


export default {
    data(){
      return {
        mapsData:{}
      }
    },
    created(){
      this.fetchData()
    },
    methods: {
      fetchData() {
        let para = {}
        queryNodeDistribute(para).then(response => {
          this.mapsData = response.data.distribute.map(function (item) {
            return [item.city,item.latitude,item.longitude,item.count]
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
        var myChartMap = this.$echarts.init(document.getElementById('maps'));

        var rawData = this.mapsData

        function makeMapData(rawData) {
            var mapData = [];
            for (var i = 0; i < rawData.length; i++) {
                mapData.push({
                    name: rawData[i][0],
                    value: [rawData[i][1],rawData[i][2],rawData[i][3]],
                })
            }
            return mapData;
        };

        myChartMap.setOption({
          geo: {
            map: 'world',
              silent: true,
              label: {
                  emphasis: {
                      show: false,
                      areaColor: '#eee'
                  }
              },
              itemStyle: {
                  normal: {
                      borderWidth: 0.2,
                      borderColor: '#404a59'
                  }
              },
              left: '1%',
              top: '1%',
              bottom: '1%',
              right: '1%',
              roam: true 
          },
          // hover显示目标数据
          tooltip : {
            trigger: 'item',
            formatter: function (params) {
                var value = params.value[2];
                return params.name + ' : ' + value;
                //return params.seriesName + '<br/>' + params.name + ' : ' + value;
            }
          },
          // 图表背景色
          backgroundColor: '#404a59',  
          // 标志颜色
          color:'red',
          // 新建散点图series
          series:[{
            name: 'Count',
            type:'scatter',//为散点类型
            coordinateSystem: 'geo',// series坐标系类型
            data:makeMapData(rawData),
            activeOpacity: 1,            
            symbolSize: 14,
            itemStyle: {
                normal: {
                    borderColor: '#fff',
                    color: '#5bc1e8',
                }
            }
          }],
 
         
        })
      },
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
#maps { 
  width: 100%;
  height: 550px;
}
@media only screen and (max-width:1000px){
	#maps { 
    height: 450px;
  }
}
@media only screen and (max-width:750px){
	#maps { 
    height: 400px;
  }
}
@media only screen and (max-width:650px){
	#maps { 
    height: 350px;
  }
}
@media only screen and (max-width:550px){
	#maps { 
    height: 250px;
  }
}
@media only screen and (max-width:450px){
	#maps { 
    height: 200px;
  }
}
</style>
