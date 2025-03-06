<template>
  <div class="h-screen w-screen flex flex-col bg-gray-50 text-gray-800 overflow-hidden relative">
    <ServerBar :servers="servers" :activeServer="activeServer" @change-server="setActiveServer" />
    <Dashboard v-if="isDashboardOpen || activeChannel == null"></Dashboard>

    <div class="flex-1 flex overflow-hidden relative h-full">
      <!-- Mobile overlay backdrop when sidebars are open -->
      <div
        v-if="(isChannelSidebarOpen || isMembersSidebarOpen) && isMobileView"
        class="fixed inset-0 bg-black bg-opacity-50 z-20"
        @click="closeMobileSidebars"
      ></div>

      <ChannelSidebar
        :channels="channels"
        :directMessages="directMessages"
        :activeChannel="activeChannel"
        :activeUser="activeUser"
        :isOpen="isChannelSidebarOpen"
        :isMobileView="isMobileView"
        @toggle="toggleChannelSidebar"
        @change-channel="setActiveChannel"
        @change-user="setActiveUser"
      />
      <ChatArea
        v-if="!isMobileView || (isMobileView && activeChannel == null)"
        :messages="messagesStore.messages"
        :channels="channels"
        :activeChannel="activeChannel"
        :isChannelSidebarOpen="isChannelSidebarOpen"
        :isMobileView="isMobileView"
        @toggle-channel-sidebar="toggleChannelSidebar"
        @toggle-members-sidebar="toggleMembersSidebar"
        @send-message="sendMessage"
      />
      <div class="w-full h-full" v-else>
        <div class="sticky left-0 w-full bg-background h-full">
          <div class="p-4">
            <div class="flex items-center justify-between mb-6">
              <h2
                class="text-lg font-bold text-primary"
                :class="{ hidden: !isOpen && !isMobileView }"
              >
                Channels
              </h2>
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
      </div>

      <MembersSidebar
        :onlineMembers="onlineMembers"
        :offlineMembers="offlineMembers"
        :isOpen="isMembersSidebarOpen"
        :isMobileView="isMobileView"
        @toggle="toggleMembersSidebar"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import { HashIcon, VolumeXIcon } from 'lucide-vue-next'
import axios from 'axios'
import ServerBar from './components/ServerBar.vue'
import ChannelSidebar from './components/ChannelSidebar.vue'
import ChatArea from './components/ChatArea.vue'
import MembersSidebar from './components/MembersSidebar.vue'
import Dashboard from './components/Dashboard.vue'
import UserVoiceControlPanel from './components/UserVoiceControlPanel.vue'
import { useMessagesStore } from './stores/messages_store'

// Active states
const activeServer = ref(null)
const activeChannel = ref(null)
const activeUser = ref(null)
const servers = ref([])
const channels = ref([])
const isDashboardOpen = ref(false)
let socket
const messagesStore = useMessagesStore()

// UI state
const isChannelSidebarOpen = ref(false)
const isMembersSidebarOpen = ref(false)
const loading = ref(false)
const windowWidth = ref(window.innerWidth)
const isMobileView = computed(() => windowWidth.value < 768)

// Resize handler
const handleResize = () => {
  windowWidth.value = window.innerWidth
  if (windowWidth.value >= 768) {
    // On desktop, we want the sidebar to be either fully open or collapsed to icon mode
    // but never completely hidden
    isChannelSidebarOpen.value = true
  } else {
    // On mobile, we want to hide the sidebars by default
    isChannelSidebarOpen.value = false
    isMembersSidebarOpen.value = false
  }
}

// Close both sidebars on mobile
const closeMobileSidebars = () => {
  if (isMobileView.value) {
    isChannelSidebarOpen.value = false
    isMembersSidebarOpen.value = false
  }
}

// Toggle functions
const toggleChannelSidebar = () => {
  if (isMobileView.value) {
    if (isMembersSidebarOpen.value) {
      isMembersSidebarOpen.value = false
    }
    isChannelSidebarOpen.value = !isChannelSidebarOpen.value
  } else {
    // For desktop, we toggle between expanded and collapsed states
    isChannelSidebarOpen.value = !isChannelSidebarOpen.value
  }
}

const toggleMembersSidebar = () => {
  if (isMobileView.value && isChannelSidebarOpen.value) {
    isChannelSidebarOpen.value = false
  }
  isMembersSidebarOpen.value = !isMembersSidebarOpen.value
}

// Change active items
const setActiveServer = (id) => {
  if (id == null) {
    isDashboardOpen.value = true
    activeServer.value = null
    return
  }
  isDashboardOpen.value = false
  activeServer.value = id
  messagesStore.clearMessages()
  loadChannelsForServer(id)
  setActiveChannel(channels.value[0].id)

  // Close sidebars on mobile when changing server
  if (isMobileView.value) {
    isChannelSidebarOpen.value = false
    isMembersSidebarOpen.value = false
  }
}

