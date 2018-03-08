package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/mattn/go-oci8"
)

func getDSN() string {
	// same as "sqlplus sys/syspwd@tnsentry as sysdba"
	return "sys/Welcome1@?as=sysdba"
}
func main() {
	os.Setenv("NLS_LANG", "")
	os.Setenv("ORACLE_SID", "db_1")

	db, err := sql.Open("oci8", getDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Println()
	var user string
	err = db.QueryRow("select user from dual").Scan(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successful 'as sysdba' connection. Current user is: %v\n", user)

	out, err := exec.Command("bash", "-c", "ps -ef | grep ora_smon | grep -v grep | tr -s ' '| cut -d' ' -f8").Output()

	fmt.Println(string(out))

	lista := strings.Replace(string(out), "ora_smon_", "", 1)
	fmt.Println(lista)
}
