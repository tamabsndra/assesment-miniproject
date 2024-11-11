import axios from 'axios'
import type {
  VerifyCookieTokenResponse,
  LoginRequest,
  RegisterRequest,
  AuthResponse,
  RegisterResponse,
  User
} from '@/types/auth'
import type {
  CreatePostData,
  UpdatePostData,
  Post,
  PostWithAuthor
} from '@/types/post'

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
  headers: { 'Content-Type': 'application/json' },
  withCredentials: true,
})

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export const authApi = {
  login: async (data: LoginRequest): Promise<AuthResponse> => {
    const response = await api.post<AuthResponse>('/login', data)
    return response.data
  },

  register: async (data: RegisterRequest): Promise<RegisterResponse> => {
    const response = await api.post<RegisterResponse>('/register', data)
    return response.data
  },

  logout: async (): Promise<void> => {
    const response = await api.post('/logout')
    return response.data
  },

  getCurrentUser: async (): Promise<User> => {
    const response = await api.get<User>('/me')
    return response.data
  },

  verifyCookieToken: async (): Promise<VerifyCookieTokenResponse> => {
    const response = await api.get<VerifyCookieTokenResponse>('/verify-cookie-token')
    return response.data
  },
}

export const postApi = {
  getAll: async (): Promise<Post[]> => {
    const response = await api.get<Post[]>('/posts')
    return response.data
  },

  getWithUser: async (): Promise<PostWithAuthor[]> => {
    const response = await api.get<PostWithAuthor[]>('/post-detail')
    return response.data
  },

  getById: async (id: number): Promise<Post> => {
    const response = await api.get<Post>(`/posts/${id}`)
    return response.data
  },

  getMyPosts: async (): Promise<Post[]> => {
    const response = await api.get<Post[]>(`/posts/my`)
    return response.data
  },

  create: async (data: CreatePostData): Promise<Post> => {
    const response = await api.post<Post>('/posts', data)
    return response.data
  },

  update: async (id: number, data: UpdatePostData): Promise<Post> => {
    const response = await api.put<Post>(`/posts/${id}`, data)
    return response.data
  },

  delete: async (id: number): Promise<void> => {
    await api.delete(`/posts/${id}`)
  },
}

export default api
