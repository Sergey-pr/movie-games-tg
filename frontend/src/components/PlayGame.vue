<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <CardPage v-if="loaded && state === 'play'" :card="currentCard" @emit-win="played"></CardPage>
  <CardInfoPage v-if="loaded && state === 'info'" :card="currentCard" :points="points" @emit-next="next"></CardInfoPage>
  <TheEndPage v-if="loaded && state === 'end'" :points="points"></TheEndPage>
</template>

<script>

import {useCards, useUsers} from "@/services/adapter";
import CardPage from "@/components/CardPage.vue";
import LoadingComponent from "@/components/LoadingComponent.vue";
import CardInfoPage from "@/components/CardInfoPage.vue";
import TheEndPage from "@/components/TheEndPage.vue";

export default {
  name: 'PlayGame',
  components: {CardInfoPage, LoadingComponent, CardPage, TheEndPage},
  data() {
    return {
      state: "play", // 3 states "play", "info", "end"
      cards: [],
      currentCard: {},
      currentCardIndex: 0,
      loaded: false,
      points: 0,
    }
  },
  mounted() {
    this.init();
  },
  methods: {
    async init() {
      window.Telegram.WebApp.onEvent('backButtonClicked', this.onClickBack)
      window.Telegram.WebApp.BackButton.show()

      let response = await useCards().cardsList(this.$store.state.jwt);
      this.cards = response.data
      this.shuffleArray(this.cards)
      this.currentCard = this.cards[this.currentCardIndex]
      this.loaded = true;
    },
    onClickBack() {
      this.$router.push('/')
    },
    played(points) {
      useUsers().processAnswer(this.$store.state.jwt, points, this.currentCard.id);
      this.points += points
      this.state = "info"
      window.scrollTo({ top: 0, behavior: 'smooth' });
    },
    next() {
      this.state = "play"
      this.currentCardIndex += 1
      if (this.currentCardIndex >= this.cards.length) {
        this.state = "end"
      } else {
        this.currentCard = this.cards[this.currentCardIndex]
      }
      window.scrollTo({ top: 0, behavior: 'smooth' });
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

<style scoped>
.loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 10px;
}
</style>