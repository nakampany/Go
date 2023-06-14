package models

import (
	"database/sql"
	"fmt"
	// postgres ドライバ
	"crypto/sha1"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	// DB接続
	Db, err := sql.Open("postgres", "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("----------------------------------")
	fmt.Println(Db)
	fmt.Println(Db.Stats())
	fmt.Println("----------------------------------")

	// Userテーブル作成
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)
}

// idをuuidに変換する関数
func createUUID() (uuid_obj uuid.UUID) {
	uuid_obj, _ = uuid.NewUUID()
	return uuid_obj
}

// パスワードをハッシュ化する関数
func Encrypt(plaintext string) (cryp_text string) {
	cryp_text = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryp_text
}
