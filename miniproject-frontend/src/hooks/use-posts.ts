"use client"
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { postApi } from '@/lib/api'
import { useToast } from '@/hooks/use-toast'
import type { CreatePostData, UpdatePostData } from '@/types/post'

export function useGetPosts() {
  return useQuery({
    queryKey: ['posts'],
    queryFn: postApi.getAll,
  })
}

export function useGetPost(id: number) {
  return useQuery({
    queryKey: ['posts', id],
    queryFn: () => postApi.getById(id),
    enabled: !!id,
  })
}

export function useGetMyPosts() {
  return useQuery({
    queryKey: ['posts', 'my-posts'],
    queryFn: postApi.getMyPosts,
  })
}

export function useCreatePost() {
  const queryClient = useQueryClient()
  const router = useRouter()
  const { toast } = useToast()

  return useMutation({
    mutationFn: (data: CreatePostData) => postApi.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['posts'] })
      toast({
        title: "Success",
        description: "Post created successfully",
      })
      router.push('/posts')
    },
    onError: (error: any) => {
      toast({
        variant: "destructive",
        title: "Error",
        description: error.response?.data?.error || "Failed to create post",
      })
    },
  })
}

export function useUpdatePost(id: number) {
  const queryClient = useQueryClient()
  const router = useRouter()
  const { toast } = useToast()

  return useMutation({
    mutationFn: (data: UpdatePostData) => postApi.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['posts', id] })
      queryClient.invalidateQueries({ queryKey: ['posts'] })
      toast({
        title: "Success",
        description: "Post updated successfully",
      })
      router.push('/posts')
    },
    onError: (error: any) => {
      toast({
        variant: "destructive",
        title: "Error",
        description: error.response?.data?.error || "Failed to update post",
      })
    },
  })
}

export function usePostWithUser() {
  return useQuery({
    queryKey: ['posts', 'postsWithUser'],
    queryFn: postApi.getWithUser,
  })
}

export function useDeletePost() {
  const queryClient = useQueryClient()
  const { toast } = useToast()

  return useMutation({
    mutationFn: postApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['posts'] })
      toast({
        title: "Success",
        description: "Post deleted successfully",
      })
    },
    onError: (error: any) => {
      toast({
        variant: "destructive",
        title: "Error",
        description: error.response?.data?.error || "Failed to delete post",
      })
    },
  })
}
