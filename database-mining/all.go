/*
Developer => ArkAngeL43

Type      => CLI

Defualt:
	Simple program demonstrating database mining and pillaging
	Support for mongoDB and MySQL




*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

// declaring variables
// this will be different because we will need to specify variables that are public to all classes
// since the built off code runs in different files and im rebuilding it and well merging it
// two vars will be shown one type dedicated to MySQL the other MONGODB

var (
	MySQL_Host   string
	MQL          *MySQLMiner
	MQL_schema   = new(Schema)
	c_MQL_schema string
	c_MQL_table  string
	c_MQL_column string
	MQL_table    Table
	MQL_db       Database
	s            = new(Schema)
	// previous
	MQL_Previous_schema string
	MQL_Previous_table  string
)

type DatabaseMiner interface {
	GetSchema() (*Schema, error)
}

type Schema struct {
	Databases []Database
}

type Database struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name    string
	Columns []string
}

func Search(m DatabaseMiner) error {
	s, err := m.GetSchema()
	if err != nil {
		return err
	}

	re := getRegex()
	for _, database := range s.Databases {
		for _, table := range database.Tables {
			for _, field := range table.Columns {
				for _, r := range re {
					if r.MatchString(field) {
						fmt.Println("[ INFO ] Database Name => ", database)
						fmt.Printf("[+] HIT FIELD ->  %s\n", field)
					}
				}
			}
		}
	}
	return nil
}

func getRegex() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)social`),
		regexp.MustCompile(`(?i)ssn`),
		regexp.MustCompile(`(?i)pass(word)?`),
		regexp.MustCompile(`(?i)hash`),
		regexp.MustCompile(`(?i)ccnum`),
		regexp.MustCompile(`(?i)card`),
		regexp.MustCompile(`(?i)security`),
		regexp.MustCompile(`(?i)key`),
	}
}

func (s Schema) String() string {
	var ret string
	for _, database := range s.Databases {
		ret += fmt.Sprint(database.String() + "\n")
	}
	return ret
}

func (d Database) String() string {
	ret := fmt.Sprintf("[DB] = %+s\n", d.Name)
	for _, table := range d.Tables {
		ret += table.String()
	}
	return ret
}

func (t Table) String() string {
	ret := fmt.Sprintf("    [TABLE] = %+s\n", t.Name)
	for _, field := range t.Columns {
		ret += fmt.Sprintf("       [COL] = %+s\n", field)
	}
	return ret
}

// function to handel errors
func check(err error, msg string, exit int, typer int) bool {
	if err != nil {
		if typer == 1 {
			fmt.Println("[ ERR ] FATAL: ", err, msg)
			return true
		}
		if typer == 421 {
			log.Fatal("\033[31m......[ ERR: ", err, msg)
			return true
		} else {
			panic(err)
		}
	}
	return false
}

// to take out the amount of error handeling names there will be
// im going to add a extra function that will handel the errrors

/*
With this module embeded in the file lets go ahead and merge the mongo db miner and the MySQL miner
together in the same file and add flags to specify what type of miner we will be using
*/

type MySQLMiner struct {
	Host string
	Db   sql.DB
}

func New(host string) (*MySQLMiner, error) {
	m := MySQLMiner{Host: host}
	err := MQL.connect()
	// check
	check(err, "Error during exeution for MySQL CONNECT DB LINE 79 on type module 1", 1, 1)
	return &m, nil
}

// initate a connection between you and the database
func (m *MySQLMiner) connect() error {

	db, err := sql.Open("mysql", fmt.Sprintf("root:password@tcp(%s:3306)/information_schema", MQL.Host))
	check(err, "Error occured during execution for MySQL OPEN function DB line 88 on type module 2", 1, 1)
	m.Db = *db
	return nil
}

// main MYSQL miner
func (m *MySQLMiner) GetSchema() (*Schema, error) {

	sql := `SELECT TABLE_SCHEMA, TABLE_NAME, COLUMN_NAME FROM columns
	WHERE TABLE_SCHEMA NOT IN ('mysql', 'information_schema', 'performance_schema', 'sys')
	ORDER BY TABLE_SCHEMA, TABLE_NAME`
	schemarows, err := m.Db.Query(sql)
	check(err, "Error occured during execution for SQL Query row DB func m MySQLMINER for module 2 sys 1", 1, 1)

	defer schemarows.Close()
	for schemarows.Next() {
		if err := schemarows.Scan(&c_MQL_schema, &c_MQL_table, &c_MQL_column); err != nil {
			return nil, err
		}

		if c_MQL_schema != MQL_Previous_schema {
			if MQL_Previous_schema != "" {
				MQL_db.Tables = append(MQL_db.Tables, MQL_table)
				MQL_schema.Databases = append(MQL_schema.Databases, MQL_db)
			}
			MQL_db = Database{Name: c_MQL_schema, Tables: []Table{}}
			MQL_Previous_schema = c_MQL_schema
			MQL_Previous_table = ""
		}

		if c_MQL_schema != MQL_Previous_table {
			if MQL_Previous_table != "" {
				MQL_db.Tables = append(MQL_db.Tables, MQL_table)
			}
			MQL_table = Table{Name: c_MQL_schema, Columns: []string{}}
			MQL_Previous_table = c_MQL_schema
		}
		MQL_table.Columns = append(MQL_table.Columns, c_MQL_column)
	}
	MQL_db.Tables = append(MQL_db.Tables, MQL_table)
	s.Databases = append(s.Databases, MQL_db)
	if err := schemarows.Err(); err != nil {
		return nil, err
	}

	return MQL_schema, nil
}

func main() {
	mm, err := New(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer mm.Db.Close()

	if err := Search(mm); err != nil {
		panic(err)
	}
}
