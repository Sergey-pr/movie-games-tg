<template>
  <div class="question-card" ref="questionCard" tabIndex="1">
    <div class="card-block">
      <div class="content-block">
        <img alt="drawing" class="drawing" :src="imgUrlPrefix + card.drawing_id">
      </div>
      <p class="block-label">{{ drawingTitle }}</p>
    </div>
    <div id="section1" class="card-block" v-if="state >= 1">
      <div class="content-block">
        <p class="quote" :style="{color: card.text_color}">{{ quote }}</p>
      </div>
      <p class="block-label">{{ quoteTitle }}</p>
    </div>
    <div id="section2" class="card-block" v-if="state >= 2">
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

export default {
  name: 'CardPage',
  components: {},
  props: {
    card: {
      type: Object,
      default() {
        return {}
      }
    }
  },
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
      drawingTitle: "3 points",
      quoteTitle: "2 points",
      pixelatedTitle: "1 point",
      user: {},
      loaded: false,
      quote: "",
      answers: {},
      name: "",
      state: 0,
      winPopupConfig: {},
      losePopupConfig: {},
      imgUrlPrefix: "",
      star1color: "#FFD024",
      star2color: "#FFD024",
      star3color: "#FFD024"
    }
  },
  created() {
    if (process.env.VUE_APP_BASE_URL !== undefined) {
      this.imgUrlPrefix = process.env.VUE_APP_BASE_URL + "/api/public/bot-image/";
    } else {
      this.imgUrlPrefix = "/api/public/bot-image/";
    }
  },
  mounted() {
  },
  methods: {
    async init() {
      this.user = this.$store.state.user
      window.Telegram.WebApp.MainButton.hide()
      this.background = 'linear-gradient(45deg, ' + this.card.bg_color_1 + ', ' + this.card.bg_color_2 + ')'
      this.setLanguage();
    },
    setLanguage() {
      this.language = this.user.language
      if (this.language === "ru") {
        switch (this.card.category) {
          case "movie":
            this.category = "Фильм"
        }
        this.shuffleArray(this.card.answers_ru)
        this.card.answers_ru.forEach((element) => {
          this.answers[element] = false
        })
        this.name = this.card.name_ru
        this.quote = this.card.quote_ru
        this.drawingTitle = "3 балла"
        this.quoteTitle = "2 балла"
        this.pixelatedTitle = "1 балл"
        this.winPopupConfig = {
          title: "Верно",
          message: "Вы правильно угадали название:\n" + this.name,
          buttons: [
            {
              text: "Дальше"
            }
          ]
        }
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
        this.shuffleArray(this.card.answers_en)
        this.card.answers_en.forEach((element) => {
          this.answers[element] = false
        })
        this.name = this.card.name_en
        this.quote = this.card.quote_en
        this.winPopupConfig = {
          title: "Correct",
          message: "You answered correctly, movie name is:\n" + this.name,
          buttons: [
            {
              text: "Next"
            }
          ]
        }
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
      if (item === this.name) {
        this.win()
      } else if (this.state <= 1) {
        switch (this.state) {
          case 0:
            this.star1color = "#ff0000"
                break
          case 1:
            this.star2color = "#ff0000"
                break
        }
        this.state += 1
        this.answers[item] = true
        await this.delay(100)
        this.scrollToId()
      } else {
        this.star3color = "#ff0000"
        this.lose()
      }
    },
    scrollToId() {
      let access = document.getElementById("section" + this.state);
      access.scrollIntoView({behavior: 'smooth'}, true);
    },
    delay(time) {
      return new Promise(resolve => setTimeout(resolve, time));
    },
    win() {
      window.Telegram.WebApp.showPopup(this.winPopupConfig, this.emitWinState)
    },
    lose() {
      window.Telegram.WebApp.showPopup(this.losePopupConfig, this.emitWinState)
    },
    emitWinState() {
      this.$emit("emit-win")
    },
    shuffleArray(array) {
      for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
      }
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

.card-block {
  height: 325px;
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
  height: 100%;
}

.content-block {
  height: 250px;
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
  color: #ffffff;
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