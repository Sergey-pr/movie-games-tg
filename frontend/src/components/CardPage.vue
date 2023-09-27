<template>
  <div>CardPage: {{ card.name_en }}</div>
</template>

<script>

// import {useCards} from "@/services/adapter";


export default {
  name: 'CardPage',
  components: {},
  props: {
    card: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
    }
  },
  mounted() {
    this.init();
  },
  methods: {
    async init() {
      console.log(this.card)
    }
  }
}
</script>

<style scoped>

</style>