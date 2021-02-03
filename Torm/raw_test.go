package Torm_test

import (
	"Torm"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

var db *sql.DB

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func TestMain(m *testing.M) {
	db, _ = sql.Open("mysql", "root:yxwdmysql@tcp(127.0.0.1:3306)/orm?charset=utf8mb4")
	code := m.Run()
	_ = db.Close()
	os.Exit(code)

}

func New() *Torm.Session {
	return Torm.NewSession(db)

}

func TestSession_QueryRow(t *testing.T) {
	s := New()
	var userName string
	var age int
	s = s.Raw("select user_name,age from user where user_name = ?", "yixingwei")
	res := s.QueryRow()
	if err := res.Scan(&userName, &age); err != nil {
		t.Fatal("failed to query db", err)

	}
	log.Info("userName--", userName)
	log.Info("age--", age)

}

func TestSession_Exec(t *testing.T) {
	s := New()
	key := "yixingwei"
	s = s.Raw("insert into user(user_name, age) values(?, ?)", key, 22)
	_, err := s.Exec()
	if err != nil {
		t.Fatal("failed to insert db", err)

	}

}

func TestSession_Query(t *testing.T) {
	s := New()
	var userName string
	var age int
	s = s.Raw("select user_name, age from user")
	rows, err := s.Query()
	if err != nil {
		t.Fatal("fialed to query db", err)

	}
	for rows.Next() {
		err = rows.Scan(&userName, &age)
		if err != nil {
			t.Fatal("fialed to query db", err)

		}
		log.Info("userName--", userName)
		log.Info("age--", age)

	}

}
