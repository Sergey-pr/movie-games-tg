<template>
  <div class="card">
    <img id="screenshot" alt="screenshot" class="screenshot" :src="imgUrlPrefix + card.screenshot_id">
    <h2 class="name">{{ name }}</h2>
    <h3 class="name-ru">{{ name_ru }}</h3>
    <p class="desc">{{ desc }}</p>
    <p class="facts-label">{{ factsLabel }}</p>
    <div class="fact" :key="index" v-for="(item, index) in facts" :style="{color: card.text_color}">
      {{ item }}
      <hr v-if="index + 1 < facts.length">
    </div>
  </div>
  <div class="stars-block">
    <svg class="starSVG" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" width="30" height="30">
      <polygon fill=#FFD024 points="15 0.001 20.25 8.987 30 11.46 23.495 19.487 24.27 30.001 15 25.976 5.73 30.001 6.505 19.487 0 11.46 9.75 8.987" />
    </svg>
    <p class="points-label">{{ pointsText }}: {{ points }}</p>
    <svg class="starSVG" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" width="30" height="30">
      <polygon fill=#FFD024 points="15 0.001 20.25 8.987 30 11.46 23.495 19.487 24.27 30.001 15 25.976 5.73 30.001 6.505 19.487 0 11.46 9.75 8.987" />
    </svg>
  </div>
</template>

<script>

export default {
  name: 'CardInfoPage',
  components: {},
  emits: ["emit-next"],
  props: {
    card: {
      type: Object,
      default() {
        return {}
      }
    },
    points: Number
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
      factsLabel: "INTERESTING FACTS",
      imgUrlPrefix: "",
      textColor: "",
      facts: [],
      loaded: false,
      name: "",
      name_ru: "",
      desc: "",
      state: 0,
      pointsText: "POINTS",
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
      window.Telegram.WebApp.onEvent('mainButtonClicked', this.emitNext)
      window.Telegram.WebApp.MainButton.show()
      this.background = 'linear-gradient(45deg, ' + this.card.bg_color_1 + ', ' + this.card.bg_color_2 + ')'
      this.textColor = this.card.text_color
      this.setLanguage();
    },
    // Sets translation
    setLanguage() {
      if (this.$store.state.user.language === "ru") {
        window.Telegram.WebApp.MainButton.text = "Дальше"
        this.factsLabel = "ИНТЕРЕСНЫЕ ФАКТЫ"
        this.pointsText = "БАЛЛЫ"
        this.name = this.card.name_en
        this.name_ru = this.card.name_ru
        this.desc = this.card.desc_ru
        this.facts = this.card.facts_ru
      } else {
        window.Telegram.WebApp.MainButton.text = "Next"
        this.name = this.card.name_en
        this.desc = this.card.desc_en
        this.facts = this.card.facts_en
      }
    },
    // Emits action to go to the next card in the PlayGame component
    emitNext() {
      this.$emit("emit-next")
    },
  }
}
</script>

<style>

body {
  scroll-behavior: smooth;
}
hr {
  border: 1px solid #433789;
}

h1 {
  color: var(--tg-theme-text-color);
  background-color: var(--tg-theme-secondary-bg-color);
}

.card {
  background: v-bind(background);
  margin: 0 60px 20px;
  border-radius: 0 0 25px 25px;
  overflow: auto;
  box-shadow: 15px 15px 30px rgba(0, 0, 0, .3);
}
.screenshot {
  margin-top: 20px;
  max-width: 250px;
  max-height: 250px;
}

.name {
  color: v-bind(textColor);
  text-align: left;
  margin: 15px;
}

.name-ru {
  color: #433789;
  text-align: left;
  margin: 15px;
}

.desc {
  color: v-bind(textColor);
  font-weight: normal;
  font-size: 18px;
  margin: 15px;
  text-align: start;
}

.facts-label {
  color: #ffffff;
  background-color: #433789;
  font-size: 20px;
}

.fact {
  text-align: start;
  color: v-bind(textColor);
  font-weight: normal;
  font-size: 18px;
  margin: 15px;
}

.stars-block {
  display: flex;
  align-items: center;
  justify-content: center;
}

.starSVG {
  margin: 0 10px 0 10px;
}

.points-label {
  font-size: 20px;
  color: var(--tg-theme-hint-color);
}

</style>