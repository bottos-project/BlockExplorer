import request from './request'
//搜索
export function queryFuzzyInfo(data) {
  return request({
    url: '/common/queryFuzzyInfoAuto',
    method: 'post',
    data
  })
}
//主要信息统计
export function PrimarySummary(data) {
  return request({
    url: '/comm/queryMajSummaryAuto',
    method: 'post',
    data
  })
}


//过去14天交易数
export function queryLastFTTradeSummary(data) {
  return request({
    url: 'trade/queryLastFTTradeSummaryAuto',
    method: 'post',
    data
  })
}
//账户增长
export function queryRigiterSummary(data) {
  return request({
    url: 'cust/queryRigiterSummaryAuto',
    method: 'post',
    data
  })
}


//区块列表
export function queryBlockList(data) {
  return request({
    url: 'block/queryBlockListAuto',
    method: 'post',
    data
  })
}
//转账列表
export function queryTransferlList(data) {
  return request({
    url: 'transferl/queryTransferlListAuto',
    method: 'post',
    data
  })
}