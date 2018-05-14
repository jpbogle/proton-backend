package controllers

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:sharpiepop@/proton?allowNativePasswords=true")
	// db, err = sql.Open("mysql", "root:sharpiepop@/proton?charset=utf8mb4&collation=utf8mb4_unicode_ci")
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}
	db.SetMaxIdleConns(100)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to 'proton' mysqldb")
	createTables()
}

func createTables() {
	ALREADY_CREATED_ERR := uint16(1050)
	var err error

	_, err = db.Query("CREATE TABLE sheets (id INT unsigned NOT NULL AUTO_INCREMENT, name varchar(30) NOT NULL, owner varchar(30) NOT NULL, PRIMARY KEY (id), UNIQUE KEY (name, owner));")
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok && sqlErr.Number != ALREADY_CREATED_ERR {
			log.Printf("Problem creating table 'sheets'... %s", err)
		}
	}

	_, err = db.Query("CREATE TABLE transactions (id INT unsigned NOT NULL AUTO_INCREMENT, sheetid INT unsigned NOT NULL, target varchar(30) NOT NULL, source varchar(30) NOT NULL, timestamp INT unsigned, amount INT unsigned NOT NULL, note varchar(200), action varchar(30) NOT NULL, completed BOOLEAN NOT NULL, PRIMARY KEY(id), FOREIGN KEY (sheetid) REFERENCES sheets(id));")
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok && sqlErr.Number != ALREADY_CREATED_ERR {
			log.Printf("Problem creating table 'transcations'... %s", err)
		}
	}
}

func DropTables() {
	db.Query("drop table sheets;")
	db.Query("drop table transcations;")
	db.Close()
}

func Close() {
	db.Close()
}
