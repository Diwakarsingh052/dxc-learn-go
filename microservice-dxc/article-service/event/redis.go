package event

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func CreateArticleEvent(r *redis.Client, articleId string, title string, authorEmail string) error {

	val := redis.XAddArgs{
		Stream: "article:add",                                       //stream name
		Values: []interface{}{"email", authorEmail, "title", title}, // event payload
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// adding event in the redis stream
	err := r.XAdd(ctx, &val).Err()

	if err != nil {
		return err
	}
	return nil
}
