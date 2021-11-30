// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// UserBlock is an object representing the database table.
type UserBlock struct {
	ID          int `boil:"id" json:"id" toml:"id" yaml:"id"`
	RequestorID int `boil:"requestor_id" json:"requestor_id" toml:"requestor_id" yaml:"requestor_id"`
	TargetID    int `boil:"target_id" json:"target_id" toml:"target_id" yaml:"target_id"`

	R *userBlockR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userBlockL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserBlockColumns = struct {
	ID          string
	RequestorID string
	TargetID    string
}{
	ID:          "id",
	RequestorID: "requestor_id",
	TargetID:    "target_id",
}

var UserBlockTableColumns = struct {
	ID          string
	RequestorID string
	TargetID    string
}{
	ID:          "user_blocks.id",
	RequestorID: "user_blocks.requestor_id",
	TargetID:    "user_blocks.target_id",
}

// Generated where

var UserBlockWhere = struct {
	ID          whereHelperint
	RequestorID whereHelperint
	TargetID    whereHelperint
}{
	ID:          whereHelperint{field: "\"user_blocks\".\"id\""},
	RequestorID: whereHelperint{field: "\"user_blocks\".\"requestor_id\""},
	TargetID:    whereHelperint{field: "\"user_blocks\".\"target_id\""},
}

// UserBlockRels is where relationship names are stored.
var UserBlockRels = struct {
	Requestor string
	Target    string
}{
	Requestor: "Requestor",
	Target:    "Target",
}

// userBlockR is where relationships are stored.
type userBlockR struct {
	Requestor *User `boil:"Requestor" json:"Requestor" toml:"Requestor" yaml:"Requestor"`
	Target    *User `boil:"Target" json:"Target" toml:"Target" yaml:"Target"`
}

// NewStruct creates a new relationship struct
func (*userBlockR) NewStruct() *userBlockR {
	return &userBlockR{}
}

// userBlockL is where Load methods for each relationship are stored.
type userBlockL struct{}

var (
	userBlockAllColumns            = []string{"id", "requestor_id", "target_id"}
	userBlockColumnsWithoutDefault = []string{"requestor_id", "target_id"}
	userBlockColumnsWithDefault    = []string{"id"}
	userBlockPrimaryKeyColumns     = []string{"id"}
)

type (
	// UserBlockSlice is an alias for a slice of pointers to UserBlock.
	// This should almost always be used instead of []UserBlock.
	UserBlockSlice []*UserBlock
	// UserBlockHook is the signature for custom UserBlock hook methods
	UserBlockHook func(context.Context, boil.ContextExecutor, *UserBlock) error

	userBlockQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userBlockType                 = reflect.TypeOf(&UserBlock{})
	userBlockMapping              = queries.MakeStructMapping(userBlockType)
	userBlockPrimaryKeyMapping, _ = queries.BindMapping(userBlockType, userBlockMapping, userBlockPrimaryKeyColumns)
	userBlockInsertCacheMut       sync.RWMutex
	userBlockInsertCache          = make(map[string]insertCache)
	userBlockUpdateCacheMut       sync.RWMutex
	userBlockUpdateCache          = make(map[string]updateCache)
	userBlockUpsertCacheMut       sync.RWMutex
	userBlockUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userBlockBeforeInsertHooks []UserBlockHook
var userBlockBeforeUpdateHooks []UserBlockHook
var userBlockBeforeDeleteHooks []UserBlockHook
var userBlockBeforeUpsertHooks []UserBlockHook

var userBlockAfterInsertHooks []UserBlockHook
var userBlockAfterSelectHooks []UserBlockHook
var userBlockAfterUpdateHooks []UserBlockHook
var userBlockAfterDeleteHooks []UserBlockHook
var userBlockAfterUpsertHooks []UserBlockHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserBlock) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserBlock) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserBlock) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserBlock) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserBlock) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserBlock) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserBlock) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserBlock) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserBlock) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBlockAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserBlockHook registers your hook function for all future operations.
func AddUserBlockHook(hookPoint boil.HookPoint, userBlockHook UserBlockHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userBlockBeforeInsertHooks = append(userBlockBeforeInsertHooks, userBlockHook)
	case boil.BeforeUpdateHook:
		userBlockBeforeUpdateHooks = append(userBlockBeforeUpdateHooks, userBlockHook)
	case boil.BeforeDeleteHook:
		userBlockBeforeDeleteHooks = append(userBlockBeforeDeleteHooks, userBlockHook)
	case boil.BeforeUpsertHook:
		userBlockBeforeUpsertHooks = append(userBlockBeforeUpsertHooks, userBlockHook)
	case boil.AfterInsertHook:
		userBlockAfterInsertHooks = append(userBlockAfterInsertHooks, userBlockHook)
	case boil.AfterSelectHook:
		userBlockAfterSelectHooks = append(userBlockAfterSelectHooks, userBlockHook)
	case boil.AfterUpdateHook:
		userBlockAfterUpdateHooks = append(userBlockAfterUpdateHooks, userBlockHook)
	case boil.AfterDeleteHook:
		userBlockAfterDeleteHooks = append(userBlockAfterDeleteHooks, userBlockHook)
	case boil.AfterUpsertHook:
		userBlockAfterUpsertHooks = append(userBlockAfterUpsertHooks, userBlockHook)
	}
}

