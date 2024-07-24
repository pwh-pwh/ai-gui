<script setup lang="ts">
import {computed} from "vue";
import {Message} from "../types/Message";
const props = defineProps<{ message: Message }>()
const positionClass = computed(() => {
  if (props.message.userType === "assistant") {
    return "flex-row"
  } else {
    return "flex-row-reverse"
  }
})
</script>

<template>
  <div :class="['flex align-items-start mb-4 px-3', positionClass]">
    <div class="w-3rem h-3rem flex align-items-center justify-content-center p-1 border-circle mx-2 bg-blue-500">
      <i class="pi pi-user"></i>
    </div>
    <div class="bg-black-alpha-20 border-round px-1 w-full">
      <p class="text-gray-800 text-left pl-2 text-xl white-space-normal">
        {{ message.content }}
      </p>
    </div>
  </div>
</template>

<style scoped>

</style>