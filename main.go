package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"

	"fmt"
)

type testtable struct {
	username string
	password string
}

const (
	host     = "localhost"
	port     = 54320
	user     = "user"
	password = "admin"
	dbname   = "dbtest"
)

func main() {
	fmt.Println("kafica.")

	connstr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbb, _ := sql.Open("postgres", connstr)
	dbb.SetMaxIdleConns(10)
	dbb.SetMaxOpenConns(10)
	dbb.SetConnMaxLifetime(0)
	a := time.Now()
	txn, err := dbb.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, _ := txn.Prepare(pq.CopyIn("testtable", "username", "password"))
	m := &testtable{
		username: "name",
		password: "pass",
	}
	mList := make([]*testtable, 0, 100)
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		mList = append(mList, m)
	}
	fmt.Println(m)
	for _, user := range mList {
		_, err := stmt.Exec(user.username, user.password)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
	delta := time.Now().Sub(a)
	fmt.Println(delta.Nanoseconds())
	fmt.Println("Program finished successfully")

}
