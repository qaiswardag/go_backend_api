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

const jobTitle = ref('');

const handlePostJob = async function () {
  console.log(`jobTitle is: ${jobTitle.value}`);
  console.log(`Job content here: ...`);
  return;
  await handleData(
    `https://www.google.com`,
    {
      headers: {
        'Accept-Version': 'v1',
        Authorization: 'hello world',
      },
      method: 'POST',
      body: JSON.stringify({
        title: jobTitle.value,
        content: 'content here',
      }),
    },
    { additionalCallTime: 2000 }
  );
};
</script>

<template>
  <div>
    <FullWidthElement :descriptionArea="true">
      <template #title>Dashboard </template>

      <template #content>
        <!-- Form # start -->
        <div
          class="flex min-h-full flex-1 flex-col justify-center py-12 sm:px-6 lg:px-8"
        >
          <div class="sm:mx-auto sm:w-full sm:max-w-md">
            <h2
              class="mt-6 text-center text-2xl/9 font-bold tracking-tight text-gray-900"
            >
              Post a new job
            </h2>
          </div>
          <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-7xl bg-red-50">
            <div class="px-6 py-12 shadow sm:rounded-lg sm:px-12">
              <form
                class="space-y-6"
                method="POST"
              >
                <div>
                  <label
                    for="jobTitle"
                    class="myPrimaryInputLabel"
                    >Content</label
                  >
                  <div class="mt-2">
                    <input
                      v-model="jobTitle"
                      id="jobTitle"
                      name="jobTitle"
                      type="jobTitle"
                      class="myPrimaryInput"
                    />
                  </div>
                </div>

                <div>
                  <button
                    type="button"
                    @click="handlePostJob"
                    class="myPrimaryButton w-full"
                  >
                    Submit
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
