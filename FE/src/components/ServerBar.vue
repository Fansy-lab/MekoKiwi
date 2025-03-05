<template>
  <div
    class="bg-[rgb(var(--muted))] border-b border-solid border-b-[rgb(var(--muted-foreground))] p-2 flex items-center"
  >
    <div class="flex-shrink-0 mr-4 ml-2">
      <div
        class="w-10 h-10 bg-[rgb(var(--card))] rounded-full flex items-center justify-center cursor-pointer"
      >
        <HomeIcon size="20" class="text-[rgb(var(--foreground))]" />
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
              ? 'bg-[rgb(var(--card))] ring-2 ring-[rgb(var(--primary))]'
              : 'bg-[rgb(var(--muted))] hover:bg-[rgb(var(--card))]'
          "
        >
          <img
            v-if="server.image"
            :src="server.image"
            class="w-6 h-6 rounded-full"
            :alt="server.name"
          />
          <span v-else class="text-sm font-bold text-[rgb(var(--foreground))]">{{
            server.name.substring(0, 2)
          }}</span>
        </div>

        <div
          class="absolute -top-10 left-1/2 transform -translate-x-1/2 bg-[rgb(var(--popover))] text-[rgb(var(--popover-foreground))] text-xs py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap"
        >
          {{ server.name }}
        </div>
      </div>
    </div>

    <div class="flex-shrink-0 ml-3">
      <div
        class="w-10 h-10 bg-[rgb(var(--muted))] rounded-full flex items-center justify-center cursor-pointer hover:bg-[rgb(var(--card))]"
      >
        <PlusIcon size="20" class="text-[rgb(var(--foreground))]" />
      </div>
    </div>

    <div class="ml-auto flex items-center space-x-3">
      <div class="relative group">
        <button
          class="p-2 bg-[rgb(var(--muted))] rounded-full text-[rgb(var(--foreground))] hover:bg-[rgb(var(--card))]"
        >
          <UserIcon size="18" />
        </button>
        <div
          class="absolute z-50 top-full right-0 mt-2 w-48 bg-[rgb(var(--popover))] rounded-lg shadow-lg p-3 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all"
        >
          <div class="text-sm font-medium text-[rgb(var(--muted-foreground))] mb-2">
            User Settings
          </div>
          <div class="space-y-2">
            <div
              class="flex items-center p-2 rounded-md hover:bg-[rgb(var(--muted))] cursor-pointer"
            >
              <UserIcon size="16" class="mr-2 text-[rgb(var(--primary))]" />
              <span class="text-sm">Profile</span>
            </div>
            <div
              class="flex items-center p-2 rounded-md hover:bg-[rgb(var(--muted))] cursor-pointer"
            >
              <SettingsIcon size="16" class="mr-2 text-[rgb(var(--primary))]" />
              <span class="text-sm">Settings</span>
            </div>
            <div
              class="flex items-center p-2 rounded-md hover:bg-[rgb(var(--muted))] cursor-pointer"
            >
              <LogOutIcon size="16" class="mr-2 text-[rgb(var(--primary))]" />
              <span class="text-sm">Log Out</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { HomeIcon, PlusIcon, SettingsIcon, UserIcon, LogOutIcon } from 'lucide-vue-next'

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
