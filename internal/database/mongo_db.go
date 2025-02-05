package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDBOptions struct {
	DirectConnection bool
}

// MongoDBStrategy implements DBStrategy for MongoDB.
// It holds the connected client and a reference to a specific database.
type MongoDBStrategy struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func (m *MongoDBStrategy) Connect(connString string, opts DBOptions) (DBConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(connString).
		SetMaxPoolSize(opts.MaxPoolSize).
		SetAuth(options.Credential{
			Username: opts.Username,
			Password: opts.Password,
		})

	if opts.EnableLog {
		clientOptions.SetMonitor(getLogMonitor())
	}

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err.Error())
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err.Error())
	}

	// Store the client and database reference.
	m.Client = client
	m.DB = client.Database(opts.DBName)

	return m, nil
}

func (m *MongoDBStrategy) Disconnect(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

// ANSI color codes for terminal output.
const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[34m" // For started events.
	colorGreen  = "\033[32m" // For succeeded events.
	colorRed    = "\033[31m" // For failed events.
	colorYellow = "\033[33m" // For highlighting text.
)

func getLogMonitor() *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			// Format: [MONGO][STARTED] Command: <command_name> | DB: <database> |
			// RequestID: <id> | Cmd: <command document>
			log.Printf("%s[MONGO][STARTED]%s Command: %s%s%s | DB: %s%s%s | RequestID: %s%d%s | Cmd: %v",
				colorBlue, colorReset,
				colorYellow, evt.CommandName, colorReset,
				colorYellow, evt.DatabaseName, colorReset,
				colorYellow, evt.RequestID, colorReset,
				evt.Command)
		},
		Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
			// Format:
			// [MONGO][SUCCEEDED] Command: <command_name> | RequestID: <id> | Duration: <duration>
			log.Printf("%s[MONGO][SUCCEEDED]%s Command: %s%s%s | RequestID: %s%d%s | Duration: %s%v%s",
				colorGreen, colorReset,
				colorYellow, evt.CommandName, colorReset,
				colorYellow, evt.RequestID, colorReset,
				colorYellow, evt.Duration, colorReset)
		},
		Failed: func(_ context.Context, evt *event.CommandFailedEvent) {
			// Format: [MONGO][FAILED] Command: <command_name> | RequestID: <id> | Duration:
			// <duration> | Error: <error>
			log.Printf(
				"%s[MONGO][FAILED]%s Command: %s%s%s | "+
					"RequestID: %s%d%s | Duration: %s%v%s | Error: %s%v%s",
				colorRed, colorReset,
				colorYellow, evt.CommandName, colorReset,
				colorYellow, evt.RequestID, colorReset,
				colorYellow, evt.Duration, colorReset,
				colorYellow, evt.Failure, colorReset)
		},
	}
}
