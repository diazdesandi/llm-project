export interface Message {
  id: string,
  conversationId: string,
  role: Role,
  content: string,
  createdAt: string,
}

type Role = "agent" | "user"

export interface Conversation {
  id: string,
  title?: string,
  model: string,
  messages: Message[]
  createdAt: string,
  updatedAt: string,
  userId?: string;
  isArchived: boolean
}