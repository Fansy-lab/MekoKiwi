<template>
  <div class="flex flex-col h-full w-full bg-muted text-primary overflow-y-auto">
    <!-- Dashboard Content -->
    <div class="p-6 flex flex-col gap-6">
      <!-- Welcome Section -->
      <div class="bg-gradient-to-r from-accent to-[#7289da] rounded-lg p-6 text-white">
        <h2 class="text-2xl mb-2">Welcome back, {{ currentUser.name }}!</h2>
        <p>{{ getGreeting() }}</p>
      </div>

      <!-- Stats Overview -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <div class="bg-card rounded-lg p-4 flex items-center gap-4">
          <UsersIcon size="24" class="text-accent" />
          <div>
            <h3 class="text-2xl font-bold mb-1">{{ stats.onlineFriends }}</h3>
            <p class="text-muted-foreground text-sm">Friends Online</p>
          </div>
        </div>
        <div class="bg-card rounded-lg p-4 flex items-center gap-4">
          <MessageSquareIcon size="24" class="text-accent" />
          <div>
            <h3 class="text-2xl font-bold mb-1">{{ stats.unreadMessages }}</h3>
            <p class="text-muted-foreground text-sm">Unread Messages</p>
          </div>
        </div>
        <div class="bg-card rounded-lg p-4 flex items-center gap-4">
          <BellIcon size="24" class="text-accent" />
          <div>
            <h3 class="text-2xl font-bold mb-1">{{ stats.notifications }}</h3>
            <p class="text-muted-foreground text-sm">Notifications</p>
          </div>
        </div>
        <div class="bg-card rounded-lg p-4 flex items-center gap-4">
          <ServerIcon size="24" class="text-accent" />
          <div>
            <h3 class="text-2xl font-bold mb-1">{{ stats.servers }}</h3>
            <p class="text-muted-foreground text-sm">Servers</p>
          </div>
        </div>
      </div>

      <!-- Main Dashboard Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
        <!-- Recent Activity -->
        <div class="bg-card rounded-lg p-4 flex flex-col gap-4">
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-semibold">Recent Activity</h3>
            <button class="text-accent text-sm hover:underline">View All</button>
          </div>
          <div class="flex flex-col gap-3">
            <div v-for="(activity, index) in recentActivity" :key="index" class="flex gap-3">
              <div class="w-9 h-9 rounded-full overflow-hidden flex-shrink-0">
                <img
                  :src="activity.userAvatar"
                  :alt="activity.userName"
                  class="w-full h-full object-cover"
                />
              </div>
              <div class="flex-1">
                <p class="text-sm">
                  <strong>{{ activity.userName }}</strong> {{ activity.action }}
                  <span class="text-accent">{{ activity.target }}</span>
                </p>
                <span class="text-xs text-muted-foreground">{{ activity.time }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Your Servers -->
        <div class="bg-card rounded-lg p-4 flex flex-col gap-4">
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-semibold">Your Servers</h3>
            <button class="text-accent text-sm hover:underline">View All</button>
          </div>
          <div class="flex flex-col gap-2">
            <div
              v-for="server in userServers"
              :key="server.id"
              class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-secondary-muted"
              @click="goToServer(server.id)"
            >
              <div class="w-10 h-10 rounded-lg overflow-hidden mr-3">
                <img :src="server.icon" :alt="server.name" class="w-full h-full object-cover" />
              </div>
              <div class="flex-1">
                <h4 class="text-sm font-semibold">{{ server.name }}</h4>
                <p class="text-xs text-muted-foreground">{{ server.memberCount }} members</p>
              </div>
              <div class="flex items-center">
                <span
                  v-if="server.unread"
                  class="bg-danger text-white text-xs font-bold min-w-[20px] h-[20px] rounded-full flex items-center justify-center px-1.5"
                  >{{ server.unread }}</span
                >
                <ChevronRightIcon v-else size="16" class="text-muted-foreground" />
              </div>
            </div>
          </div>
        </div>

        <!-- Direct Messages -->
        <div class="bg-card rounded-lg p-4 flex flex-col gap-4">
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-semibold">Direct Messages</h3>
            <button class="text-accent text-sm hover:underline">View All</button>
          </div>
          <div class="flex flex-col gap-2">
            <div
              v-for="dm in directMessages"
              :key="dm.id"
              class="flex items-center p-2 rounded-lg cursor-pointer hover:bg-secondary-muted"
              @click="openChat(dm.id)"
            >
              <div class="relative w-10 h-10 rounded-full overflow-hidden">
                <img :src="dm.avatar" :alt="dm.name" class="w-full h-full object-cover" />
                <span
                  class="absolute bottom-0 right-0 w-2.5 h-2.5 rounded-full border-2 border-card"
                  :class="dm.status === 'online' ? 'bg-success' : 'bg-muted-foreground'"
                ></span>
              </div>
              <div class="flex-1 ml-3">
                <h4 class="text-sm font-semibold">{{ dm.name }}</h4>
                <p class="text-xs text-muted-foreground truncate max-w-[180px]">
                  {{ dm.lastMessage }}
                </p>
              </div>
              <div class="text-xs text-muted-foreground whitespace-nowrap">{{ dm.time }}</div>
            </div>
          </div>
        </div>

        <!-- Upcoming Events -->
        <div class="bg-card rounded-lg p-4 flex flex-col gap-4">
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-semibold">Upcoming Events</h3>
            <button class="text-accent text-sm hover:underline">View All</button>
          </div>
          <div class="flex flex-col gap-3">
            <div
              v-for="event in upcomingEvents"
              :key="event.id"
              class="flex items-center p-2 rounded-lg bg-secondary-muted"
            >
              <div
                class="flex flex-col items-center justify-center w-10 h-10 bg-accent rounded-lg mr-3"
              >
                <span class="text-lg font-bold leading-none">{{ event.day }}</span>
                <span class="text-xs uppercase">{{ event.month }}</span>
              </div>
              <div class="flex-1">
                <h4 class="text-sm font-semibold">{{ event.title }}</h4>
                <p class="text-xs text-muted-foreground">{{ event.time }} • {{ event.server }}</p>
              </div>
              <button
                class="bg-accent text-white rounded-md px-3 py-1 text-xs font-semibold flex items-center gap-1 hover:bg-accent-foreground"
              >
                <CalendarIcon size="16" />
                <span>RSVP</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  BellIcon,
  UsersIcon,
  MessageSquareIcon,
  ServerIcon,
  ChevronRightIcon,
  CalendarIcon
} from 'lucide-vue-next'

