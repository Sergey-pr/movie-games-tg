<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <CardPage v-if="loaded && !playedState" :card="currentCard" @emit-win="played"></CardPage>
  <CardInfoPage v-if="loaded && playedState" :card="currentCard" :points="points" @emit-next="next"></CardInfoPage>
</template>

<script>

import {useCards} from "@/services/adapter";
import CardPage from "@/components/CardPage.vue";
import LoadingComponent from "@/components/LoadingComponent.vue";
import CardInfoPage from "@/components/CardInfoPage.vue";

export default {
  name: 'PlayGame',
  components: {CardInfoPage, LoadingComponent, CardPage},
  data() {
    return {
      playedState: false,
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
      this.currentCard = this.cards[this.currentCardIndex]
      this.loaded = true;
    },
    onClickBack() {
      this.$router.push('/')
    },
    played(points) {
      this.points += points
      this.playedState = true
    },
    next() {
      this.playedState = false
      this.currentCardIndex += 1
      if (this.currentCardIndex >= this.cards.length) {
        // The end
        this.currentCardIndex = 0
      }
      this.currentCard = this.cards[this.currentCardIndex]
    }
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