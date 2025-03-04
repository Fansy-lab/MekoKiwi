<template>
  <div class="flex-1 flex flex-col bg-white overflow-hidden relative">
    <!-- Toggle button for mobile -->
    <button
      @click="$emit('toggle-channel-sidebar')"
      class="absolute top-4 left-4 p-2 bg-gray-200 rounded-full text-teal-600 hover:bg-gray-300 md:hidden z-10"
    >
      <MenuIcon size="18" />
    </button>

    <!-- Chat header -->
    <div class="h-16 flex items-center px-4 border-b border-gray-200 bg-white">
      <div class="ml-10 md:ml-0">
        <div class="flex items-center">
          <HashIcon size="18" class="mr-2 text-teal-600" />
          <span class="font-medium text-gray-800">{{ getActiveChannelName() }}</span>
        </div>
        <div class="text-xs text-gray-500">Welcome to the {{ getActiveChannelName() }} channel</div>
      </div>

      <div class="ml-auto flex items-center space-x-3">
        <div class="relative">
          <SearchIcon
            size="18"
            class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
          />
          <input
            type="text"
            placeholder="Search..."
            class="bg-gray-100 rounded-full pl-10 pr-4 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500 w-40 md:w-60"
          />
        </div>
        <button
          @click="$emit('toggle-members-sidebar')"
          class="p-2 bg-gray-100 rounded-full text-teal-600 hover:bg-gray-200 lg:hidden"
        >
          <UsersIcon size="18" />
        </button>
      </div>
    </div>

    <!-- Messages -->
    <div class="flex-1 overflow-y-auto p-4 space-y-6">
      <Message
        v-for="(message, index) in messages"
        :key="message.id"
        :message="message"
        :showReactions="index % 2 === 0"
      />
    </div>

    <!-- Message input -->
    <div class="p-4 border-t border-gray-200">
      <div class="flex items-center bg-gray-100 rounded-xl p-1">
        <button class="p-2 text-teal-600 hover:text-teal-800 hover:bg-gray-200 rounded-lg">
          <PlusCircleIcon size="20" />
        </button>
        <input
          type="text"
          placeholder="Message #general"
          class="flex-1 bg-transparent border-none focus:outline-none px-3 py-2 text-gray-700"
          v-model="newMessage"
          @keyup.enter="sendMessage"
        />
        <button class="p-2 text-teal-600 hover:text-teal-800 hover:bg-gray-200 rounded-lg">
          <SmileIcon size="20" />
        </button>
        <button class="p-2 text-teal-600 hover:text-teal-800 hover:bg-gray-200 rounded-lg">
          <ImageIcon size="20" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import {
  HashIcon,
  SearchIcon,
  UsersIcon,
  PlusCircleIcon,
  SmileIcon,
  ImageIcon,
  MenuIcon
} from 'lucide-vue-next'
import Message from './Message.vue'

const props = defineProps({
  messages: {
    type: Array,
    required: true
  },
  activeChannel: {
    type: Number,
    required: true
  },
  channels: {
    type: Array,
    default: () => []
  },
  isChannelSidebarOpen: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['toggle-channel-sidebar', 'toggle-members-sidebar', 'send-message'])

const newMessage = ref('')

const getActiveChannelName = () => {
  const channel = props.channels.find((c) => c.id === props.activeChannel)
  return channel ? channel.name : 'general'
}

const sendMessage = () => {
  if (newMessage.value.trim()) {
    emit('send-message', newMessage.value)
    newMessage.value = ''
  }
}
</script>
