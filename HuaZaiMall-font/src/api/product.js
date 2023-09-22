import request from '@/utils/request'

export function getList() {
  return request({
    url: '/product/info',
    method: 'get'
  })
}

export function addList(data) {
  return request({
    url: '/product/add',
    method: 'post',
    data
  })
}
