// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOrders(t *testing.T) {
	t.Parallel()

	query := Orders()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrdersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrdersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Orders().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrdersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrderSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrdersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrderExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Order exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrderExists to return true, but got false.")
	}
}

func testOrdersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	orderFound, err := FindOrder(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if orderFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrdersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Orders().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrdersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Orders().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrdersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	orderOne := &Order{}
	orderTwo := &Order{}
	if err = randomize.Struct(seed, orderOne, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}
	if err = randomize.Struct(seed, orderTwo, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Orders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrdersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	orderOne := &Order{}
	orderTwo := &Order{}
	if err = randomize.Struct(seed, orderOne, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}
	if err = randomize.Struct(seed, orderTwo, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func orderBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func orderAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func orderAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func orderBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func orderAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func orderBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func orderAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func orderBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func orderAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Order) error {
	*o = Order{}
	return nil
}

func testOrdersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Order{}
	o := &Order{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, orderDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Order object: %s", err)
	}

	AddOrderHook(boil.BeforeInsertHook, orderBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	orderBeforeInsertHooks = []OrderHook{}

	AddOrderHook(boil.AfterInsertHook, orderAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	orderAfterInsertHooks = []OrderHook{}

	AddOrderHook(boil.AfterSelectHook, orderAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	orderAfterSelectHooks = []OrderHook{}

	AddOrderHook(boil.BeforeUpdateHook, orderBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	orderBeforeUpdateHooks = []OrderHook{}

	AddOrderHook(boil.AfterUpdateHook, orderAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	orderAfterUpdateHooks = []OrderHook{}

	AddOrderHook(boil.BeforeDeleteHook, orderBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	orderBeforeDeleteHooks = []OrderHook{}

	AddOrderHook(boil.AfterDeleteHook, orderAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	orderAfterDeleteHooks = []OrderHook{}

	AddOrderHook(boil.BeforeUpsertHook, orderBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	orderBeforeUpsertHooks = []OrderHook{}

	AddOrderHook(boil.AfterUpsertHook, orderAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	orderAfterUpsertHooks = []OrderHook{}
}

func testOrdersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrdersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(orderColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrderToManyBookOrders(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c BookOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookOrderDBTypes, false, bookOrderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookOrderDBTypes, false, bookOrderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.OrderID = a.ID
	c.OrderID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BookOrders().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.OrderID == b.OrderID {
			bFound = true
		}
		if v.OrderID == c.OrderID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OrderSlice{&a}
	if err = a.L.LoadBookOrders(ctx, tx, false, (*[]*Order)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BookOrders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BookOrders = nil
	if err = a.L.LoadBookOrders(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BookOrders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOrderToManyDiscounts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c Discount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, discountDBTypes, false, discountColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, discountDBTypes, false, discountColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"order_discount\" (\"order_id\", \"discount_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"order_discount\" (\"order_id\", \"discount_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Discounts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OrderSlice{&a}
	if err = a.L.LoadDiscounts(ctx, tx, false, (*[]*Order)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Discounts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Discounts = nil
	if err = a.L.LoadDiscounts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Discounts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOrderToManyAddOpBookOrders(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c, d, e BookOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BookOrder{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookOrderDBTypes, false, strmangle.SetComplement(bookOrderPrimaryKeyColumns, bookOrderColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BookOrder{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBookOrders(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.OrderID {
			t.Error("foreign key was wrong value", a.ID, first.OrderID)
		}
		if a.ID != second.OrderID {
			t.Error("foreign key was wrong value", a.ID, second.OrderID)
		}

		if first.R.Order != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Order != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BookOrders[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BookOrders[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BookOrders().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testOrderToManyAddOpDiscounts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c, d, e Discount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Discount{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discountDBTypes, false, strmangle.SetComplement(discountPrimaryKeyColumns, discountColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Discount{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDiscounts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Orders[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Orders[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Discounts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Discounts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Discounts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOrderToManySetOpDiscounts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c, d, e Discount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Discount{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discountDBTypes, false, strmangle.SetComplement(discountPrimaryKeyColumns, discountColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetDiscounts(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Discounts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetDiscounts(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Discounts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Orders) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Orders) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Orders[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Orders[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Discounts[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Discounts[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testOrderToManyRemoveOpDiscounts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c, d, e Discount

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Discount{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, discountDBTypes, false, strmangle.SetComplement(discountPrimaryKeyColumns, discountColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddDiscounts(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Discounts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveDiscounts(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Discounts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Orders) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Orders) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Orders[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Orders[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Discounts) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Discounts[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Discounts[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testOrdersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrdersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrderSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrdersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Orders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	orderDBTypes = map[string]string{`ID`: `bigint`, `CreatedAt`: `timestamp with time zone`, `Total`: `integer`, `Status`: `boolean`}
	_            = bytes.MinRead
)

func testOrdersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(orderPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(orderAllColumns) == len(orderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orderDBTypes, true, orderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrdersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(orderAllColumns) == len(orderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orderDBTypes, true, orderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(orderAllColumns, orderPrimaryKeyColumns) {
		fields = orderAllColumns
	} else {
		fields = strmangle.SetComplement(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := OrderSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrdersUpsert(t *testing.T) {
	t.Parallel()

	if len(orderAllColumns) == len(orderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Order{}
	if err = randomize.Struct(seed, &o, orderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Order: %s", err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, orderDBTypes, false, orderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Order: %s", err)
	}

	count, err = Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
