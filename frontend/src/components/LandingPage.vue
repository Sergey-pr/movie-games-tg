<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <div class="main" v-if="loaded">
    <div class="lang-changer">
      <p class="lang">RU🇷🇺</p>
      <label class="switch">
        <input type="checkbox" v-model="language" true-value="en" false-value="ru"  @change="changeLang()">
        <span class="slider round"></span>
      </label>
      <p class="lang">EN🇬🇧</p>
    </div>
    <img class="logo" alt="Game logo" src="./../assets/logo.png">
    <h1 class="welcome-message">{{ welcomeMessage }}{{ userName }}</h1>
    <h3 class="description">{{ description }}</h3>
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
    }
  },
  created() {
    this.init();
  },
  methods: {
    async init() {
      window.Telegram.WebApp.ready()
      let initData = window.Telegram.WebApp.initData
      let body = JSON.parse('{"' + initData.replace(/&/g, '","').replace(/=/g,'":"') + '"}', function(key, value) { return key===""?value:decodeURIComponent(value) });
      body.user = JSON.parse(body.user)
      let response = await useAuth().login(body);
      this.$store.commit('setJwt', response.data["token"])

      response = await useUsers().user(response.data["token"]);
      this.user = response.data
      this.userName = response.data["name"]
      this.language = response.data["language"]
      this.$store.commit('setUser', response.data)

      window.Telegram.WebApp.BackButton.hide()
      window.Telegram.WebApp.MainButton.show()
      window.Telegram.WebApp.onEvent('mainButtonClicked', this.onClickStart)
      this.setLanguage();
      this.loaded = true;
    },
    setLanguage() {
      if (this.language === "ru") {
        window.Telegram.WebApp.MainButton.text = "Начать"
        this.welcomeMessage = "Привет "
        this.description = "Давай сыграем в игру на знание известных фильмов. Здесь ты проверишь свои знания и узнаешь о новых интересных фильмах."
      } else {
        window.Telegram.WebApp.MainButton.text = "Start"
        this.welcomeMessage = "Welcome "
        this.description = "Let's see how good is your movie knowledge. You will try to guess movie names and learn about famous movies."
      }
    },
    changeLang() {
      this.setLanguage();
      useUsers().changeLang(this.$store.state.jwt, this.language)
    },
    onClickStart() {
      this.$router.push('/play')
    }
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

.welcome-message {
  margin: 10px;
  color: var(--tg-theme-text-color);
}

.description {
  margin: 10px;
  color: var(--tg-theme-hint-color);
}

body {
  background-color: var(--tg-theme-bg-color);
}

.main {
  position: absolute;
  top: 20px;
  bottom: 20px;
  right: 10px;
  left: 10px;
  border-radius: 25px;
  background-color: var(--tg-theme-secondary-bg-color);
}

.lang {
  color: var(--tg-theme-text-color);
  font-size: 25px;
  font-weight: bold;
  margin: 0 5px 5px;
}

.lang-changer {
  margin: 5px;
  display: flex;
  float: right;
}

.logo {
  margin: 30px 10px 30px 10px;
  max-width: 18em;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: var(--tg-theme-button-color);
  margin-top: 60px;
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

</style>
