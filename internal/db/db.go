package db

import(
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db     *goqu.Database
	SqlDB  *sql.DB
)

func ConnectDB(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"root", "root", "db", "3306", "grocery")

	var err error
	SqlDB , err = sql.Open("mysql",dsn)	
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	// Wait for DB to be ready
	for {
		err = SqlDB.Ping()
		if err == nil {
			break
		}
		log.Println("Waiting for MySQL to be ready...")
		time.Sleep(2 * time.Second)
	}

	Db = goqu.New("mysql", SqlDB)
	log.Println("✅ Connected to MySQL database!")
}