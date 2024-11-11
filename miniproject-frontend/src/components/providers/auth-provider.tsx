'use client'

import { useEffect, useState } from 'react'
import { useAtom } from 'jotai'
import { useRouter, usePathname } from 'next/navigation'
import { authAtom } from '@/store/auth'
import { authApi } from '@/lib/api'
import { LoadingSpinner } from '../layout/loading'

const publicRoutes = ['/login', '/register']

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [auth, setAuth] = useAtom(authAtom)
  const [isInitialized, setIsInitialized] = useState(false)
  const router = useRouter()
  const pathname = usePathname()

  useEffect(() => {
    const initializeAuth = async () => {
      if (isInitialized) return

      try {
        const data = await authApi.verifyCookieToken()
        if (data.isAuthenticated) {
          setAuth(data)
        }
        else {
          setAuth({
            user: null,
            token: null,
            isAuthenticated: false,
            isLoading: false,
          })

        }
      } catch (error) {
        setAuth({
          user: null,
          token: null,
          isAuthenticated: false,
          isLoading: false,
        })
      } finally {
        setIsInitialized(true)
      }
    }

    initializeAuth()
  }, [setAuth, isInitialized])

  useEffect(() => {
    if (!isInitialized) return

    if (!auth.isAuthenticated && !publicRoutes.includes(pathname)) {
      router.push(`/login`)
    }

    if (auth.isAuthenticated && publicRoutes.includes(pathname)) {
      router.push('/dashboard')
    }
  }, [auth.isAuthenticated, pathname, router, isInitialized])

  if (!isInitialized || auth.isLoading) {
    return (
      <div className="flex justify-center min-h-screen">
        <LoadingSpinner/>
      </div>
    )
  }

  return <>{children}</>
}
