<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <div class="container" v-if="loaded">
    <div class="top-card">
      <h1 class="welcome-message">{{ youWinLabel }}<br>{{ user.name }}</h1>
      <h3 class="description">{{ youWinText }}</h3>
    </div>
    <div class="block-pink">
      <img class="landing-drawing" alt="Game drawing" src="./../assets/drawing.png">
    </div>
  </div>
</template>

<script>
import LoadingComponent from "@/components/LoadingComponent.vue";

export default {
  name: 'TheEndPage',
  components: {
    LoadingComponent
  },
  props: {
    points: Number
  },
  data() {
    return {
      user: {},
      language: "en",
      loaded: false,
      youWinText: "",
      youWinLabel: "",
    }
  },
  mounted() {
    this.init();
    this.loaded = true
  },
  methods: {
    async init() {
      this.user = this.$store.state.user
      window.Telegram.WebApp.BackButton.show()
      window.Telegram.WebApp.onEvent('backButtonClicked', this.onClickBack)
      window.Telegram.WebApp.MainButton.hide()
      this.setLanguage();
    },
    setLanguage() {
      this.language = this.user.language
      if (this.language === "ru") {
        window.Telegram.WebApp.MainButton.text = "Назад"
        this.youWinLabel = "Поздравляю"
        this.youWinText = "Вы набрали " + this.points + " очков!"
      } else {
        window.Telegram.WebApp.MainButton.text = "Back"
        this.youWinLabel = "Congrats"
        this.youWinText = "You've got " + this.points + " points!"
      }
    },
    onClickBack() {
      this.$router.push('/')
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

body {
  background-color: var(--tg-theme-bg-color);
}

.loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 10px;
}

.welcome-message {
  padding: 20px;
  margin: 10px;
  color: #ffffff;
  background: transparent;
  font-weight: bold;
  font-size: 40px;
}

.description {
  margin: 10px;
  padding: 20px;
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

.landing-drawing {
  max-width: 225px;
  max-height: 225px;
}


</style>
