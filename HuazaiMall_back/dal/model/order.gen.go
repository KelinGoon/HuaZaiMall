// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOrder = "order"

// Order mapped from table <order>
type Order struct {
	OrderID         int64     `gorm:"column:OrderID;type:int(11);primaryKey;autoIncrement:true" json:"OrderID"`                   // 订单唯一标识符
	UserID          int64     `gorm:"column:UserID;type:int(11)" json:"UserID"`                                                   // 下订单的用户ID，外键指向用户表
	OrderDate       time.Time `gorm:"column:OrderDate;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"OrderDate"`        // 订单创建日期，默认为当前时间戳
	OrderStatus     string    `gorm:"column:OrderStatus;type:enum('未支付','已支付','已发货','已完成','已取消');default:未支付" json:"OrderStatus"` // 订单状态，可以是 "未支付"、"已支付"、"已发货"、"已完成" 或 "已取消"
	TotalAmount     float64   `gorm:"column:TotalAmount;type:decimal(10,2);not null" json:"TotalAmount"`                          // 订单的总金额
	ShippingAddress string    `gorm:"column:ShippingAddress;type:varchar(255);not null" json:"ShippingAddress"`                   // 订单的配送地址
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}