'use client'

import { useEffect } from 'react'
import { useAtom } from 'jotai'
import { useRouter, usePathname } from 'next/navigation'
import { authAtom } from '@/store/auth'
import { storage } from '@/lib/storage'

const publicRoutes = ['/login', '/register']

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [auth, setAuth] = useAtom(authAtom)
  const router = useRouter()
  const pathname = usePathname()

  useEffect(() => {
    const token = storage.getToken()
    const user = storage.getUser()

    if (token && user) {
      setAuth({ token, user, isAuthenticated: true })

      if (publicRoutes.includes(pathname)){
          router.push('/dashboard')
      }

    }

    if (!token && !user && !publicRoutes.includes(pathname)) {
      router.push('/login')
    }
  }, [setAuth, router, pathname])

  return<>{children}</>
}
