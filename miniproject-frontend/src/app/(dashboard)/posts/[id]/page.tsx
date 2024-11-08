'use client'

import { useGetPost } from '@/hooks/use-posts'
import { PageLoader } from '@/components/layout/loading'
import { ErrorDisplay } from '@/components/layout/error'
import { PostCard } from '@/components/posts/post-card'

export default function PostDetailPage({ params }: { params: { id: string } }) {
  const { data: post, isLoading, error } = useGetPost(Number(params.id))

  if (isLoading) return <PageLoader />
  if (error) return <ErrorDisplay message="Failed to load post" />
  if (!post) return <ErrorDisplay message="Post not found" />

  return <PostCard post={post} showFullContent />
}
