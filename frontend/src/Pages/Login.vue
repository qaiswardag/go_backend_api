<script setup>
import FullWidthElement from '@/Components/Layouts/FullWidthElement.vue';
</script>

<template>
  <div>
    <FullWidthElement
      :descriptionArea="true"
      class="bg-gray-50"
    >
      <template #title>Not found </template>
      <template #description>
        <p class="myPrimaryParagraph font-normal">Page not fount.</p>
      </template>
      <template #content> Lorem ipsum dolor sit amet. </template>
    </FullWidthElement>
  </div>
</template>
