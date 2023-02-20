package article

import (
	"article-service/event"
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

// Service struct is used inject dependencies inside the article package
type Service struct {
	*sql.DB
	*redis.Client
}

func (s *Service) AddArticle(na NewArticle) (Article, error) {

	//inserting the data in db for specific user
	const q = `INSERT INTO articles
		(author_email, title, content)
		VALUES ( $1, $2, $3)
		Returning article_id`

	//starting the transaction in postgres
	tx, err := s.DB.Begin()
	if err != nil {
		return Article{}, fmt.Errorf("starting transaction %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//exec the query // QueryRowContext is used when we are expecting one row back in the result
	row := tx.QueryRowContext(ctx, q, na.AuthorEmail, na.Title, na.Content) // using tx.QueryRow to exec the query inside a transaction

	var id int
	err = row.Scan(&id) // storing article_id return from postgres in id var
	if err != nil {
		log.Println(err)
		return Article{}, fmt.Errorf("inserting Article %w", err)
	}

	//converting id to string
	articleId := strconv.Itoa(id)

	//setting the values in the Article struct
	a := Article{
		ArticleId:   articleId,
		Title:       na.Title,
		Content:     na.Content,
		AuthorEmail: na.AuthorEmail,
	}

	//New article is published, so we will generate an event and store it in redis
	err = event.CreateArticleEvent(s.Client, a.ArticleId, a.Title, a.AuthorEmail)
	if err != nil {
		tx.Rollback() // if we have a problem while generating the event we will rollback the transaction
		return Article{}, fmt.Errorf("error in redis event %w", err)
	}

	//if everything is good we will commit the new post in the db
	err = tx.Commit()
	if err != nil {
		return Article{}, fmt.Errorf("error in commit %w", err)
	}

	//returning article
	return a, nil
}

func (s *Service) ListArticles(email string) ([]Article, error) {

	var articles []Article
	var a Article
	const q = "Select author_email,article_id,title,content FROM articles where author_email = $1"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//QueryContext return multiple rows at a time
	rows, err := s.DB.QueryContext(ctx, q, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//looping over the rows returned by the QueryContext
	for rows.Next() {

		//scanning the value in the article struct
		err = rows.Scan(&a.AuthorEmail, &a.ArticleId, &a.Title, &a.Content)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		//appending the article struct to the slice of the article struct to create a list of articles that user has published
		articles = append(articles, a)

	}

	//returning the list of the articles
	return articles, nil

}
