<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <div class="leaderboard-card" v-if="loaded">
    <p class="leaderboard-desc">{{ desc }}</p>
    <p class="leaderboard-label">{{ leaderboardLabel }}</p>
    <p class="leaderboard-text" :key="index" v-for="(data, index) in usersData">
      {{ index + 1 }}. {{ data.name }}  {{ data.last_name }}:&nbsp;{{ data.points }}
    </p>
  </div>
</template>

<script>
import LoadingComponent from "@/components/LoadingComponent.vue";
import {privateApi} from "@/services/api";

export default {
  name: 'LeaderboardPage',
  components: {
    LoadingComponent,
  },
  data() {
    return {
      loaded: false,
      usersData: [],
      leaderboardLabel: "",
      desc: ""
    }
  },
  mounted() {
    this.init();
    window.scrollTo({ top: 0, behavior: 'smooth' });
    this.loaded = true
  },
  methods: {
    async init() {
      window.Telegram.WebApp.BackButton.show()
      window.Telegram.WebApp.onEvent('backButtonClicked', this.onClickBack)
      window.Telegram.WebApp.MainButton.hide()
      let response = await privateApi().getLeaderboards(this.$store.state.jwt);
      if (response.status !== 200) {
        this.loaded = false
      }
      this.usersData = response.data
      this.setLanguage();
    },
    // Sets translation
    setLanguage() {
      if (this.$store.state.user.language === "ru") {
        this.leaderboardLabel = "Лучшие Игроки"
        this.desc = "В лучший счёт засчитывается только первый ответ на карточку, " +
            "чтобы нельзя было сразу переиграть и на всё ответить правильно."
      } else {
        this.leaderboardLabel = "High Scores"
        this.desc = "To the high scores goes only your first answer to card. " +
            "It is done to prevent learning all the answers for perfect points."
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

body {
  background-color: var(--tg-theme-bg-color);
}

.loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 10px;
}

.leaderboard-desc {
  color: #ffffff;
  font-weight: normal;
  font-size: 18px;
  margin: 15px;
  text-align: center;
}

.leaderboard-text {
  margin: 10px;
  color: white;
}

.leaderboard-label {
  color: #ffffff;
  background-color: #433789;
  font-weight: normal;
  font-size: 26px;
  padding: 10px;
}

.leaderboard-card {
  margin: 0 60px 20px;
  border-radius: 0 0 25px 25px;
  overflow: auto;
  box-shadow: 15px 15px 30px rgba(0, 0, 0, .3);
  background: linear-gradient(#b96977, #936e4f);
}

</style>
