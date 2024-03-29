// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

var (
	Q         = new(Query)
	Cart      *cart
	Category  *category
	Order     *order
	Orderitem *orderitem
	Product   *product
	Review    *review
	User      *user
)

func SetDefault(db *gorm.DB) {
	*Q = *Use(db)
	Cart = &Q.Cart
	Category = &Q.Category
	Order = &Q.Order
	Orderitem = &Q.Orderitem
	Product = &Q.Product
	Review = &Q.Review
	User = &Q.User
}

func Use(db *gorm.DB) *Query {
	return &Query{
		db:        db,
		Cart:      newCart(db),
		Category:  newCategory(db),
		Order:     newOrder(db),
		Orderitem: newOrderitem(db),
		Product:   newProduct(db),
		Review:    newReview(db),
		User:      newUser(db),
	}
}

type Query struct {
	db *gorm.DB

	Cart      cart
	Category  category
	Order     order
	Orderitem orderitem
	Product   product
	Review    review
	User      user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:        db,
		Cart:      q.Cart.clone(db),
		Category:  q.Category.clone(db),
		Order:     q.Order.clone(db),
		Orderitem: q.Orderitem.clone(db),
		Product:   q.Product.clone(db),
		Review:    q.Review.clone(db),
		User:      q.User.clone(db),
	}
}

type queryCtx struct {
	Cart      ICartDo
	Category  ICategoryDo
	Order     IOrderDo
	Orderitem IOrderitemDo
	Product   IProductDo
	Review    IReviewDo
	User      IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Cart:      q.Cart.WithContext(ctx),
		Category:  q.Category.WithContext(ctx),
		Order:     q.Order.WithContext(ctx),
		Orderitem: q.Orderitem.WithContext(ctx),
		Product:   q.Product.WithContext(ctx),
		Review:    q.Review.WithContext(ctx),
		User:      q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
