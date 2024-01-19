package psql

import (
	"fmt"
	"log"

	"github.com/rohitkhatri1st/Task-API/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPsql(psqlConfig *config.PsqlConfig) *gorm.DB {
	// connStr := "postgresql://postgres:gopher@localhost/todos?sslmode=disable"
	dbConfig := psqlConfig.DbConfig
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, psqlConfig.DbName, dbConfig.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDb, _ := db.DB()
	err = sqlDb.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully connected with Database %s\n", psqlConfig.DbName)
	}

	return db
}

// func TestingDbConnection() {
// query := `SELECT datname FROM pg_database;`
// // result := db.Exec(query)
// // if result.Error != nil {
// // 	log.Fatal("Failed to execute query:", result.Error)
// // }
// rows, err := db.Raw(query).Rows()
// if err != nil {
// 	log.Fatal(err)
// }
// defer rows.Close()

// var results []interface{}
// for rows.Next() {
// 	result := map[string]interface{}{}
// 	err := db.ScanRows(rows, &result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	results = append(results, result)
// }
// fmt.Println(results...)
// }
