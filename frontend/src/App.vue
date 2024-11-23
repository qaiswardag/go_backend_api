<script setup>
import Navbar from '@/Components/Homepage/Navbar.vue';
import Footer from '@/Components/Homepage/Footer.vue';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const getCookie = function (name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) {
    return parts.pop().split(';').shift();
  }
  return name + ' not found.';
};

// Get csrf_token cookie

// Get csrf_token cookie on route change
router.beforeEach((to, from, next) => {
  const session_token = getCookie('session_token');
  console.log('session Token:', csrfToken);

  const csrfToken = getCookie('csrf_token');
  console.log('CSRF Token:', csrfToken);
  next();
});
</script>

<template>
  <div>
    <Navbar></Navbar>
    <div class="mx-auto min-h-screen">
      <router-view></router-view>
    </div>
    <Footer></Footer>
  </div>
</template>
