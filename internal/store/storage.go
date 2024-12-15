package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	QueryTimeoutDuration = 5 * time.Second
	ErrorNotFound        = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exist")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetPostByID(context.Context, int64) (*Post, error)
		UpdatePost(context.Context, *Post) error
		DeletePost(context.Context, int64) error
		GetUserFeed(context.Context, int64) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByUserId(context.Context, int64) (*User, error)
	}
	Comments interface {
		GetCommentsByPostID(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}

	Followers interface {
		FollowUser(ctx context.Context, followerID, userID int64) error
		UnfollowUser(ctx context.Context, followerID, userID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
