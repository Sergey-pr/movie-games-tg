<template>
  <div class="question-card" ref="questionCard" tabIndex="1">
    <div class="card-block">
      <div class="content-block">
        <img alt="drawing" class="drawing" :src="imgUrlPrefix + card.drawing_id">
      </div>
      <p class="block-label">{{ drawingTitle }}</p>
    </div>
    <div id="section2" class="card-block" v-if="points <= 2">
      <div class="content-block">
        <p class="quote" :style="{color: card.text_color}">{{ quote }}</p>
      </div>
      <p class="block-label">{{ quoteTitle }}</p>
    </div>
    <div id="section1" class="card-block" v-if="points <= 1">
      <div class="content-block">
        <img alt="pixelated" class="pixelated" :src="imgUrlPrefix + card.pixelated_id">
      </div>
      <p class="block-label">{{ pixelatedTitle }}</p>
    </div>
  </div>
  <div class="stars-block">
  <svg class="starSVG" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" width="30" height="30">
      <polygon :fill="star1color" points="15 0.001 20.25 8.987 30 11.46 23.495 19.487 24.27 30.001 15 25.976 5.73 30.001 6.505 19.487 0 11.46 9.75 8.987" />
    </svg>
  <svg class="starSVG" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" width="30" height="30">
      <polygon :fill="star2color" points="15 0.001 20.25 8.987 30 11.46 23.495 19.487 24.27 30.001 15 25.976 5.73 30.001 6.505 19.487 0 11.46 9.75 8.987" />
    </svg>
  <svg class="starSVG" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" width="30" height="30">
      <polygon :fill="star3color" points="15 0.001 20.25 8.987 30 11.46 23.495 19.487 24.27 30.001 15 25.976 5.73 30.001 6.505 19.487 0 11.46 9.75 8.987" />
    </svg>
</div>
  <div class="answers">
    <button class="btn-answer" :key="key" v-for="(value, key) in answers" :disabled="value" @click="this.processAnswer(key)">
      <p class="btn-label">{{ key }}</p>
    </button>
  </div>
</template>

<script>

import { shuffleArray } from "@/services/utils";

