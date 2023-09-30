import request from '@/utils/request'

export function getList() {
  return request({
    url: '/product/info',
    method: 'get'
  })
}

// /' + data.pagesize + '/' + data.pagenum
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
