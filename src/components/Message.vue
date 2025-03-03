<template>
  <div class="flex">
    <UserAvatar :name="message.user.name" :image="message.user.avatar" class="mr-3" />
    <div class="flex-1">
      <div class="flex items-center">
        <span class="font-medium text-teal-800">{{ message.user.name }}</span>
        <span class="ml-2 text-xs text-gray-400">{{ message.time }}</span>
      </div>
      <p class="text-gray-700">{{ message.content }}</p>

      <!-- Reactions -->
      <div v-if="showReactions" class="flex mt-2 space-x-2">
        <div
          class="bg-gray-100 rounded-full px-2 py-1 text-xs flex items-center space-x-1 cursor-pointer hover:bg-gray-200"
        >
          <span>👍</span>
          <span>2</span>
        </div>
        <div
          class="bg-gray-100 rounded-full px-2 py-1 text-xs flex items-center space-x-1 cursor-pointer hover:bg-gray-200"
        >
          <span>❤️</span>
          <span>1</span>
        </div>
        <div
          class="bg-gray-100 rounded-full px-2 py-1 text-xs flex items-center space-x-1 cursor-pointer hover:bg-gray-200"
        >
          <span>+</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import UserAvatar from './UserAvatar.vue'

defineProps({
  message: {
    type: Object,
    required: true
  },
  showReactions: {
    type: Boolean,
    default: false
  }
})
</script>
