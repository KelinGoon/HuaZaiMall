// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCart = "cart"

// Cart mapped from table <cart>
type Cart struct {
	CartID       int64     `gorm:"column:CartID;type:int(11);primaryKey;autoIncrement:true" json:"CartID"`                    // 购物车条目唯一标识符
	UserID       int64     `gorm:"column:UserID;type:int(11)" json:"UserID"`                                                  // 用户ID，购物车所属用户的外键
	ProductID    int64     `gorm:"column:ProductID;type:int(11)" json:"ProductID"`                                            // 商品ID，购物车中商品的外键
	Quantity     int64     `gorm:"column:Quantity;type:int(11);not null" json:"Quantity"`                                     // 购物车中商品的数量
	Status       string    `gorm:"column:Status;type:enum('未下单','已下单');default:未下单" json:"Status"`                            // 购物车商品状态，可以是 "未下单" 或 "已下单"
	CreationDate time.Time `gorm:"column:CreationDate;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"CreationDate"` // 购物车条目创建日期，默认为当前时间戳
}

// TableName Cart's table name
func (*Cart) TableName() string {
	return TableNameCart
}
