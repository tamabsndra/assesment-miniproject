"use client"
import { storage } from './storage'
import type { User } from '@/types/auth'

export const isAuthenticated = (): boolean => {
  const token = storage.getToken()
  return !!token
}

export const getCurrentUser = (): User | null => {
  return storage.getUser()
}

export const setAuthData = (token: string, user: User): void => {
  storage.setToken(token)
  storage.setUser(user)
}

export const clearAuthData = (): void => {
  storage.clear()
}

export const getAuthHeader = (): { Authorization: string } | {} => {
  const token = storage.getToken()
  return token ? { Authorization: `Bearer ${token}` } : {}
}

export const validateTokenExpiry = (token: string): boolean => {
  try {
    const [, payload] = token.split('.')
    const decodedPayload = JSON.parse(atob(payload))
    const expirationTime = decodedPayload.exp * 1000 // Convert to milliseconds
    return Date.now() < expirationTime
  } catch {
    return false
  }
}
