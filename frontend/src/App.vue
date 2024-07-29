<script lang="ts" setup>

import Chat from "./components/Chat.vue";
import Toast from "primevue/toast";
import {useToast} from "primevue/usetoast";
import {onMounted, provide, ref} from "vue";
import {GetStatus} from "../wailsjs/go/main/App";
import {WindowGetSize, WindowSetSize} from "../wailsjs/runtime";
import {is} from "@babel/types";
const toast = useToast()
const isShrink = ref(false)
const showInfoToast = (msg:string,summary:string = 'info',duration:number = 3000) => {
  toast.add({ severity: 'info', summary: summary, detail: msg, life: duration });
}

const shrink = () => {
  isShrink.value = true
  WindowSetSize(40,40)
  WindowGetSize().then((res) => {
    console.log(res)
  })
}

const enlarge = () => {
  isShrink.value = false
  WindowSetSize(500,668)
  WindowGetSize().then((res) => {
    console.log(res)
  })
}

provide('isShrink',isShrink)
provide('shrink',shrink)
provide('showInfoToast',showInfoToast)

onMounted(() => {
  GetStatus().then((res: string) => {
    showInfoToast(res, '初始化信息', 6000)
  })
})
</script>

<template>
  <div v-if="!isShrink" class="border-round-3xl p-3 h-full">
    <Toast />
      <Chat />
  </div>
  <div v-else class="absolute top-0 left-0 pointer-events-auto overflow-hidden" @click="enlarge"
       style="--wails-draggable:drag">
    <img src="./assets/images/img2.png" alt="" class="w-2rem h-2rem cursor-move select-none pointer-events-none"
         style="--wails-draggable:drag"/>
  </div>
</template>

<style>

</style>
