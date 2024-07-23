<script setup lang="ts">

import MsgItem from "./MsgItem.vue";
import {reactive, ref} from "vue";
import {Message} from "../types/Message";

const msgList = reactive<Message[]>([{content: 'hello', userType: 'user', id: '1'},
  {content: 'how are you', userType: 'assistant', id: '2'}])
const inputMsg = ref('')
</script>

<template>
  <div class="flex text-center flex-column h-full">
    <div class="flex-shrink-0 text-center">
      <div class="text-white-alpha-80 text-3xl font-bold">
        Ai-Gui
      </div>

    </div>

    <div class="flex-1 overflow-auto bg-blue-400 border-2 border-white-alpha-40 border-round p-2">
      <ScrollPanel class="w-full h-full">
      <MsgItem v-for="(item,idx) in msgList" :message="item" :key="idx"/>
      </ScrollPanel>
    </div>

    <div class="flex-grow flex flex-column mt-2">
      <div class="flex-shrink-0 text-center flex w-full">
        <InputText type="text" class="flex-grow-1 bg-black-alpha-50" v-model="inputMsg"/>
        <Button icon="pi pi-send" class="ml-2"/>
      </div>

    </div>
  </div>
</template>

<style scoped>

</style>