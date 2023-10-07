<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <CardPage v-if="loaded && state === 'play'" :card="currentCard" @emit-win="cardPlayedCallback"></CardPage>
  <CardInfoPage v-if="loaded && state === 'info'" :card="currentCard" :points="points" @emit-next="infoPlayedCallback"></CardInfoPage>
  <TheEndPage v-if="loaded && state === 'end'" :points="points"></TheEndPage>
  <p class="cards-counter" v-if="state !== 'end'">{{ cardText }}: {{ currentCardIndex + 1}}/{{cards.length}}</p>
</template>

<script>

import {privateApi} from "@/services/api";
import { shuffleArray } from "@/services/utils";
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
      cardText: "Card"
    }
  },
  mounted() {
    this.init();
  },
  methods: {
    async init() {
      // Enable back button at the top of the screen
      window.Telegram.WebApp.onEvent('backButtonClicked', this.onClickBack)
      window.Telegram.WebApp.BackButton.show()
      // Get all cards objects
      let response = await privateApi().getCards(this.$store.state.jwt);
      if (response.status !== 200) {
        this.loaded = false
      }
      this.cards = response.data
      // Shuffle cards
      shuffleArray(this.cards)
      // Set current card
      this.currentCard = this.cards[this.currentCardIndex]
      this.setLanguage()
      this.loaded = true;
    },
    // Sets translation
    setLanguage() {
      if (this.$store.state.user.language === "ru") {
        this.cardText = "Карта"
      }
    },
    // Returns to the landing page
    onClickBack() {
      this.$router.push('/')
    },
    // Sends points to backend to save result in the database, adds points, and changes state
    async cardPlayedCallback(points) {
      let response = await privateApi().processAnswer(this.$store.state.jwt, points, this.currentCard.id);
      if (response.status !== 200) {
        this.loaded = false
      }
      this.points += points
      this.state = "info"
      window.scrollTo({ top: 0, behavior: 'smooth' });
    },
    // Changes state to play or end if there are no more cards
    infoPlayedCallback() {
      this.state = "play"
      this.currentCardIndex += 1
      if (this.currentCardIndex >= this.cards.length) {
        this.state = "end"
      } else {
        this.currentCard = this.cards[this.currentCardIndex]
      }
      window.scrollTo({ top: 0, behavior: 'smooth' });
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

.cards-counter {
  color: #565473;
}
</style>