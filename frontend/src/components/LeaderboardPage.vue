<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <div class="leaderboard-card" v-if="loaded">
    <h1 id="leaderboard-label">{{ leaderboardLabel }}</h1>
    <p class="leaderboard-text" :key="index" v-for="(data, index) in usersData">
      {{ data.name }}:&nbsp;{{ data.points }}
    </p>
  </div>
</template>

<script>
import LoadingComponent from "@/components/LoadingComponent.vue";
import {useUsers} from "@/services/adapter";

export default {
  name: 'LeaderboardPage',
  components: {
    LoadingComponent,
  },
  data() {
    return {
      user: {},
      language: "en",
      loaded: false,
      usersData: [],
      leaderboardLabel: "",
    }
  },
  mounted() {
    this.init();
    window.scrollTo({ top: 0, behavior: 'smooth' });
    this.loaded = true
  },
  methods: {
    async init() {
      this.user = this.$store.state.user
      window.Telegram.WebApp.BackButton.show()
      window.Telegram.WebApp.onEvent('backButtonClicked', this.onClickBack)
      window.Telegram.WebApp.MainButton.hide()
      let response = await useUsers().getLeaderboards(this.$store.state.jwt);
      this.usersData = response.data
      this.setLanguage();
    },
    setLanguage() {
      this.language = this.user.language
      if (this.language === "ru") {
        this.leaderboardLabel = "Лучшие Игроки"
             } else {
        this.leaderboardLabel = "High Scores"}
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

.leaderboard-text {
  margin: 10px;
  color: white;
}

#leaderboard-label {
  margin: 10px;
  color: white;
  background: transparent;
}

body {
  background-color: var(--tg-theme-bg-color);
}

.leaderboard-card {
  margin: 0 60px 20px;
  border-radius: 0 0 25px 25px;
  overflow: auto;
  box-shadow: 15px 15px 30px rgba(0, 0, 0, .3);
  background: linear-gradient(#b96977, #936e4f);
}

</style>
