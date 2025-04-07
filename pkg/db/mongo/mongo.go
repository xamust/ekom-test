package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/xamust/ekom-test/config"
	"github.com/xamust/ekom-test/internal/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	_defaultHost           = "localhost"
	_defaultPort           = "27017"
	_defaultMaxPoolSize    = 10
	_defaultConnectTimeout = 5 * time.Second
)

type Mongo struct {
	URL            string
	Port           string
	MaxPoolSize    uint64
	ConnectTimeout time.Duration

	log    interfaces.Logger
	Client *mongo.Client
}

func NewMongoDBConnectionWithCtx(
	ctx context.Context,
	log interfaces.Logger,
	cfg *config.Config,
	opts ...Option,
) (*Mongo, error) {
	m := &Mongo{
		URL:            _defaultHost,
		Port:           _defaultPort,
		MaxPoolSize:    _defaultMaxPoolSize,
		ConnectTimeout: _defaultConnectTimeout,
		log:            log,
	}
	for _, option := range opts {
		option(m)
	}

	if err := m.newClient(ctx, cfg); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Mongo) dsn() string {
	return "mongodb://" + m.URL + ":" + m.Port
}

func (m *Mongo) newClient(ctx context.Context, cfg *config.Config) error {
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(m.dsn()).
		SetMaxPoolSize(m.MaxPoolSize).
		SetConnectTimeout(m.ConnectTimeout).
		SetAuth(options.Credential{
			Username: cfg.DB.Mongo.Username,
			Password: cfg.DB.Mongo.Password,
		}))
	if err != nil {
		return fmt.Errorf("ошибка подключения к Mongo: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("ошибка подключения к Mongo: %w", err)
	}
	m.Client = client
	m.log.Debug("подключение к mongoDB установлено")
	return nil
}
