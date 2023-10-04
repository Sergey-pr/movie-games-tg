<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <div class="main" v-if="loaded">
    <h1 class="rules-label">{{ rulesLabel }}</h1>
    <h3 class="rules-sub-label">{{ rulesLabel1 }}</h3>
    <p class="rules-text">{{ rules1 }}</p>
    <h3 class="rules-sub-label">{{ rulesLabel2 }}</h3>
    <p class="rules-text">{{ rules2 }}</p>
    <h3 class="rules-sub-label">{{ rulesLabel3 }}</h3>
    <p class="rules-text">{{ rules3 }}</p>
    <div class="buttons">
      <button class="btn" @click="this.onClickBack()">
        <p class="btn-label">{{ backLabel }}</p>
      </button>
    </div>
  </div>
</template>

<script>
import LoadingComponent from "@/components/LoadingComponent.vue";

export default {
  name: 'RulesComponent',
  components: {
    LoadingComponent,
  },
  data() {
    return {
      user: {},
      language: "en",
      loaded: false,
      rulesLabel1: "",
      rules1: "",
      rulesLabel2: "",
      rules2: "",
      rulesLabel3: "",
      rules3: "",
      rulesLabel: "",
      backLabel: ""
    }
  },
  created() {
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
        this.backLabel = "Назад"
        this.rulesLabel = "Правила"
        this.rulesLabel1 = "1 часть"
        this.rules1 = "Нужно отгадать фильм по стилистическому рисунку. "+
            "Это может быть ключевой предмет сюжета, персонаж, пейзаж или даже природное явление. " +
            "За угаданный фильм по рисунку участник получает 3 балла." +
            "И он может взять другую карточку. " +
            "Если участник не угадал, то он переходит на 2 часть."
        this.rulesLabel2 = "2 часть"
        this.rules2 = "Нужно угадать цитату из фильма. " +
            "За угаданный фильм по Цитате участник получает 2 балла. И он может взять другую карточку. " +
            "Если участник не угадал, то он переходит на 3 часть."
        this.rulesLabel3 = "3 часть"
        this.rules3 = "Последняя попытка! Нужно угадать стилизованный под мозайку кадо из фильма. " +
            "За угаданный фильм по стилизованному кадру участник получает 1 балла. И он может перейти на другую карточку. " +
            "Если фильм не угадан, то участник может перевернуть карточку и посмотреть ответ."
      } else {
        this.backLabel = "Back"
        this.rulesLabel = "Rules"
        this.rulesLabel1 = "Part 1"
        this.rules1 = "You need to guess the film based on its stylistic drawing. "+
            "This could be a key plot object, a character, a landscape, or even a natural phenomenon. " +
            "For guessing the film based on the picture, the participant receives 3 points." +
            "And he can take another card. " +
            "If the participant does not guess correctly, then he goes to part 2."
        this.rulesLabel2 = "Part 2"
        this.rules2 = "You need to guess the quote from the movie. " +
            "For guessing the movie based on the Quote, the participant receives 2 points. And he can take another card. " +
            "If the participant does not guess correctly, then he goes to part 3."
        this.rulesLabel3 = "Part 3"
        this.rules3 = "Last try! You need to guess a stylized mosaic frame from the film. " +
            "For guessing the film based on the stylized frame, the participant receives 1 point. And he can switch to another card. " +
            "If the movie is not guessed correctly, the participant can turn over the card and see the answer."
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

.rules-text {
  margin: 10px;
  color: var(--tg-theme-hint-color);
}

.rules-label {
  margin: 10px;
  color: var(--tg-theme-text-color);
}

.rules-sub-label {
  padding-top: 5px;
  padding-bottom: 5px;
  color: #ffffff;
  background-color: #000000;
}

body {
  background-color: var(--tg-theme-bg-color);
}

.main {
  position: absolute;
  top: 20px;
  right: 10px;
  left: 10px;
  margin-bottom: 20px;
  border-radius: 25px;
  background-color: var(--tg-theme-secondary-bg-color);
}

.buttons {
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
}

.btn {
  background-color: #000000;
  border: none;
  color: #ffffff;
  font-weight: bold;
  font-size: 25px;
  margin: 10px;
  border-radius: 5px;
}

.btn-label {
  margin: 5px;
}

</style>
