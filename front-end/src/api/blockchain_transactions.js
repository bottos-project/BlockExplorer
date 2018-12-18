import request from './request'

//区块列表
export function queryTradeList(data) {
  return request({
    url: 'trade/queryTradeListAuto',
    method: 'post',
    data
  })
}
//区块详情
export function queryTradeDetail(data) {
  return request({
    url: 'trade/queryTradeDetailAuto',
    method: 'post',
    data
  })
}