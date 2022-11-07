package postgresql

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postgre struct {
	DBHOST     string
	DBUSERNAME string
	DBPASSWORD string
	DBNAME     string
	DBPORT     string
}

func PostgreSQLConnect() string {

	psl := &postgre{}
	psl.DBHOST = os.Getenv("DB_HOST")
	psl.DBUSERNAME = os.Getenv("DB_USERNAME")
	psl.DBPASSWORD = os.Getenv("DB_PASSWORD")
	psl.DBNAME = os.Getenv("DB_NAME")
	psl.DBPORT = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", psl.DBHOST, psl.DBUSERNAME, psl.DBPASSWORD, psl.DBNAME, psl.DBPORT)

	return dsn

}
