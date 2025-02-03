package database

import "context"

type DBOptions struct {
	Username    string
	Password    string
	DBName      string
	MaxPoolSize uint64
	EnableLog   bool
}

// DBConnection is a generic type to hold any kind of DB connection/client.
type DBConnection interface{}

// DBStrategy defines an interface for connecting to a database.
type DBStrategy interface {
	Connect(connString string, opts DBOptions) (DBConnection, error)
	Disconnect(ctx context.Context) error
}

// DBContext holds a DBStrategy. It serves as the "context" in the strategy pattern.
type DBContext struct {
	strategy DBStrategy
}

// NewDBContext creates a new DBContext with the specified strategy.
func NewDBContext(strategy DBStrategy) *DBContext {
	return &DBContext{strategy: strategy}
}

// Connect uses the selected strategy to establish a connection.
func (c *DBContext) Connect(connString string, opts DBOptions) (DBConnection, error) {
	return c.strategy.Connect(connString, opts)
}

// Disconnect uses the selected strategy to close the connection.
func (c *DBContext) Disconnect(ctx context.Context) error {
	return c.strategy.Disconnect(ctx)
}
