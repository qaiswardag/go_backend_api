<script setup>
import Navbar from '@/Components/Homepage/Navbar.vue';
import Footer from '@/Components/Homepage/Footer.vue';
</script>

<template>
  <Navbar></Navbar>
  <div class="myPrimaryWidthScreenModule">
    <router-view></router-view>
  </div>
  <Footer></Footer>
</template>
