"use client"
import { useMutation, useQueryClient, useQuery } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { useAtom } from 'jotai'
import { authApi } from '@/lib/api'
import { authAtom } from '@/store/auth'
import { useToast } from '@/hooks/use-toast'
import type { LoginRequest, RegisterRequest } from '@/types/auth'

export function useLogin() {
  const [_, setAuth] = useAtom(authAtom)
  const router = useRouter()
  const { toast } = useToast()
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: LoginRequest) => authApi.login(data),
    onSuccess: (response) => {
      setAuth({
        user: response.user,
        token: response.token,
        isAuthenticated: true,
      })
      queryClient.invalidateQueries()
      toast({
        title: "Success",
        description: "Successfully logged in",
      })
      router.push('/dashboard')
    },
    onError: (error: any) => {
      toast({
        variant: "destructive",
        title: "Error",
        description: error.response?.data?.error || "Invalid credentials",
      })
    },
  })
}

export function useRegister() {
  const router = useRouter()
  const { toast } = useToast()

  return useMutation({
    mutationFn: (data: RegisterRequest) => authApi.register(data),
    onSuccess: () => {
      toast({
        title: "Success",
        description: "Account created successfully. Please login.",
      })
      router.push('/login')
    },
    onError: (error: any) => {
      toast({
        variant: "destructive",
        title: "Error",
        description: error.response?.data?.error || "Registration failed",
      })
    },
  })
}

export function useLogout() {
  const [_, setAuth] = useAtom(authAtom)
  const router = useRouter()
  const { toast } = useToast()
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: authApi.logout,
    onSuccess: () => {
      setAuth({ user: null, token: null, isAuthenticated: false })
      queryClient.clear()
      toast({
        title: "Success",
        description: "Successfully logged out",
      })
      router.push('/login')
    },
    onError: (error: any) => {
      toast({
        variant: "destructive",
        title: "Error",
        description: "Failed to logout",
      })
    },
  })
}
