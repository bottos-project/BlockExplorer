import request from './request'

//超级节点统计
export function queryNodeSummary(data) {
  return request({
    url: '/superNode/queryNodeSummaryAuto',
    method: 'post',
    data
  })
}

//节点列表
export function queryNodeList(data) {
  return request({
    url: '/superNode/queryNodeListAuto',
    method: 'post',
    data
  })
}

//节点详情
export function queryNodeDetail(data) {
  return request({
    url: '/superNode/queryNodeDetailAuto',
    method: 'post',
    data
  })
}


//出块列表
export function queryBlockList(data) {
  return request({
    url: 'block/queryBlockListAuto',
    method: 'post',
    data
  })
}

//投票记录列表
export function queryVoterTransaction(data) {
  return request({
    url: 'superNode/queryVoteListAuto',
    method: 'post',
    data
  })
}


export function queryVoterList(data){
  return request({
    url: 'superNode/queryVoterList',
    method: 'post',
    data
  })
}


