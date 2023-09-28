<template>
  <div class="card" :style="{background: 'linear-gradient(45deg, ' + card.bg_color_1 + ', ' + card.bg_color_2 + ')'}">
    <img alt="screenshot" class="screenshot" :src="imgUrlPrefix + card.screenshot_id">
    <h2 class="name" :style="{color: card.text_color}">{{ name }}</h2>
    <h3 class="name-ru" :style="{color: card.text_color}">{{ name_ru }}</h3>
    <p class="desc" :style="{color: card.text_color}">{{ desc }}</p>
    <h1>{{ factsLabel }}</h1>
    <p class="fact" :key="index" v-for="(item, index) in facts" :style="{color: card.text_color}">{{ item }}</p>
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
      factsLabel: "Interesting facts",
      language: "en",
      user: {},
      facts: [],
      loaded: false,
      name: "",
      name_ru: "",
      desc: "",
      state: 0,
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
      window.Telegram.WebApp.MainButton.show()
      window.Telegram.WebApp.onEvent('mainButtonClicked', this.emitNext)
      this.setLanguage();
    },
    setLanguage() {
      this.language = this.user.language
      if (this.language === "ru") {
        window.Telegram.WebApp.MainButton.text = "Дальше"
        this.factsLabel = "Интересные факты"
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
  position: absolute;
  top: 20px;
  right: 10px;
  left: 10px;
  margin-bottom: 20px;
  border-radius: 25px;
}
.screenshot {
  margin-top: 20px;
  max-width: 250px;
  max-height: 250px;
}

.name {
  color: white;
}

.name-ru {
  color: white;
}

.desc {
  color: white;
}

.fact {
  color: white;
}

body {
  scroll-behavior: smooth;
}

</style>