<template>
  <div class="relative" :class="[className]">
    <div
      :class="[
        sizeClasses[size],
        'rounded-full flex items-center justify-center overflow-hidden',
        bgColor || 'bg-muted',
        { [`opacity-${opacity}`]: opacity < 100 }
      ]"
    >
      <img v-if="image" :src="image" :alt="name" class="w-full h-full object-cover" />
      <span v-else :class="['text-xs font-bold', textColor || 'text-muted-foreground']">{{
        name.substring(0, 2).toUpperCase()
      }}</span>
    </div>
    <div
      v-if="status"
      class="absolute bottom-0 right-0 w-3 h-3 rounded-full border-2 border-background"
      :class="[statusClasses[status]]"
    ></div>
  </div>
</template>

<script setup lang="ts">
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
    type: String as () => 'online' | 'offline' | 'away' | 'dnd' | null,
    default: null
  },
  size: {
    type: String as () => 'sm' | 'md' | 'lg',
    default: 'sm'
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
    default: 100,
    validator: (value: number) => value >= 0 && value <= 100
  }
})

const sizeClasses = {
  sm: 'w-8 h-8',
  md: 'w-9 h-9',
  lg: 'w-10 h-10'
}

const statusClasses = {
  online: 'bg-status-online',
  offline: 'bg-status-offline',
  away: 'bg-status-away',
  dnd: 'bg-status-dnd'
}
</script>
