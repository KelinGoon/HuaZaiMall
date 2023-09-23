package service

import (
	"github.com/gin-gonic/gin"
	"singo/dal/model"
	"singo/dao"
	"singo/module/response"
	"strconv"
)

// ProductInfo 查询货物信息
func ProductInfo(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var models dao.SitePaginationSelectModel
	// 从查询参数中获取页数和每页记录数，默认值为1和10
	pagesize, _ := strconv.Atoi(c.Param("pagesize"))
	pagenum, _ := strconv.Atoi(c.Param("pagenum"))

	// 设置模型的分页参数
	models.PageSize = pagesize
	models.PageNum = pagenum
	models.UserId = int(*userId.(*uint))

	var ProductList []*model.Product
	total, product_list, err := dao.ProductInfoById(models)
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
	} else {
		// 将查询到的商品信息添加到 ProductList 切片中
		ProductList = append(ProductList, product_list...)
		response.OKRespWithData(c, map[string]interface{}{
			"data":  ProductList,
			"total": total,
		})
	}
}

// AddProduct 添加商品
func AddProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBind(&product); err != nil { //进行数据校验和绑定操作
		response.ErrRespWithMsg(c, err.Error())
	} else {
		userId, _ := c.Get("user_id")
		product.SellerID = int64(*userId.(*uint))
		err := dao.AddProductInfo(product)
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
		} else {
			response.OKResp(c)
		}
	}
}

// DeleteProduct 根据ID进行删除操作
func DeleteProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBind(&product); err != nil { //进行数据校验和绑定操作
		response.ErrRespWithMsg(c, err.Error())
	} else {
		err := dao.DeleteProductInfo(product.ProductID)
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		response.OKResp(c)
	}
}

func UpdateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBind(&product); err != nil { //进行数据校验和绑定操作
		response.ErrRespWithMsg(c, err.Error())
	} else {
		err := dao.ChangeProductInfo(product)
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
			return
		}
		response.OKResp(c)
	}
}
