
import { authApi } from '@/lib/api'
import { redirect } from 'next/navigation'


export default async function HomePage() {
  const session = await authApi.verifyCookieToken()
  if (session.isAuthenticated === false) {
    redirect('/login')
  } else {
    redirect('/dashboard')
  }
}
