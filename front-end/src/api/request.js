import axios from 'axios';

// 创建axios实例
const service = axios.create({
    baseURL: process.env.BASE_API, // api的base_url
    timeout: 15000 // 请求超时时间
})
// request拦截器
service.interceptors.request.use(config => {
    //if (store.getters.token) {
      // config.headers['X-Token'] = "23423sd223dsdsd223sdf2323"
      // config.headers['Authorization'] = "Basic dXNlcjpzZWNyZXQ="
      // config.headers['project'] = "slb"
      // config.headers['appSource'] = "web"
    //}
    return config
  }, error => {
    console.log("request拦截器报错："+error)
    Promise.reject(error)
  })

// respone拦截器
service.interceptors.response.use(
    response => {
      const res = response.data
      if (res.success !== true) {
        Message({
          message: response.message,
          type: 'error',
          duration: 5 * 1000
        })
        return Promise.reject('error')
      } else {
        return response.data.data
      }
    },
    error => {
      console.log('err' + error)// for debug
      Message({
        message: error.message,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(error)
    }
  )
  
  export default service








