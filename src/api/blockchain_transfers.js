import request from './request'

//转账列表
export function queryTransferlList(data) {
  return request({
    url: 'transferl/queryTransferlListAuto',
    method: 'post',
    data
  })
}

//转账详情
export function queryTradeDetail(data) {
  return request({
    url: 'trade/queryTradeDetailAuto',
    method: 'post',
    data
  })
}

