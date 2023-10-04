<template>
  <div class="card">
    <img alt="screenshot" class="screenshot" :src="imgUrlPrefix + card.screenshot_id">
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
    <p class="points-label">{{ pointsText }}: 30</p>
    <svg class="starSVG" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" width="30" height="30">
      <polygon fill=#FFD024 points="15 0.001 20.25 8.987 30 11.46 23.495 19.487 24.27 30.001 15 25.976 5.73 30.001 6.505 19.487 0 11.46 9.75 8.987" />
    </svg>
  </div>
</template>

<script>

export default {
  name: 'CardInfoPage',
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
      factsLabel: "INTERESTING FACTS",
      language: "en",
      imgUrlPrefix: "",
      textColor: "",
      user: {},
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
    if (process.env.VUE_APP_BASE_URL !== "") {
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
      window.Telegram.WebApp.MainButton.show()
      window.Telegram.WebApp.onEvent('mainButtonClicked', this.emitNext)
      this.background = 'linear-gradient(45deg, ' + this.card.bg_color_1 + ', ' + this.card.bg_color_2 + ')'
      this.textColor = this.card.text_color
      this.setLanguage();
    },
    setLanguage() {
      this.language = this.user.language
      if (this.language === "ru") {
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
    emitNext() {
      this.$emit("emit-next")
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
  text-align: justify;
}

.facts-label {
  color: #ffffff;
  background-color: #433789;
  font-size: 20px;
}

.fact {
  text-align: justify;
  color: v-bind(textColor);
  font-weight: normal;
  font-size: 18px;
  margin: 15px;
}

body {
  scroll-behavior: smooth;
}
hr {
  border: 1px solid #433789;
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