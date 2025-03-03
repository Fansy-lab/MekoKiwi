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
      <div class="relative group">
        <button class="p-2 bg-teal-800 rounded-full text-teal-100 hover:bg-teal-900">
          <UserIcon size="18" />
        </button>
        <div
          class="absolute bottom-full right-0 mb-2 w-48 bg-white rounded-lg shadow-lg p-3 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all"
        >
          <div class="text-sm font-medium text-gray-800 mb-2">User Settings</div>
          <div class="space-y-2">
            <div class="flex items-center p-2 rounded-md hover:bg-gray-100 cursor-pointer">
              <UserIcon size="16" class="mr-2 text-teal-600" />
              <span class="text-sm">Profile</span>
            </div>
            <div class="flex items-center p-2 rounded-md hover:bg-gray-100 cursor-pointer">
              <SettingsIcon size="16" class="mr-2 text-teal-600" />
              <span class="text-sm">Settings</span>
            </div>
            <div class="flex items-center p-2 rounded-md hover:bg-gray-100 cursor-pointer">
              <LogOutIcon size="16" class="mr-2 text-teal-600" />
              <span class="text-sm">Log Out</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { MicIcon, HeadphonesIcon, UserIcon, SettingsIcon, LogOutIcon } from 'lucide-vue-next'
import UserAvatar from './UserAvatar.vue'
</script>
