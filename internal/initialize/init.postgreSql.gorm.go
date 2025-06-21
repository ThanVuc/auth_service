// Obsolete
package initialize

// It is init the PostgreSQL database connection and set up the connection pool to gorm
// But we not use gorm, we use sqlc to generate the database queries and use pgx to connect to the PostgreSQL database

// import (
// 	"authorization_service/global"
// 	po "authorization_service/internal/PO"
// 	"authorization_service/pkg/response"
// 	"fmt"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func checkInitError(err error) {
// 	if err != nil {
// 		panic("failed to initialize PostgreSQL: " + err.Error())
// 	}
// }

// func InitPostgreSQL() {
// 	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
// 	configs := global.Config.Postgres
// 	var connectString = fmt.Sprintf(dsn, configs.Host, configs.User, configs.Password, configs.Database, configs.Port)
// 	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{
// 		SkipDefaultTransaction: false,
// 	})
// 	checkInitError(err)
// 	global.PostgresDb = db

// 	// set pool
// 	setPool()
// 	migrateTables()
// 	go trackConnection()
// }

// func setPool() {
// 	postConfig := global.Config.Postgres
// 	postgresDb, err := global.PostgresDb.DB()
// 	checkInitError(err)
// 	postgresDb.SetMaxIdleConns(postConfig.MaxIdleConns)
// 	postgresDb.SetConnMaxIdleTime(time.Duration(postConfig.ConnMaxIdleTime) * time.Minute) // in seconds
// 	postgresDb.SetMaxOpenConns(postConfig.MaxOpenConns)
// 	postgresDb.SetConnMaxLifetime(time.Duration(postConfig.MaxLifetime) * time.Second.Abs()) // in seconds
// }

// func migrateTables() {
// 	err := global.PostgresDb.AutoMigrate(
// 		&po.User{},
// 	)

// 	checkInitError(err)
// 	logger := global.Logger
// 	logger.InfoString("PostgreSQL tables migrated successfully")
// }

// func trackConnection() {
// 	postgresDb, err := global.PostgresDb.DB()
// 	logger := global.Logger
// 	if err != nil {
// 		fmt.Println("failed to get PostgreSQL DB instance:", err)
// 		return
// 	}

// 	statTicker := time.NewTicker(5 * time.Second)
// 	defer statTicker.Stop()

// 	for {
// 		select {
// 		case <-statTicker.C:
// 			stats := postgresDb.Stats()
// 			if stats.OpenConnections > stats.MaxOpenConnections {
// 				logger.Error(response.InternalServerError("Overload the connections"), "routine check", nil)
// 			}
// 		}
// 	}
// }
