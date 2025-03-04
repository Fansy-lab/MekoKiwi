<template>
  <div class="bg-teal-700 p-2 flex items-center overflow-x-auto">
    <div class="flex-shrink-0 mr-4 ml-2">
      <div
        class="w-10 h-10 bg-teal-900 rounded-full flex items-center justify-center cursor-pointer"
      >
        <HomeIcon size="20" class="text-teal-100" />
      </div>
    </div>

    <div class="flex space-x-3 py-1">
      <div
        v-for="server in servers"
        :key="server.id"
        class="flex-shrink-0 transition-all duration-200 relative group"
        :class="server.id === activeServer ? 'scale-110' : ''"
        @click="$emit('change-server', server.id)"
      >
        <div
          class="w-10 h-10 rounded-full flex items-center justify-center cursor-pointer"
          :class="
            server.id === activeServer
              ? 'bg-teal-900 ring-2 ring-teal-300'
              : 'bg-teal-800 hover:bg-teal-900'
          "
        >
          <img
            v-if="server.image"
            :src="server.image"
            class="w-6 h-6 rounded-full"
            :alt="server.name"
          />
          <span v-else class="text-sm font-bold text-teal-100">{{
            server.name.substring(0, 2)
          }}</span>
        </div>
        <div
          class="absolute -left-1 top-1/2 transform -translate-y-1/2 h-0 w-1 bg-teal-300 rounded-r transition-all duration-200"
          :class="server.id === activeServer ? 'h-5' : 'group-hover:h-3'"
        ></div>
        <div
          class="absolute -top-10 left-1/2 transform -translate-x-1/2 bg-gray-900 text-white text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap"
        >
          {{ server.name }}
        </div>
      </div>
    </div>

    <div class="flex-shrink-0 ml-3">
      <div
        class="w-10 h-10 bg-teal-800 rounded-full flex items-center justify-center cursor-pointer hover:bg-teal-900"
      >
        <PlusIcon size="20" class="text-teal-100" />
      </div>
    </div>

    <div class="ml-auto flex items-center space-x-3 mr-4">
      <button class="p-2 bg-teal-800 rounded-full text-teal-100 hover:bg-teal-900">
        <BellIcon size="18" />
      </button>
      <button class="p-2 bg-teal-800 rounded-full text-teal-100 hover:bg-teal-900">
        <SettingsIcon size="18" />
      </button>
      <UserAvatar
        name="ME"
        size="lg"
        status="online"
        bgColor="bg-teal-900"
        textColor="text-teal-100"
      />
    </div>
  </div>
</template>

<script setup>
import { HomeIcon, PlusIcon, BellIcon, SettingsIcon } from 'lucide-vue-next'
import UserAvatar from './UserAvatar.vue'

defineProps({
  servers: {
    type: Array,
    required: true
  },
  activeServer: {
    type: Number,
    required: true
  }
})

defineEmits(['change-server'])
</script>