// One returns a single userBlock record from the query.
func (q userBlockQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserBlock, error) {
	o := &UserBlock{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_blocks")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserBlock records from the query.
func (q userBlockQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserBlockSlice, error) {
	var o []*UserBlock

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserBlock slice")
	}

	if len(userBlockAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserBlock records in the query.
func (q userBlockQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_blocks rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userBlockQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_blocks exists")
	}

	return count > 0, nil
}

// Requestor pointed to by the foreign key.
func (o *UserBlock) Requestor(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RequestorID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Target pointed to by the foreign key.
func (o *UserBlock) Target(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TargetID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadRequestor allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userBlockL) LoadRequestor(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserBlock interface{}, mods queries.Applicator) error {
	var slice []*UserBlock
	var object *UserBlock

	if singular {
		object = maybeUserBlock.(*UserBlock)
	} else {
		slice = *maybeUserBlock.(*[]*UserBlock)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userBlockR{}
		}
		args = append(args, object.RequestorID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userBlockR{}
			}

			for _, a := range args {
				if a == obj.RequestorID {
					continue Outer
				}
			}

			args = append(args, obj.RequestorID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
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

	if len(userBlockAfterSelectHooks) != 0 {
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
		object.R.Requestor = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.RequestorUserBlocks = append(foreign.R.RequestorUserBlocks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RequestorID == foreign.ID {
				local.R.Requestor = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.RequestorUserBlocks = append(foreign.R.RequestorUserBlocks, local)
				break
			}
		}
	}

	return nil
}

// LoadTarget allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userBlockL) LoadTarget(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserBlock interface{}, mods queries.Applicator) error {
	var slice []*UserBlock
	var object *UserBlock

	if singular {
		object = maybeUserBlock.(*UserBlock)
	} else {
		slice = *maybeUserBlock.(*[]*UserBlock)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userBlockR{}
		}
		args = append(args, object.TargetID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userBlockR{}
			}

			for _, a := range args {
				if a == obj.TargetID {
					continue Outer
				}
			}

			args = append(args, obj.TargetID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
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

	if len(userBlockAfterSelectHooks) != 0 {
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
		object.R.Target = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.TargetUserBlocks = append(foreign.R.TargetUserBlocks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TargetID == foreign.ID {
				local.R.Target = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.TargetUserBlocks = append(foreign.R.TargetUserBlocks, local)
				break
			}
		}
	}

	return nil
}

// SetRequestor of the userBlock to the related item.
// Sets o.R.Requestor to related.
// Adds o to related.R.RequestorUserBlocks.
func (o *UserBlock) SetRequestor(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_blocks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"requestor_id"}),
		strmangle.WhereClause("\"", "\"", 2, userBlockPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RequestorID = related.ID
	if o.R == nil {
		o.R = &userBlockR{
			Requestor: related,
		}
	} else {
		o.R.Requestor = related
	}

	if related.R == nil {
		related.R = &userR{
			RequestorUserBlocks: UserBlockSlice{o},
		}
	} else {
		related.R.RequestorUserBlocks = append(related.R.RequestorUserBlocks, o)
	}

	return nil
}

// SetTarget of the userBlock to the related item.
// Sets o.R.Target to related.
// Adds o to related.R.TargetUserBlocks.
func (o *UserBlock) SetTarget(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_blocks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"target_id"}),
		strmangle.WhereClause("\"", "\"", 2, userBlockPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TargetID = related.ID
	if o.R == nil {
		o.R = &userBlockR{
			Target: related,
		}
	} else {
		o.R.Target = related
	}

	if related.R == nil {
		related.R = &userR{
			TargetUserBlocks: UserBlockSlice{o},
		}
	} else {
		related.R.TargetUserBlocks = append(related.R.TargetUserBlocks, o)
	}

	return nil
}

// UserBlocks retrieves all the records using an executor.
func UserBlocks(mods ...qm.QueryMod) userBlockQuery {
	mods = append(mods, qm.From("\"user_blocks\""))
	return userBlockQuery{NewQuery(mods...)}
}

// FindUserBlock retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserBlock(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*UserBlock, error) {
	userBlockObj := &UserBlock{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_blocks\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userBlockObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_blocks")
	}

	if err = userBlockObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userBlockObj, err
	}

	return userBlockObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserBlock) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_blocks provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userBlockColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userBlockInsertCacheMut.RLock()
	cache, cached := userBlockInsertCache[key]
	userBlockInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userBlockAllColumns,
			userBlockColumnsWithDefault,
			userBlockColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userBlockType, userBlockMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userBlockType, userBlockMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_blocks\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_blocks\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_blocks")
	}

	if !cached {
		userBlockInsertCacheMut.Lock()
		userBlockInsertCache[key] = cache
		userBlockInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserBlock.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserBlock) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userBlockUpdateCacheMut.RLock()
	cache, cached := userBlockUpdateCache[key]
	userBlockUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userBlockAllColumns,
			userBlockPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_blocks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_blocks\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userBlockPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userBlockType, userBlockMapping, append(wl, userBlockPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_blocks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_blocks")
	}

	if !cached {
		userBlockUpdateCacheMut.Lock()
		userBlockUpdateCache[key] = cache
		userBlockUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userBlockQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_blocks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_blocks")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserBlockSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBlockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_blocks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userBlockPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userBlock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userBlock")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserBlock) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_blocks provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userBlockColumnsWithDefault, o)

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

	userBlockUpsertCacheMut.RLock()
	cache, cached := userBlockUpsertCache[key]
	userBlockUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userBlockAllColumns,
			userBlockColumnsWithDefault,
			userBlockColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userBlockAllColumns,
			userBlockPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_blocks, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userBlockPrimaryKeyColumns))
			copy(conflict, userBlockPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_blocks\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userBlockType, userBlockMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userBlockType, userBlockMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert user_blocks")
	}

	if !cached {
		userBlockUpsertCacheMut.Lock()
		userBlockUpsertCache[key] = cache
		userBlockUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserBlock record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserBlock) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserBlock provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userBlockPrimaryKeyMapping)
	sql := "DELETE FROM \"user_blocks\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_blocks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_blocks")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userBlockQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userBlockQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_blocks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_blocks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserBlockSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userBlockBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBlockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_blocks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userBlockPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userBlock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_blocks")
	}

	if len(userBlockAfterDeleteHooks) != 0 {
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
func (o *UserBlock) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserBlock(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserBlockSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserBlockSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userBlockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_blocks\".* FROM \"user_blocks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userBlockPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserBlockSlice")
	}

	*o = slice

	return nil
}

// UserBlockExists checks if the UserBlock row exists.
func UserBlockExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_blocks\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_blocks exists")
	}

	return exists, nil
}