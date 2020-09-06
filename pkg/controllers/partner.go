package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/db/models"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/external"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/security"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var s = external.NewSlackNotifier()
var db = external.NewWriterDB("canberktest")

func CreatePartner(c echo.Context) error {

	req := new(models.Partner)
	if err := c.Bind(req); err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}

	var err error
	var exists bool

	err = validation.ValidateStruct(req,
		validation.Field(&req.EncryptedPassword, validation.Required, validation.Length(8, 32)),
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Surname, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.PhoneNumber, validation.Required, is.E164),
		validation.Field(&req.Birthday, validation.Required),
	)
	if err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}

	err = utils.ValidatePasswordInput(req.EncryptedPassword)
	if err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}

	hashedPassword, err := security.HashPassword(req.EncryptedPassword)
	if err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}

	exists, err = IsPartnerEmailExists(req.Email, db)
	if err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}
	if exists {
		s.Notify(fmt.Sprintf("Someone tried to register with the same email.\n%s is already exists.", req.Email))
		return c.NoContent(http.StatusBadRequest)
	}

	exists, err = IsPartnerPhoneNumberExists(req.PhoneNumber, db)
	if err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}
	if exists {
		s.Notify(fmt.Sprintf("Someone tried to register with the same phone number.\n%s is already exists.", req.PhoneNumber))
		return c.NoContent(http.StatusBadRequest)
	}

	realIP := utils.GetIP(c)

	partner := &models.Partner{
		Email:             req.Email,
		PhoneNumber:       req.PhoneNumber,
		EncryptedPassword: string(hashedPassword),
		Name:              req.Name,
		Surname:           req.Surname,
		Birthday:          req.Birthday,
		LastIP:            null.NewString(realIP, true),
	}

	err = partner.Insert(db, boil.Infer())
	if err != nil {
		s.Notify(fmt.Sprint(err))
		return c.NoContent(http.StatusBadRequest)
	}

	s.Notify("New Partner Has Been Created.")
	s.Notify(utils.PrettyString(partner))

	return c.NoContent(http.StatusOK)
}

func IsPartnerEmailExists(email string, db *sql.DB) (bool, error) {
	return models.Partners(
		models.PartnerWhere.Email.EQ(email),
	).Exists(db)
}

func IsPartnerPhoneNumberExists(phoneNumber string, db *sql.DB) (bool, error) {
	return models.Partners(
		models.PartnerWhere.PhoneNumber.EQ(phoneNumber),
	).Exists(db)
}
