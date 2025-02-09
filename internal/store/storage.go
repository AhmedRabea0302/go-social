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
		GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
		Activate(context.Context, string) error
		GetByUserId(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		CreateAndInvite(context.Context, *User, string, time.Duration) error
		Delete(context.Context, int64) error
	}
	Comments interface {
		GetCommentsByPostID(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}

	Followers interface {
		FollowUser(ctx context.Context, followerID, userID int64) error
		UnfollowUser(ctx context.Context, followerID, userID int64) error
	}

	Roles interface {
		GetRoleByName(ctx context.Context, name string) (*Role, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
		Roles:     &RoleStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
