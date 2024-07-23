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
    <div class="w-2 h-12 border-circle mr-4">
      <i class="pi pi-user"></i>
    </div>
    <div class="bg-gray-200 border-round p-1 w-full">
      <p class="text-gray-800 text-left pl-2">
        {{ message.content }}
      </p>
    </div>
  </div>
</template>

<style scoped>

</style>