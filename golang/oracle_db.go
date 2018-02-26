package main
    
import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-oci8"
)
    

type registro struct {
    instancia string
    servidor  string
}

func main(){
    db, err := sql.Open("oci8", "amarti/Welcome1@clouddb.cbjkxisxciuj.us-east-2.rds.amazonaws.com:1521/ORCL")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
    
    if err = db.Ping(); err != nil {
        fmt.Printf("Error connecting to the database: %s\n", err)
        return
    }
    
    //rows,err := db.Query("select 2+2 from dual")
    rows,err := db.Query("select instance_name,host_name from v$instance")
    if err != nil {
        fmt.Println("Error fetching addition")
        fmt.Println(err)
        return
    }
    defer rows.Close()
    
    for rows.Next() {
        //var sum int
        var u registro
        //rows.Scan(&sum)
        rows.Scan(&u.instancia, &u.servidor)
        //fmt.Printf("2 + 2 always equals: %d\n", sum)
        fmt.Printf("Instance %s no Servidor %s.\n", u.instancia, u.servidor)
    }
}
