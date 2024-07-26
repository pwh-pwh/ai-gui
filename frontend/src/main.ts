import {createApp} from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config'
import Aura from '@primevue/themes/aura'
import Tooltip from "primevue/tooltip";
import 'primeflex/primeflex.css';
import './style.css';
import 'primeicons/primeicons.css'
import ToastService from 'primevue/toastservice';


const app = createApp(App)
app.use(PrimeVue, {
    theme: {
        preset: Aura
    }
})
app.directive('tooltip', Tooltip)
app.use(ToastService)
app.mount('#app')
