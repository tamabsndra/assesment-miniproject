import { z } from 'zod'

export const PostSchema = z.object({
  title: z.string()
    .min(3, 'Title must be at least 3 characters')
    .max(100, 'Title must be less than 100 characters'),
  content: z.string()
    .min(10, 'Content must be at least 10 characters')
    .max(5000, 'Content must be less than 5000 characters'),
  is_published: z.boolean(),
})

export type PostFormData = z.infer<typeof PostSchema>
