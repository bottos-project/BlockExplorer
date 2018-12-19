import request from './request'

//区块列表
export function queryBlockList(data) {
  return request({
    url: 'block/queryBlockListAuto',
    method: 'post',
    data
  })
}

//区块详情
export function queryBlockDetail(data) {
  return request({
    url: 'block/queryBlockDetailAuto',
    method: 'post',
    data
  })
}

//区块交易记录
export function queryTradeList(data) {
  return request({
    url: 'trade/queryTradeListAuto',
    method: 'post',
    data
  })
}

//区块转账记录
export function queryTransferlList(data) {
  return request({
    url: 'transferl/queryTransferlListAuto',
    method: 'post',
    data
  })
}

