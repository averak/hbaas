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

// Echo is an object representing the database table.
type Echo struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Message   string    `boil:"message" json:"message" toml:"message" yaml:"message"`
	Timestamp time.Time `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *echoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L echoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EchoColumns = struct {
	ID        string
	Message   string
	Timestamp string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Message:   "message",
	Timestamp: "timestamp",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var EchoTableColumns = struct {
	ID        string
	Message   string
	Timestamp string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "echos.id",
	Message:   "echos.message",
	Timestamp: "echos.timestamp",
	CreatedAt: "echos.created_at",
	UpdatedAt: "echos.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
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

var EchoWhere = struct {
	ID        whereHelperstring
	Message   whereHelperstring
	Timestamp whereHelpertime_Time
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "\"echos\".\"id\""},
	Message:   whereHelperstring{field: "\"echos\".\"message\""},
	Timestamp: whereHelpertime_Time{field: "\"echos\".\"timestamp\""},
	CreatedAt: whereHelpertime_Time{field: "\"echos\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"echos\".\"updated_at\""},
}

// EchoRels is where relationship names are stored.
var EchoRels = struct {
}{}

// echoR is where relationships are stored.
type echoR struct {
}

// NewStruct creates a new relationship struct
func (*echoR) NewStruct() *echoR {
	return &echoR{}
}

// echoL is where Load methods for each relationship are stored.
type echoL struct{}

var (
	echoAllColumns            = []string{"id", "message", "timestamp", "created_at", "updated_at"}
	echoColumnsWithoutDefault = []string{"id", "message", "timestamp", "created_at", "updated_at"}
	echoColumnsWithDefault    = []string{}
	echoPrimaryKeyColumns     = []string{"id"}
	echoGeneratedColumns      = []string{}
)

type (
	// EchoSlice is an alias for a slice of pointers to Echo.
	// This should almost always be used instead of []Echo.
	EchoSlice []*Echo
	// EchoHook is the signature for custom Echo hook methods
	EchoHook func(context.Context, boil.ContextExecutor, *Echo) error

	echoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	echoType                 = reflect.TypeOf(&Echo{})
	echoMapping              = queries.MakeStructMapping(echoType)
	echoPrimaryKeyMapping, _ = queries.BindMapping(echoType, echoMapping, echoPrimaryKeyColumns)
	echoInsertCacheMut       sync.RWMutex
	echoInsertCache          = make(map[string]insertCache)
	echoUpdateCacheMut       sync.RWMutex
	echoUpdateCache          = make(map[string]updateCache)
	echoUpsertCacheMut       sync.RWMutex
	echoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var echoAfterSelectMu sync.Mutex
var echoAfterSelectHooks []EchoHook

var echoBeforeInsertMu sync.Mutex
var echoBeforeInsertHooks []EchoHook
var echoAfterInsertMu sync.Mutex
var echoAfterInsertHooks []EchoHook

var echoBeforeUpdateMu sync.Mutex
var echoBeforeUpdateHooks []EchoHook
var echoAfterUpdateMu sync.Mutex
var echoAfterUpdateHooks []EchoHook

var echoBeforeDeleteMu sync.Mutex
var echoBeforeDeleteHooks []EchoHook
var echoAfterDeleteMu sync.Mutex
var echoAfterDeleteHooks []EchoHook

var echoBeforeUpsertMu sync.Mutex
var echoBeforeUpsertHooks []EchoHook
var echoAfterUpsertMu sync.Mutex
var echoAfterUpsertHooks []EchoHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Echo) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Echo) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Echo) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Echo) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Echo) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Echo) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Echo) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Echo) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Echo) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range echoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEchoHook registers your hook function for all future operations.
func AddEchoHook(hookPoint boil.HookPoint, echoHook EchoHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		echoAfterSelectMu.Lock()
		echoAfterSelectHooks = append(echoAfterSelectHooks, echoHook)
		echoAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		echoBeforeInsertMu.Lock()
		echoBeforeInsertHooks = append(echoBeforeInsertHooks, echoHook)
		echoBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		echoAfterInsertMu.Lock()
		echoAfterInsertHooks = append(echoAfterInsertHooks, echoHook)
		echoAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		echoBeforeUpdateMu.Lock()
		echoBeforeUpdateHooks = append(echoBeforeUpdateHooks, echoHook)
		echoBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		echoAfterUpdateMu.Lock()
		echoAfterUpdateHooks = append(echoAfterUpdateHooks, echoHook)
		echoAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		echoBeforeDeleteMu.Lock()
		echoBeforeDeleteHooks = append(echoBeforeDeleteHooks, echoHook)
		echoBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		echoAfterDeleteMu.Lock()
		echoAfterDeleteHooks = append(echoAfterDeleteHooks, echoHook)
		echoAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		echoBeforeUpsertMu.Lock()
		echoBeforeUpsertHooks = append(echoBeforeUpsertHooks, echoHook)
		echoBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		echoAfterUpsertMu.Lock()
		echoAfterUpsertHooks = append(echoAfterUpsertHooks, echoHook)
		echoAfterUpsertMu.Unlock()
	}
}

// One returns a single echo record from the query.
func (q echoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Echo, error) {
	o := &Echo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dao: failed to execute a one query for echos")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Echo records from the query.
func (q echoQuery) All(ctx context.Context, exec boil.ContextExecutor) (EchoSlice, error) {
	var o []*Echo

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dao: failed to assign all query results to Echo slice")
	}

	if len(echoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Echo records in the query.
func (q echoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to count echos rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q echoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dao: failed to check if echos exists")
	}

	return count > 0, nil
}

// Echos retrieves all the records using an executor.
func Echos(mods ...qm.QueryMod) echoQuery {
	mods = append(mods, qm.From("\"echos\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"echos\".*"})
	}

	return echoQuery{q}
}

// FindEcho retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEcho(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Echo, error) {
	echoObj := &Echo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"echos\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, echoObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dao: unable to select from echos")
	}

	if err = echoObj.doAfterSelectHooks(ctx, exec); err != nil {
		return echoObj, err
	}

	return echoObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Echo) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dao: no echos provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(echoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	echoInsertCacheMut.RLock()
	cache, cached := echoInsertCache[key]
	echoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			echoAllColumns,
			echoColumnsWithDefault,
			echoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(echoType, echoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(echoType, echoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"echos\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"echos\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dao: unable to insert into echos")
	}

	if !cached {
		echoInsertCacheMut.Lock()
		echoInsertCache[key] = cache
		echoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Echo.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Echo) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	echoUpdateCacheMut.RLock()
	cache, cached := echoUpdateCache[key]
	echoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			echoAllColumns,
			echoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dao: unable to update echos, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"echos\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, echoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(echoType, echoMapping, append(wl, echoPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dao: unable to update echos row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by update for echos")
	}

	if !cached {
		echoUpdateCacheMut.Lock()
		echoUpdateCache[key] = cache
		echoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q echoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to update all for echos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to retrieve rows affected for echos")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EchoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), echoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"echos\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, echoPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to update all in echo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to retrieve rows affected all in update all echo")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Echo) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("dao: no echos provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(echoColumnsWithDefault, o)

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

	echoUpsertCacheMut.RLock()
	cache, cached := echoUpsertCache[key]
	echoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			echoAllColumns,
			echoColumnsWithDefault,
			echoColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			echoAllColumns,
			echoPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dao: unable to upsert echos, could not build update column list")
		}

		ret := strmangle.SetComplement(echoAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(echoPrimaryKeyColumns) == 0 {
				return errors.New("dao: unable to upsert echos, could not build conflict column list")
			}

			conflict = make([]string, len(echoPrimaryKeyColumns))
			copy(conflict, echoPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"echos\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(echoType, echoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(echoType, echoMapping, ret)
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
		return errors.Wrap(err, "dao: unable to upsert echos")
	}

	if !cached {
		echoUpsertCacheMut.Lock()
		echoUpsertCache[key] = cache
		echoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Echo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Echo) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dao: no Echo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), echoPrimaryKeyMapping)
	sql := "DELETE FROM \"echos\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete from echos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by delete for echos")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q echoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dao: no echoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete all from echos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by deleteall for echos")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EchoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(echoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), echoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"echos\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, echoPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete all from echo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by deleteall for echos")
	}

	if len(echoAfterDeleteHooks) != 0 {
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
func (o *Echo) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEcho(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EchoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EchoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), echoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"echos\".* FROM \"echos\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, echoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dao: unable to reload all in EchoSlice")
	}

	*o = slice

	return nil
}

// EchoExists checks if the Echo row exists.
func EchoExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"echos\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dao: unable to check if echos exists")
	}

	return exists, nil
}

// Exists checks if the Echo row exists.
func (o *Echo) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EchoExists(ctx, exec, o.ID)
}

// /////////////////////////////// BEGIN EXTENSIONS /////////////////////////////////
// Expose table columns
var (
	EchoAllColumns            = echoAllColumns
	EchoColumnsWithoutDefault = echoColumnsWithoutDefault
	EchoColumnsWithDefault    = echoColumnsWithDefault
	EchoPrimaryKeyColumns     = echoPrimaryKeyColumns
	EchoGeneratedColumns      = echoGeneratedColumns
)

// InsertAll inserts all rows with the specified column values, using an executor.
func (o EchoSlice) InsertAll(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
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
			echoAllColumns,
			echoColumnsWithDefault,
			echoColumnsWithoutDefault,
			queries.NonZeroDefaultSet(echoColumnsWithDefault, row),
		)
		if i == 0 {
			sql = "INSERT INTO \"echos\" " + "(\"" + strings.Join(wl, "\",\"") + "\")" + " VALUES "
		}
		sql += strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), len(vals)+1, len(wl))
		if i != len(o)-1 {
			sql += ","
		}
		valMapping, err := queries.BindMapping(echoType, echoMapping, wl)
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
		return 0, errors.Wrap(err, "dao: unable to insert all from echo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by insertall for echos")
	}

	if len(echoAfterInsertHooks) != 0 {
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
func (o EchoSlice) UpsertAll(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	nzDefaults := queries.NonZeroDefaultSet(echoColumnsWithDefault, o[0])

	insert, _ := insertColumns.InsertColumnSet(
		echoAllColumns,
		echoColumnsWithDefault,
		echoColumnsWithoutDefault,
		nzDefaults,
	)
	update := updateColumns.UpdateColumnSet(
		echoAllColumns,
		echoPrimaryKeyColumns,
	)

	if updateOnConflict && len(update) == 0 {
		return 0, errors.New("dao: unable to upsert echos, could not build update column list")
	}

	conflict := conflictColumns
	if len(conflict) == 0 {
		conflict = make([]string, len(echoPrimaryKeyColumns))
		copy(conflict, echoPrimaryKeyColumns)
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
		"\"echos\"",
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
	valueMapping, err := queries.BindMapping(echoType, echoMapping, insert)
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
		return 0, errors.Wrap(err, "dao: unable to upsert for echos")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by upsert for echos")
	}

	if len(echoAfterUpsertHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterUpsertHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

///////////////////////////////// END EXTENSIONS /////////////////////////////////
