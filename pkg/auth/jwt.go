package auth

import (
	"fmt"
	"time"

	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/db/models"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type tokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	ATExpires    int64
	RTExpires    int64
}

type jwtCustomClaimsAT struct {
	ID         uint   `json:"id"`
	FirmID     int    `json:"firm_id"`
	AccessUUID string `json:"access_uuid"`
	jwt.StandardClaims
}

type jwtCustomClaimsRT struct {
	ID          uint   `json:"id"`
	RefreshUUID string `json:"refresh_uuid"`
	jwt.StandardClaims
}

func CreateToken(p *models.Partner) (*tokenDetails, error) {

	td := new(tokenDetails)
	var err error

	td.ATExpires = time.Now().Add(time.Minute * 15).Unix() // Access Token expires after 15 min
	td.RTExpires = time.Now().Add(time.Hour * 48).Unix()   // Refresh Token expires after 2 days

	td.AccessUUID = uuid.New().String()
	td.RefreshUUID = fmt.Sprintf("%s-----%d", td.AccessUUID, p.ID)

	atClaims := &jwtCustomClaimsAT{
		ID:         p.ID,
		FirmID:     p.FirmID.Int,
		AccessUUID: td.AccessUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.ATExpires,
		},
	}

	if !p.FirmID.Valid {
		atClaims.FirmID = -9
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.Config.JWTAccessSecret))
	if err != nil {
		return nil, err
	}

	rtClaims := &jwtCustomClaimsRT{
		ID:          p.ID,
		RefreshUUID: td.RefreshUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.ATExpires,
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(config.Config.JWTAccessSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

// var Authenticate = middleware.JWTWithConfig(middleware.JWTConfig{
// 	Claims:     new(jwtCustomClaims),
// 	SigningKey: []byte(config.Config.JWTAccessSecret),
// })
