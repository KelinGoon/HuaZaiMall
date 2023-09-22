// get user info
import { addList } from '@/api/product'

const actions = {
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
  }
}
export default {
  namespaced: true,
  actions
}
