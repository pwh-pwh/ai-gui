<script setup lang="ts">

import MsgItem from "./MsgItem.vue";
import {nextTick, onMounted, reactive, ref, watch} from "vue";
import {Message} from "../types/Message";
import ScrollPanel from "primevue/scrollpanel";

const msgList = reactive<Message[]>([{content: 'hello', userType: 'user', id: '1'},
  {content: 'how are you', userType: 'assistant', id: '2'}])
const inputMsg = ref('')
const doChat = () => {
  //生成id
  const id = Date.now().toString(16)
  msgList.push({content: inputMsg.value, userType: 'user', id: id})
}
const scrollPanelRef = ref()
const scrollToBottom = () => {
  nextTick(() => {
    scrollPanelRef.value.$refs.content.scrollTop = scrollPanelRef.value.$refs.content.scrollHeight
  })
}
watch(msgList, () => {
  scrollToBottom()
})

onMounted(() => {
  scrollToBottom()
})

</script>

<template>
  <div class="flex text-center flex-column h-full">
    <div class="flex-shrink-0 text-center">
      <div class="text-white-alpha-80 text-3xl font-bold">
        Ai-Gui
      </div>

    </div>

    <div class="flex-1 overflow-auto border-2 border-white-alpha-40 border-round p-2">
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
</template>

<style scoped>

</style>