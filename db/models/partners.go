// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Partner is an object representing the database table.
type Partner struct {
	ID                  uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Email               string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	PhoneNumber         string      `boil:"phone_number" json:"phone_number" toml:"phone_number" yaml:"phone_number"`
	EncryptedPassword   string      `boil:"encrypted_password" json:"encrypted_password" toml:"encrypted_password" yaml:"encrypted_password"`
	DepartmentID        int         `boil:"department_id" json:"department_id" toml:"department_id" yaml:"department_id"`
	Title               string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Name                string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Surname             string      `boil:"surname" json:"surname" toml:"surname" yaml:"surname"`
	PPURL               null.String `boil:"pp_url" json:"pp_url,omitempty" toml:"pp_url" yaml:"pp_url,omitempty"`
	Birthday            time.Time   `boil:"birthday" json:"birthday" toml:"birthday" yaml:"birthday"`
	FirmID              null.Int    `boil:"firm_id" json:"firm_id,omitempty" toml:"firm_id" yaml:"firm_id,omitempty"`
	PhoneNumberS        null.String `boil:"phone_number_s" json:"phone_number_s,omitempty" toml:"phone_number_s" yaml:"phone_number_s,omitempty"`
	CreatedAt           time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt           time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt           null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	ReactivatedAt       null.Time   `boil:"reactivated_at" json:"reactivated_at,omitempty" toml:"reactivated_at" yaml:"reactivated_at,omitempty"`
	CurrentSignInAt     null.Time   `boil:"current_sign_in_at" json:"current_sign_in_at,omitempty" toml:"current_sign_in_at" yaml:"current_sign_in_at,omitempty"`
	IsActive            bool        `boil:"is_active" json:"is_active" toml:"is_active" yaml:"is_active"`
	IsConfirmed         bool        `boil:"is_confirmed" json:"is_confirmed" toml:"is_confirmed" yaml:"is_confirmed"`
	ConfirmedBy         int         `boil:"confirmed_by" json:"confirmed_by" toml:"confirmed_by" yaml:"confirmed_by"`
	LastIP              null.String `boil:"last_ip" json:"last_ip,omitempty" toml:"last_ip" yaml:"last_ip,omitempty"`
	SignInCount         int         `boil:"sign_in_count" json:"sign_in_count" toml:"sign_in_count" yaml:"sign_in_count"`
	ForcePasswordChange bool        `boil:"force_password_change" json:"force_password_change" toml:"force_password_change" yaml:"force_password_change"`
	IsBlocked           bool        `boil:"is_blocked" json:"is_blocked" toml:"is_blocked" yaml:"is_blocked"`
	IsVerified          bool        `boil:"is_verified" json:"is_verified" toml:"is_verified" yaml:"is_verified"`

	R *partnerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L partnerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PartnerColumns = struct {
	ID                  string
	Email               string
	PhoneNumber         string
	EncryptedPassword   string
	DepartmentID        string
	Title               string
	Name                string
	Surname             string
	PPURL               string
	Birthday            string
	FirmID              string
	PhoneNumberS        string
	CreatedAt           string
	UpdatedAt           string
	DeletedAt           string
	ReactivatedAt       string
	CurrentSignInAt     string
	IsActive            string
	IsConfirmed         string
	ConfirmedBy         string
	LastIP              string
	SignInCount         string
	ForcePasswordChange string
	IsBlocked           string
	IsVerified          string
}{
	ID:                  "id",
	Email:               "email",
	PhoneNumber:         "phone_number",
	EncryptedPassword:   "encrypted_password",
	DepartmentID:        "department_id",
	Title:               "title",
	Name:                "name",
	Surname:             "surname",
	PPURL:               "pp_url",
	Birthday:            "birthday",
	FirmID:              "firm_id",
	PhoneNumberS:        "phone_number_s",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
	ReactivatedAt:       "reactivated_at",
	CurrentSignInAt:     "current_sign_in_at",
	IsActive:            "is_active",
	IsConfirmed:         "is_confirmed",
	ConfirmedBy:         "confirmed_by",
	LastIP:              "last_ip",
	SignInCount:         "sign_in_count",
	ForcePasswordChange: "force_password_change",
	IsBlocked:           "is_blocked",
	IsVerified:          "is_verified",
}

// Generated where

type whereHelperuint struct{ field string }

func (w whereHelperuint) EQ(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint) NEQ(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint) LT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint) LTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint) GT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint) GTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint) IN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint) NIN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var PartnerWhere = struct {
	ID                  whereHelperuint
	Email               whereHelperstring
	PhoneNumber         whereHelperstring
	EncryptedPassword   whereHelperstring
	DepartmentID        whereHelperint
	Title               whereHelperstring
	Name                whereHelperstring
	Surname             whereHelperstring
	PPURL               whereHelpernull_String
	Birthday            whereHelpertime_Time
	FirmID              whereHelpernull_Int
	PhoneNumberS        whereHelpernull_String
	CreatedAt           whereHelpertime_Time
	UpdatedAt           whereHelpertime_Time
	DeletedAt           whereHelpernull_Time
	ReactivatedAt       whereHelpernull_Time
	CurrentSignInAt     whereHelpernull_Time
	IsActive            whereHelperbool
	IsConfirmed         whereHelperbool
	ConfirmedBy         whereHelperint
	LastIP              whereHelpernull_String
	SignInCount         whereHelperint
	ForcePasswordChange whereHelperbool
	IsBlocked           whereHelperbool
	IsVerified          whereHelperbool
}{
	ID:                  whereHelperuint{field: "`partners`.`id`"},
	Email:               whereHelperstring{field: "`partners`.`email`"},
	PhoneNumber:         whereHelperstring{field: "`partners`.`phone_number`"},
	EncryptedPassword:   whereHelperstring{field: "`partners`.`encrypted_password`"},
	DepartmentID:        whereHelperint{field: "`partners`.`department_id`"},
	Title:               whereHelperstring{field: "`partners`.`title`"},
	Name:                whereHelperstring{field: "`partners`.`name`"},
	Surname:             whereHelperstring{field: "`partners`.`surname`"},
	PPURL:               whereHelpernull_String{field: "`partners`.`pp_url`"},
	Birthday:            whereHelpertime_Time{field: "`partners`.`birthday`"},
	FirmID:              whereHelpernull_Int{field: "`partners`.`firm_id`"},
	PhoneNumberS:        whereHelpernull_String{field: "`partners`.`phone_number_s`"},
	CreatedAt:           whereHelpertime_Time{field: "`partners`.`created_at`"},
	UpdatedAt:           whereHelpertime_Time{field: "`partners`.`updated_at`"},
	DeletedAt:           whereHelpernull_Time{field: "`partners`.`deleted_at`"},
	ReactivatedAt:       whereHelpernull_Time{field: "`partners`.`reactivated_at`"},
	CurrentSignInAt:     whereHelpernull_Time{field: "`partners`.`current_sign_in_at`"},
	IsActive:            whereHelperbool{field: "`partners`.`is_active`"},
	IsConfirmed:         whereHelperbool{field: "`partners`.`is_confirmed`"},
	ConfirmedBy:         whereHelperint{field: "`partners`.`confirmed_by`"},
	LastIP:              whereHelpernull_String{field: "`partners`.`last_ip`"},
	SignInCount:         whereHelperint{field: "`partners`.`sign_in_count`"},
	ForcePasswordChange: whereHelperbool{field: "`partners`.`force_password_change`"},
	IsBlocked:           whereHelperbool{field: "`partners`.`is_blocked`"},
	IsVerified:          whereHelperbool{field: "`partners`.`is_verified`"},
}

