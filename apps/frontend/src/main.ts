import '@/assets/scss/main.scss';

import { createApp, watch } from 'vue'
import { createPinia, storeToRefs } from 'pinia'

import { useUserStore } from '@/stores/userStore'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

const userStore = useUserStore()

watch(() => userStore.isLoggedIn, v => {
    router.push(v ? "/" : "/login")
})

app.mount('#app')
