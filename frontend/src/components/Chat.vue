<script setup lang="ts">

import MsgItem from "./MsgItem.vue";
import {nextTick, onMounted, reactive, ref, toRaw, watch} from "vue";
import ScrollPanel from "primevue/scrollpanel";
import {BrowserOpenURL, Quit} from "../../wailsjs/runtime";
import {DoChat} from "../../wailsjs/go/main/App";
import {types} from "../../wailsjs/go/models";
import Message = types.Message;

const msgList = reactive<Message[]>([{content: 'hello', userType: 'user', id: '1'},
  {content: 'how are you', userType: 'assistant', id: '2'}])
const inputMsg = ref('')
const isShrink = ref(false)
const helpMsg = ref('/help: show this message\n/clear: clear messages\n')
const doChat = () => {
  if (!dispatchMsg(inputMsg.value)) {
    inputMsg.value = ''
    return
  }

  //生成id
  const id = Date.now().toString(16)
  msgList.push({content: inputMsg.value, userType: 'user', id: id})
  const cList = msgList.filter(msg => msg.id !== 'tool')
  msgList.push({content: '正在思考...', userType: 'assistant', id: id})
  DoChat(cList).then(res => {
    msgList[msgList.length - 1].content = res
  })
  inputMsg.value = ''
}

const dispatchMsg = (msg: string):boolean =>  {
  switch (msg) {
    case '/clear':
      msgList.length = 0
      return false
    case '/help':
      msgList.push({content: inputMsg.value, userType: 'user', id: 'tool'})
      msgList.push({content: helpMsg.value, userType: 'assistant', id: 'tool'})
      return false
    default:
      return true
  }
}

const scrollPanelRef = ref()
const scrollToBottom = () => {
  nextTick(() => {
    scrollPanelRef.value.$refs.content.scrollTop = scrollPanelRef.value.$refs.content.scrollHeight
  })
}

const quitApp = () => {
  Quit()
}

const toMyGithub = () => {
  BrowserOpenURL('https://github.com/pwh-pwh/ai-gui')
}

watch(msgList, () => {
  scrollToBottom()
})

onMounted(() => {
  scrollToBottom()
})

</script>

<template>
  <div class="flex text-center flex-column h-full" v-if="!isShrink">
    <div class="flex-shrink-0 text-center cursor-move select-none" style="--wails-draggable:drag">
      <div class="flex-row flex justify-content-between scalein animation-delay-800 animation-iteration-1 animation-ease-in">
        <div>
          <i @click="isShrink = true" class="pi pi-arrow-down-left-and-arrow-up-right-to-center cursor-pointer text-gray-800" style="font-size: 1.5rem"></i>
        </div>
        <div class="cool-title text-3xl font-bold cursor-move select-none" @dblclick="quitApp" >
          Ai-Gui
        </div>
        <div>
          <i class="pi pi-github cursor-pointer text-gray-800" style="font-size: 1.5rem" @click="toMyGithub"></i>
        </div>
      </div>

    </div>

    <div class="flex-1 overflow-auto border-round p-2">
      <ScrollPanel ref="scrollPanelRef" class="w-full h-full">
      <MsgItem v-for="(item,idx) in msgList" :message="item" :key="idx"/>
      </ScrollPanel>
    </div>

    <div class="flex-grow flex flex-column mt-2">
      <div class="flex-shrink-0 text-center flex w-full">
        <InputText type="text" class="flex-grow-1 bg-white-alpha-10" v-model="inputMsg" @keyup.enter="doChat"/>
        <Button icon="pi pi-send" class="ml-2" @click="doChat"/>
      </div>

    </div>
  </div>
  <div v-else class="absolute top-0 left-0 ml-5 mt-4 pointer-events-auto" @click="isShrink=false" style="--wails-draggable:drag">
    <img src="../assets/images/img2.png" alt="" class="w-2rem h-2rem cursor-move select-none pointer-events-none" style="--wails-draggable:drag"/>
  </div>
</template>

<style scoped>

</style>