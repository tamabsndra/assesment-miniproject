import { atom } from 'jotai'
import type { Post } from '@/types/post'

export const postsAtom = atom<Post[]>([])
export const selectedPostAtom = atom<Post | null>(null)
