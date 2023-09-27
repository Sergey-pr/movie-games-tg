<template>
  <div>PlayGame</div>
  <CardPage :card="currentCard"></CardPage>
</template>

<script>

import {useCards} from "@/services/adapter";
import CardPage from "@/components/CardPage.vue";

export default {
  name: 'PlayGame',
  components: {CardPage},
  data() {
    return {
      cards: [],
      currentCard: {}
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
    },
    onClickBack() {
      this.$router.push('/')
    }
  }
}
</script>

<style scoped>

</style>