'use client'

import { useGetPost } from '@/hooks/use-posts'
import { PageLoader } from '@/components/layout/loading'
import { ErrorDisplay } from '@/components/layout/error'
import { PostForm } from '@/components/posts/post-form'
import { use } from 'react';

export default function EditPostPage({ params }: { params: { id: string } }) {
  const resolvedParams = use(params);
  const { data: post, isLoading, error } = useGetPost(Number(resolvedParams.id));

  if (isLoading) return <PageLoader />
  if (error) return <ErrorDisplay message="Failed to load post" />
  if (!post) return <ErrorDisplay message="Post not found" />

  return (
    <div className="max-w-2xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Edit Post</h1>
      <PostForm initialData={post} postId={Number(resolvedParams.id)} />
    </div>
  )
}
