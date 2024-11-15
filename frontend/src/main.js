import './css/app.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import Home from './Pages/Home.vue';
import Login from './Pages/Login.vue';
import NotFound from './Pages/NotFound.vue';
import { createWebHistory, createRouter } from 'vue-router';

const newRoutes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/login', name: 'Login', component: Login },
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound },
];

const router = createRouter({
  history: createWebHistory(),
  routes: newRoutes,
});

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.mount('#app');
