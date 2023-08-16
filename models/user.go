package models

import "time"

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values (?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Email, user.Id, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, email, user_id, created_at from sessions where uuid = ?")
	if err != nil {
		return
	}
	defer stmtin.Close()
	err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// get user session
func (user *User) Session() (session Session, err error) {
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from sessions where user_id = ?", user.Id).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// create new user
func (user *User) Create() (err error) {
	statement := "insert into users (uuid, name, email, password, created_at) values (?,?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())

	stmtout, err := Db.Prepare("select id, uuid, created_at from users where uuid = ?")
	if err != nil {
		return err
	}
	defer stmtout.Close()
	row := stmtout.QueryRow(uuid)
	err = row.Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// delete an user
func (user *User) Delete() (err error) {
	statement := "delete from users where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

// update a user
func (user *User) Update() (err error) {
	statement := "update users set name = ?, email = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return
}

// delete all users
func (user *User) DeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// get all users
func Users() (users []User, err error) {
	rows, err := Db.Query("select id,uuid,name,email,password,created_at from users")
	if err != nil {
		return
	}
	for rows.Next() { //(*sql.Rows).Next() return false when next is nothing
		user := User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// get user by email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.
		QueryRow("select id,uuid,name,email,password,created_at from users where email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// get user by uuid
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.
		QueryRow("select id,uuid,name,email,password,created_at from users where uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// create a new thread
func (user *User) CreateThread(topic string) (thread Thread, err error) {
	statement := "insert into threads (uuid,topic,user_id,created_at) value (?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, topic, user.Id, time.Now())

	stmtout, err := Db.Prepare("select id,uuid,topic,user_id,created_at from threads where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	err = stmtout.
		QueryRow(uuid).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	return
}

// create a new post to a thread
func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid,body,user_id,thread_id,created_at) value (?,?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, body, user.Id, thread.Id, time.Now())

	stmtout, err := Db.Prepare("select id,uuid,body,user_id,thread_id,created_at from posts where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	err = stmtout.
		QueryRow(uuid).
		Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}
