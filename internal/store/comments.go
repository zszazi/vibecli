package store

import (
	"context"
	"database/sql"
	"time"
)

type CommentStore struct {
	db *sql.DB
}

type Comment struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	UserId    int64     `json:"user_id"`
	PostId    int64     `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}

func (s *CommentStore) GetByPostId(ctx context.Context, postId int64) ([]Comment, error) {

	query := `
	SELECT c.id, c.content, c.user_id, c.post_id, c.created_at, users.username, users.id FROM comments c JOIN users on users.id = c.user_id WHERE post_id = $1 ORDER BY c.created_at DESC;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, postId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var c Comment
		c.User = User{}
		err := rows.Scan(&c.Id, &c.Content, &c.UserId, &c.PostId, &c.CreatedAt, &c.User.Username, &c.User.Id)

		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
