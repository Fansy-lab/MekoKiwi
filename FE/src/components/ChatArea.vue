<template>
  <div class="flex-1 flex flex-col bg-background overflow-hidden relative">
    <!-- Toggle button for mobile -->
    <button
      @click="$emit('toggle-channel-sidebar')"
      class="absolute top-4 left-4 p-2 bg-muted rounded-full text-primary hover:bg-accent-foreground md:hidden z-10"
    >
      <ArrowLeftIcon size="18" />
    </button>

    <!-- Chat header -->
    <div class="h-16 flex items-center px-4 border-b border-border bg-background">
      <div class="ml-10 md:ml-0">
        <div class="flex items-center">
          <HashIcon size="18" class="mr-2 text-primary" />
          <span class="font-medium text-foreground">{{ activeChannelName }}</span>
        </div>
        <div class="text-xs text-muted-foreground">
          Welcome to the {{ activeChannelName }} channel
        </div>
      </div>

      <div class="ml-auto flex items-center space-x-3">
        <div class="relative hidden sm:block">
          <SearchIcon
            size="18"
            class="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground"
          />
          <input
            type="text"
            placeholder="Search..."
            class="bg-muted text-foreground rounded-full pl-10 pr-4 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-ring w-40 md:w-60"
          />
        </div>
        <button
          @click="$emit('toggle-members-sidebar')"
          class="p-2 bg-muted rounded-full text-primary hover:bg-accent lg:hidden"
        >
          <UsersIcon size="18" />
        </button>
      </div>
    </div>

    <!-- Messages -->
    <div ref="messagesContainer" class="flex-1 overflow-y-auto p-4 space-y-6">
      <Message
        v-for="(message, index) in messagesStore.messages"
        :key="message.id"
        :message="message"
        :showReactions="index % 2 === 0"
        :isMobileView="isMobileView"
      />
    </div>

    <!-- Message input -->
    <div class="p-2 sm:p-4 border-t border-border">
      <div class="flex items-center bg-muted rounded-xl p-1">
        <button class="p-2 text-primary hover:text-primary-emphasis hover:bg-accent rounded-lg">
          <PlusCircleIcon size="20" />
        </button>
        <input
          type="text"
          placeholder="Message #general"
          class="flex-1 bg-transparent border-none focus:outline-none px-2 sm:px-3 py-2 text-foreground text-sm sm:text-base"
          v-model="newMessage"
          @keyup.enter="sendMessage"
        />
        <button
          class="p-2 text-primary hover:text-primary-emphasis hover:bg-accent rounded-lg hidden sm:block"
        >
          <SmileIcon size="20" />
        </button>
        <button
          class="p-2 text-primary hover:text-primary-emphasis hover:bg-accent rounded-lg hidden sm:block"
        >
          <ImageIcon size="20" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, nextTick, onBeforeUnmount, computed } from 'vue'
import { useMessagesStore } from '../stores/messages_store'

import {
  HashIcon,
  SearchIcon,
  UsersIcon,
  PlusCircleIcon,
  SmileIcon,
  ImageIcon,
  MenuIcon,
  ArrowLeftIcon
} from 'lucide-vue-next'
import Message from './Message.vue'

const props = defineProps({
  messages: {
    type: Array,
    required: true
  },
  activeChannel: {
    type: Number,
    required: false
  },
  channels: {
    type: Array,
    default: () => []
  },
  isChannelSidebarOpen: {
    type: Boolean,
    default: true
  },
  isMobileView: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['toggle-channel-sidebar', 'toggle-members-sidebar', 'send-message'])

const newMessage = ref('')
const messagesStore = useMessagesStore()
const messagesContainer = ref(null)

//computed activechannelname
const activeChannelName = computed(() => {
  const channel = props.channels.find((channel) => channel.id === props.activeChannel)
  return channel ? channel.name : 'Unknown channel'
})

const scrollToBottom = () => {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
    }
  })
}

const sendMessage = () => {
  if (newMessage.value.trim()) {
    emit('send-message', newMessage.value)
    newMessage.value = ''
    scrollToBottom()
  }
}

messagesStore.$subscribe((mutation, state) => {
  scrollToBottom()
})

const unsubscribe = messagesStore.$subscribe((mutation, state) => {
  scrollToBottom()
})

onBeforeUnmount(() => {
  unsubscribe()
})
</script>
