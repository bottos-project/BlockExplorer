import request from './request'

//地图
export function queryNodeDistribute(data) {
  return request({
    url: '/node/queryNodeDistributeAuto',
    method: 'post',
    data
  })
}
//节点排名
export function queryNodeRank(data) {
  return request({
    url: '/node/queryNodeRankAuto',
    method: 'post',
    data
  })
}

//节点统计
export function queryTopNodeSummary(data) {
  return request({
    url: '/node/queryTopNodeSummaryAuto',
    method: 'post',
    data
  })
}