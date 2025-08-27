package api

type Extra struct {
	RepayUnit   string   `json:"repay_unit"`   //还款单位
	RepayImages []string `json:"repay_images"` //还款证明图
}
