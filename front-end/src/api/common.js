import request from './request'

//查询币种
export function querParamListAuto(data) {
    return request({
      url: 'param/querParamListAuto',
      method: 'post',
      data
    })
  }