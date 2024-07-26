<script setup lang="ts">
import {computed, inject} from "vue";
import Avatar from 'primevue/avatar';
import {types} from "../../wailsjs/go/models";
import Message = types.Message;


const props = defineProps<{ message: Message }>()
const positionClass = computed(() => {
  if (props.message.userType === "assistant") {
    return "flex-row"
  } else {
    return "flex-row-reverse"
  }
})
const avatarImg = computed(() => {
  if(props.message.userType === "user") {
    return "https://s3.bmp.ovh/imgs/2024/07/25/b82d96fabe6f83e1.png"
  } else {
    return "https://s3.bmp.ovh/imgs/2024/07/25/f214b677e17f25f3.png"
  }
})
const showInfoToast = inject<any>('showInfoToast')

const copyText = () => {
  navigator.clipboard.writeText(props.message.content)
  showInfoToast('复制到剪切板啦~')
}
</script>

<template>
  <div :class="['flex align-items-start mb-4 px-3', positionClass]">
    <div class="w-3rem h-3rem flex align-items-center justify-content-center p-1 border-circle mx-2">
      <Avatar :image="avatarImg" shape="circle" size="large" @click="copyText"/>
    </div>
    <div class="dybg border-round px-1 w-full">
      <p class="text-red-800 text-left pl-2 text-xl white-space-normal">
        {{ message.content }}
      </p>
      <i @click="copyText" v-if="message.userType === 'assistant'" v-tooltip="'copy text'" class="pi pi-copy text-gray-800 fixed right-0 mr-5" style="font-size: 1rem"></i>
    </div>
  </div>
</template>

<style scoped>

</style>