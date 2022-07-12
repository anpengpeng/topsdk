package domain

import (
	"github.com/liu8534584/topsdk/util"
)

type TaobaoTbkScPunishOrderGetTopApiAfOrderOption struct {
	/*
	   渠道关系id     */
	RelationId *int64 `json:"relation_id,omitempty" `

	/*
	   子订单号     */
	TbTradeId *int64 `json:"tb_trade_id,omitempty" `

	/*
	   此参数不再使用，请勿入参     */
	TbTradeParentId *int64 `json:"tb_trade_parent_id,omitempty" `

	/*
	   pagesize     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   pageNo     */
	PageNo *int64 `json:"page_no,omitempty" `

	/*
	   查询时间跨度，不超过30天，单位是天     */
	Span *int64 `json:"span,omitempty" `

	/*
	   查询开始时间，以taoke订单创建时间开始     */
	StartTime *util.LocalTime `json:"start_time,omitempty" `

	/*
	   此参数不再使用，请勿入参     */
	SpecialId *int64 `json:"special_id,omitempty" `

	/*
	   pid中的第三段，adzoneId     */
	AdzoneId *int64 `json:"adzone_id,omitempty" `

	/*
	   pid中的第二段，siteId     */
	SiteId *int64 `json:"site_id,omitempty" `

	/*
	   此参数不再使用，请勿入参     */
	ViolationType *int64 `json:"violation_type,omitempty" `

	/*
	   此参数不再使用，请勿入参     */
	PunishStatus *int64 `json:"punish_status,omitempty" `
}

func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetRelationId(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetTbTradeId(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.TbTradeId = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetTbTradeParentId(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.TbTradeParentId = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetPageSize(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetPageNo(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetSpan(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.Span = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetStartTime(v util.LocalTime) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.StartTime = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetSpecialId(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetAdzoneId(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetSiteId(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.SiteId = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetViolationType(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.ViolationType = &v
	return s
}
func (s *TaobaoTbkScPunishOrderGetTopApiAfOrderOption) SetPunishStatus(v int64) *TaobaoTbkScPunishOrderGetTopApiAfOrderOption {
	s.PunishStatus = &v
	return s
}
