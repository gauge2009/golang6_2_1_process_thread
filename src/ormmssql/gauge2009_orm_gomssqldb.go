package ormmssql

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/golang-sql/civil"
	"log"
	"net/url"
	"strconv"
	"time"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "sparksubmit666", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "192.168.1.7\\hive", "the database server")
	user          = flag.String("user", "sa", "the database user")
	instance      = "hive"
	database      = flag.String("database", "ai_company", "the table name")
)

const (
	createTableSql      = "CREATE TABLE TestAnsiNull (bitcol bit, charcol char(1));"
	dropTableSql        = "IF OBJECT_ID('TestAnsiNull', 'U') IS NOT NULL DROP TABLE TestAnsiNull;"
	insertQuery1        = "INSERT INTO TestAnsiNull VALUES (0, NULL);"
	insertQuery2        = "INSERT INTO TestAnsiNull VALUES (1, 'a');"
	selectNullFilter    = "SELECT bitcol FROM TestAnsiNull WHERE charcol = NULL;"
	selectNotNullFilter = "SELECT bitcol FROM TestAnsiNull WHERE charcol <> NULL;"
)

func makeConnURL() *url.URL {
	return &url.URL{
		Scheme: "sqlserver",
		//Host:   *server+"/"+ instance + ":" + strconv.Itoa(*port),
		Host: *server + "/" + instance + ":" + strconv.Itoa(*port),
		User: url.UserPassword(*user, *password),
	}
}

//

//func main() {
//
//
//	//connector_demo()
//
//    Query_table();
//
//
//}

///█ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ 查询数据域操作数据  █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
///This example shows how to insert and retrieve date and time types data
func Query_table() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	//connString := makeConnURL().String()
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", *server, *port, *database, *user, *password)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	// Create a new connector object by calling NewConnector
	connector, err := mssql.NewConnector(connString)
	if err != nil {
		log.Println(err)
		return
	}

	// Use SessionInitSql to set any options that cannot be set with the dsn string
	// With ANSI_NULLS set to ON, compare NULL data with = NULL or <> NULL will return 0 rows
	connector.SessionInitSQL = "SET ANSI_NULLS ON"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()

	//createDateTime(db)
	//insertDateTime(db)
	UpdateDateTime(db)
	retrieveDateTime(db)
	//createStoreprodure(db)
	retrieveDateTimeOutParam(db)

}

func createDateTime(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE datetimeTable (timeCol TIME, dateCol DATE, smalldatetimeCol SMALLDATETIME, datetimeCol DATETIME, datetime2Col DATETIME2, datetimeoffsetCol DATETIMEOFFSET)")
	if err != nil {
		log.Fatal(err)
	}

}

//█ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ Insert   █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
func insertDateTime(db *sql.DB) {

	stmt, err := db.Prepare("INSERT INTO datetimeTable VALUES(@p1, @p2, @p3, @p4, @p5, @p6)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	tin, err := time.Parse(time.RFC3339, "2006-01-02T22:04:05.787-07:00")
	if err != nil {
		log.Fatal(err)
	}
	var timeCol civil.Time = civil.TimeOf(tin)
	var dateCol civil.Date = civil.DateOf(tin)
	var smalldatetimeCol string = "2006-01-02 22:04:00"
	var datetimeCol mssql.DateTime1 = mssql.DateTime1(tin)
	var datetime2Col civil.DateTime = civil.DateTimeOf(tin)
	var datetimeoffsetCol mssql.DateTimeOffset = mssql.DateTimeOffset(tin)
	_, err = stmt.Exec(timeCol, dateCol, smalldatetimeCol, datetimeCol, datetime2Col, datetimeoffsetCol)
	if err != nil {
		log.Fatal(err)
	}
}

//█ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ Insert   █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
func UpdateDateTime(db *sql.DB) {

	stmt, err := db.Prepare("UPDATE datetimeTable set  smalldatetimeCol =@p1, datetimeCol =  @p2, datetime2Col = @p3")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	tin, err := time.Parse(time.RFC3339, "2021-08-13T14:55:05.000-08:00")
	if err != nil {
		log.Fatal(err)
	}
	//var timeCol civil.Time = civil.TimeOf(tin)
	//var dateCol civil.Date = civil.DateOf(tin)
	var smalldatetimeCol string = "2021-08-13 14:55:00"
	var datetimeCol mssql.DateTime1 = mssql.DateTime1(tin)
	var datetime2Col civil.DateTime = civil.DateTimeOf(tin)
	//var datetimeoffsetCol mssql.DateTimeOffset = mssql.DateTimeOffset(tin)
	_, err = stmt.Exec(smalldatetimeCol, datetimeCol, datetime2Col)
	if err != nil {
		log.Fatal(err)
	}
}

//█ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ select  █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
func retrieveDateTime(db *sql.DB) {
	rows, err := db.Query("SELECT timeCol, dateCol, smalldatetimeCol, datetimeCol, datetime2Col, datetimeoffsetCol FROM datetimeTable")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var c1, c2, c3, c4, c5, c6 time.Time
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("c1: %+v; c2: %+v; c3: %+v; c4: %+v; c5: %+v; c6: %+v;\n", c1, c2, c3, c4, c5, c6)
	}
}

