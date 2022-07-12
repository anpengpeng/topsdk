package domain

import (
	"github.com/liu8534584/topsdk/util"
)

type TaobaoTbkScRelationRefundTopApiRefundRptOption struct {
	/*
	   pagesize     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   1-维权发起时间，2-订单结算时间（正向订单），3-维权完成时间，4-订单创建时间，5-订单更新时间     */
	SearchType *int64 `json:"search_type,omitempty" `

	/*
	   1 表示2方，2表示3方，0表示不限     */
	RefundType *int64 `json:"refund_type,omitempty" `

	/*
	   开始时间     */
	StartTime *util.LocalTime `json:"start_time,omitempty" `

	/*
	   pagenumber     */
	PageNo *int64 `json:"page_no,omitempty" `

	/*
	   1代表渠道关系id，2代表会员关系id     */
	BizType *int64 `json:"biz_type,omitempty" `
}

func (s *TaobaoTbkScRelationRefundTopApiRefundRptOption) SetPageSize(v int64) *TaobaoTbkScRelationRefundTopApiRefundRptOption {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkScRelationRefundTopApiRefundRptOption) SetSearchType(v int64) *TaobaoTbkScRelationRefundTopApiRefundRptOption {
	s.SearchType = &v
	return s
}
func (s *TaobaoTbkScRelationRefundTopApiRefundRptOption) SetRefundType(v int64) *TaobaoTbkScRelationRefundTopApiRefundRptOption {
	s.RefundType = &v
	return s
}
func (s *TaobaoTbkScRelationRefundTopApiRefundRptOption) SetStartTime(v util.LocalTime) *TaobaoTbkScRelationRefundTopApiRefundRptOption {
	s.StartTime = &v
	return s
}
func (s *TaobaoTbkScRelationRefundTopApiRefundRptOption) SetPageNo(v int64) *TaobaoTbkScRelationRefundTopApiRefundRptOption {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkScRelationRefundTopApiRefundRptOption) SetBizType(v int64) *TaobaoTbkScRelationRefundTopApiRefundRptOption {
	s.BizType = &v
	return s
}
