// get user info
import { addList, deleteList, getList, updateList } from '@/api/product'

const actions = {
  // 这里是get商品信息
  getList({ commit }, data) {
    console.log(this.state)
    return new Promise((resolve, reject) => {
      getList(data).then(response => {
        console.log(response)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // 这里是add商品信息
  addList({ commit }, data) {
    console.log(this.state)
    return new Promise((resolve, reject) => {
      addList(data).then(response => {
        console.log(response)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // 这里是修改商品信息
  UpdateList({ commit }, data) {
    console.log(this.state)
    return new Promise((resolve, reject) => {
      updateList(data).then(response => {
        console.log(response)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // 这里删除商品信息
  DeleteList({ commit }, data) {
    console.log(this.state)
    return new Promise((resolve, reject) => {
      deleteList(data).then(response => {
        console.log(response)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  }
}
export default {
  namespaced: true,
  actions
}
