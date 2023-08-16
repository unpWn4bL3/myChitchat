package models

import "time"

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// check session validation in db
func (session *Session) Check() (valid bool, err error) {
	err = Db.
		QueryRow("select id,uuid,email,user_id,created_at from sessions where uuid = ?", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// delete session from db
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// get user from session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.
		QueryRow("select id,uuid,name,email,created_at from users where id = ?", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// delete all sessions
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}
