<script setup>
import FullWidthElement from '@/Components/Layouts/FullWidthElement.vue';
import { clearCookie } from '@/composables/clearCookie';
import { getCookie } from '@/composables/getCookie';
import { vueFetch } from '@/composables/vueFetch';
import { ref } from 'vue';

const {
  handleData,
  fetchedData,
  isError,
  error,
  errors,
  isLoading,
  isSuccess,
} = vueFetch();

const email = ref('qais.wardag@outlook.com');
const password = ref('123456');

const handleForm = async function () {
  clearCookie('session_token');
  clearCookie('csrf_token');

  try {
    const data = await handleData(
      `http://localhost:5555/login`,
      {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
          'Accept-Version': '',
          Authorization: '',
        },
        body: JSON.stringify({
          email: email.value,
          password: password.value,
        }),
      },
      {
        additionalCallTime: 1000,
      }
    );

    console.log('Session Token:', getCookie('session_token'));
    console.log('CSRF Token:', getCookie('csrf_token'));
    console.log(`data:`, data);
  } catch (error) {
    console.log(`error:`, error);
  }
};
</script>

<template>
  <div>
    <FullWidthElement :descriptionArea="true">
      <template #title>Login</template>

      <template #content>
        <!-- Form # start -->
        <div
          class="flex min-h-full flex-1 flex-col justify-center py-12 sm:px-6 lg:px-8"
        >
          <div class="sm:mx-auto sm:w-full sm:max-w-md">
            <h2
              class="mt-6 text-center text-2xl/9 font-bold tracking-tight text-gray-900"
            >
              Sign up today!
            </h2>
            <p class="myPrimaryParagraph my-6">
              type of fetchedDataGet: {{ typeof fetchedData }}
              <br />
              fetchedData: {{ JSON.stringify(fetchedData) }}
            </p>
            <p class="myPrimaryParagraph my-6">
              type of error:
              {{ typeof error }}
              <br />
              error: {{ JSON.stringify(error) }}
            </p>
          </div>

          <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-[480px]">
            <div class="bg-white px-6 py-12 shadow rounded-lg sm:px-12">
              <form
                @submit.prevent
                class="space-y-6"
              >
                <div>
                  <label
                    for="email"
                    class="myPrimaryInputLabel"
                    >Email address</label
                  >
                  <div class="mt-2">
                    <input
                      v-model="email"
                      id="email"
                      name="email"
                      type="email"
                      autocomplete="email"
                      class="myPrimaryInput"
                    />
                  </div>
                </div>

                <div>
                  <label
                    for="password"
                    class="myPrimaryInputLabel"
                    >Password</label
                  >
                  <div class="mt-2">
                    <input
                      v-model="password"
                      id="password"
                      name="password"
                      type="password"
                      class="myPrimaryInput"
                    />
                  </div>
                </div>

                <div class="flex items-center justify-between">
                  <div class="flex items-center">
                    <input
                      id="remember-me"
                      name="remember-me"
                      type="checkbox"
                      class="myPrimaryCheckbox"
                    />
                    <label
                      for="remember-me"
                      class="ml-3 block text-sm/6 text-gray-900"
                      >Remember me</label
                    >
                  </div>
                </div>

                <div>
                  <button
                    type="button"
                    :disabled="isLoading"
                    @click="handleForm"
                    :class="{
                      'opacity-25 cursor-default': isLoading,
                    }"
                    class="myPrimaryButton w-full"
                  >
                    <template v-if="!isLoading">
                      <span> Submit </span>
                    </template>
                    <template v-if="isLoading">Loading.. </template>
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
        <!-- Form # end -->
      </template>
    </FullWidthElement>
  </div>
</template>
