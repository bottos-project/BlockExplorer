import request from './request'

//账户列表
export function queryCustList(data) {
  return request({
    url: '/cust/queryCustListAuto',
    method: 'post',
    data
  })
}

//实时账户总数
export function queryCustSummary(data) {
  return request({
    url: '/cust/queryCustSummaryAuto',
    method: 'post',
    data
  })
}

//账户详情
export function queryCustDetail(data) {
  return request({
    url: '/cust/queryCustDetailAuto',
    method: 'post',
    data
  })
}


//地址交易信息
export function queryTradeList(data) {
  return request({
    url: 'trade/queryTradeListAuto',
    method: 'post',
    data
  })
}

//地址转账信息
export function queryTransferlList(data) {
  return request({
    url: 'transferl/queryPersonalTransferlListAuto',
    method: 'post',
    data
  })
}



