package constants

import "time"

const (
	PostgresDefaultMySQLDataSourceName = "root:mysql@tcp(localhost:3308)/mysqldbd?parseTime=true"
	PostgresDefaultPort                = "5432"
	PostgresDataSourceFormat           = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	PostgresOpenDBErrorFormat          = "can't open db: %v"
	PostgresConnectedMessage           = "connected to postgres"
	MySQLConnectedMessage              = "connected to mysql"
	PostgresSKUWhereClause             = "sku = ?"
	PostgresQueryTimeout               = 15 * time.Second
)
