package models

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/unpWn4bl3/myChitchat/config"
)

var Db *sql.DB

func init() {
	config := LoadConfig()
	var err error
	driver := config.Db.Driver
	connection := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=true",
		config.Db.User,
		config.Db.Password,
		config.Db.Address,
		config.Db.Database)
	// Db, err = sql.Open("mysql", "admin:password@/chitchat?charset=utf8&parseTime=true")
	Db, err = sql.Open(driver, connection)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// create UUID in RFC 4122
func createUUID() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) string {
	cryptext := fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
