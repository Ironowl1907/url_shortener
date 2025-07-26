export interface User {
  id: number;
  user_name: string;
  user_email: string;
  created_at?: string;
  updated_at?: string;
}

export interface AuthData {
  user: User | null;
  isAuthenticated: boolean;
}

export interface ShortenedUrl {
  id: number,
  originalURl: string,
  ShortCode: string,
  CreatedAt: Date,
  UpdatedAt: Date,
  Title: string,
  Description: string,
  OwnerID: number,
}
