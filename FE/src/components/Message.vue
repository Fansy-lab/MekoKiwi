<template>
  <div class="flex">
    <UserAvatar
      :name="message.user.name"
      :image="message.user.avatar"
      class="mr-3"
      :size="isMobileView ? 'sm' : 'md'"
    />
    <div class="flex-1 min-w-0">
      <div class="flex items-center flex-wrap">
        <span class="font-medium text-primary">{{ message.user.name }}</span>
        <span class="ml-2 text-xs text-muted-foreground">{{
          formatTimestamp(message.timestamp)
        }}</span>
      </div>
      <p class="text-foreground break-words">{{ message.content }}</p>

      <!-- Reactions -->
      <div v-if="showReactions" class="flex mt-2 space-x-2 flex-wrap">
        <div
          class="bg-muted text-foreground rounded-full px-2 py-1 text-xs flex items-center space-x-1 cursor-pointer hover:bg-accent mb-1"
        >
          <span>👍</span>
          <span>2</span>
        </div>
        <div
          class="bg-muted text-foreground rounded-full px-2 py-1 text-xs flex items-center space-x-1 cursor-pointer hover:bg-accent mb-1"
        >
          <span>❤️</span>
          <span>1</span>
        </div>
        <div
          class="bg-muted text-foreground rounded-full px-2 py-1 text-xs flex items-center space-x-1 cursor-pointer hover:bg-accent mb-1"
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
  },
  isMobileView: {
    type: Boolean,
    default: false
  }
})

const formatTimestamp = (timestamp) => {
  const date = new Date(timestamp * 1000) // Convert to milliseconds

  const hours = date.getHours().toString().padStart(2, '0')
  const minutes = date.getMinutes().toString().padStart(2, '0')

  return `${hours}:${minutes}`
}
</script>
