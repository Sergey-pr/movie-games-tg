<template>
  <router-view></router-view>
</template>

<script>
export default {
  name: 'App',
  components: {},
  mounted() {
    window.Telegram.WebApp.enableClosingConfirmation()
  }
}
</script>


<style>

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: var(--tg-theme-button-color);
}

</style>
