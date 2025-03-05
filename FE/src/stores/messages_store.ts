import { ref } from 'vue'
import { defineStore } from 'pinia'

interface User {
  name: string
  avatar: string | null
}

interface Message {
  id: number
  userId: string
  user: User
  timestamp: number
  messageType: number
  data: string
}

export const useMessagesStore = defineStore('messages', () => {
  const messages = ref<Message[]>([])

  function setMessages(newMessages: Message[]) {
    messages.value = newMessages
  }

  function addMessage(message: Message) {
    var user = { name: message.userId, avatar: null }
    message.user = user
    messages.value.push(message)
  }
  function clearMessages() {
    messages.value = []
  }

  return { messages, setMessages, addMessage, clearMessages }
})
