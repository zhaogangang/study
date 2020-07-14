package data

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	uuid "github.com/nu7hatch/gouuid"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp password=gwp dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() string {
	// use third package
	u1, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	return u1.String()
	/*
		 * original
		u := new([16]byte)
		_, err := rand.Read(u[:])
		if err != nil {
			log.Fatalln("Cannot generate UUID", err)
		}

		// 0x40 is reserved variant from RFC 4122
		u[8] = (u[8] | 0x40) & 0x7F
		// Set the four most significant bits (bits 12 throught 15) of the
		// time_hi_and_version field to the 4-bit version number.
		u[6] = (u[6] & 0xF) | 0x40
		uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
		return
	*/
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
