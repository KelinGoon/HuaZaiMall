package dao

import (
	"singo/dal/model"
	"singo/dal/query"
)

// ProductInfoById 根据店铺id查询商品列表
func ProductInfoById(page SitePaginationSelectModel) (int64, []*model.Product, error) {
	q := query.Product
	total, _ := q.Where(q.SellerID.Eq(int64(page.UserId))).Count()
	product, err := q.Where(q.SellerID.Eq(int64(page.UserId))).Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find()
	return total, product, err
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
	_, err = q.Where(q.ProductID.Eq(product.ProductID)).Updates(model.Product{
		ProductName:        product.ProductName,
		ProductDescription: product.ProductDescription,
		ProductPrice:       product.ProductPrice,
		ProductImageURL:    product.ProductImageURL,
		ProductStock:       product.ProductStock,
		CreationDate:       product.CreationDate,
		UpdateDate:         product.UpdateDate,
		CategoryID:         product.CategoryID,
	})
	return err
}
