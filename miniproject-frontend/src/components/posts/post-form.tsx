"use client"

import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { PostSchema, PostFormData } from '@/validations/post'
import { useCreatePost, useUpdatePost } from '@/hooks/use-posts'
import { Button } from '@/components/ui/button'
import { RadioGroup, RadioGroupItem } from '../ui/radio-group'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { Label } from '@radix-ui/react-label'

interface PostFormProps {
  initialData?: PostFormData
  postId?: Number
}

export function PostForm({ initialData, postId }: PostFormProps) {
  const { mutate: createPost, isPending: isCreating } = useCreatePost()
  const { mutate: updatePost, isPending: isUpdating } = useUpdatePost(Number(postId || ''))

  const form = useForm<PostFormData>({
    resolver: zodResolver(PostSchema),
    defaultValues: initialData || {
      title: '',
      content: '',
      is_published: false,
    },
  })

  const onSubmit = (data: PostFormData) => {
    if (postId) {
      updatePost(data)
    } else {
      createPost(data)
    }
  }

  const isLoading = isCreating || isUpdating

  return (
    <Card>
      <CardHeader>
        <CardTitle>{postId ? 'Edit Post' : 'Create Post'}</CardTitle>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="title"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Title</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter post title" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="content"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Content</FormLabel>
                  <FormControl>
                    <Textarea
                      placeholder="Enter post content"
                      className="h-32"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
                control={form.control}
                name="is_published"
                render={({ field }) => (
                    <FormItem>
                    <FormLabel>Is Published</FormLabel>
                    <FormControl>
                        <RadioGroup
                        defaultValue={field.value ? "true" : "false"}
                        onValueChange={(value) => field.onChange(value === "true")}
                        value={field.value ? "true" : "false"}
                        >
                        <div className="flex items-center space-x-2">
                            <RadioGroupItem value="false" id="draft" />
                            <Label htmlFor="draft">Draft</Label>
                        </div>
                        <div className="flex items-center space-x-2">
                            <RadioGroupItem value="true" id="published" />
                            <Label htmlFor="published">Publish</Label>
                        </div>
                        </RadioGroup>
                    </FormControl>
                    <FormMessage />
                    </FormItem>
                )}
                />

            < Button
              type="submit"
              className="w-full"
              disabled={isLoading}
            >
              {isLoading ? 'Saving...' : postId ? 'Update Post' : 'Create Post'}
            </Button>
          </form>
        </Form>
      </CardContent>
    </Card>
  )
}
