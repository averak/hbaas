// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dao

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserProfile is an object representing the database table.
type UserProfile struct {
	UserID    string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Content   []byte    `boil:"content" json:"content" toml:"content" yaml:"content"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userProfileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userProfileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserProfileColumns = struct {
	UserID    string
	Content   string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserProfileTableColumns = struct {
	UserID    string
	Content   string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_profiles.user_id",
	Content:   "user_profiles.content",
	CreatedAt: "user_profiles.created_at",
	UpdatedAt: "user_profiles.updated_at",
}

// Generated where

var UserProfileWhere = struct {
	UserID    whereHelperstring
	Content   whereHelper__byte
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	UserID:    whereHelperstring{field: "\"user_profiles\".\"user_id\""},
	Content:   whereHelper__byte{field: "\"user_profiles\".\"content\""},
	CreatedAt: whereHelpertime_Time{field: "\"user_profiles\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"user_profiles\".\"updated_at\""},
}

// UserProfileRels is where relationship names are stored.
var UserProfileRels = struct {
	User string
}{
	User: "User",
}

// userProfileR is where relationships are stored.
type userProfileR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userProfileR) NewStruct() *userProfileR {
	return &userProfileR{}
}

func (r *userProfileR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userProfileL is where Load methods for each relationship are stored.
type userProfileL struct{}

var (
	userProfileAllColumns            = []string{"user_id", "content", "created_at", "updated_at"}
	userProfileColumnsWithoutDefault = []string{"user_id", "content", "created_at", "updated_at"}
	userProfileColumnsWithDefault    = []string{}
	userProfilePrimaryKeyColumns     = []string{"user_id"}
	userProfileGeneratedColumns      = []string{}
)

type (
	// UserProfileSlice is an alias for a slice of pointers to UserProfile.
	// This should almost always be used instead of []UserProfile.
	UserProfileSlice []*UserProfile
	// UserProfileHook is the signature for custom UserProfile hook methods
	UserProfileHook func(context.Context, boil.ContextExecutor, *UserProfile) error

	userProfileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userProfileType                 = reflect.TypeOf(&UserProfile{})
	userProfileMapping              = queries.MakeStructMapping(userProfileType)
	userProfilePrimaryKeyMapping, _ = queries.BindMapping(userProfileType, userProfileMapping, userProfilePrimaryKeyColumns)
	userProfileInsertCacheMut       sync.RWMutex
	userProfileInsertCache          = make(map[string]insertCache)
	userProfileUpdateCacheMut       sync.RWMutex
	userProfileUpdateCache          = make(map[string]updateCache)
	userProfileUpsertCacheMut       sync.RWMutex
	userProfileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userProfileAfterSelectMu sync.Mutex
var userProfileAfterSelectHooks []UserProfileHook

var userProfileBeforeInsertMu sync.Mutex
var userProfileBeforeInsertHooks []UserProfileHook
var userProfileAfterInsertMu sync.Mutex
var userProfileAfterInsertHooks []UserProfileHook

var userProfileBeforeUpdateMu sync.Mutex
var userProfileBeforeUpdateHooks []UserProfileHook
var userProfileAfterUpdateMu sync.Mutex
var userProfileAfterUpdateHooks []UserProfileHook

var userProfileBeforeDeleteMu sync.Mutex
var userProfileBeforeDeleteHooks []UserProfileHook
var userProfileAfterDeleteMu sync.Mutex
var userProfileAfterDeleteHooks []UserProfileHook

var userProfileBeforeUpsertMu sync.Mutex
var userProfileBeforeUpsertHooks []UserProfileHook
var userProfileAfterUpsertMu sync.Mutex
var userProfileAfterUpsertHooks []UserProfileHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserProfile) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserProfile) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserProfile) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserProfile) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserProfile) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserProfile) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserProfile) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserProfile) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserProfile) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userProfileAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserProfileHook registers your hook function for all future operations.
func AddUserProfileHook(hookPoint boil.HookPoint, userProfileHook UserProfileHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userProfileAfterSelectMu.Lock()
		userProfileAfterSelectHooks = append(userProfileAfterSelectHooks, userProfileHook)
		userProfileAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		userProfileBeforeInsertMu.Lock()
		userProfileBeforeInsertHooks = append(userProfileBeforeInsertHooks, userProfileHook)
		userProfileBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		userProfileAfterInsertMu.Lock()
		userProfileAfterInsertHooks = append(userProfileAfterInsertHooks, userProfileHook)
		userProfileAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		userProfileBeforeUpdateMu.Lock()
		userProfileBeforeUpdateHooks = append(userProfileBeforeUpdateHooks, userProfileHook)
		userProfileBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		userProfileAfterUpdateMu.Lock()
		userProfileAfterUpdateHooks = append(userProfileAfterUpdateHooks, userProfileHook)
		userProfileAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		userProfileBeforeDeleteMu.Lock()
		userProfileBeforeDeleteHooks = append(userProfileBeforeDeleteHooks, userProfileHook)
		userProfileBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		userProfileAfterDeleteMu.Lock()
		userProfileAfterDeleteHooks = append(userProfileAfterDeleteHooks, userProfileHook)
		userProfileAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		userProfileBeforeUpsertMu.Lock()
		userProfileBeforeUpsertHooks = append(userProfileBeforeUpsertHooks, userProfileHook)
		userProfileBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		userProfileAfterUpsertMu.Lock()
		userProfileAfterUpsertHooks = append(userProfileAfterUpsertHooks, userProfileHook)
		userProfileAfterUpsertMu.Unlock()
	}
}

// One returns a single userProfile record from the query.
func (q userProfileQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserProfile, error) {
	o := &UserProfile{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dao: failed to execute a one query for user_profiles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserProfile records from the query.
func (q userProfileQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserProfileSlice, error) {
	var o []*UserProfile

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dao: failed to assign all query results to UserProfile slice")
	}

	if len(userProfileAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserProfile records in the query.
func (q userProfileQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to count user_profiles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userProfileQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dao: failed to check if user_profiles exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserProfile) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userProfileL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserProfile interface{}, mods queries.Applicator) error {
	var slice []*UserProfile
	var object *UserProfile

	if singular {
		var ok bool
		object, ok = maybeUserProfile.(*UserProfile)
		if !ok {
			object = new(UserProfile)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserProfile)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserProfile))
			}
		}
	} else {
		s, ok := maybeUserProfile.(*[]*UserProfile)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserProfile)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserProfile))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &userProfileR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userProfileR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserProfile = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserProfile = local
				break
			}
		}
	}

	return nil
}

// SetUser of the userProfile to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserProfile.
func (o *UserProfile) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_profiles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userProfilePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userProfileR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserProfile: o,
		}
	} else {
		related.R.UserProfile = o
	}

	return nil
}

// UserProfiles retrieves all the records using an executor.
func UserProfiles(mods ...qm.QueryMod) userProfileQuery {
	mods = append(mods, qm.From("\"user_profiles\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_profiles\".*"})
	}

	return userProfileQuery{q}
}

// FindUserProfile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserProfile(ctx context.Context, exec boil.ContextExecutor, userID string, selectCols ...string) (*UserProfile, error) {
	userProfileObj := &UserProfile{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_profiles\" where \"user_id\"=$1", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(ctx, exec, userProfileObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dao: unable to select from user_profiles")
	}

	if err = userProfileObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userProfileObj, err
	}

	return userProfileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserProfile) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dao: no user_profiles provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userProfileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userProfileInsertCacheMut.RLock()
	cache, cached := userProfileInsertCache[key]
	userProfileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userProfileAllColumns,
			userProfileColumnsWithDefault,
			userProfileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userProfileType, userProfileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userProfileType, userProfileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_profiles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_profiles\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dao: unable to insert into user_profiles")
	}

	if !cached {
		userProfileInsertCacheMut.Lock()
		userProfileInsertCache[key] = cache
		userProfileInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserProfile.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserProfile) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userProfileUpdateCacheMut.RLock()
	cache, cached := userProfileUpdateCache[key]
	userProfileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userProfileAllColumns,
			userProfilePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dao: unable to update user_profiles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_profiles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userProfilePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userProfileType, userProfileMapping, append(wl, userProfilePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to update user_profiles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by update for user_profiles")
	}

	if !cached {
		userProfileUpdateCacheMut.Lock()
		userProfileUpdateCache[key] = cache
		userProfileUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userProfileQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to update all for user_profiles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to retrieve rows affected for user_profiles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserProfileSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dao: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_profiles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userProfilePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to update all in userProfile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to retrieve rows affected all in update all userProfile")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserProfile) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("dao: no user_profiles provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userProfileColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	userProfileUpsertCacheMut.RLock()
	cache, cached := userProfileUpsertCache[key]
	userProfileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			userProfileAllColumns,
			userProfileColumnsWithDefault,
			userProfileColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userProfileAllColumns,
			userProfilePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dao: unable to upsert user_profiles, could not build update column list")
		}

		ret := strmangle.SetComplement(userProfileAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(userProfilePrimaryKeyColumns) == 0 {
				return errors.New("dao: unable to upsert user_profiles, could not build conflict column list")
			}

			conflict = make([]string, len(userProfilePrimaryKeyColumns))
			copy(conflict, userProfilePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_profiles\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(userProfileType, userProfileMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userProfileType, userProfileMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dao: unable to upsert user_profiles")
	}

	if !cached {
		userProfileUpsertCacheMut.Lock()
		userProfileUpsertCache[key] = cache
		userProfileUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserProfile record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserProfile) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dao: no UserProfile provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userProfilePrimaryKeyMapping)
	sql := "DELETE FROM \"user_profiles\" WHERE \"user_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete from user_profiles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by delete for user_profiles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userProfileQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dao: no userProfileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete all from user_profiles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by deleteall for user_profiles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserProfileSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userProfileBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_profiles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userProfilePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete all from userProfile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by deleteall for user_profiles")
	}

	if len(userProfileAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserProfile) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserProfile(ctx, exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserProfileSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserProfileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_profiles\".* FROM \"user_profiles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userProfilePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dao: unable to reload all in UserProfileSlice")
	}

	*o = slice

	return nil
}

// UserProfileExists checks if the UserProfile row exists.
func UserProfileExists(ctx context.Context, exec boil.ContextExecutor, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_profiles\" where \"user_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID)
	}
	row := exec.QueryRowContext(ctx, sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dao: unable to check if user_profiles exists")
	}

	return exists, nil
}

// Exists checks if the UserProfile row exists.
func (o *UserProfile) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserProfileExists(ctx, exec, o.UserID)
}

// /////////////////////////////// BEGIN EXTENSIONS /////////////////////////////////
// Expose table columns
var (
	UserProfileAllColumns            = userProfileAllColumns
	UserProfileColumnsWithoutDefault = userProfileColumnsWithoutDefault
	UserProfileColumnsWithDefault    = userProfileColumnsWithDefault
	UserProfilePrimaryKeyColumns     = userProfilePrimaryKeyColumns
	UserProfileGeneratedColumns      = userProfileGeneratedColumns
)

// InsertAll inserts all rows with the specified column values, using an executor.
func (o UserProfileSlice) InsertAll(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var sql string
	vals := []interface{}{}
	for i, row := range o {
		if !boil.TimestampsAreSkipped(ctx) {
			currTime := time.Now().In(boil.GetLocation())
			if row.CreatedAt.IsZero() {
				row.CreatedAt = currTime
			}
			if row.UpdatedAt.IsZero() {
				row.UpdatedAt = currTime
			}
		}

		if err := row.doBeforeInsertHooks(ctx, exec); err != nil {
			return 0, err
		}

		wl, _ := columns.InsertColumnSet(
			userProfileAllColumns,
			userProfileColumnsWithDefault,
			userProfileColumnsWithoutDefault,
			queries.NonZeroDefaultSet(userProfileColumnsWithDefault, row),
		)
		if i == 0 {
			sql = "INSERT INTO \"user_profiles\" " + "(\"" + strings.Join(wl, "\",\"") + "\")" + " VALUES "
		}
		sql += strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), len(vals)+1, len(wl))
		if i != len(o)-1 {
			sql += ","
		}
		valMapping, err := queries.BindMapping(userProfileType, userProfileMapping, wl)
		if err != nil {
			return 0, err
		}

		value := reflect.Indirect(reflect.ValueOf(row))
		vals = append(vals, queries.ValuesFromMapping(value, valMapping)...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, vals)
	}

	result, err := exec.ExecContext(ctx, sql, vals...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to insert all from userProfile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by insertall for user_profiles")
	}

	if len(userProfileAfterInsertHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterInsertHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// UpsertAll inserts or updates all rows
// Currently it doesn't support "NoContext" and "NoRowsAffected"
func (o UserProfileSlice) UpsertAll(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	nzDefaults := queries.NonZeroDefaultSet(userProfileColumnsWithDefault, o[0])

	insert, _ := insertColumns.InsertColumnSet(
		userProfileAllColumns,
		userProfileColumnsWithDefault,
		userProfileColumnsWithoutDefault,
		nzDefaults,
	)
	update := updateColumns.UpdateColumnSet(
		userProfileAllColumns,
		userProfilePrimaryKeyColumns,
	)

	if updateOnConflict && len(update) == 0 {
		return 0, errors.New("dao: unable to upsert user_profiles, could not build update column list")
	}

	conflict := conflictColumns
	if len(conflict) == 0 {
		conflict = make([]string, len(userProfilePrimaryKeyColumns))
		copy(conflict, userProfilePrimaryKeyColumns)
	}

	buf := strmangle.GetBuffer()
	defer strmangle.PutBuffer(buf)

	columns := "DEFAULT VALUES"
	if len(insert) != 0 {
		columns = fmt.Sprintf("(%s) VALUES %s",
			strings.Join(insert, ", "),
			strmangle.Placeholders(dialect.UseIndexPlaceholders, len(insert)*len(o), 1, len(insert)),
		)
	}

	fmt.Fprintf(
		buf,
		"INSERT INTO %s %s ON CONFLICT ",
		"\"user_profiles\"",
		columns,
	)

	if !updateOnConflict || len(update) == 0 {
		buf.WriteString("DO NOTHING")
	} else {
		buf.WriteByte('(')
		buf.WriteString(strings.Join(conflict, ", "))
		buf.WriteString(") DO UPDATE SET ")

		for i, v := range update {
			if i != 0 {
				buf.WriteByte(',')
			}
			quoted := strmangle.IdentQuote(dialect.LQ, dialect.RQ, v)
			buf.WriteString(quoted)
			buf.WriteString(" = EXCLUDED.")
			buf.WriteString(quoted)
		}
	}

	query := buf.String()
	valueMapping, err := queries.BindMapping(userProfileType, userProfileMapping, insert)
	if err != nil {
		return 0, err
	}

	var vals []interface{}
	for _, row := range o {
		if !boil.TimestampsAreSkipped(ctx) {
			currTime := time.Now().In(boil.GetLocation())
			if row.CreatedAt.IsZero() {
				row.CreatedAt = currTime
			}

			row.UpdatedAt = currTime
		}

		if err := row.doBeforeUpsertHooks(ctx, exec); err != nil {
			return 0, err
		}

		value := reflect.Indirect(reflect.ValueOf(row))
		vals = append(vals, queries.ValuesFromMapping(value, valueMapping)...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, vals)
	}

	result, err := exec.ExecContext(ctx, query, vals...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to upsert for user_profiles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by upsert for user_profiles")
	}

	if len(userProfileAfterUpsertHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterUpsertHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

///////////////////////////////// END EXTENSIONS /////////////////////////////////
