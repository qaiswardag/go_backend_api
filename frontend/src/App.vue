<script setup>
import Navbar from '@/Components/Homepage/Navbar.vue';
import Footer from '@/Components/Homepage/Footer.vue';
import { useRouter } from 'vue-router';
import { getCookie } from '@/composables/getCookie';

const router = useRouter();

// Get cookies
// router.beforeEach((to, from, next) => {
//   console.log('Session Token:', getCookie('session_token'));
//   console.log('CSRF Token:', getCookie('csrf_token'));
//   next();
// });
</script>

<template>
  <div class="bg-orange-200">
    <Navbar></Navbar>
    <div class="mx-auto min-h-screen">
      <router-view></router-view>
    </div>
    <Footer></Footer>
  </div>
</template>
