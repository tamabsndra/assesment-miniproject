"use client"

import { useAtom } from 'jotai'
import { useLogout } from '@/hooks/use-auth'
import { authAtom } from '@/store/auth'
import { Button } from '@/components/ui/button'
import { LogOut, User } from 'lucide-react'
import Link from 'next/link'

export function Header() {
  const [auth] = useAtom(authAtom)
  const { mutate: logout } = useLogout()

  return (
    <header className="bg-white border-b">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <Link href="/dashboard">
              <h1 className="text-xl font-bold">Dashboard</h1>
            </Link>
          </div>

          <div className="flex items-center space-x-4">
            <div className="flex items-center space-x-2">
              <User className="h-5 w-5 text-gray-500" />
              <span className="text-gray-700">{auth.user?.name}</span>
            </div>

            <Button
              variant="outline"
              size="sm"
              onClick={() => logout()}
            >
              <LogOut className="h-4 w-4 mr-2" />
              Logout
            </Button>
          </div>
        </div>
      </div>
    </header>
  )
}
