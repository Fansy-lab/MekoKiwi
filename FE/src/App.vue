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
        :messages="messages"
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

    <UserControlPanel />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import ServerBar from './components/ServerBar.vue'
import ChannelSidebar from './components/ChannelSidebar.vue'
import ChatArea from './components/ChatArea.vue'
import MembersSidebar from './components/MembersSidebar.vue'
import UserControlPanel from './components/UserControlPanel.vue'

// Active states
const activeServer = ref(1)
const activeChannel = ref(1)
const activeUser = ref(null)

// UI state
const isChannelSidebarOpen = ref(true)
const isMembersSidebarOpen = ref(false)

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
}

const setActiveChannel = (id) => {
  activeChannel.value = id
}

const setActiveUser = (id) => {
  activeUser.value = id
}

// Mock data
const servers = ref([
  { id: 1, name: 'Main Server', image: null },
  { id: 2, name: 'Gaming', image: '/placeholder.svg?height=40&width=40' },
  { id: 3, name: 'Study Group', image: null },
  { id: 4, name: 'Friends', image: '/placeholder.svg?height=40&width=40' },
  { id: 5, name: 'Music', image: null },
  { id: 6, name: 'Art', image: '/placeholder.svg?height=40&width=40' },
  { id: 7, name: 'Movies', image: null }
])

const channels = ref([
  { id: 1, name: 'general', type: 'text' },
  { id: 2, name: 'announcements', type: 'text' },
  { id: 3, name: 'voice-chat', type: 'voice' },
  { id: 4, name: 'resources', type: 'text' },
  { id: 5, name: 'random', type: 'text' }
])

const directMessages = ref([
  { id: 1, name: 'Jane Smith', avatar: '/placeholder.svg?height=40&width=40', online: true },
  { id: 2, name: 'John Doe', avatar: null, online: true },
  { id: 3, name: 'Alice Johnson', avatar: '/placeholder.svg?height=40&width=40', online: false },
  { id: 4, name: 'Bob Williams', avatar: null, online: true }
])

const messages = ref([
  {
    id: 1,
    user: { name: 'Jane Smith', avatar: '/placeholder.svg?height=40&width=40' },
    content: 'Hey everyone! Welcome to our new server with the cool teal theme!',
    time: 'Today at 10:30 AM'
  },
  {
    id: 2,
    user: { name: 'John Doe', avatar: null },
    content:
      'This layout is so different from Discord! I love the horizontal server bar at the top.',
    time: 'Today at 10:32 AM'
  },
  {
    id: 3,
    user: { name: 'Alice Johnson', avatar: '/placeholder.svg?height=40&width=40' },
    content: 'The teal color scheme is refreshing and easy on the eyes.',
    time: 'Today at 10:35 AM'
  },
  {
    id: 4,
    user: { name: 'Bob Williams', avatar: null },
    content: 'I like how the channels can be collapsed to save space. Very intuitive!',
    time: 'Today at 10:40 AM'
  },
  {
    id: 5,
    user: { name: 'Jane Smith', avatar: '/placeholder.svg?height=40&width=40' },
    content: 'And the rounded design elements give it a modern feel. Great job on the redesign!',
    time: 'Today at 10:42 AM'
  },
  {
    id: 6,
    user: { name: 'John Doe', avatar: null },
    content: 'The reactions feature is also a nice touch. Makes it easy to respond quickly.',
    time: 'Today at 10:45 AM'
  }
])

const onlineMembers = ref([
  { id: 1, name: 'Jane Smith', avatar: '/placeholder.svg?height=40&width=40' },
  { id: 2, name: 'John Doe', avatar: null },
  { id: 4, name: 'Bob Williams', avatar: null },
  { id: 6, name: 'Mike Johnson', avatar: '/placeholder.svg?height=40&width=40' },
  { id: 7, name: 'Sarah Parker', avatar: null }
])

const offlineMembers = ref([
  { id: 3, name: 'Alice Johnson', avatar: '/placeholder.svg?height=40&width=40' },
  { id: 5, name: 'Emma Davis', avatar: null },
  { id: 8, name: 'Tom Wilson', avatar: '/placeholder.svg?height=40&width=40' },
  { id: 9, name: 'Lisa Brown', avatar: null },
  { id: 10, name: 'David Miller', avatar: '/placeholder.svg?height=40&width=40' },
  { id: 11, name: 'Chris Taylor', avatar: null },
  { id: 12, name: 'Olivia White', avatar: '/placeholder.svg?height=40&width=40' }
])

// Methods
const sendMessage = (content) => {
  if (content.trim()) {
    messages.value.push({
      id: messages.value.length + 1,
      user: { name: 'YourUsername', avatar: null },
      content: content,
      time: 'Just now'
    })
  }
}
</script>