// Current user
const currentUser = ref({
  id: 1,
  name: 'Jane Smith',
  avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=1'
})

// Dashboard stats
const stats = ref({
  onlineFriends: 12,
  unreadMessages: 8,
  notifications: 3,
  servers: 7
})

// Recent activity
const recentActivity = ref([
  {
    userName: 'John Doe',
    userAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=2',
    action: 'posted in',
    target: '#gaming',
    time: '10 min ago'
  },
  {
    userName: 'Bob Williams',
    userAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=4',
    action: 'started a voice chat in',
    target: 'Gaming Lounge',
    time: '25 min ago'
  },
  {
    userName: 'Alice Johnson',
    userAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=3',
    action: 'mentioned you in',
    target: '#announcements',
    time: '1 hour ago'
  },
  {
    userName: 'Mike Johnson',
    userAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=5',
    action: 'shared a file in',
    target: 'Project Alpha',
    time: '3 hours ago'
  }
])

// User servers
const userServers = ref([
  {
    id: 1,
    name: 'Gaming Community',
    icon: 'https://api.dicebear.com/7.x/identicon/svg?seed=gaming',
    memberCount: 1243,
    unread: 5
  },
  {
    id: 2,
    name: 'Design Hub',
    icon: 'https://api.dicebear.com/7.x/identicon/svg?seed=design',
    memberCount: 856,
    unread: 0
  },
  {
    id: 3,
    name: 'Developers',
    icon: 'https://api.dicebear.com/7.x/identicon/svg?seed=dev',
    memberCount: 1892,
    unread: 12
  },
  {
    id: 4,
    name: 'Movie Club',
    icon: 'https://api.dicebear.com/7.x/identicon/svg?seed=movies',
    memberCount: 421,
    unread: 0
  }
])

// Direct messages
const directMessages = ref([
  {
    id: 1,
    name: 'John Doe',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=2',
    status: 'online',
    lastMessage: 'Are you coming to the meeting?',
    time: '10:45 AM'
  },
  {
    id: 2,
    name: 'Alice Johnson',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=3',
    status: 'offline',
    lastMessage: 'I sent you the files you requested',
    time: 'Yesterday'
  },
  {
    id: 3,
    name: 'Bob Williams',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=4',
    status: 'online',
    lastMessage: "Let me know when you're free to chat",
    time: 'Yesterday'
  }
])

// Upcoming events
const upcomingEvents = ref([
  {
    id: 1,
    title: 'Game Night: Valorant',
    day: '15',
    month: 'MAR',
    time: '8:00 PM',
    server: 'Gaming Community'
  },
  {
    id: 2,
    title: 'UI/UX Workshop',
    day: '18',
    month: 'MAR',
    time: '6:30 PM',
    server: 'Design Hub'
  },
  {
    id: 3,
    title: 'Movie Watch Party',
    day: '20',
    month: 'MAR',
    time: '9:00 PM',
    server: 'Movie Club'
  }
])

// Methods
const getGreeting = () => {
  const hour = new Date().getHours()
  if (hour < 12) return 'Good morning! Ready to start your day?'
  if (hour < 18) return "Good afternoon! What's happening in your communities?"
  return 'Good evening! Catch up with what you missed today.'
}

const goToServer = (serverId) => {
  console.log(`Navigating to server ${serverId}`)
  // Implementation would depend on your routing setup
}

const openChat = (userId) => {
  console.log(`Opening chat with user ${userId}`)
  // Implementation would depend on your routing setup
}
</script>

<style scoped>
/* Add any additional styles if needed */
</style>
