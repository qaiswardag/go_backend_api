<script setup>
import { RouterLink } from 'vue-router';
</script>

<template>
  <nav
    class="bg-slate-200 px-6 py-6 mx-auto flex items-center justify-between lg:h-[10vh] h-[16vh]"
    aria-label="Global"
  >
    <RouterLink :to="{ name: 'Home' }">
      <img
        class="h-6"
        src="/logo/logo.svg"
        alt="Logo"
      />
    </RouterLink>

    <div class="flex lg:gap-x-8 myPrimaryGap">
      <p
        class="text-xs font-medium text-myPrimaryDarkGrayColor flex gap-2 items-center cursor-pointer"
      >
        <RouterLink
          class="text-myPrimaryDarkGrayColor"
          :to="{ name: 'Home' }"
        >
          Home
        </RouterLink>
      </p>
      <p
        class="text-xs font-medium text-myPrimaryDarkGrayColor flex gap-2 items-center cursor-pointer"
      >
        <RouterLink
          class="text-myPrimaryDarkGrayColor"
          :to="{ name: 'Jobs' }"
        >
          Jobs
        </RouterLink>
      </p>
      <p
        class="text-xs font-medium text-myPrimaryDarkGrayColor flex gap-2 items-center cursor-pointer"
      >
        <RouterLink
          class="text-myPrimaryDarkGrayColor"
          :to="{ name: 'Dashboard' }"
        >
          Dashboard
        </RouterLink>
      </p>
      <p
        class="text-xs font-medium text-myPrimaryDarkGrayColor flex gap-2 items-center cursor-pointer"
      >
        <RouterLink
          class="text-myPrimaryDarkGrayColor"
          :to="{ name: 'Login' }"
        >
          Login
        </RouterLink>
      </p>
      <p
        class="text-xs font-medium text-myPrimaryDarkGrayColor flex gap-2 items-center cursor-pointer"
      >
        <RouterLink
          class="text-myPrimaryDarkGrayColor"
          :to="{ name: 'Register' }"
        >
          Sign up
        </RouterLink>
      </p>
    </div>
  </nav>
</template>
