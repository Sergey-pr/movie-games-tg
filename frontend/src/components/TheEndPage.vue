<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <div class="end-game-container" v-if="loaded">
    <div class="end-game-card">
      <h1 class="you-win-message">{{ youWinLabel }}<br>{{ user.name }}</h1>
      <h3 class="you-win-description">{{ youWinText }}</h3>
    </div>
    <div>
      <img class="end-game-drawing" alt="Game drawing" src="./../assets/end_drawing.png">
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
    // Sets translation
    setLanguage() {
      if (this.user.language === "ru") {
        window.Telegram.WebApp.MainButton.text = "Назад"
        this.youWinLabel = "Поздравляю"
        this.youWinText = "Вы набрали " + this.points + " очков!"
      } else {
        window.Telegram.WebApp.MainButton.text = "Back"
        this.youWinLabel = "Congrats"
        this.youWinText = "You've got " + this.points + " points!"
      }
    },
    // Returns to the landing page
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

.you-win-message {
  padding: 20px;
  margin: 10px;
  color: #ffffff;
  background: transparent;
  font-weight: bold;
  font-size: 40px;
}

.you-win-description {
  margin: 10px;
  padding: 20px;
  color: #999999;
}

body {
  background-color: var(--tg-theme-bg-color);
  margin: 0;
}

.end-game-container {
  background: linear-gradient(#cc7676, #8b60a8);
  margin: 0 45px 20px;
  border-radius: 0 0 25px 25px;
  box-shadow: 15px 15px 30px rgba(0, 0, 0, .3);
}

.end-game-card {
  background-color: #433789;
  border-radius: 0 0 50% 50% / 85% 85% 15% 15% ;
}

.end-game-drawing {
  max-width: 225px;
  max-height: 225px;
}


</style>
