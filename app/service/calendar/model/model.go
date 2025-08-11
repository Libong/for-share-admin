package model

import "gorm.io/gorm"

type FinanceBill struct {
	gorm.Model
	BillId              string  `gorm:"column:bill_id" json:"bill_id"`                             //账单编号
	DayIncome           float64 `gorm:"column:day_income" json:"day_income"`                       //今日收入
	Cash                float64 `gorm:"column:cash" json:"cash"`                                   //现金收入
	OnlineIncome        float64 `gorm:"column:online_income" json:"online_income"`                 //在线收入
	GoodsExpenses       float64 `gorm:"column:goods_expenses" json:"goods_expenses"`               //货物支出
	DrinksExpenses      float64 `gorm:"column:drinks_expenses" json:"drinks_expenses"`             //酒水支出
	OtherExpenses       float64 `gorm:"column:other_expenses" json:"other_expenses"`               //其他支出
	OtherExpensesRemark string  `gorm:"column:other_expenses_remark" json:"other_expenses_remark"` //其他支出备注
	IsDraw              float64 `gorm:"column:is_draw" json:"is_draw"`                             //是否提款
	DrawIncome          float64 `gorm:"column:draw_income" json:"draw_income"`                     //提款收入
	CardIncome          float64 `gorm:"column:card_income" json:"card_income"`                     //卡收入
	Remark              string  `gorm:"column:remark" json:"remark"`                               //备注
	Timestamp           int64   `gorm:"column:timestamp" json:"timestamp"`                         //时间戳
	Owner               string  `gorm:"column:owner" json:"owner"`                                 //拥有者
}
type UpdateFinanceBillReq struct {
	DayIncome           float64 `json:"day_income"`            //今日收入
	Cash                float64 `json:"cash"`                  //现金收入
	OnlineIncome        float64 `json:"online_income"`         //在线收入
	GoodsExpenses       float64 `json:"goods_expenses"`        //货物支出
	DrinksExpenses      float64 `json:"drinks_expenses"`       //酒水支出
	OtherExpenses       float64 `json:"other_expenses"`        //其他支出
	OtherExpensesRemark string  `json:"other_expenses_remark"` //其他支出备注
	IsDraw              uint32  `json:"is_draw"`               //是否提款
	DrawIncome          float64 `json:"draw_income"`           //提款收入
	CardIncome          float64 `json:"card_income"`           //卡收入
	Remark              string  `json:"remark"`                //备注
	Timestamp           int64   `json:"timestamp"`             //时间戳
}
