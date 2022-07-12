package domain


type TaobaoTbkScRelationRefundPageResult struct {
    /*
        pageNo     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        pageSize     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        总值     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        订单列表     */
    Results  *[]TaobaoTbkScRelationRefundResult `json:"results,omitempty" `

}

func (s *TaobaoTbkScRelationRefundPageResult) SetPageNo(v int64) *TaobaoTbkScRelationRefundPageResult {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkScRelationRefundPageResult) SetPageSize(v int64) *TaobaoTbkScRelationRefundPageResult {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkScRelationRefundPageResult) SetTotalCount(v int64) *TaobaoTbkScRelationRefundPageResult {
    s.TotalCount = &v
    return s
}
func (s *TaobaoTbkScRelationRefundPageResult) SetResults(v []TaobaoTbkScRelationRefundResult) *TaobaoTbkScRelationRefundPageResult {
    s.Results = &v
    return s
}
