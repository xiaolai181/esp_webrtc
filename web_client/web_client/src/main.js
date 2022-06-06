import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from '@/plugins/axiosInstance.js'
import './mock/'


import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"

const app =createApp(App)
app.config.globalProperties.$axios=axios
app.use(store).use(router).mount('#app')
