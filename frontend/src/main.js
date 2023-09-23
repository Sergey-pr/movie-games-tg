import { createApp } from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config';
import Card from 'primevue/card';

const app = createApp(App)
app.use(PrimeVue);
app.component('PrimeCard', Card);

app.mount('#app')