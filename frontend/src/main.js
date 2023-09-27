import { createApp } from 'vue'
import { createStore } from 'vuex'
import { createRouter, createWebHistory } from "vue-router"

import App from './App.vue'
import PlayGame from './components/PlayGame.vue'
import LandingPage from './components/LandingPage.vue'


const routes = [
    { path: '/', component: LandingPage },
    { path: '/play', component: PlayGame },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

const store = createStore({
    state () {
        return {
            jwt: "",
            user: {}
        }
    },
    mutations: {
        setJwt (state, jwt) {
            state.jwt = jwt
        },
        setUser (state, user) {
            state.user = user
        }
    }
})

const app = createApp(App)
app.use(router)
app.use(store)

app.mount('#app')