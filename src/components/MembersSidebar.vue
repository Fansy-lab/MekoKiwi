<template>
  <div
    class="bg-gray-100 w-64 flex-shrink-0 overflow-y-auto transition-all duration-300"
    :class="isOpen ? 'translate-x-0' : 'translate-x-full lg:translate-x-0'"
  >
    <div class="p-4">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-lg font-bold text-teal-700">Members</h2>
        <button
          @click="$emit('toggle')"
          class="p-2 bg-gray-200 rounded-full text-teal-600 hover:bg-gray-300 lg:hidden"
        >
          <XIcon size="18" />
        </button>
      </div>

      <div class="mb-6">
        <h3 class="text-xs font-semibold text-gray-500 uppercase mb-3">
          Online - {{ onlineMembers.length }}
        </h3>
        <div class="space-y-3">
          <div
            v-for="member in onlineMembers"
            :key="member.id"
            class="flex items-center p-2 rounded-lg hover:bg-gray-200 cursor-pointer"
          >
            <UserAvatar
              :name="member.name"
              :image="member.avatar"
              status="online"
              size="md"
              class="mr-3"
            />
            <div>
              <div class="font-medium text-teal-800">{{ member.name }}</div>
              <div class="text-xs text-gray-500">Online</div>
            </div>
          </div>
        </div>
      </div>

      <div>
        <h3 class="text-xs font-semibold text-gray-500 uppercase mb-3">
          Offline - {{ offlineMembers.length }}
        </h3>
        <div class="space-y-3">
          <div
            v-for="member in offlineMembers"
            :key="member.id"
            class="flex items-center p-2 rounded-lg hover:bg-gray-200 cursor-pointer"
          >
            <UserAvatar
              :name="member.name"
              :image="member.avatar"
              status="offline"
              size="md"
              class="mr-3"
              :opacity="70"
            />
            <div>
              <div class="font-medium text-gray-500">{{ member.name }}</div>
              <div class="text-xs text-gray-400">Offline</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { XIcon } from 'lucide-vue-next'
import UserAvatar from './UserAvatar.vue'

defineProps({
  onlineMembers: {
    type: Array,
    required: true
  },
  offlineMembers: {
    type: Array,
    required: true
  },
  isOpen: {
    type: Boolean,
    default: true
  }
})

defineEmits(['toggle'])
</script>
