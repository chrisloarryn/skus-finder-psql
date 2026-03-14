package postgres

import "time"

const (
	defaultMySQLDataSourceName = "root:mysql@tcp(localhost:3308)/mysqldbd?parseTime=true"
	defaultPostgresPort        = "5432"
	postgresDataSourceFormat   = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	openDBErrorFormat          = "can't open db: %v"
	postgresConnectedMessage   = "connected to postgres"
	mysqlConnectedMessage      = "connected to mysql"
	skuWhereClause             = "sku = ?"
	queryTimeout               = 15 * time.Second
)
