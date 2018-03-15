package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	_ "github.com/mattn/go-oci8"
)

var global_user = "system"
var global_pass = "Welcome1"

type registro struct {
	instancia string
	servidor  string
}

func connDB(url string) *sql.DB {
	db, err := sql.Open("oci8", url)
	if err != nil {
		//fmt.Println(err)
		return nil
	}
	//defer db.Close()

	if err = db.Ping(); err != nil {
		//fmt.Printf("Erro connectando ao database: %s\n", err)
		return nil
	}
	return db
}

func connectToDB(urlDB string, db string) *sql.DB {
	urlConn := fmt.Sprintf("%s/%s/%s", urlDB, strings.Split(db, "_")[0], db)
	dbconn := connDB(urlConn)
	if dbconn == nil {

		urlConn := fmt.Sprintf("%s/%s.sicredi.net/%s", urlDB, strings.Split(db, "_")[0], db)
		dbconn = connDB(urlConn)
		if dbconn == nil {

			urlConn = fmt.Sprintf("c##%s/%s/%s", urlDB, strings.Split(db, "_")[0], db)
			dbconn = connDB(urlConn)
			if dbconn == nil {

				urlConn = fmt.Sprintf("c##%s/%s.sicredi.net/%s", urlDB, strings.Split(db, "_")[0], db)
				dbconn = connDB(urlConn)
			}
		}
	}

	return dbconn

}

func getParam(db *sql.DB, url string, param string, scope string) string {

	var instancia string
	var servidor string
	err := db.QueryRow("select instance_name,host_name from v$instance").Scan(&instancia, &servidor)
	if err != nil {
		return fmt.Sprint(err)
	}

	//fmt.Printf("Instance %s no Servidor %s.\n", instancia, servidor)

	var valor string
	if scope == "SPFILE" {
		err = db.QueryRow("select value from v$spparameter where name=:1", param).Scan(&valor)
	} else {
		err = db.QueryRow("select value from v$parameter where name=:1", param).Scan(&valor)
	}
	if err != nil {
		//return fmt.Sprint("Aqui", err)
		valor = "NULL"
	}
	return valor

}

//
// Retorna uma lista utilizando contendo o DB e o path do oracle_home cadastrados no oratab
//
func getHomeList() map[string]string {

	dbs := make(map[string]string)

	file, err := os.Open("/etc/oratab")
	if err != nil {
		log.Fatal(err)
		return dbs
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" && strings.Index(scanner.Text(), "#") != 0 {
			dbs[strings.ToUpper(strings.Split(scanner.Text(), ":")[0])] = strings.Split(scanner.Text(), ":")[1]
			//fmt.Printf("posicao->%d/n", strings.Index(scanner.Text(), "#"))
		}

	}

	return dbs

	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}
}

// Retorna um array com SID como chave e o path do oracle_home como valor
// exemplo
//   lista := getSIDList()
//   lista['oemdb_1'] = '/oracle/product/....'
func getSIDList() map[string]string {

	instances := make(map[string]string)

	// Lista de databases do /etc/oratab
	homeList := getHomeList()

	// Lista de SIDs no DB
	out, err := exec.Command("bash", "-c", "ps -ef | grep ora_smon | grep -v grep | tr -s ' '| cut -d' ' -f8").Output()

	if err != nil {
		fmt.Println(err)
		return instances
	}

	SIDs := strings.Split(strings.Replace(string(out), "ora_smon_", "", -1), "\n")

	for _, SID := range SIDs {
		if SID != "" {
			instances[SID] = homeList[strings.ToUpper(strings.Split(SID, "_")[0])]
			//fmt.Printf("DB->%s\n", SID)
		}
	}

	return instances

}

func printSep(cnt int) {
	for i := 1; i <= cnt; i++ {
		fmt.Print("-")
	}
	fmt.Println("")

}

func printLineParam(pDb, pInst, pHome, pParam, pValMem, pValSpf string) {
	fmt.Printf("| %-10s | %-10s | %-50s | %-30s | %-30s | %-30s |\n", pDb, pInst, pHome, pParam, pValMem, pValSpf)
}

func printHeaderParam() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	nmServer, _ := os.Hostname()
	fmt.Printf("Servidor: %s\n", nmServer)
	printSep(179)
	printLineParam("DATABASE", "INSTANCE", "ORACLE HOME", "PARAMETRO", "IN MEMORY", "IN SPFILE")
	printSep(179)
}

