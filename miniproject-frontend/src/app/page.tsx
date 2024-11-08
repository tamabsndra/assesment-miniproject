"use client"
import { redirect } from 'next/navigation'
import { isAuthenticated } from '@/lib/auth'

export default async function HomePage() {
  const session = await isAuthenticated

  if (!session) {
    redirect('/login')
  } else {
    redirect('/dashboard')
  }
}
