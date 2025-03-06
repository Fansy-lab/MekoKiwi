<template>
  <div
    class="bg-muted flex-shrink-0 overflow-y-auto transition-all duration-300 fixed md:relative z-30"
    :class="[
      isMobileView
        ? isOpen
          ? 'translate-x-0 w-[80%] max-w-[280px] h-screen'
          : '-translate-x-full w-0'
        : isOpen
          ? 'w-64'
          : 'w-20',
      'md:translate-x-0 md:h-[calc(100vh-3.5rem)] fixed top-0' // Full height minus ServerBar height on desktop
    ]"
  >
    <div class="p-4">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-lg font-bold text-primary" :class="{ hidden: !isOpen && !isMobileView }">
          Channels
        </h2>
        <button
          @click="$emit('toggle')"
          class="p-2 bg-background rounded-full text-card hover:text-primary hover:bg-card"
          :class="{ 'md:hidden': isMobileView }"
        >
          <ChevronLeftIcon class="text-accent" v-if="isOpen" size="18" />
          <ChevronRightIcon class="text-accent" v-else size="18" />
        </button>
      </div>

      <div class="space-y-1">
        <div
          v-for="channel in channels"
          :key="channel.id"
          class="flex items-center p-2 rounded-lg cursor-pointer transition-colors"
          :class="
            channel.id === activeChannel
              ? 'bg-secondary-muted text-primary'
              : 'text-muted-foreground hover:bg-card hover:text-primary'
          "
          @click="$emit('change-channel', channel.id)"
        >
          <HashIcon v-if="channel.type === 'text'" size="18" class="flex-shrink-0" />
          <VolumeXIcon v-else-if="channel.type === 'voice'" size="18" class="flex-shrink-0" />
          <span
            v-if="isOpen"
            :class="channel.id === activeChannel ? 'font-bold' : ''"
            class="ml-2 truncate"
            >{{ channel.name }}</span
          >
        </div>
      </div>

      <div class="mt-8 mb-4" v-if="isOpen">
        <h2 class="text-lg font-bold text-primary mb-4">Direct Messages</h2>
        <div class="space-y-2">
          <div
            v-for="user in directMessages"
            :key="user.id"
            class="flex items-center p-2 rounded-lg cursor-pointer transition-colors"
            :class="
              user.id === activeUser
                ? 'bg-secondary-muted text-primary'
                : 'text-muted-foreground hover:bg-card hover:text-primary'
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
    required: false,
    default: null
  },
  activeUser: {
    type: Number,
    default: null
  },
  isOpen: {
    type: Boolean,
    default: true
  },
  isMobileView: {
    type: Boolean,
    default: false
  }
})

defineEmits(['toggle', 'change-channel', 'change-user'])
</script>
