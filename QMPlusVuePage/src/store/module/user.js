import { login } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
export const user = {
    namespaced: true,
    state: {
        userInfo: {
            uuid: "",
            userName: "",
            nickName: "",
            headerImg: "",
            authority: "",
        },
        token: "",
        expiresAt: "",
        enPriseId:"",
    },
    mutations: {
        setUserInfo(state, userInfo) {
            // 这里的 `state` 对象是模块的局部状态
            state.userInfo = userInfo
        },
        setToken(state, token) {
            // 这里的 `state` 对象是模块的局部状态
            state.token = token
        },
        setExpiresAt(state, expiresAt) {
            // 这里的 `state` 对象是模块的局部状态
            state.expiresAt = expiresAt
        },
        setEnPriseId(state, enPriseId) {
            // 这里的 `state` 对象是模块的局部状态
            state.enPriseId = enPriseId
        },
        LoginOut(state) {
            state.userInfo = {}
            state.token = ""
            state.expiresAt = ""
            state.enPriseId = ""
            router.push({ name: 'login', replace: true })
            sessionStorage.clear()
            window.location.reload()
        },
        ResetUserInfo(state, userInfo = {}) {
            state.userInfo = {...state.userInfo,
                ...userInfo
            }
        }
    },
    actions: {
        async LoginIn({ commit }, loginInfo) {
            const res = await login(loginInfo)
            commit('setUserInfo', res.data.user)
            commit('setToken', res.data.token)
            commit('setExpiresAt', res.data.expiresAt)
            commit('setEnPriseId', res.data.enPriseId)
            if (res.success) {
                const redirect = router.history.current.query.redirect
                if (redirect) {
                    router.push({ path: redirect })
                } else {
                    router.push({ path: '/layout/person' })
                }
            }
        },
        async LoginOut({ commit }) {
            const res = await jsonInBlacklist()
            if (res.success) {
                commit("LoginOut")
            }
        }
    },
    getters: {
        userInfo(state) {
            return state.userInfo
        },
        token(state) {
            return state.token
        },
        expiresAt(state) {
            return state.expiresAt
        },
        enPriseId(state){
            return state.enPriseId
        }
    }
}