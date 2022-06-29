package dbconstants

var (
	User             string            // Set a DB owner/username
	Dbn              string            // Set a DB name
	Dbt              string            // Set database type EG PostgreSQL, MySQL, MONGO
	Dbh              string            // Set DB host
	Wordlist         string            // Set a wordlist
	Login            func(string)      // Set a login function
	Thread_count     = 0               // Set a modified thread count
	Max_Thread_Count = 1000            // Set a Max thread count
	Databasechan     = make(chan bool) // Set a channel signal
)
