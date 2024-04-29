package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/qrave1/BlazingMoneyControlBot/internal/domain"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const defaultTTL = 10 * time.Minute

var ErrNotFound = errors.New("not found")

type WalletRedisRepository struct {
	rdb *redis.Client
}

func (wr *WalletRedisRepository) Create(ctx context.Context, wallet domain.Wallet) error {
	dbo, err := json.Marshal(wallet)
	if err != nil {
		return err
	}

	return wr.rdb.Set(ctx, strconv.Itoa(wallet.Id), string(dbo), defaultTTL).Err()
}

func (wr *WalletRedisRepository) Read(ctx context.Context, id int) (domain.Wallet, error) {
	result, err := wr.rdb.Get(ctx, strconv.Itoa(id)).Result()
	if errors.Is(err, redis.Nil) {
		return domain.Wallet{}, ErrNotFound
	} else if err != nil {
		return domain.Wallet{}, err
	} else {
		var wallet domain.Wallet
		if err = json.Unmarshal([]byte(result), &wallet); err != nil {
			return domain.Wallet{}, err
		}

		return wallet, nil
	}
}

func (wr *WalletRedisRepository) UpdateBalance(ctx context.Context, id, value int) error {
	//TODO implement me
	panic("implement me")
}

func (wr *WalletRedisRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
