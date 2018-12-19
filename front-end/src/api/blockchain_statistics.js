import request from './request'

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