const setActiveChannel = (id) => {
  activeChannel.value = id
  messagesStore.clearMessages()
  loadMessagesForChannel(activeChannel.value)

  // Close sidebars on mobile when changing channel
  if (isMobileView.value) {
    isChannelSidebarOpen.value = false
    isMembersSidebarOpen.value = false
  }
}

const setActiveUser = (id) => {
  activeUser.value = id

  // Close sidebars on mobile when changing user
  if (isMobileView.value) {
    isChannelSidebarOpen.value = false
    isMembersSidebarOpen.value = false
  }
}

// Mock data
const directMessages = ref([
  { id: 1, name: 'Jane Smith', avatar: 'https://i.pravatar.cc/125', online: true },
  { id: 2, name: 'John Doe', avatar: 'https://i.pravatar.cc/112', online: true },
  { id: 3, name: 'Alice Johnson', avatar: 'https://i.pravatar.cc/125', online: false },
  { id: 4, name: 'Bob Williams', avatar: 'https://i.pravatar.cc/154', online: true }
])

const onlineMembers = ref([
  { id: 1, name: 'Jane Smith', avatar: 'https://i.pravatar.cc/115' },
  { id: 2, name: 'John Doe', avatar: 'https://i.pravatar.cc/133' },
  { id: 4, name: 'Bob Williams', avatar: 'https://i.pravatar.cc/111' },
  { id: 6, name: 'Mike Johnson', avatar: 'https://i.pravatar.cc/125' },
  { id: 7, name: 'Sarah Parker', avatar: 'https://i.pravatar.cc/100' }
])

const offlineMembers = ref([
  { id: 3, name: 'Alice Johnson', avatar: 'https://i.pravatar.cc/125' },
  { id: 5, name: 'Emma Davis', avatar: 'https://i.pravatar.cc/33' },
  { id: 8, name: 'Tom Wilson', avatar: 'https://i.pravatar.cc/125' },
  { id: 9, name: 'Lisa Brown', avatar: 'https://i.pravatar.cc/167' },
  { id: 10, name: 'David Miller', avatar: 'https://i.pravatar.cc/125' },
  { id: 11, name: 'Chris Taylor', avatar: 'https://i.pravatar.cc/134' },
  { id: 12, name: 'Olivia White', avatar: 'https://i.pravatar.cc/125' }
])

// Methods
const sendMessage = (content) => {
  if (content.trim()) {
    const messagePayload = {
      channelId: activeChannel.value,
      content: content
    }
    socket.send(JSON.stringify(messagePayload))
  }
}

onMounted(async () => {
  loading.value = true
  socket = new WebSocket('ws://localhost:8032/ws')
  await loadServersAndChannels() // Load servers on component mount
  await loadChannelsForServer(activeServer.value) // Load channels for the active server
  if (!isMobileView.value && activeServer.value != null) {
    activeChannel.value = channels.value[0].id
  }
  socket.onopen = () => {
    if (activeServer.value == null || activeChannel.value == null) return
    socket.send(
      JSON.stringify({
        type: 'get_messages',
        channelId: channels.value[0].id
      })
    )
  }

  socket.onmessage = (event) => {
    const message = JSON.parse(event.data)
    messagesStore.addMessage(message)
  }

  socket.onclose = () => {
    console.log('Real Client: WebSocket connection closed')
  }

  // Add resize event listener
  window.addEventListener('resize', handleResize)
  // Initial check
  handleResize()

  loading.value = false
})

const loadMessagesForChannel = (channelId) => {
  socket.send(
    JSON.stringify({
      type: 'get_messages',
      channelId: channelId
    })
  )
}

const loadChannelsForServer = (serverId) => {
  //get the channels from the servers
  const server = servers.value.find((server) => server.serverId === serverId)
  if (server) {
    channels.value = server.channels
  }
}

const loadServersAndChannels = async () => {
  try {
    const response = await axios.get('http://localhost:8032/get_servers')
    servers.value = response.data
    //loop through the servers and set static image
    servers.value.forEach((server) => {
      server.image = `https://api.lorem.space/image/game?w=32&h=48`
    })
  } catch (error) {
    console.error('Error loading servers:', error)
  }
}

onBeforeUnmount(() => {
  if (socket) {
    socket.close()
  }
  // Remove resize event listener
  window.removeEventListener('resize', handleResize)
})
</script>

<style>
/* Add global styles for mobile */
@media (max-width: 767px) {
  .h-screen {
    height: 100dvh; /* Use dynamic viewport height for mobile */
  }
}
</style>