// PartnerRels is where relationship names are stored.
var PartnerRels = struct {
}{}

// partnerR is where relationships are stored.
type partnerR struct {
}

// NewStruct creates a new relationship struct
func (*partnerR) NewStruct() *partnerR {
	return &partnerR{}
}

// partnerL is where Load methods for each relationship are stored.
type partnerL struct{}

var (
	partnerAllColumns            = []string{"id", "email", "phone_number", "encrypted_password", "department_id", "title", "name", "surname", "pp_url", "birthday", "firm_id", "phone_number_s", "created_at", "updated_at", "deleted_at", "reactivated_at", "current_sign_in_at", "is_active", "is_confirmed", "confirmed_by", "last_ip", "sign_in_count", "force_password_change", "is_blocked", "is_verified"}
	partnerColumnsWithoutDefault = []string{"email", "phone_number", "encrypted_password", "department_id", "title", "name", "surname", "pp_url", "birthday", "firm_id", "phone_number_s", "deleted_at", "reactivated_at", "current_sign_in_at", "last_ip"}
	partnerColumnsWithDefault    = []string{"id", "created_at", "updated_at", "is_active", "is_confirmed", "confirmed_by", "sign_in_count", "force_password_change", "is_blocked", "is_verified"}
	partnerPrimaryKeyColumns     = []string{"id"}
)

