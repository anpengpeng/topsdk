package domain


type TaobaoTbkScVegasSendReportRightsSendRelationRptDto struct {
    /*
        统计日期     */
    BizDate  *string `json:"biz_date,omitempty" `

    /*
        渠道关系id     */
    RelationId  *int64 `json:"relation_id,omitempty" `

    /*
        红包发放数量     */
    FundNum  *int64 `json:"fund_num,omitempty" `

}

func (s *TaobaoTbkScVegasSendReportRightsSendRelationRptDto) SetBizDate(v string) *TaobaoTbkScVegasSendReportRightsSendRelationRptDto {
    s.BizDate = &v
    return s
}
func (s *TaobaoTbkScVegasSendReportRightsSendRelationRptDto) SetRelationId(v int64) *TaobaoTbkScVegasSendReportRightsSendRelationRptDto {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScVegasSendReportRightsSendRelationRptDto) SetFundNum(v int64) *TaobaoTbkScVegasSendReportRightsSendRelationRptDto {
    s.FundNum = &v
    return s
}
