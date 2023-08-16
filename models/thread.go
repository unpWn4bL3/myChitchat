package models

import (
	"fmt"
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// get thread's date
func (thread *Thread) CreatedAtDate() (date string) {
	date = thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
	return
}

// get number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("select count(*) from posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	return
}

// get thread's posts
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id,uuid,body,user_id,thread_id,created_at from posts where thread_id = ?", thread.Id)
	if err != nil {
		log.Println("Cannot get posts: " + err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

// get all threads
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id,uuid,topic,user_id,created_at from threads order by created_at desc")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		thread := Thread{}
		err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
		if err != nil {
			return
		}
		threads = append(threads, thread)
	}
	return
}

// get a thread by uuid
func ThreadByUUID(uuid string) (thread Thread, err error) {
	thread = Thread{}
	err = Db.
		QueryRow("select id,uuid,topic,user_id,created_at from threads where uuid = ?", uuid).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	return
}

// get thread's owner
func (thread *Thread) User() (user User, err error) {
	user = User{}
	err = Db.
		QueryRow("select id,uuid,name,email,created_at from users where id = ?", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		fmt.Println("Cannot read thread's user: " + err.Error())
	}
	return
}
