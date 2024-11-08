// components/posts/post-card.tsx
import Link from 'next/link'
import { useDeletePost } from '@/hooks/use-posts'
import { formatDate } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { Pencil, Trash2 } from 'lucide-react'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import type { Post } from '@/types/post'

interface PostCardProps {
  post: Post
  showFullContent?: boolean
}

export function PostCard({ post, showFullContent = false }: PostCardProps) {
  const { mutate: deletePost, isLoading: isDeleting } = useDeletePost()

  const content = showFullContent
    ? post.content
    : post.content.slice(0, 200) + (post.content.length > 200 ? '...' : '')

  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between">
          <CardTitle>
            {showFullContent ? (
              post.title
            ) : (
              <Link
                href={`/posts/${post.id}`}
                className="hover:text-blue-600 transition-colors"
              >
                {post.title}
              </Link>
            )}
          </CardTitle>
          <div className="flex items-center space-x-2">
            <Link href={`/posts/${post.id}/edit`}>
              <Button variant="outline" size="sm">
                <Pencil className="h-4 w-4" />
              </Button>
            </Link>

            <AlertDialog>
              <AlertDialogTrigger asChild>
                <Button variant="outline" size="sm" disabled={isDeleting}>
                  <Trash2 className="h-4 w-4 text-red-500" />
                </Button>
              </AlertDialogTrigger>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>Delete Post</AlertDialogTitle>
                  <AlertDialogDescription>
                    Are you sure you want to delete this post? This action cannot be undone.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Cancel</AlertDialogCancel>
                  <AlertDialogAction
                    onClick={() => deletePost(post.id)}
                    className="bg-red-500 hover:bg-red-600"
                  >
                    {isDeleting ? 'Deleting...' : 'Delete'}
                  </AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        <div className="prose prose-sm max-w-none">
          <p>{content}</p>
        </div>
        <div className="mt-4 text-sm text-gray-500">
          Created at: {formatDate(post.created_at)}
        </div>
      </CardContent>
    </Card>
  )
}
