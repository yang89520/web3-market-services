package database

import (
	"context"
	"fmt"
	"github.com/yang89520/web3-merket-services/config"
	"gorm.io/gorm"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	//"github.com/yang89520/web3-merket-services/common/retry"
	//"github.com/yang89520/web3-merket-services/config"
)

type DB struct {
	gorm             *gorm.DB
	MarketPrice      MarketPriceDB
	OfficialCoinRate OfficialCoinRateDB
}

func NewDB(ctx context.Context, dbConfig config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Name)
	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf(" port=%s", dbConfig.Port)
	}
	if dbConfig.User != "" {
		dsn += fmt.Sprintf(" user=%s ", dbConfig.User)
	}
	if dbConfig.Password != "" {
		dsn += fmt.Sprintf(" password=%s ", dbConfig.Password)
	}

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}

	&retry.E
}