export default {
  name: 'CardPage',
  components: {},
  emits: ["emit-win"],
  props: {
    card: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  // Wait until props are given to component before init page
  watch: {
    $props: {
      handler(newVal) {
        if (Object.keys(newVal).length !== 0) {
          this.init();
        }
      },
      deep: true,
      immediate: true,
    },
  },
  data() {
    return {
      background: "",
      language: "en",
      drawingTitle: "GUESS BY DRAWING",
      quoteTitle: "GUESS BY QUOTE",
      pixelatedTitle: "GUESS BY FRAME",
      user: {},
      loaded: false,
      quote: "",
      answers: {},
      name: "",
      points: 3,
      winPopupConfig: {},
      losePopupConfig: {},
      imgUrlPrefix: "",
      star1color: "#FFD024",
      star2color: "#FFD024",
      star3color: "#FFD024"
    }
  },
  created() {
    // Set image urls prefix for dev/prod
    if (process.env.VUE_APP_BASE_URL !== undefined) {
      this.imgUrlPrefix = process.env.VUE_APP_BASE_URL + "/api/public/bot-image/";
    } else {
      this.imgUrlPrefix = "/api/public/bot-image/";
    }
  },
  methods: {
    async init() {
      this.user = this.$store.state.user
      // Hide main telegram button at the bottom
      window.Telegram.WebApp.MainButton.hide()
      // Set dynamic background for card
      this.background = 'linear-gradient(45deg, ' + this.card.bg_color_1 + ', ' + this.card.bg_color_2 + ')'
      this.setLanguage();
    },
    // Sets translation
    setLanguage() {
      this.language = this.user.language
      if (this.language === "ru") {
        switch (this.card.category) {
          case "movie":
            this.category = "Фильм"
        }
        // Shuffle answers so the order would be different every playthrough
        shuffleArray(this.card.answers_ru)
        // Set false to all the answers, if the answer is chosen we set this answer to true
        this.card.answers_ru.forEach((element) => {
          this.answers[element] = false
        })
        this.name = this.card.name_ru
        this.quote = this.card.quote_ru
        this.drawingTitle = "УГАДАЙ ПО РИСУНКУ"
        this.quoteTitle = "УГАДАЙ ПО ЦИТАТЕ"
        this.pixelatedTitle = "УГАДАЙ ПО КАДРУ"
        // Set telegram popup config to use with window.Telegram.WebApp.showPopup
        let pointsText = "балл"
        if (this.points > 1) {
          pointsText = "балла"
        }
        this.winPopupConfig = {
          title: "Верно",
          message: "Вы правильно угадали название:\n" + this.name +
              "\nВы заработали " + this.points + " " + pointsText,
          buttons: [
            {
              text: "Дальше"
            }
          ]
        }
        // Set telegram popup config to use with window.Telegram.WebApp.showPopup
        this.losePopupConfig = {
          title: "Неверно",
          message: "Правильный ответ:\n" + this.name,
          buttons: [
            {
              text: "Дальше"
            }
          ]
        }
      } else {
        switch (this.card.category) {
          case "movie":
            this.category = "Movie"
        }
        // Shuffle answers so the order would be different every playthrough
        shuffleArray(this.card.answers_en)
        // Set false to all the answers, if the answer is chosen we set this answer to true
        this.card.answers_en.forEach((element) => {
          this.answers[element] = false
        })
        this.name = this.card.name_en
        this.quote = this.card.quote_en
        // Set telegram popup config to use with window.Telegram.WebApp.showPopup
        this.winPopupConfig = {
          title: "Correct",
          message: "You answered correctly, movie name is:\n" + this.name +
              "\nYou've earned " + this.points + " points",
          buttons: [
            {
              text: "Next"
            }
          ]
        }
        // Set telegram popup config to use with window.Telegram.WebApp.showPopup
        this.losePopupConfig = {
          title: "Wrong",
          message: "The correct answer is:\n" + this.name,
          buttons: [
            {
              text: "Next"
            }
          ]
        }
      }
    },
    async processAnswer(item) {
      // If answer is correct we win
      if (item === this.name) {
        this.win()
        return
      }
      this.points -= 1
      // If there are 1 or more points we lose points
      // else we lose
      if (this.points >= 1) {
        // If answer is not correct we lose 1 point
        // and set current point star color to red
        switch (this.points) {
          case 2:
            this.star1color = "#ff0000"
                break
          case 1:
            this.star2color = "#ff0000"
                break
        }
        // Set this answer value to true
        this.answers[item] = true
        await this.delay(100)
        // Scroll to the new hint
        this.scrollToCurrentSection()
      } else {
        this.star3color = "#ff0000"
        this.lose()
      }
    },
    scrollToCurrentSection() {
      // Get section by our point
      let access = document.getElementById("section" + this.points);
      if (access !== null) {
        access.scrollIntoView({behavior: 'smooth'}, true);
      }
    },
    delay(time) {
      return new Promise(resolve => setTimeout(resolve, time));
    },
    win() {
      window.Telegram.WebApp.showPopup(this.winPopupConfig, this.emitPlayedState)
    },
    lose() {
      window.Telegram.WebApp.showPopup(this.losePopupConfig, this.emitPlayedState)
    },
    // Emits action to go to show card info in the PlayGame component
    emitPlayedState() {
      this.$emit("emit-win", this.points)
    },
  }
}
</script>

<style>

.question-card {
  background: v-bind(background);
  margin: 0 60px 20px;
  border-radius: 0 0 25px 25px;
  overflow: auto;
  box-shadow: 15px 15px 30px rgba(0, 0, 0, .3);
}

.drawing {
  margin-top: 25px;
  max-width: 225px;
  max-height: 225px;
}

.block-label {
  color: #ffffff;
  background-color: #433789;
  font-weight: normal;
  font-size: 25px;
}

.quote {
  padding: 10px;
}

.pixelated {
  width: 100%;
}

.content-block {
  display: block;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 25px;
}

.answers {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stars-block {
  display: flex;
  align-items: center;
  justify-content: center;
}

.starSVG {
  margin: 0 10px 0 10px;
}

.btn-answer {
  border: none;
  color: var(--tg-theme-text-color);
  font-weight: normal;
  font-size: 25px;
  margin: 10px;
  border-radius: 5px;
  background-color: transparent;
}
.btn-answer:disabled {
  color: #ff0000;
}

.btn-label {
  margin: 5px;
}

body {
  scroll-behavior: smooth;
}

</style>