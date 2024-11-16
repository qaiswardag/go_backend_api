import './css/app.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import Home from './Pages/Home.vue';
import Login from './Pages/Login.vue';
import Register from './Pages/Register.vue';
import Dashboard from './Pages/Dashboard.vue';
import NotFound from './Pages/NotFound.vue';
import Jobs from './Pages/Jobs.vue';
import { createWebHistory, createRouter } from 'vue-router';

const newRoutes = [
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound },
  { path: '/', name: 'Home', component: Home },
  { path: '/login', name: 'Login', component: Login },
  { path: '/register', name: 'Register', component: Register },
  { path: '/dashboard', name: 'Dashboard', component: Dashboard },
  { path: '/jobs', name: 'Jobs', component: Jobs },
];

const router = createRouter({
  history: createWebHistory(),
  routes: newRoutes,
});

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.mount('#app');
