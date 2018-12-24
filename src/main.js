import 'babel-polyfill'//解决IE不支持“Promise”未定义"问题
import Vue from 'vue'
import App from './App'
import router from './router'

Vue.config.productionTip = false

//全局消息组件
import VueMessage from './components/message/index.js';
Vue.use(VueMessage);

//多语言
import VueI18n from 'vue-i18n'
Vue.use(VueI18n)
import messages from './language/index'

// 将选择的语言存在localStorage中
function lang () {
  let t = window.localStorage.getItem('language')
  if (t) return t
  else return 'en'
}
const language = lang()
// 自定义 window 的 lang 属性
window.lang = lang()
// 创建一个 i18n 实例
const i18n = new VueI18n({
  locale: language,    // 语言标识
  messages
})


import locale from 'element-ui/lib/locale';
import {Row,Col,Table,TableColumn,pagination,Loading,Tabs,TabPane} from 'element-ui';
Vue.use(Row);
Vue.use(Col);
Vue.use(Table);
Vue.use(TableColumn);
Vue.use(pagination);
Vue.use(Loading);
Vue.use(Tabs);
Vue.use(TabPane);
locale.i18n((key, value) => i18n.t(key, value))




import '@/styles/index.less' //全局CSS


// 加载echarts基础包
import echarts from 'echarts/lib/echarts'
//按需加载
import 'echarts/lib/chart/line'
import 'echarts/lib/component/tooltip'
import 'echarts/lib/component/dataZoom'

Vue.prototype.$echarts = echarts




import moment from 'moment'
require('moment/locale/zh-cn');

//日期语言
moment.locale(lang());
Vue.filter('dateformat', function(dataStr, pattern = 'YYYY-MM-DD HH:mm:ss') {
  //return moment(dataStr).format(pattern)
  return moment(dataStr * 1000).startOf('second').fromNow();
})
Vue.filter('dateformatNow', function(dataStr, pattern = 'YYYY/MM/DD HH:mm:ss') {
  return moment(dataStr * 1000).format(pattern)
  //return moment(dataStr).startOf('hour').fromNow();
})


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  i18n,
  template: '<App/>',
  components: { App }
})
