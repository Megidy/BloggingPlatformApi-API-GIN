package models

import (
	"database/sql"
	"log"

	"github.com/Megidy/BloggingPlatform-Api/pkj/config"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDb()
}
func CreatePost(post *Post) (*Post, error) {
	_, err := db.Exec("INSERT INTO posts (title, content, category) VALUES (?, ?, ?)",
		post.Title, post.Content, post.Category)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return post, nil
}
func GetAllPosts() ([]Post, error) {
	var posts []Post
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Category)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
func GetPostById(Id int64) (*Post, *sql.DB, error) {
	var post Post
	row := db.QueryRow("SELECT * FROM posts WHERE ID=?", Id)

	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Category)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	return &post, nil, nil
}

func DeletePost(Id int64) (*Post, error) {
	var post Post
	row := db.QueryRow("SELECT * FROM posts WHERE ID=?", Id)
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Category)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	_, err = db.Query("DELETE FROM posts WHERE ID=?", Id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &post, nil
}
