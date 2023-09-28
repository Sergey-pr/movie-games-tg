<template>
  <div class="card" :style="{background: 'linear-gradient(45deg, ' + card.bg_color_1 + ', ' + card.bg_color_2 + ')'}">
    <h1 class="category">{{ category }}</h1>
    <img class="drawing" :src="imgUrlPrefix + card.drawing_id">
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
      this.setPopup();
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
    showPopup() {
      window.Telegram.WebApp.showPopup({title: "title", message: "message", buttons: [
          {
            "id": 1,
            "text": "btn1"
          },
          {
            "id": 2,
            "text": "btn2"
          },
          {
            "id": 3,
            "text": "btn3"
          }
        ]}, this.callback)
    },
    callback(id) {
      console.log(id)

    }
  }
}
</script>

<style scoped>
.category {
  color: var(--tg-theme-text-color);
  background-color: var(--tg-theme-secondary-bg-color);
}

.card {
  position: absolute;
  top: 20px;
  right: 10px;
  left: 10px;
  border-radius: 25px;
  height: 1200px;
}

.drawing {
  max-width: 300px;
  max-height: 300px;
}
</style>