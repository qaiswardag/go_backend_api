<script setup>
import FullWidthElement from '@/Components/Layouts/FullWidthElement.vue';
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

const handleGetJobs = async function () {
  try {
    await handleData(
      `http://localhost:7070`,
      {
        headers: {
          'Accept-Version': 'v1',
          Authorization: 'hello world',
        },
        method: 'POST',
        body: JSON.stringify({
          email: email.value,
          password: password.value,
        }),
      },
      { additionalCallTime: 2000 }
    );
  } catch (error) {
    console.log(`error:`, error);
  }
};
</script>

<template>
  <div>
    <FullWidthElement :descriptionArea="true">
      <template #title>Jobs</template>

      <template #content>
        <div>
          <ul class="flex flex-col gap-8">
            <li
              class="rounded bg-red-200 overflow-hidden whitespace-pre-line flex-1 h-auto px-4 py-12"
            >
              <div>
                <p class="myPrimaryParagraph">Title here</p>
              </div>
            </li>
            <li
              class="rounded bg-red-200 overflow-hidden whitespace-pre-line flex-1 h-auto px-4 py-12"
            >
              <div>
                <p class="myPrimaryParagraph">Title here</p>
              </div>
            </li>
          </ul>
        </div>
      </template>
    </FullWidthElement>
  </div>
</template>