func createStoreprodure(db *sql.DB) {
	CreateProcSql := `
	CREATE PROCEDURE OutDatetimeProc
		@timeOutParam TIME OUTPUT,
		@dateOutParam DATE OUTPUT,
		@smalldatetimeOutParam SMALLDATETIME OUTPUT,
		@datetimeOutParam DATETIME OUTPUT,
		@datetime2OutParam DATETIME2 OUTPUT,
		@datetimeoffsetOutParam DATETIMEOFFSET OUTPUT
	AS
		SET NOCOUNT ON
		SET @timeOutParam = '22:04:05.7870015'
		SET @dateOutParam = '2006-01-02'
		SET @smalldatetimeOutParam = '2006-01-02 22:04:00'
		SET @datetimeOutParam = '2006-01-02 22:04:05.787'
		SET @datetime2OutParam = '2006-01-02 22:04:05.7870015'
		SET @datetimeoffsetOutParam = '2006-01-02 22:04:05.7870015 -07:00'`
	_, err := db.Exec(CreateProcSql)
	if err != nil {
		log.Fatal(err)
	}
}

func retrieveDateTimeOutParam(db *sql.DB) {
	var (
		timeOutParam, datetime2OutParam, datetimeoffsetOutParam mssql.DateTimeOffset
		dateOutParam, datetimeOutParam                          mssql.DateTime1
		smalldatetimeOutParam                                   string
	)
	_, err := db.Exec("OutDatetimeProc",
		sql.Named("timeOutParam", sql.Out{Dest: &timeOutParam}),
		sql.Named("dateOutParam", sql.Out{Dest: &dateOutParam}),
		sql.Named("smalldatetimeOutParam", sql.Out{Dest: &smalldatetimeOutParam}),
		sql.Named("datetimeOutParam", sql.Out{Dest: &datetimeOutParam}),
		sql.Named("datetime2OutParam", sql.Out{Dest: &datetime2OutParam}),
		sql.Named("datetimeoffsetOutParam", sql.Out{Dest: &datetimeoffsetOutParam}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("timeOutParam: %+v; dateOutParam: %+v; smalldatetimeOutParam: %s; datetimeOutParam: %+v; datetime2OutParam: %+v; datetimeoffsetOutParam: %+v;\n",
		time.Time(timeOutParam), time.Time(dateOutParam), smalldatetimeOutParam, time.Time(datetimeOutParam), time.Time(datetime2OutParam), time.Time(datetimeoffsetOutParam))
}

/// █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ 连接数据库、运行DDL 语言建表
/// █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ 连接数据库、运行DDL 语言建表
/// █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ 连接数据库、运行DDL 语言建表
/// █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ 连接数据库、运行DDL 语言建表
/// █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ 连接数据库、运行DDL 语言建表
/// █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ 连接数据库、运行DDL 语言建表
func connector_demo() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	//connString := makeConnURL().String()
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", *server, *port, *database, *user, *password)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	// Create a new connector object by calling NewConnector
	connector, err := mssql.NewConnector(connString)
	if err != nil {
		log.Println(err)
		return
	}

	// Use SessionInitSql to set any options that cannot be set with the dsn string
	// With ANSI_NULLS set to ON, compare NULL data with = NULL or <> NULL will return 0 rows
	connector.SessionInitSQL = "SET ANSI_NULLS ON"

	// Pass connector to sql.OpenDB to get a sql.DB object
	db := sql.OpenDB(connector)
	defer db.Close()

	// Create and populate table
	_, err = db.Exec(createTableSql)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Exec(dropTableSql)
	_, err = db.Exec(insertQuery1)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = db.Exec(insertQuery2)
	if err != nil {
		log.Println(err)
		return
	}

	var bitval bool
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// (*Row) Scan should return ErrNoRows since ANSI_NULLS is set to ON
	err = db.QueryRowContext(ctx, selectNullFilter).Scan(&bitval)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println(err)
			return
		}
	} else {
		log.Println("Expects an ErrNoRows error. No error is returned")
		return
	}

	// (*Row) Scan should return ErrNoRows since ANSI_NULLS is set to ON
	err = db.QueryRowContext(ctx, selectNotNullFilter).Scan(&bitval)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println(err)
			return
		}
	} else {
		log.Println("Expects an ErrNoRows error. No error is returned")
		return
	}

	// Set ANSI_NULLS to OFF
	connector.SessionInitSQL = "SET ANSI_NULLS OFF"

	// (*Row) Scan should copy data to bitval
	err = db.QueryRowContext(ctx, selectNullFilter).Scan(&bitval)
	if err != nil {
		log.Println(err)
		return
	}
	if bitval != false {
		log.Println("Incorrect value retrieved.")
		return
	}

	// (*Row) Scan should copy data to bitval
	err = db.QueryRowContext(ctx, selectNotNullFilter).Scan(&bitval)
	if err != nil {
		log.Println(err)
		return
	}
	if bitval != true {
		log.Println("Incorrect value retrieved.")
		return
	}

}
func execute_sp() {
	//	sqltextcreate := `
	//CREATE PROCEDURE spwithoutputandrows
	//	@bitparam BIT OUTPUT
	//AS BEGIN
	//	SET @bitparam = 1
	//	SELECT 'Row 1'
	//END
	username := "sa"
	password := "sparksubmit666"
	hostname := "192.168.1.7\\hive"
	query := url.Values{}
	query.Add("app name", "MyAppName")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%d", hostname, port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		fmt.Println("sql open failed: ", err)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bitout int64
	rows, err := db.QueryContext(ctx, "spwithoutputandrows", sql.Named("bitparam", sql.Out{Dest: &bitout}))
	var strrow string
	for rows.Next() {
		err = rows.Scan(&strrow)
	}
	fmt.Printf("bitparam is %d", bitout)
}
