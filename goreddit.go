package goreddit

import "github.com/google/uuid"

// Thread is the basic struct for a thread
type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

// Post is the basic struct for a post
type Post struct {
	ID            uuid.UUID `db:"id"`
	ThreadID      uuid.UUID `db:"thread_id"`
	Title         string    `db:"title"`
	Content       string    `db:"content"`
	Votes         int       `db:"votes"`
	CommentsCount int       `db:"comments_count"`
}

// Comment is the basic struct for a comment
type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

// ThreadStore is the basic interface for postgres.ThreadStore
type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}

// PostStore is the basic interface for postgres.ostStore
type PostStore interface {
	Post(id uuid.UUID) (Post, error)
	PostsByThread(threadID uuid.UUID) ([]Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id uuid.UUID) error
}

// CommentStore is the basic interface for postgres.CommentStore
type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	CommentsByPost(postID uuid.UUID) ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id uuid.UUID) error
}

// Store is the wrapper for ThreadStore, PostStore, and CommentStore interfaces
type Store interface {
	ThreadStore
	PostStore
	CommentStore
}
