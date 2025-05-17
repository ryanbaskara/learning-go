package connection

// Mysql interface of mysql database
type Mysql interface {
	NamedCreate(query string, instance interface{}) (int64, error)
	Create(query string, params ...interface{}) (int64, error)
	NamedExec(query string, instance interface{}) (int64, error)
	Exec(query string, params ...interface{}) (int64, error)
	Select(result interface{}, query string, params ...interface{}) error
	Find(result interface{}, query string, params ...interface{}) error
	Rebind(query string) string
	Close() error
}
