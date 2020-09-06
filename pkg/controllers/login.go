package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/db/models"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/auth"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/external"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/security"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {

	s := external.NewSlackNotifier()

	req := new(loginReq)
	if err := c.Bind(req); err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}

	err := validation.ValidateStruct(req, validation.Field(&req.Email, validation.Required, validation.Length(5, 50)))

	if err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}
	if !utils.ValidateEmailInput(req.Email) {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}

	p, err := models.Partners(
		qm.Select(models.PartnerColumns.ID,
			models.PartnerColumns.Email,
			models.PartnerColumns.EncryptedPassword,
			models.PartnerColumns.FirmID,
			models.PartnerColumns.IsConfirmed,
			models.PartnerColumns.IsActive,
			models.PartnerColumns.IsBlocked),
		models.PartnerWhere.Email.EQ(req.Email),
	).One(external.NewWriterDB("canberktest"))

	if err == sql.ErrNoRows {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusUnauthorized)
	} else if err != nil {
		s.Notify(fmt.Sprintf("Something went wrong!\n`Login`\n%s\n", err.Error()))
		return err
	}

	if err := security.VerifyPassword(p.EncryptedPassword, req.Password); err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusUnauthorized)
	}

	td, err := auth.CreateToken(p)
	if err != nil {
		s.Notify(fmt.Sprintf("Something went wrong!\n`CreateToken`\n%s\n", err.Error()))
		return err
	}

	if err := auth.AddJWTToRedis(p.ID, td); err != nil {
		s.Notify(fmt.Sprintf("Something went wrong!\n`AddAccessTokenToRedis`\n%s\n", err.Error()))
		return err
	}

	s.Notify(fmt.Sprint(td))

	tm := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}


	return c.JSON(http.StatusOK, tm)
}
