<template>
  <div class="h-16 bg-teal-700 flex items-center px-4">
    <div class="flex items-center">
      <UserAvatar
        name="ME"
        status="online"
        size="lg"
        bgColor="bg-teal-900"
        textColor="text-teal-100"
        class="mr-3"
      />
      <div>
        <div class="font-medium text-white">YourUsername</div>
        <div class="text-xs text-teal-200">Online</div>
      </div>
    </div>

    <div class="ml-auto flex items-center space-x-3">
      <button class="p-2 bg-teal-800 rounded-full text-teal-100 hover:bg-teal-900">
        <MicIcon size="18" />
      </button>
      <button class="p-2 bg-teal-800 rounded-full text-teal-100 hover:bg-teal-900">
        <HeadphonesIcon size="18" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { MicIcon, HeadphonesIcon, UserIcon, SettingsIcon, LogOutIcon } from 'lucide-vue-next'
import UserAvatar from './UserAvatar.vue'
</script>
