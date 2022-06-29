/*
Commet 1####

ArkAngeL43 -> Gen host
File       -> sql.go
Template   -> SUPER_SQL_AUTO_GENERARTED_FILE_MODS_FSUPER_MODS_L6_D9_ID237000000000.tmpl
Package    -> SUPER_SQL_AUTO_GENERARTED_FILE_MODS_FSUPER_MODS_L6_D9_ID237000000000
File perm  -> *
Desc       -> nil
Time       -> Tue May 17 11:09:35 PM EDT 2022
Section    -> SQL
Type       -> Go
Sig        -> 000000e0

# WARNING: IF YOU ARE NOT A DEV DO NOT TOUCH THIS FILE

*/
package SUPER_SQL_AUTO_GENERARTED_FILE_MODS_FSUPER_MODS_L6_D9_ID237000000000

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	v "main/modg/colors"
	ex "main/modg/warnings"
	SUPER_VARS "main/modules/go-main/SUPER-CONSTS"

	"gopkg.in/mgo.v2"
)

var CONNMSG_ERROR = "<RR6> TCP Error: Could not load or connect to the port, this might be good given the port is occupied by a SQL server of sort..."
var CONNMSG_PASS = "<RR6> TCP Data: Was able to load port, might not be able to ping the database, exiting....."
var Port string
var Database string
var Database_Arg sql.DB
var Connect_Str string
var Method = "tcp"
var Waiter = time.Second

func Return_DB_1(db, host string) {
	switch db {
	case "PostGreSQL":
		Port = "5432"
		Database = "PostGreSQL"
	case "MySQL":
		Port = "3306"
		Database = "MySQL"
	case "MSSQL":
		Port = "1433"
		Database = "Microsoft SQL Server"
	case "Mongo":
		Port = "27017"
		Database = "Mongo Database"
	}
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting:   Port      | ", Port)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting:   Database  | ", Database)
	connect, x := net.DialTimeout(Method, net.JoinHostPort(host, Port), Waiter)
	if x != nil {
		fmt.Println(CONNMSG_ERROR)
	} else {
		fmt.Println(CONNMSG_PASS)
	}
	defer connect.Close()
}

// Comment 2
// U: GOT ISSUE -> WARNING -> THIS IS A DATABASE PING FUNCTION, USES ACTUALL CONNECTION STRINGS WHICH MAY USE EXTRA CONFIGURATION TO SET, IF YOU DO NOT ANT TO USE RAW CONNECTION STRINGS DO NOT USE FUNCTION, IF YOU DONT HAVE A STATE TO USE THIS THEN USE 3
func Ping_DB_2(hostname, username, password, dbname, db string) {
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Pinging DB  | ", Database)
	switch db {
	case "PostGreSQL":
		Port = "5432"
		Database = "postgresql"
		parser := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)
		d, x := sql.Open("postgres", parser)
		if x != nil {
			fmt.Println("<RR6> SQL Module: Could not ping the server or even open it, got error -> ", x)
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information: Got a good code on the database ping, database is alive and well...")
		}
		defer d.Close()
	case "MySQL":
		Port = "3306"
		Database = "mysql"
		parser := username + ":" + password + "@tcp(" + hostname + ":" + Port + ")"
		d, x := sql.Open("mysql", parser)
		if x != nil {
			fmt.Println("<RR6> SQL Module: Could not ping the server or even open it, got error -> ", x)
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information: Got a good code on the database ping, database is alive and well...")
		}
		defer d.Close()
	}
}

func Ping_DB_3(hostname, db string) {
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Pinging Host| ", hostname)
	switch db {
	case "PostGreSQl":
		Port = "5432"
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Port        | ", Port)
		parser := hostname + ":" + Port
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Pinging host| ", parser)
		conn, x := net.Dial(Method, parser)
		if x != nil {
			fmt.Println("<RR6> SQL Module: Got error, could not ping the host because??? -> ", x)
			os.Exit(0)
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information: Got a good ping on the host and dial, database is alive")
		}
		defer conn.Close()
	case "MySQL":
		Port = "3306"
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Port        | ", Port)
		parser := hostname + ":" + Port
		conn, x := net.Dial(Method, parser)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Pinging host| ", parser)
		if x != nil {
			fmt.Println("<RR6> SQL Module: Got error, could not ping the host because??? -> ", x)
			os.Exit(0)
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information: Got a good ping on the host and dial, database is alive")
		}
		defer conn.Close()
	case "Mogo":
		Port = "27017"
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Port        | ", Port)
		parser := hostname + ":" + Port
		conn, x := net.Dial(Method, parser)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Pinging host| ", parser)
		if x != nil {
			fmt.Println("<RR6> SQL Module: Got error, could not ping the host because??? -> ", x)
			os.Exit(0)
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information: Got a good ping on the host and dial, database is alive")
		}
		defer conn.Close()
	case "MSSQL":
		Port = "1433"
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Port        | ", Port)
		parser := hostname + ":" + Port
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Pinging host| ", parser)
		conn, x := net.Dial(Method, parser)
		if x != nil {
			fmt.Println("<RR6> SQL Module: Got error, could not ping the host because??? -> ", x)
			os.Exit(0)
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Information: Got a good ping on the host and dial, database is alive")
		}
		defer conn.Close()
	}
}

