package repository

import "fmt"

type Config struct {
	User           string
	Password       string
	Host           string
	Port           int
	DatabaseName   string
	MaxConnections int
	SSL            bool
}

func (c Config) connectionString() string {
	// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
	confStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?pool_max_conns=%d&sslmode=",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DatabaseName,
		c.MaxConnections,
	)

	// if c.SSL {
	// 	// TODO: implement SSL connection
	// 	// confStr += ""
	// }

	return confStr
}