type (
	// PartnerSlice is an alias for a slice of pointers to Partner.
	// This should generally be used opposed to []Partner.
	PartnerSlice []*Partner

	partnerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	partnerType                 = reflect.TypeOf(&Partner{})
	partnerMapping              = queries.MakeStructMapping(partnerType)
	partnerPrimaryKeyMapping, _ = queries.BindMapping(partnerType, partnerMapping, partnerPrimaryKeyColumns)
	partnerInsertCacheMut       sync.RWMutex
	partnerInsertCache          = make(map[string]insertCache)
	partnerUpdateCacheMut       sync.RWMutex
	partnerUpdateCache          = make(map[string]updateCache)
	partnerUpsertCacheMut       sync.RWMutex
	partnerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single partner record from the query.
func (q partnerQuery) One(exec boil.Executor) (*Partner, error) {
	o := &Partner{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for partners")
	}

	return o, nil
}

// All returns all Partner records from the query.
func (q partnerQuery) All(exec boil.Executor) (PartnerSlice, error) {
	var o []*Partner

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Partner slice")
	}

	return o, nil
}

// Count returns the count of all Partner records in the query.
func (q partnerQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count partners rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q partnerQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if partners exists")
	}

	return count > 0, nil
}

// Partners retrieves all the records using an executor.
func Partners(mods ...qm.QueryMod) partnerQuery {
	mods = append(mods, qm.From("`partners`"))
	return partnerQuery{NewQuery(mods...)}
}

// FindPartner retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPartner(exec boil.Executor, iD uint, selectCols ...string) (*Partner, error) {
	partnerObj := &Partner{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `partners` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, partnerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from partners")
	}

	return partnerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Partner) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no partners provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(partnerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	partnerInsertCacheMut.RLock()
	cache, cached := partnerInsertCache[key]
	partnerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			partnerAllColumns,
			partnerColumnsWithDefault,
			partnerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(partnerType, partnerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(partnerType, partnerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `partners` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `partners` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `partners` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, partnerPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into partners")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == partnerMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for partners")
	}

CacheNoHooks:
	if !cached {
		partnerInsertCacheMut.Lock()
		partnerInsertCache[key] = cache
		partnerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Partner.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Partner) Update(exec boil.Executor, columns boil.Columns) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	partnerUpdateCacheMut.RLock()
	cache, cached := partnerUpdateCache[key]
	partnerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			partnerAllColumns,
			partnerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update partners, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `partners` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, partnerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(partnerType, partnerMapping, append(wl, partnerPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update partners row")
	}

	if !cached {
		partnerUpdateCacheMut.Lock()
		partnerUpdateCache[key] = cache
		partnerUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q partnerQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for partners")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PartnerSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partnerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `partners` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, partnerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in partner slice")
	}

	return nil
}

var mySQLPartnerUniqueColumns = []string{
	"id",
	"email",
	"phone_number",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Partner) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no partners provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(partnerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPartnerUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	partnerUpsertCacheMut.RLock()
	cache, cached := partnerUpsertCache[key]
	partnerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			partnerAllColumns,
			partnerColumnsWithDefault,
			partnerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			partnerAllColumns,
			partnerPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert partners, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "partners", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `partners` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(partnerType, partnerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(partnerType, partnerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for partners")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == partnerMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(partnerType, partnerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for partners")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for partners")
	}

CacheNoHooks:
	if !cached {
		partnerUpsertCacheMut.Lock()
		partnerUpsertCache[key] = cache
		partnerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Partner record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Partner) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Partner provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), partnerPrimaryKeyMapping)
	sql := "DELETE FROM `partners` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from partners")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q partnerQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("models: no partnerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from partners")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PartnerSlice) DeleteAll(exec boil.Executor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partnerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `partners` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, partnerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from partner slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Partner) Reload(exec boil.Executor) error {
	ret, err := FindPartner(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PartnerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PartnerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partnerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `partners`.* FROM `partners` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, partnerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PartnerSlice")
	}

	*o = slice

	return nil
}

// PartnerExists checks if the Partner row exists.
func PartnerExists(exec boil.Executor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `partners` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if partners exists")
	}

	return exists, nil
}
