// get user info
import { getInfo } from '@/api/user'

const actions = {
  // 这里是add商品信息
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        console.log(response)
        const { user_list } = response
        console.log(user_list)

        if (!user_list) {
          return reject('Verification failed, please Login again.')
        }

        const { name, avatar } = user_list

        commit('SET_NAME', name)
        commit('SET_AVATAR', avatar)
        resolve(user_list)
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
