package databases

import (
	"fmt"
	"net/url"
	"os"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/mysql"
)

func NewMigration() (db *dbmate.DB, err error) {
	dbUrl, err := url.Parse(
		fmt.Sprintf("mysql://%s:%s@%s:%s/%s?charset=utf8mb4&parseTime=True&loc=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_TIMEZONE"),
		),
	)
	if err != nil {
		return
	}

	db = dbmate.New(dbUrl)
	db.MigrationsDir = []string{"./migrations"}

	return
}
