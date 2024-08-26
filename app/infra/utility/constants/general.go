package constants

const FAILED_CONNECT_DB = "Failed to connect to database!"

type MongoTypeOrder int

const (
	MongoOrderAsc  MongoTypeOrder = 1
	MongoOrderDesc MongoTypeOrder = -1
)

const (
	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)
