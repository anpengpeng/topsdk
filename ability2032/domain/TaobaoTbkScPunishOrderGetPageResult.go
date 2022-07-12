package domain


type TaobaoTbkScPunishOrderGetPageResult struct {
    /*
        处罚订单列表     */
    Results  *[]TaobaoTbkScPunishOrderGetResult `json:"results,omitempty" `

    /*
        翻页的pageno     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        翻页的pagesie     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        一共能查询出来的结果总数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

}

func (s *TaobaoTbkScPunishOrderGetPageResult) SetResults(v []TaobaoTbkScPunishOrderGetResult) *TaobaoTbkScPunishOrderGetPageResult {
    s.Results = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetPageResult) SetPageNo(v int64) *TaobaoTbkScPunishOrderGetPageResult {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetPageResult) SetPageSize(v int64) *TaobaoTbkScPunishOrderGetPageResult {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetPageResult) SetTotalCount(v int64) *TaobaoTbkScPunishOrderGetPageResult {
    s.TotalCount = &v
    return s
}
