package domain

import (
	"github.com/liu8534584/topsdk/util"
)

type TaobaoTbkScRelationRefundResult struct {
	/*
	   淘宝订单编号     */
	TbTradeParentId *int64 `json:"tb_trade_parent_id,omitempty" `

	/*
	   会员关系id     */
	SpecialId *int64 `json:"special_id,omitempty" `

	/*
	   渠道关系id     */
	RelationId *int64 `json:"relation_id,omitempty" `

	/*
	   第三方推广者memberid     */
	Tk3rdPubId *int64 `json:"tk3rd_pub_id,omitempty" `

	/*
	   推广者memberid     */
	TkPubId *int64 `json:"tk_pub_id,omitempty" `

	/*
	   第三方应该返还的补贴     */
	TkSubsidyFeeRefund3rdPub *string `json:"tk_subsidy_fee_refund3rd_pub,omitempty" `

	/*
	   第三方应该返还的佣金     */
	TkCommissionFeeRefund3rdPub *string `json:"tk_commission_fee_refund3rd_pub,omitempty" `

	/*
	   第二方应该返还的补贴(不包括技术服务费)     */
	TkSubsidyFeeRefundPub *string `json:"tk_subsidy_fee_refund_pub,omitempty" `

	/*
	   第二方应该返还的佣金(不包括技术服务费)     */
	TkCommissionFeeRefundPub *string `json:"tk_commission_fee_refund_pub,omitempty" `

	/*
	   维权完成时间     */
	TkRefundSuitTime *util.LocalTime `json:"tk_refund_suit_time,omitempty" `

	/*
	   维权创建时间     */
	TkRefundTime *util.LocalTime `json:"tk_refund_time,omitempty" `

	/*
	   订单结算时间     */
	EarningTime *util.LocalTime `json:"earning_time,omitempty" `

	/*
	   订单创建时间     */
	TbTradeCreateTime *util.LocalTime `json:"tb_trade_create_time,omitempty" `

	/*
	   维权创建(淘客结算回执) 4,维权成功(淘客结算回执) 2,维权失败(淘客结算回执) 3,发生多次维权，待处理      11,从淘客处补扣（钱已结给淘客） 等待扣款 12,从淘客处补扣（钱已结给淘客） 扣款成功 13,从卖家处补扣（钱已结给卖家） 等待扣款 14,从卖家处补扣（钱已结给卖家） 扣款成功 15     */
	RefundStatus *int64 `json:"refund_status,omitempty" `

	/*
	   宝贝标题     */
	TbAuctionTitle *string `json:"tb_auction_title,omitempty" `

	/*
	   淘宝子订单编号     */
	TbTradeId *int64 `json:"tb_trade_id,omitempty" `

	/*
	   维权金额     */
	RefundFee *string `json:"refund_fee,omitempty" `

	/*
	   结算金额     */
	TbTradeFinishPrice *string `json:"tb_trade_finish_price,omitempty" `

	/*
	   应返商家金额(二方)     */
	TkPubShowReturnFee *string `json:"tk_pub_show_return_fee,omitempty" `

	/*
	   应返商家金额(三方)     */
	Tk3rdPubShowReturnFee *string `json:"tk3rd_pub_show_return_fee,omitempty" `

	/*
	   1 表示2方，2表示3方     */
	RefundType *int64 `json:"refund_type,omitempty" `

	/*
	   （口碑订单）口碑父订单号     */
	AlscPid *string `json:"alsc_pid,omitempty" `

	/*
	   （口碑订单）口碑子订单号     */
	AlscId *string `json:"alsc_id,omitempty" `

	/*
	   更新时间     */
	ModifiedTime *util.LocalTime `json:"modified_time,omitempty" `
}

func (s *TaobaoTbkScRelationRefundResult) SetTbTradeParentId(v int64) *TaobaoTbkScRelationRefundResult {
	s.TbTradeParentId = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetSpecialId(v int64) *TaobaoTbkScRelationRefundResult {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetRelationId(v int64) *TaobaoTbkScRelationRefundResult {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTk3rdPubId(v int64) *TaobaoTbkScRelationRefundResult {
	s.Tk3rdPubId = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTkPubId(v int64) *TaobaoTbkScRelationRefundResult {
	s.TkPubId = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTkSubsidyFeeRefund3rdPub(v string) *TaobaoTbkScRelationRefundResult {
	s.TkSubsidyFeeRefund3rdPub = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTkCommissionFeeRefund3rdPub(v string) *TaobaoTbkScRelationRefundResult {
	s.TkCommissionFeeRefund3rdPub = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTkSubsidyFeeRefundPub(v string) *TaobaoTbkScRelationRefundResult {
	s.TkSubsidyFeeRefundPub = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTkCommissionFeeRefundPub(v string) *TaobaoTbkScRelationRefundResult {
	s.TkCommissionFeeRefundPub = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTkRefundSuitTime(v util.LocalTime) *TaobaoTbkScRelationRefundResult {
	s.TkRefundSuitTime = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTkRefundTime(v util.LocalTime) *TaobaoTbkScRelationRefundResult {
	s.TkRefundTime = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetEarningTime(v util.LocalTime) *TaobaoTbkScRelationRefundResult {
	s.EarningTime = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTbTradeCreateTime(v util.LocalTime) *TaobaoTbkScRelationRefundResult {
	s.TbTradeCreateTime = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetRefundStatus(v int64) *TaobaoTbkScRelationRefundResult {
	s.RefundStatus = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTbAuctionTitle(v string) *TaobaoTbkScRelationRefundResult {
	s.TbAuctionTitle = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTbTradeId(v int64) *TaobaoTbkScRelationRefundResult {
	s.TbTradeId = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetRefundFee(v string) *TaobaoTbkScRelationRefundResult {
	s.RefundFee = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTbTradeFinishPrice(v string) *TaobaoTbkScRelationRefundResult {
	s.TbTradeFinishPrice = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTkPubShowReturnFee(v string) *TaobaoTbkScRelationRefundResult {
	s.TkPubShowReturnFee = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetTk3rdPubShowReturnFee(v string) *TaobaoTbkScRelationRefundResult {
	s.Tk3rdPubShowReturnFee = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetRefundType(v int64) *TaobaoTbkScRelationRefundResult {
	s.RefundType = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetAlscPid(v string) *TaobaoTbkScRelationRefundResult {
	s.AlscPid = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetAlscId(v string) *TaobaoTbkScRelationRefundResult {
	s.AlscId = &v
	return s
}
func (s *TaobaoTbkScRelationRefundResult) SetModifiedTime(v util.LocalTime) *TaobaoTbkScRelationRefundResult {
	s.ModifiedTime = &v
	return s
}
