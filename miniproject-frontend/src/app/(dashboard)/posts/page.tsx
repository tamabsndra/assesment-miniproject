import Link from 'next/link'
import { PostList } from '@/components/posts/post-list'
import { Button } from '@/components/ui/button'
import { Plus } from 'lucide-react'

export default function PostsPage() {
  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Posts</h1>
        <Link href="/posts/create">
          <Button>
            <Plus className="mr-2 h-4 w-4" />
            Create Post
          </Button>
        </Link>
      </div>
      <PostList />
    </div>
  )
}
