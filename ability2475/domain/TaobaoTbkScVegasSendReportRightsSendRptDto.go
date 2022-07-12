package domain


type TaobaoTbkScVegasSendReportRightsSendRptDto struct {
    /*
        渠道关系id的发放数据的数据集     */
    RelationRptList  *[]TaobaoTbkScVegasSendReportRightsSendRelationRptDto `json:"relation_rpt_list,omitempty" `

}

func (s *TaobaoTbkScVegasSendReportRightsSendRptDto) SetRelationRptList(v []TaobaoTbkScVegasSendReportRightsSendRelationRptDto) *TaobaoTbkScVegasSendReportRightsSendRptDto {
    s.RelationRptList = &v
    return s
}
