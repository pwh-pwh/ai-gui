<script lang="ts" setup>

import Chat from "./components/Chat.vue";
import Toast from "primevue/toast";
import {useToast} from "primevue/usetoast";
import {provide} from "vue";
const toast = useToast()
const showInfoToast = (msg:string,summary:string = 'info',duration:number = 3000) => {
  toast.add({ severity: 'info', summary: summary, detail: msg, life: duration });
}
provide('showInfoToast',showInfoToast)
</script>

<template>
  <div class="border-round-3xl p-3 h-full">
    <Toast />
      <Chat />
  </div>
</template>

<style>

</style>
