<script setup>
import FullWidthElement from '@/Components/Layouts/FullWidthElement.vue';
import { vueFetch } from '@/composables/vueFetch';
import { ca } from 'date-fns/locale';
import { onMounted, ref } from 'vue';

const {
  handleData,
  fetchedData,
  isError,
  error,
  errors,
  isLoading,
  isSuccess,
} = vueFetch();

const {
  handleData: handleDataGet,
  fetchedData: fetchedDataGet,
  isError: isErrorGet,
  error: errorGet,
  errors: errorsGet,
  isLoading: isLoadingGet,
  isSuccess: isSuccessGet,
} = vueFetch();

const jobTitle = ref('');

const handlePostJob = async function () {
  console.log(`jobTitle is: ${jobTitle.value}`);
  console.log(`Job content here: ...`);
  return;
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
          title: jobTitle.value,
          content: 'content here',
        }),
      },
      { additionalCallTime: 2000 }
    );
  } catch (error) {
    console.log(`error:`, error);
  }
};

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
const handleSubmit = async function () {
  try {
    const data = await handleDataGet(
      `http://localhost:5555/user/settings`,
      {
        credentials: 'include',
      },
      {
        additionalCallTime: 1000,
      }
    );
  } catch (error) {
    console.log(`error:`, error);
  }
};
</script>

<template>
  <div>
    <FullWidthElement :descriptionArea="true">
      <template #title>Dashboard</template>

      <template #content>
        <div class="flex gap-8 w-full justify-center">
          <!-- Data for logged in users # start -->
          <div
            class="w-1/2 bg-gray-300 border-2 border-gray-600 py-8 px-4 rounded-lg"
          >
            <h2
              class="mt-6 text-center text-2xl/9 font-bold tracking-tight text-gray-900"
            >
              List of data only for logged in users
            </h2>
            <div class="my-4">
              <button
                type="button"
                :disabled="isLoadingGet"
                @click="handleSubmit"
                :class="{
                  'opacity-25 cursor-default': isLoadingGet,
                }"
                class="myPrimaryButton w-full"
              >
                <template v-if="!isLoadingGet">
                  <span> Submit </span>
                </template>
                <template v-if="isLoadingGet">Loading.. </template>
              </button>

              <p class="myPrimaryParagraph my-6">
                type of fetchedDataGet: {{ typeof fetchedDataGet }}
                <br />
                fetchedDataGet: {{ JSON.stringify(fetchedDataGet) }}
              </p>
              <p class="myPrimaryParagraph my-6">
                type of error:
                {{ typeof errorGet }}
                <br />
                error: {{ JSON.stringify(errorGet) }}
              </p>
            </div>
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
          <!-- Data for logged in users # end -->
          <!-- Form # start -->
          <div
            class="w-1/2 bg-gray-300 border-2 border-gray-600 py-8 px-4 rounded-lg"
          >
            <h2
              class="mt-6 text-center text-2xl/9 font-bold tracking-tight text-gray-900"
            >
              Post a new job
            </h2>
            <div>
              <div>
                <form class="space-y-6">
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
        </div>
      </template>
    </FullWidthElement>
  </div>
</template>
