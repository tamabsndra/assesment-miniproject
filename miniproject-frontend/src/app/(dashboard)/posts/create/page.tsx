import { PostForm } from '@/components/posts/post-form'

export default function CreatePostPage() {
  return (
    <div className="max-w-2xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Create Post</h1>
      <PostForm />
    </div>
  )
}
