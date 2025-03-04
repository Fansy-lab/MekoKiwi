<template>
  <div
    class="bg-gray-100 w-64 flex-shrink-0 overflow-y-auto transition-all duration-300"
    :class="isOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0 md:w-20'"
  >
    <div class="p-4">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-lg font-bold text-teal-700" :class="isOpen ? '' : 'md:hidden'">Channels</h2>
        <button
          @click="$emit('toggle')"
          class="p-2 bg-gray-200 rounded-full text-teal-600 hover:bg-gray-300 md:hidden"
        >
          <ChevronLeftIcon v-if="isOpen" size="18" />
          <ChevronRightIcon v-else size="18" />
        </button>
        <button
          @click="$emit('toggle')"
          class="hidden md:block p-2 bg-gray-200 rounded-full text-teal-600 hover:bg-gray-300"
        >
          <ChevronLeftIcon v-if="isOpen" size="18" />
          <ChevronRightIcon v-else size="18" />
        </button>
      </div>

      <div class="space-y-1">
        <div
          v-for="channel in channels"
          :key="channel.id"
          class="flex items-center p-2 rounded-lg cursor-pointer transition-colors"
          :class="
            channel.id === activeChannel
              ? 'bg-teal-100 text-teal-900'
              : 'text-gray-600 hover:bg-gray-200 hover:text-teal-700'
          "
          @click="$emit('change-channel', channel.id)"
        >
          <HashIcon v-if="channel.type === 'text'" size="18" class="flex-shrink-0" />
          <VolumeXIcon v-else-if="channel.type === 'voice'" size="18" class="flex-shrink-0" />
          <span v-if="isOpen" class="ml-2 truncate">{{ channel.name }}</span>
        </div>
      </div>

      <div class="mt-8 mb-4" v-if="isOpen">
        <h2 class="text-lg font-bold text-teal-700 mb-4">Direct Messages</h2>
        <div class="space-y-2">
          <div
            v-for="user in directMessages"
            :key="user.id"
            class="flex items-center p-2 rounded-lg cursor-pointer transition-colors"
            :class="
              user.id === activeUser
                ? 'bg-teal-100 text-teal-900'
                : 'text-gray-600 hover:bg-gray-200 hover:text-teal-700'
            "
            @click="$emit('change-user', user.id)"
          >
            <UserAvatar
              :name="user.name"
              :image="user.avatar"
              :status="user.online ? 'online' : 'offline'"
            />
            <span class="ml-2 truncate">{{ user.name }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { HashIcon, VolumeXIcon, ChevronLeftIcon, ChevronRightIcon } from 'lucide-vue-next'
import UserAvatar from './UserAvatar.vue'

defineProps({
  channels: {
    type: Array,
    required: true
  },
  directMessages: {
    type: Array,
    required: true
  },
  activeChannel: {
    type: Number,
    required: true
  },
  activeUser: {
    type: Number,
    default: null
  },
  isOpen: {
    type: Boolean,
    default: true
  }
})

defineEmits(['toggle', 'change-channel', 'change-user'])
</script>
