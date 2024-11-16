<script setup>
import FullWidthElement from '@/Components/Layouts/FullWidthElement.vue';
import { ref, computed, nextTick, onBeforeMount } from 'vue';
import PageBuilderModal from '@/Components/Modals/PageBuilderModal.vue';

import PageBuilderView from '@/PageBuilder/PageBuilder.vue';
import PageBuilder from '@/composables/PageBuilder';
import FullScreenSpinner from '@/Components/Loaders/FullScreenSpinner.vue';
import { usePageBuilderStateStore } from '@/stores/page-builder-state';
import { useUserStore } from '@/stores/user';
import { useMediaLibraryStore } from '@/stores/media-library';

const features = [
  {
    name: 'Live Drag & Drop Builder',
    description:
      'Click & Drop content on a page and watch your pages come to life. Bring your vision to life and create impressive pages using a click & drop Page Builder',
  },
  {
    name: 'True Visual Editing',
    description:
      'See your changes in real-time as you make them. Elevate your creative vision and create pages using an intuitive click & drop page builder. Break free from design limitations and turn your visions into reality.',
  },
  {
    name: 'Features',
    description:
      'Click & Drop, Reordering, True Visual Editing, Responsive Editing, Font Customization, Undo & Redo, Text Editing, Media Library, Unsplash Integration, YouTube Videos.',
  },
  {
    name: 'Technologies',
    description:
      'Developed with JavaScript, Vue 3, Composition API, Pinia, CSS, Tailwind CSS and HTML.',
  },
  {
    name: 'Set Brand and Link Colors one place',
    description:
      'Global Styles for fonts, designs, & colors. Set Brand and Link Colors once and apply them across the entire Platform with ease.',
  },
  {
    name: 'Mobile-First Approach',
    description:
      'Developed with Mobile-First approach. The Page Builder even works on mobile phones.',
  },
  {
    name: 'Media Library',
    description:
      'A beautiful and user-friendly media library that allows you to change and update images. Unsplash Integration is included.',
  },
  {
    name: 'Minimal and Intuitive Design',
    description:
      'Beautiful, elegant and intuitive design. Enhance user engagement with amazing visual experience.',
  },
];

const mediaLibraryStore = useMediaLibraryStore();
const pageBuilderStateStore = usePageBuilderStateStore();
const userStore = useUserStore();
const openPageBuilder = ref(false);

const pageBuilderPrimaryHandler = ref(null);
const pageBuilderSecondaryHandler = ref(null);
const pageBuilder = new PageBuilder(pageBuilderStateStore, mediaLibraryStore);
const formType = ref('create');

const getIsLoading = computed(() => {
  return userStore.getIsLoading;
});

const pathPageBuilderStorageCreate = `page-builder-create-post`;
const pathPageBuilderStorageUpdate = `page-builder-update-post-id-1`;

const handlePageBuilder = async function () {
  userStore.setIsLoading(true);

  await nextTick();
  openPageBuilder.value = true;

  if (formType.value === 'create') {
    pageBuilderStateStore.setComponents([]);
    pageBuilder.areComponentsStoredInLocalStorage();
  }

  // handle click
  pageBuilderPrimaryHandler.value = async function () {
    userStore.setIsLoading(true);

    if (formType.value === 'update') {
      await nextTick();
      pageBuilder.saveComponentsLocalStorageUpdate();
    }

    openPageBuilder.value = false;
    userStore.setIsLoading(false);
  };

  // handle click
  pageBuilderSecondaryHandler.value = async function () {
    userStore.setIsLoading(true);

    // save to local storage if new resource
    if (formType.value === 'create') {
      await nextTick();
      pageBuilder.saveComponentsLocalStorage();
      await nextTick();
    }
    // save to local storage if update
    if (formType.value === 'update') {
      await nextTick();
      pageBuilder.synchronizeDOMAndComponents();
      await nextTick();
    }

    openPageBuilder.value = false;

    userStore.setIsLoading(false);
  };

  userStore.setIsLoading(false);

  // end modal
};
// Builder # End
const handleDraftForUpdate = async function () {
  userStore.setIsLoading(true);

  if (formType.value === 'update') {
    await nextTick();
    pageBuilder.areComponentsStoredInLocalStorageUpdate();
    await nextTick();
    pageBuilder.setEventListenersForElements();

    userStore.setIsLoading(false);
  }
};

onBeforeMount(() => {
  // Define local storage key name before on mount
  pageBuilderStateStore.setLocalStorageItemName(pathPageBuilderStorageCreate);

  // Define local storage key name before on mount
  pageBuilderStateStore.setLocalStorageItemNameUpdate(
    pathPageBuilderStorageUpdate
  );
});
</script>

<template>
  <div>
    <teleport to="body">
      <FullScreenSpinner v-if="getIsLoading"></FullScreenSpinner>
    </teleport>
    <PageBuilderModal
      :show="openPageBuilder"
      updateOrCreate="create"
      @pageBuilderPrimaryHandler="pageBuilderPrimaryHandler"
      @pageBuilderSecondaryHandler="pageBuilderSecondaryHandler"
      @handleDraftForUpdate="handleDraftForUpdate"
    >
      <PageBuilderView></PageBuilderView>
    </PageBuilderModal>

    <FullWidthElement :descriptionArea="true">
      <template #title>Home</template>
      <template #description>
        <button
          @click="handlePageBuilder"
          class="myPrimaryButton lg:text-sm text-[10px] lg:py-2 py-2 min-h-2 mt-4"
          type="button"
        >
          Start Builder
        </button>
      </template>
      <template #content>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Cumque incidunt
        illo autem corrupti voluptas, inventore nihil quae quas cum quisquam
        dolorem unde perspiciatis totam nisi, distinctio reprehenderit enim
        doloribus ea. Vitae, facilis repellat? Architecto eius maiores facere
        cupiditate quisquam? Explicabo est ipsum veniam in, blanditiis
        similique, molestiae quae maxime tempore aperiam quam officiis, nobis
        voluptatibus. Voluptatibus, maxime eum. Reprehenderit quidem facilis
        necessitatibus modi esse, unde nobis blanditiis veniam quia ipsa error
        sit quisquam cumque totam possimus voluptates dicta neque illum officiis
        rem eos voluptatibus doloribus. Consequuntur non ea quaerat animi
        aspernatur, ab iure quod exercitationem nihil autem maiores molestiae
        voluptatum obcaecati corporis est ad unde modi ullam provident ipsum?
        Alias quae voluptate aperiam deserunt quaerat saepe at ratione nulla.
        Quis, doloribus optio eveniet distinctio odit commodi! Cum officia
        fugiat rerum exercitationem, molestiae architecto vel amet, ut ullam
        vero non quo temporibus quia repudiandae ea maiores voluptatem nihil,
        maxime facere eum obcaecati. Quia harum omnis consequuntur deserunt
        dolores accusantium, neque fugiat mollitia debitis alias voluptatem fuga
        quas incidunt odio culpa dolorem adipisci hic eaque sequi saepe
        assumenda corporis iusto. Quidem beatae autem ducimus, cum impedit hic
        laboriosam nostrum quam asperiores veritatis illum expedita vero
        repellat aperiam mollitia porro! Ad, rem voluptatum.
      </template>
    </FullWidthElement>
  </div>
</template>
