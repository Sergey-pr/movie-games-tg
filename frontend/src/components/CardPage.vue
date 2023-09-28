<template>
  <div class="card" :style="{background: 'linear-gradient(45deg, ' + card.bg_color_1 + ', ' + card.bg_color_2 + ')'}">
    <div class="card-block ">
      <h1 class="drawing-label">{{ drawingTitle }}</h1>
      <img alt="drawing" class="drawing" :src="imgUrlPrefix + card.drawing_id">
    </div>
    <div id="section1" class="card-block" v-if="state >= 1">
      <h1 class="quote-label">{{ quoteTitle }}</h1>
      <div class="quote">
        <p :style="{color: card.text_color}">{{ quote }}</p>
      </div>
    </div>
    <div id="section2" class="card-block" v-if="state >= 2">
      <h1>{{ pixelatedTitle }}</h1>
      <img alt="pixelated" class="pixelated" :src="imgUrlPrefix + card.pixelated_id">
    </div>
  </div>
  <div class="answers">
    <PrimeButton class="btn" :key="key" :label="key" v-for="(value, key) in answers" :disabled="value" @click="this.processAnswer(key)"></PrimeButton>
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
      category: "Movie",
      language: "en",
      drawingTitle: "Drawing",
      quoteTitle: "Quote",
      pixelatedTitle: "Screenshot",
      user: {},
      loaded: false,
      quote: "",
      answers: {},
      name: "",
      state: 0,
      winPopupConfig: {},
      losePopupConfig: {},
    }
  },
  created() {
    this.imgUrlPrefix = "/api/public/bot-image/";
  },
  mounted() {
  },
  methods: {
    async init() {
      this.user = this.$store.state.user
      window.Telegram.WebApp.MainButton.hide()
      this.setLanguage();
    },
    setLanguage() {
      this.language = this.user.language
      if (this.language === "ru") {
        switch (this.card.category) {
          case "movie":
            this.category = "Фильм"
        }
        this.card.answers_ru.forEach((element) => {
          this.answers[element] = false
        })
        this.name = this.card.name_ru
        this.quote = this.card.quote_ru
        this.drawingTitle = "Рисунок"
        this.quoteTitle = "Цитата"
        this.pixelatedTitle = "Кадр из фильма"
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
        this.state += 1
        this.answers[item] = true
        await this.delay(100)
        this.scrollToId()
      } else {
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
    }
  }
}
</script>

<style>

h1 {
  color: var(--tg-theme-text-color);
  background-color: var(--tg-theme-secondary-bg-color);
}

.card {
  position: absolute;
  top: 20px;
  right: 10px;
  left: 10px;
  margin-bottom: 20px;
  border-radius: 25px;
}

.card-block {
  height: 500px;
  scroll-snap-align: start;
}

.drawing {
  max-width: 250px;
  max-height: 250px;
}

.pixelated {
  max-width: 250px;
  max-height: 250px;
}

.quote-label {
}

.quote {
  color: white;
  font-weight: bold;
  font-size: 25px;
}

.answers {
  position: fixed;
  left: 10px;
  right: 10px;
  bottom: 10px;
}

.p-button {
  background-color: var(--tg-theme-button-color);
  border: none;
}

.p-button:enabled:hover {
  background-color: var(--tg-theme-button-color);
  border: none;
}

.p-button-label {
  color: var(--tg-theme-button-text-color);
}

body {
  scroll-behavior: smooth;
}

</style>