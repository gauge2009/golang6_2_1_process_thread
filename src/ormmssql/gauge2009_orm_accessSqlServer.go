package ormmssql

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-adodb"
)

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA
}

type SA struct {
	user   string
	passwd string
	port   int
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.dataSource)
	if m.windows {
		// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器(需要在安装sqlserver时候设置)，
		// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "Initial Catalog="+m.database)
	conf = append(conf, "user id="+m.sa.user)
	conf = append(conf, "password="+m.sa.passwd)
	conf = append(conf, "port="+fmt.Sprint(m.sa.port))

	//	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}

//func main() {
//	//access_mssql();
//	access_adodb()
//	execute_sp()
//}

func Access_adodb() {
	db := Mssql{
		dataSource: "192.168.1.7\\HIVE",
		database:   "base_inspection_company",
		// windwos: true 为windows身份验证，false 必须设置sa账号和密码
		windows: false,
		sa: SA{
			user:   "SA",
			passwd: "sparksubmit666",
			port:   1433,
		},
	}
	// 连接数据库
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}
	defer db.Close()

	// 执行SQL语句
	rows, err := db.Query("SELECT '█ █ █ ats_result' as tb_name,* FROM  dbo.ats_result")
	if err != nil {
		fmt.Println("query: ", err)
		return
	}
	for rows.Next() {
		var key_id string
		var emp_id string
		//var ats_date mssql.DateTime1
		rows.Scan(&key_id, &emp_id)
		fmt.Printf("key_id: %s \t emp_id: %d\n", key_id, emp_id)
	}
}

func Access_mssql() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")
}
