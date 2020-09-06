package auth

import (
	"strconv"
	"time"

	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/external"
)

func AddJWTToRedis(userID uint, td *tokenDetails) error {

	at := time.Unix(td.ATExpires, 0)
	rt := time.Unix(td.RTExpires, 0)
	now := time.Now()

	var err error
	rClient := external.NewSecurityRedisClient()

	err = rClient.Set(td.AccessUUID, strconv.Itoa(int(userID)), at.Sub(now)).Err()
	if err != nil {
		return err
	}

	err = rClient.Set(td.RefreshUUID, strconv.Itoa(int(userID)), rt.Sub(now)).Err()
	if err != nil {
		return err
	}

	return nil
}
