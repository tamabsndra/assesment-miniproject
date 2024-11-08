import { redirect } from 'next/navigation'
import { isAuthenticated } from '@/lib/auth'
import {Header} from '@/components/layout/header'
import {Sidebar} from '@/components/layout/sidebar'

export default async function DashboardLayout({
  children,
}: {
  children: React.ReactNode
}) {

  return (
    <div className="min-h-screen bg-gray-100">
      <Header />
      <div className="flex">
        <Sidebar />
        <main className="flex-1 p-6">{children}</main>
      </div>
    </div>
  )
}
