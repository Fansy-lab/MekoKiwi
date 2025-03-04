<template>
  <div class="relative" :class="[className]">
    <div
      :class="[
        sizeClasses[size],
        'rounded-full flex items-center justify-center overflow-hidden',
        bgColor || 'bg-gray-300',
        { 'opacity-70': opacity < 100 }
      ]"
    >
      <img v-if="image" :src="image" :alt="name" class="w-full h-full object-cover" />
      <span v-else :class="['text-xs font-bold', textColor || 'text-gray-600']">{{
        name.substring(0, 2)
      }}</span>
    </div>
    <div
      v-if="status"
      class="absolute bottom-0 right-0 w-3 h-3 rounded-full border-2 border-gray-100"
      :class="statusClasses[status]"
    ></div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  name: {
    type: String,
    required: true
  },
  image: {
    type: String,
    default: null
  },
  status: {
    type: String,
    default: null,
    validator: (value) => ['online', 'offline', 'away', 'dnd', null].includes(value)
  },
  size: {
    type: String,
    default: 'sm',
    validator: (value) => ['sm', 'md', 'lg'].includes(value)
  },
  bgColor: {
    type: String,
    default: ''
  },
  textColor: {
    type: String,
    default: ''
  },
  className: {
    type: String,
    default: ''
  },
  opacity: {
    type: Number,
    default: 100
  }
})

const sizeClasses = {
  sm: 'w-8 h-8',
  md: 'w-9 h-9',
  lg: 'w-10 h-10'
}

const statusClasses = {
  online: 'bg-green-500',
  offline: 'bg-gray-400',
  away: 'bg-yellow-500',
  dnd: 'bg-red-500'
}
</script>
