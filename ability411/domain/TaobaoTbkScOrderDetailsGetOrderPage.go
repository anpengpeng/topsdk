package domain


type TaobaoTbkScOrderDetailsGetOrderPage struct {
    /*
        PublisherOrderDto     */
    Results  *[]TaobaoTbkScOrderDetailsGetPublisherOrderDto `json:"results,omitempty" `

    /*
        是否还有上一页     */
    HasPre  *bool `json:"has_pre,omitempty" `

    /*
        位点字段，由调用方原样传递     */
    PositionIndex  *string `json:"position_index,omitempty" `

    /*
        是否还有下一页     */
    HasNext  *bool `json:"has_next,omitempty" `

    /*
        页码     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *TaobaoTbkScOrderDetailsGetOrderPage) SetResults(v []TaobaoTbkScOrderDetailsGetPublisherOrderDto) *TaobaoTbkScOrderDetailsGetOrderPage {
    s.Results = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetOrderPage) SetHasPre(v bool) *TaobaoTbkScOrderDetailsGetOrderPage {
    s.HasPre = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetOrderPage) SetPositionIndex(v string) *TaobaoTbkScOrderDetailsGetOrderPage {
    s.PositionIndex = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetOrderPage) SetHasNext(v bool) *TaobaoTbkScOrderDetailsGetOrderPage {
    s.HasNext = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetOrderPage) SetPageNo(v int64) *TaobaoTbkScOrderDetailsGetOrderPage {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetOrderPage) SetPageSize(v int64) *TaobaoTbkScOrderDetailsGetOrderPage {
    s.PageSize = &v
    return s
}
