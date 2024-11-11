import { atom } from 'jotai'
import type { User } from '@/types/auth'

interface AuthState {
  user: User | null
  token: string | null
  isAuthenticated: boolean
  isLoading: boolean
}

export const authAtom = atom<AuthState>({
  user: null,
  token: null,
  isAuthenticated: false,
  isLoading: true,  // Start with loading true
})
