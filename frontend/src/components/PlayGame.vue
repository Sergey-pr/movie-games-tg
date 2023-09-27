<template>
  <LoadingComponent class="loading" v-if="!loaded"></LoadingComponent>
  <CardPage v-if="loaded" :card="currentCard"></CardPage>
</template>

<script>

import {useCards} from "@/services/adapter";
import CardPage from "@/components/CardPage.vue";
import LoadingComponent from "@/components/LoadingComponent.vue";

export default {
  name: 'PlayGame',
  components: {LoadingComponent, CardPage},
  data() {
    return {
      cards: [],
      currentCard: {},
      loaded: false
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
      this.currentCard = this.cards[0]
      this.loaded = true;
    },
    onClickBack() {
      this.$router.push('/')
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