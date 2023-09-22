package dao

import (
	"singo/dal/model"
	"singo/dal/query"
)

// ProductInfoById 根据店铺id查询商品列表
func ProductInfoById(ID int64) ([]*model.Product, error) {
	//var product model.Product
	q := query.Product
	product, err := q.Where(q.SellerID.Eq(ID)).Find()
	return product, err
}

func AddProductInfo(product model.Product) error {
	q := query.Product
	err := q.Create(&product)
	return err
}

func DeleteProductInfo(ID int64) (err error) {
	q := query.Product
	_, err = q.Where(q.ProductID.Eq(ID)).Delete()
	return err
}

func ChangeProductInfo(product model.Product) (err error) {
	q := query.Product
	_, err = q.Where(q.ProductID.Eq(product.ProductID)).Update(q.ALL, product)
	return err
}
