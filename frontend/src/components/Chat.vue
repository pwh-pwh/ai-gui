<script setup lang="ts">

import MsgItem from "./MsgItem.vue";
import {inject, nextTick, onMounted, reactive, ref, toRaw, watch} from "vue";
import ScrollPanel from "primevue/scrollpanel";
import {BrowserOpenURL, Quit, WindowGetSize, WindowSetSize} from "../../wailsjs/runtime";
import {DoChat, GetStatus, OpenConfigFolder, ReloadConf} from "../../wailsjs/go/main/App";
import {types} from "../../wailsjs/go/models";
import Message = types.Message;


const inputMsg = ref('')
const helpMsg = ref('/help: 显示帮助信息\n/clear: 清除历史信息\n/loadconf: 重新加载配置文件\n/opconf: 打开配置文件夹\n' +
    '/quit: 退出应用')
const msgList = reactive<Message[]>([{content: '欢迎使用Ai-Gui,用法介绍如下', userType: 'user', id: 'tool'},
  {content: helpMsg.value, userType: 'assistant', id: 'tool'}])
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

const dispatchMsg = (msg: string): boolean => {
  switch (msg) {
    case '/clear':
      msgList.length = 0
      return false
    case '/help':
      msgList.push({content: inputMsg.value, userType: 'user', id: 'tool'})
      msgList.push({content: helpMsg.value, userType: 'assistant', id: 'tool'})
      return false
    case '/loadconf':
      msgList.push({content: inputMsg.value, userType: 'user', id: 'tool'})
      ReloadConf().then(res => {
        msgList.push({content: res, userType: 'assistant', id: 'tool'})
      })
      return false
    case '/opconf':
      msgList.push({content: inputMsg.value, userType: 'user', id: 'tool'})
        OpenConfigFolder().then(() => {
          showInfoToast('打开配置文件夹成功', '提示', 3000)
        })
      return false
    case '/quit':
      quitApp()
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

const showInfoToast = inject<any>('showInfoToast')

const shrink = inject<any>('shrink')


watch(msgList, () => {
  scrollToBottom()
})

onMounted(() => {
  scrollToBottom()
})

</script>

<template>
  <div class="flex text-center flex-column h-full">
    <div class="flex-shrink-0 text-center cursor-move select-none" style="--wails-draggable:drag">
      <div
          class="flex-row flex justify-content-between fadeindown animation-delay-800 animation-iteration-1 animation-ease-in">
        <div>
          <i @click="shrink"
             class="pi pi-arrow-down-left-and-arrow-up-right-to-center cursor-pointer text-gray-800"
             style="font-size: 1.5rem"></i>
        </div>
        <div class="cool-title text-3xl font-bold cursor-move select-none" @dblclick="quitApp">
          Ai-Gui
        </div>
        <div>
          <i class="pi pi-github cursor-pointer text-gray-800" style="font-size: 1.5rem" @click="toMyGithub"></i>
        </div>
      </div>

    </div>

    <div class="flex-1 overflow-auto border-round p-2 grow">
      <ScrollPanel ref="scrollPanelRef" class="w-full h-full overflow-x-hidden">
        <MsgItem v-for="(item,idx) in msgList" :message="item" :key="idx"/>
      </ScrollPanel>
    </div>

    <div
        class="flex-grow flex flex-column mt-2 fadeoutdown animation-delay-800 animation-iteration-1 animation-ease-in">
      <div class="flex-shrink-0 text-center flex w-full">
        <InputText type="text" class="flex-grow-1 bg-white-alpha-10" v-model="inputMsg" @keyup.enter="doChat"/>
        <Button icon="pi pi-send" class="ml-2" @click="doChat"/>
      </div>

    </div>
  </div>
</template>

<style scoped>

.grow {
  animation: grow 0.8s linear;
}

@keyframes grow {
  0% {
    transform: scale(0.2) rotate(0deg);
    opacity: 0.3;
  }
  100% {
    transform: scale(1) rotate(360deg);
    opacity: 1;
  }
}
</style>