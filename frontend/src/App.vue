<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import {useAuth} from "@/services/adapter";

export default {
  name: 'App',
  components: {
    HelloWorld
  },
  mounted() {
    window.Telegram.WebApp.ready()
    let initData = window.Telegram.WebApp.initData
    let body = JSON.parse('{"' + initData.replace(/&/g, '","').replace(/=/g,'":"') + '"}', function(key, value) { return key===""?value:decodeURIComponent(value) });
    body.user = JSON.parse(body.user)
    useAuth().login(body)
  },
  methods: {
  }

}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
