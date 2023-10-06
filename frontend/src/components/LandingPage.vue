<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <div class="container" v-if="loaded">
    <div class="top-card">
      <div class="lang-changer">
        <p class="lang">RU🇷🇺</p>
        <label class="switch">
          <input type="checkbox" v-model="language" true-value="en" false-value="ru"  @change="changeLang()">
          <span class="slider round"></span>
        </label>
        <p class="lang">EN🇬🇧</p>
      </div>
      <img v-if="language === 'ru'" class="logo" alt="Game logo" src="./../assets/logo_ru.png">
      <img v-if="language === 'en'" class="logo" alt="Game logo" src="./../assets/logo_en.png">
      <h1 class="welcome-message-landing">{{ welcomeMessage }}{{ userName }}</h1>
      <h3 class="description-landing">{{ description }}</h3>
      <button class="btn-start" @click="this.onClickStart()">
        <p class="btn-start-label">{{ startLabel }}</p>
      </button>
    </div>
    <div class="block-pink">
      <img class="landing-drawing" alt="Game drawing" src="./../assets/drawing.png">
      <button class="btn-landing" @click="this.onClickRules()">
        <p class="btn-landing-label">{{ rulesLabel }}</p>
      </button>
      <button class="btn-landing" @click="this.onClickLeaderboard()">
        <p class="btn-landing-label">{{ leaderboardLabel }}</p>
      </button>
    </div>
  </div>
</template>

<script>
import {useAuth, useUsers} from "@/services/adapter";
import LoadingComponent from "@/components/LoadingComponent.vue";

export default {
  name: 'LandingPage',
  components: {
    LoadingComponent,
  },
  data() {
    return {
      welcomeMessage: "",
      description: "",
      user: {},
      userName: "",
      language: "en",
      loaded: false,
      startLabel: "",
      rulesLabel: "",
      leaderboardLabel: ""
    }
  },
  created() {
    this.init();
  },
  methods: {
    async init() {
      // Send login request to log in and verify initData hash
      let response = await useAuth().login(
          {init_data: window.Telegram.WebApp.initData}
      );
      if (response.status !== 200) {
        return
      }
      // Set token to store
      this.$store.commit('setJwt', response.data["token"])
      // Get user data
      response = await useUsers().user(response.data["token"]);
      this.user = response.data
      this.userName = response.data["name"]
      this.language = response.data["language"]
      this.$store.commit('setUser', response.data)
      // Hide default telegram buttons
      window.Telegram.WebApp.BackButton.hide()
      window.Telegram.WebApp.MainButton.hide()
      // Set up translations
      this.setLanguage();
      // Tell telegram that the app is ready
      await window.Telegram.WebApp.ready()
      this.loaded = true;
    },
    setLanguage() {
      if (this.language === "ru") {
        this.welcomeMessage = "Привет "
        this.description = "Давай сыграем в игру на знание известных фильмов. Здесь ты проверишь свои знания и узнаешь о новых интересных фильмах."
        this.startLabel = "СТАРТ"
        this.rulesLabel = "ПРАВИЛА"
        this.leaderboardLabel = "ЛУЧШИЕ ИГРОКИ"
      } else {
        this.welcomeMessage = "Welcome "
        this.description = "Let's see how good is your movie knowledge. You will try to guess movie names and learn about famous movies."
        this.startLabel = "START"
        this.rulesLabel = "RULES"
        this.leaderboardLabel = "LEADERBOARD"
      }
    },
    changeLang() {
      this.setLanguage();
      this.$store.commit('setLang', this.language)
      useUsers().changeLang(this.$store.state.jwt, this.language)
    },
    onClickStart() {
      this.$router.push('/play')
    },
    onClickRules() {
      this.$router.push('/rules')
    },
    onClickLeaderboard() {
      this.$router.push('/leaderboard')
    },
  }
}
</script>

<style>
.loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 10px;
}

.welcome-message-landing {
  margin: 10px;
  color: #ffffff;
  background: transparent;
  font-weight: bold;
  font-size: 40px;
}

.description-landing {
  margin: 10px;
  color: #999999;
}

body {
  background-color: var(--tg-theme-bg-color);
  margin: 0;
}

.container {
  background: linear-gradient(#cc7676, #d0a67f);
  margin: 0 45px 20px;
  border-radius: 0 0 25px 25px;
  box-shadow: 15px 15px 30px rgba(0, 0, 0, .3);
}

.top-card {
  background-color: #433789;
  border-radius: 0 0 50% 50% / 85% 85% 15% 15% ;
}

.lang {
  color: var(--tg-theme-text-color);
  font-size: 25px;
  font-weight: bold;
  margin: 0 5px 5px;
}

.lang-changer {
  padding-top: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo {
  margin: 30px 10px 30px 10px;
  max-width: 250px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}

/* The switch - the box around the slider */
.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
}

/* Hide default HTML checkbox */
.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

/* The slider */
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: var(--tg-theme-button-color);
  -webkit-transition: .4s;
  transition: .4s;
  border-radius: 34px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  -webkit-transition: .4s;
  transition: .4s;
  border-radius: 50%;
}

input:focus + .slider {
  box-shadow: 0 0 1px #2196F3;
}

input:checked + .slider:before {
  -webkit-transform: translateX(26px);
  -ms-transform: translateX(26px);
  transform: translateX(26px);
}

.btn-start {
  background: transparent;
  border: none;
}

.btn-start-label {
  font-weight: bold;
  font-size: 40px;
  color: #ffffff;
}

.btn-landing {
  color: #ffffff;
  background-color: #433789;
  font-weight: normal;
  font-size: 25px;
  border: none;
  display: block;
  width: 100%;
  height: 50px;
  padding: 0;
  margin: 20px 0 0;
}

.btn-landing-label {
  padding: 0;
  margin: 0;
}

.landing-drawing {
  max-width: 225px;
  max-height: 225px;
}

.block-pink {
  padding-bottom: 20px;
}

</style>
