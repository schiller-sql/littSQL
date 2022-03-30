package sqlExecutor

type Repository interface {
	ExecuteSQLite(sqlString string) (*[]byte, error)
}