func printFoot() {
	fmt.Println("")
	fmt.Println("*** DBACTL 1.0 (ALEXANDRE MARTI)")
	fmt.Println("")
}

func main() {

	// Valida parametros de entrada
	var isLocal bool
	flag.BoolVar(&isLocal, "local", true, "Informa para usar os bancos registrados no oratab do servidor. Ignora server, port e service")

	var servidor string
	flag.StringVar(&servidor, "server", "", "Nome ou IP do servidor. Obrigatorio quando --local=false")

	var porta string
	flag.StringVar(&porta, "port", "1521", "Porta do listener")

	var serviceList string
	flag.StringVar(&serviceList, "servicenames", "", "Lista de service names. Obrigatorio quando --local=false. Ex: srv1,srv2,srv3...")

	var param string
	flag.StringVar(&param, "param", "cpu_count", "Parametro do banco para listar. Ex: --param=sga_target")

	flag.Parse()

	if (!isLocal && (servidor == "" || serviceList == "")) || param == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	servername, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	//os.Setenv("ORACLE_HOME", "/grid/product/12.1.0")
	//os.Setenv("LD_LIBRARY_PATH", "/grid/product/12.1.0/lib")
	servername = strings.Split(servername, ".")[0]
	//fmt.Println(servername)

	printHeaderParam()
	if isLocal {
		urlDB := fmt.Sprintf("%s/%s@%s:%s", global_user, global_pass, servername, porta)
		//urlCDB := fmt.Sprintf("c##%s/%s@%s:%s", global_user, global_pass, servername, porta)
		//fmt.Println(urlDB, urlCDB)
		// Pega a lista de databases do /etc/oratab
		//homeList := getHomeList()
		//fmt.Println(getSIDList())
		for db, ohome := range getSIDList() {
			//fmt.Printf("instance-> %s (%s)\n", db, ohome)
			urlConn := fmt.Sprintf("%s/%s/%s", urlDB, strings.Split(db, "_")[0], db)
			os.Setenv("ORACLE_HOME", ohome)
			os.Setenv("LD_LIBRARY_PATH", ohome+"/lib")

			dbconn := connDB(urlConn)
			if dbconn == nil {
				//fmt.Println("Falhou", urlConn)
				urlConn = fmt.Sprintf("%s/%s.sicredi.net/%s", urlDB, strings.Split(db, "_")[0], db)
				dbconn = connDB(urlConn)
				if dbconn == nil {
					//fmt.Println("Falhou", urlConn)
					urlConn = fmt.Sprintf("c##%s/%s/%s", urlDB, strings.Split(db, "_")[0], db)
					dbconn = connDB(urlConn)
					if dbconn == nil {
						//fmt.Println("Falhou", urlConn)
						urlConn = fmt.Sprintf("c##%s/%s.sicredi.net/%s", urlDB, strings.Split(db, "_")[0], db)
						dbconn = connDB(urlConn)
					}
				}
			}
			if dbconn != nil {
				memValue := getParam(dbconn, urlConn, param, "MEMORY")
				spfValue := getParam(dbconn, urlConn, param, "SPFILE")
				//fmt.Printf("%s=%s(memory),%s(spfile) (DB: %s, SID: %s, OHOME: %s)\n", param, memValue, spfValue, strings.Split(db, "_")[0], db, ohome)
				printLineParam(strings.Split(db, "_")[0], db, ohome, param, memValue, spfValue)
				dbconn.Close()
			} else {
				fmt.Printf("DB: %s, SID: %s (Falhou nas tentativas de conectar!!)\n", strings.Split(db, "_")[0], db)
			}
		}

	} else {
		urlDB := fmt.Sprintf("%s/%s@%s:%s/", global_user, global_pass, servidor, porta)
		urlConn := fmt.Sprintf("%s%s", urlDB, serviceList)
		dbconn := connectToDB(urlConn, "")
		if dbconn != nil {
			memValue := getParam(dbconn, urlConn, param, "MEMORY")
			spfValue := getParam(dbconn, urlConn, param, "SPFILE")
			//fmt.Printf("%s=%s(memory),%s(spfile) (DB: %s, SID: %s, OHOME: %s)\n", param, memValue, spfValue, strings.Split(db, "_")[0], db, ohome)
			printLineParam("", serviceList, "", param, memValue, spfValue)
			dbconn.Close()
		} else {
			fmt.Printf("DB: %s, SID: %s (Falhou nas tentativas de conectar!!)\n", "", serviceList)
		}

	}

	printFoot()
}
