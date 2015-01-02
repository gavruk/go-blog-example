package routes

import (
	"fmt"

	"github.com/gavruk/go-blog-example/db/documents"
	"github.com/gavruk/go-blog-example/models"
	"github.com/gavruk/go-blog-example/session"

	"github.com/martini-contrib/render"

	"labix.org/v2/mgo"
)

func IndexHandler(rnd render.Render, s *session.Session, db *mgo.Database) {
	fmt.Println(s.Username)

	postDocuments := []documents.PostDocument{}
	postsCollection := db.C("posts")
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts = append(posts, post)
	}

	model := models.PostListModel{}
	model.IsAuthorized = s.IsAuthorized
	model.Posts = posts

	rnd.HTML(200, "index", model)
}
