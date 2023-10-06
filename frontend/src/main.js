import { createApp } from 'vue'
import { createStore } from 'vuex'
import { createRouter, createWebHistory } from "vue-router"

import App from './App.vue'
import PlayGame from './components/PlayGame.vue'
import LandingPage from './components/LandingPage.vue'
import RulesPage from "@/components/RulesPage.vue";
import LeaderboardPage from "@/components/LeaderboardPage.vue";

import 'primevue/resources/themes/bootstrap4-light-blue/theme.css';
import 'primevue/resources/primevue.min.css';


const routes = [
    { path: '/', component: LandingPage },
    { path: '/play', component: PlayGame },
    { path: '/rules', component: RulesPage },
    { path: '/leaderboard', component: LeaderboardPage },
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
        },
        setLang(state, lang) {
            state.user.language = lang
        }
    }
})

const app = createApp(App)
app.use(router)
app.use(store)

app.mount('#app')