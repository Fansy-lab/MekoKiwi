<template>
  <div class="h-16 bg-primary flex items-center px-4">
    <div class="flex items-center">
      <UserAvatar
        name="ME"
        status="online"
        size="lg"
        bgColor="bg-primary-emphasis"
        textColor="text-secondary-muted"
        class="mr-3"
      />
      <div>
        <div class="font-medium text-primary-foreground">YourUsername</div>
        <div class="text-xs text-secondary-muted">Online</div>
      </div>
    </div>

    <div class="ml-auto flex items-center space-x-3">
      <button
        class="p-2 bg-primary-muted rounded-full text-secondary-muted hover:bg-primary-emphasis"
      >
        <MicIcon size="18" />
      </button>
      <button
        class="p-2 bg-primary-muted rounded-full text-secondary-muted hover:bg-primary-emphasis"
      >
        <HeadphonesIcon size="18" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { MicIcon, HeadphonesIcon, UserIcon, SettingsIcon, LogOutIcon } from 'lucide-vue-next'
import UserAvatar from './UserAvatar.vue'
</script>
