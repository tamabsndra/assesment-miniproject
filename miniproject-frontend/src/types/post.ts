export interface Post {
  id: number;
  title: string;
  content: string;
  is_published: boolean;
  user_id: number;
  created_at: string;
  updated_at: string;
}

export interface PostWithAuthor extends Post {
  user: {
    id: number;
    name: string;
    email: string;
    created_at: string;
    updated_at: string;
  }
}


export interface CreatePostData {
  title: string;
  content: string;
  is_published: boolean;
}

export interface UpdatePostData {
  title: string;
  content: string;
  is_published: boolean;
}
