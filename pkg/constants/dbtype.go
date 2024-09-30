// Package constants @Author hubo 2024/9/29 18:18:00
package constants

type DBType int

const (
	MySQL DBType = iota
	PostgreSQL
	SQLite
	SQLServer
)

func (db DBType) String() string {
	return [...]string{
		"mysql",
		"pgsql",
		"sqlite",
		"mssql",
	}[db]
}
