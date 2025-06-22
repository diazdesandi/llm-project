export interface Message {
  role: Role,
  content: string,
}

type Role = "agent" | "user"