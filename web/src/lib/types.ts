export interface User {
  id: number;
  user_name: string;
  user_email: string;
  created_at?: Date;
  updated_at?: Date;
}

export interface AuthData {
  user: User | null;
  isAuthenticated: boolean;
}

export interface ShortenedUrl {
  id: number,
  OriginalURL: string,
  ShortCode: string,
  CreatedAt: Date,
  UpdatedAt: Date,
  Title: string,
  Description: string,
  OwnerID: number,
}
