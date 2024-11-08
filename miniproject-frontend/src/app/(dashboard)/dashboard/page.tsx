"use client"
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { useGetMyPosts, usePostWithUser, useDeletePost } from '@/hooks/use-posts'
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
import Link from 'next/link'
import { Button } from '@/components/ui/button'
import { Pencil, Trash2 } from 'lucide-react'
import { useAtom } from 'jotai'
import { authAtom } from '@/store/auth'

export default function DashboardPage() {
  const { data: postWithUser } = usePostWithUser()
  const { data: myPost} = useGetMyPosts()
  const { mutate: deletePost } = useDeletePost()
  const [auth] = useAtom(authAtom)

  const todayPosts = postWithUser?.filter((postTday) => {
    const today = new Date()
    const postDate = new Date(postTday.created_at)
    return today.toDateString() === postDate.toDateString()
  })


  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      <Card>
        <CardHeader>
          <CardTitle className="text-sm font-medium">Total Posts</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{postWithUser?.length || 0}</div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className="text-sm font-medium">Total My Post</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{myPost?.length || 0}</div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className="text-sm font-medium">Posts Today</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{todayPosts?.length || 0}</div>
        </CardContent>
      </Card>

      <div className='col-span-3'>
        <h1 className="text-2xl font-bold mb-4">All Posts</h1>
        {postWithUser?.map((postdetail) => (
        <Card className='mb-4' key={postdetail.id}>
            <CardHeader>
            <div className="flex items-center justify-between">
                <CardTitle>{postdetail.title}</CardTitle>

                {postdetail.user_id === auth.user?.id && (
                <div className="flex items-center space-x-2">
                    <Link href={`/posts/${postdetail.id}/edit`}>
                    <Button variant="outline" size="sm">
                        <Pencil className="h-4 w-4" />
                    </Button>
                    </Link>

                    <AlertDialog>
                    <AlertDialogTrigger asChild>
                        <Button variant="outline" size="sm">
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
                            onClick={() => deletePost(postdetail.id)}
                            className="bg-red-500 hover:bg-red-600"
                        >
                            Delete
                        </AlertDialogAction>
                        </AlertDialogFooter>
                    </AlertDialogContent>
                    </AlertDialog>
                </div>
                )}
            </div>
            </CardHeader>
            <CardContent>
            <p className="text-gray-600">{postdetail.content}</p>
            <div className="mt-2 text-sm text-gray-500">
                <p>Author: {postdetail.user.name}</p>
                <p>Created at: {new Date(postdetail.created_at).toLocaleDateString()}</p>
            </div>
            </CardContent>
        </Card>
        ))}

      </div>


    </div>
  )
}
