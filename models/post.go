package models

import (
	"log"
	"time"
)

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (post *Post) CreatedAtDate() (date string) {
	date = post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
	return
}

// get post's owner
func (post *Post) User() (user User, err error) {
	user = User{}
	err = Db.
		QueryRow("select id,uuid,name,email,created_at from users where id = ?", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Println("Cannot get post's user: " + err.Error())
	}
	return
}
