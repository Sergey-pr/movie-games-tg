<template>
  <div class="card">
    <h1 class="category">{{ category }}</h1>
    <img :src="imgUrlPrefix + card.drawing_id">
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
      user: {},
      loaded: false
    }
  },
  created() {
    this.imgUrlPrefix = process.env.VUE_APP_BASE_URL + "/api/public/bot-image/";
  },
  mounted() {
    // this.init();
  },
  methods: {
    async init() {
      this.user = this.$store.state.user
      this.setLanguage();
    },
    setLanguage() {
      this.language = this.user.language
      if (this.language === "ru") {
        window.Telegram.WebApp.MainButton.text = "Ответить"
        switch (this.card.category) {
          case "movie":
            this.category = "Фильм"
        }
      } else {
        window.Telegram.WebApp.MainButton.text = "Answer"
        switch (this.card.category) {
          case "movie":
            this.category = "Movie"
        }
      }
    },
  }
}
</script>

<style scoped>
.category {
  color: var(--tg-theme-text-color);
}

.card {
  background: linear-gradient(45deg, #6262e0, #bb4d4d);
  position: absolute;
  top: 20px;
  right: 10px;
  left: 10px;
  border-radius: 25px;
  height: 1200px;
}
</style>