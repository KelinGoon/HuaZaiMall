import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/product/info/' + data.pagesize + '/' + data.pagenum,
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

export function updateList(data) {
  return request({
    url: '/product/update',
    method: 'post',
    data
  })
}

export function deleteList(data) {
  return request({
    url: '/product/delete',
    method: 'post',
    data
  })
}
