import { atom } from 'jotai'
import type { User } from '@/types/auth'

interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
}

const initialState: AuthState = {
  user: null,
  token: null,
  isAuthenticated: false,
  isLoading: true,
}

export const authAtom = atom<AuthState>(initialState)
