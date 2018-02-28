package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func main() {
	db, err := sql.Open("mysql", "root:Welcome1@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 0)
	defer rows.Close()
	// var usu usuario

	var usu []usuario
	for rows.Next() {
		var u usuario
		rows.Scan(&u.ID, &u.Nome)
		fmt.Println(u)
		// usu = usuario{u.Id, u.Nome}
		usu = append(usu, u)
	}

	p1Json, _ := json.Marshal(usu)
	fmt.Println(string(p1Json))

}
