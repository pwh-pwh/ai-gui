<script lang="ts" setup>

import Chat from "./components/Chat.vue";
import Toast from "primevue/toast";
import {useToast} from "primevue/usetoast";
import {onMounted, provide, reactive, ref} from "vue";
import {GetStatus} from "../wailsjs/go/main/App";
import {WindowGetPosition, WindowSetPosition, WindowSetSize} from "../wailsjs/runtime";
import {types} from "../wailsjs/go/models";
import Message = types.Message;

const toast = useToast()
const isShrink = ref(false)
const helpMsg = ref('/help: 显示帮助信息\n/clear: 清除历史信息\n/loadconf: 重新加载配置文件\n/opconf: 打开配置文件夹\n' +
    '/quit: 退出应用')
const msgList = reactive<Message[]>([{content: '欢迎使用Ai-Gui,用法介绍如下', userType: 'user', id: 'tool'},
  {content: helpMsg.value, userType: 'assistant', id: 'tool'}])

const showInfoToast = (msg:string,summary:string = 'info',duration:number = 3000) => {
  toast.add({ severity: 'info', summary: summary, detail: msg, life: duration });
}

const shrink = () => {
  isShrink.value = true
  WindowSetSize(40,40)
  WindowGetPosition().then((res) => {
    WindowSetPosition(-15,res.y)
  })
}

const enlarge = () => {
  isShrink.value = false
  WindowSetSize(500,668)
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
      <Chat :msgList="msgList" :helpMsg="helpMsg"/>
  </div>
  <div v-else class="absolute top-0 left-0 pointer-events-auto overflow-hidden" @click="enlarge"
       style="--wails-draggable:drag">
    <img src="./assets/images/img2.png" alt="" class="w-2rem h-2rem cursor-move select-none pointer-events-none"
         style="--wails-draggable:drag"/>
  </div>
</template>

<style>

</style>