func wordlist(list string) (string, error) {
	l, e := os.Open(list)
	if e != nil {
		return "Could not open wordlist", e
	}
	scanner := bufio.NewScanner(l)
	for scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", nil
}

func PostGreSQL_Brute(db_name, password, host, user string) {
	conn := "postgres://"
	conn += user + ":" + password
	conn += "@" + host + "/" + db_name
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Connection String: -> ", conn)
	db, e := sql.Open("postgres", conn)
	ex.Ce(e, v.REDHB, "<RR6> Database: Could not make a postgres connection", 1)
	e = db.Ping()
	if e == nil {
		fmt.Println("<RR6> Database: Was able to login with password -> ", password)
	} else {
		fmt.Println("<RR6> Database: Was not able to process a correct and direct authentication request, auth came back failed")
	}
	SUPER_VARS.Oschan <- true
}

func Mongo_Brute(password, host, user string) {
	mgo_inf := &mgo.DialInfo{
		Addrs:    []string{host},
		Timeout:  10 * time.Second,
		Username: user,
		Password: password,
	}
	_, err := mgo.DialWithInfo(mgo_inf)
	if err == nil {
		fmt.Println("<RR6> Database: Able to login with password - ", password)
	} else {
		log.Println("Error connecting to Mongo. ", err)
	}
	SUPER_VARS.Oschan <- true
}

func MySQL_Brute(dbname, host, password, user string) {
	conn := user + ":" + password
	conn += "@tcp(" + host + ")/" + dbname
	conn += "?charse=utf8"
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Connection String: -> ", conn)
	db, e := sql.Open("mysql", conn)
	ex.Ce(e, v.REDHB, "<RR6> Database: Could not make a MySQL connection", 1)
	e = db.Ping()
	if e == nil {
		fmt.Println("<RR6> Database: Was able to log into a MYSQL database with password - ", password)
	} else {
		fmt.Println("<RR6> Database: Could not authenticate password -> ", password)
	}
	SUPER_VARS.Oschan <- true
}

func brute_s(wl, user, dbtype, host, dbname string, threads int) {
	SUPER_VARS.Os_USR_BASED_MAX_THREAD_COUNT_SERVICES_BASED = threads
	pass, err := wordlist(wl)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Got return code -> ", pass)
	ex.Ce(err, v.REDHB, "<RR6> Database: Could not open the wordlist", 1)
	if SUPER_VARS.Os_STD_THREAD_COUNT >= SUPER_VARS.Os_USR_BASED_MAX_THREAD_COUNT_SERVICES_BASED {
		SUPER_VARS.Oschan <- true
		SUPER_VARS.Os_STD_THREAD_COUNT -= 1
		fmt.Println("<RR6> Hardware - digital: Subtracting thread count by 1, thread count -> ", SUPER_VARS.Os_STD_THREAD_COUNT)
	}
	switch dbtype {
	case "mongo":
		go Mongo_Brute(pass, host, user)
	case "mysql":
		go MySQL_Brute(dbname, host, pass, user)
	case "postgresql":
		go PostGreSQL_Brute(dbname, pass, host, user)
	default:
		fmt.Println("[-] Unexpected error: expecting mysql, mongo, or postgresql firm inputs")
	}
	SUPER_VARS.Os_STD_THREAD_COUNT++
}

func Parse_Server(username, wordlist, hostname, dbname, dbtype string, threads int) {
	SUPER_VARS.Oschan = make(chan bool)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: Username -> ", username)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: Wordlist -> ", wordlist)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: DB Name  -> ", dbname)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: DB Type  -> ", dbtype)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: DB Host  -> ", hostname)
	time.Sleep(1 * time.Second)
	brute_s(wordlist, username, dbtype, hostname, dbname, threads)
}
