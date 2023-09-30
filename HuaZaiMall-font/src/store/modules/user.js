import { getInfo, login } from '@/api/user'
import { getToken, removeToken, setToken } from '@/utils/auth'

const getDefaultState = () => {
  return {
    token: getToken(),
    name: '',
    avatar: '5.svg'
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = '\n' + 'https://avatarkelin.oss-cn-shenzhen.aliyuncs.com/g5YQiqjdMM.png'
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { account, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: account.trim(), password: password }).then(response => {
        // const { data } = response
        commit('SET_TOKEN', response.data)
        setToken(response.data)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo(state.token).then(response => {
        console.log(response)
        const res = response.data
        console.log(res)

        if (!res) {
          return reject('Verification failed, please Login again.')
        }
        console.log(111)
        const name = res.fullName
        const avatar = null
        commit('SET_NAME', name)
        commit('SET_AVATAR', avatar)
        resolve(res)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

