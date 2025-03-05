<template>
  <div class="h-screen w-screen flex flex-col bg-gray-50 text-gray-800">
    <ServerBar :servers="servers" :activeServer="activeServer" @change-server="setActiveServer" />

    <div class="flex-1 flex overflow-hidden">
      <ChannelSidebar
        :channels="channels"
        :directMessages="directMessages"
        :activeChannel="activeChannel"
        :activeUser="activeUser"
        :isOpen="isChannelSidebarOpen"
        @toggle="toggleChannelSidebar"
        @change-channel="setActiveChannel"
        @change-user="setActiveUser"
      />

      <ChatArea
        v-if="!loading"
        :messages="messagesStore.messages"
        :channels="channels"
        :activeChannel="activeChannel"
        :isChannelSidebarOpen="isChannelSidebarOpen"
        @toggle-channel-sidebar="toggleChannelSidebar"
        @toggle-members-sidebar="toggleMembersSidebar"
        @send-message="sendMessage"
      />

      <MembersSidebar
        :onlineMembers="onlineMembers"
        :offlineMembers="offlineMembers"
        :isOpen="isMembersSidebarOpen"
        @toggle="toggleMembersSidebar"
      />
    </div>

    <!-- <UserVoiceControlPanel /> -->
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import axios from 'axios'
import ServerBar from './components/ServerBar.vue'
import ChannelSidebar from './components/ChannelSidebar.vue'
import ChatArea from './components/ChatArea.vue'
import MembersSidebar from './components/MembersSidebar.vue'
import UserVoiceControlPanel from './components/UserVoiceControlPanel.vue'
import { useMessagesStore } from './stores/messages_store'

// Active states
const activeServer = ref(1)
const activeChannel = ref(1)
const activeUser = ref(null)
let socket

// UI state
const isChannelSidebarOpen = ref(true)
const isMembersSidebarOpen = ref(false)
const loading = ref(false)

// Toggle functions
const toggleChannelSidebar = () => {
  isChannelSidebarOpen.value = !isChannelSidebarOpen.value
}

const toggleMembersSidebar = () => {
  isMembersSidebarOpen.value = !isMembersSidebarOpen.value
}

// Change active items
const setActiveServer = (id) => {
  activeServer.value = id
  messagesStore.clearMessages()
  loadMessagesForServer(id)
}

const setActiveChannel = (id) => {
  activeChannel.value = id
}

const setActiveUser = (id) => {
  activeUser.value = id
}

// Mock data
const channels = ref([
  { id: 1, name: 'general', type: 'text' },
  { id: 2, name: 'announcements', type: 'text' },
  { id: 3, name: 'voice-chat', type: 'voice' },
  { id: 4, name: 'resources', type: 'text' },
  { id: 5, name: 'random', type: 'text' }
])

const directMessages = ref([
  { id: 1, name: 'Jane Smith', avatar: 'https://i.pravatar.cc/125', online: true },
  { id: 2, name: 'John Doe', avatar: null, online: true },
  { id: 3, name: 'Alice Johnson', avatar: 'https://i.pravatar.cc/125', online: false },
  { id: 4, name: 'Bob Williams', avatar: null, online: true }
])

const onlineMembers = ref([
  { id: 1, name: 'Jane Smith', avatar: 'https://i.pravatar.cc/125' },
  { id: 2, name: 'John Doe', avatar: null },
  { id: 4, name: 'Bob Williams', avatar: null },
  { id: 6, name: 'Mike Johnson', avatar: 'https://i.pravatar.cc/125' },
  { id: 7, name: 'Sarah Parker', avatar: null }
])

const offlineMembers = ref([
  { id: 3, name: 'Alice Johnson', avatar: 'https://i.pravatar.cc/125' },
  { id: 5, name: 'Emma Davis', avatar: null },
  { id: 8, name: 'Tom Wilson', avatar: 'https://i.pravatar.cc/125' },
  { id: 9, name: 'Lisa Brown', avatar: null },
  { id: 10, name: 'David Miller', avatar: 'https://i.pravatar.cc/125' },
  { id: 11, name: 'Chris Taylor', avatar: null },
  { id: 12, name: 'Olivia White', avatar: 'https://i.pravatar.cc/125' }
])

// Methods
const sendMessage = (content) => {
  if (content.trim()) {
    const messagePayload = {
      serverId: activeServer.value,
      content: content
    }
    socket.send(JSON.stringify(messagePayload))
  }
}

const messagesStore = useMessagesStore()

onMounted(() => {
  loading.value = true
  socket = new WebSocket('ws://localhost:8032/ws')
  loadServers() // Load servers on component mount

  socket.onopen = () => {
    // Request messages for the active server
    socket.send(
      JSON.stringify({
        type: 'get_messages',
        serverId: activeServer.value
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
  loading.value = false
})

const loadMessagesForServer = (serverId) => {
  socket.send(
    JSON.stringify({
      type: 'get_messages',
      serverId: serverId
    })
  )
}
const loadServers = async () => {
  try {
    const response = await axios.get('http://localhost:8032/get_servers')
    servers.value = response.data
    console.log('Servers loaded:', response.data)
  } catch (error) {
    console.error('Error loading servers:', error)
  }
}

onBeforeUnmount(() => {
  if (socket) {
    socket.close()
  }
})
</script>
