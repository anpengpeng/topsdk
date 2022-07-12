package domain


type TaobaoTbkScVegasSendReportResult struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        model     */
    Model  *TaobaoTbkScVegasSendReportRightsSendRptDto `json:"model,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

}

func (s *TaobaoTbkScVegasSendReportResult) SetSuccess(v bool) *TaobaoTbkScVegasSendReportResult {
    s.Success = &v
    return s
}
func (s *TaobaoTbkScVegasSendReportResult) SetModel(v TaobaoTbkScVegasSendReportRightsSendRptDto) *TaobaoTbkScVegasSendReportResult {
    s.Model = &v
    return s
}
func (s *TaobaoTbkScVegasSendReportResult) SetMsgInfo(v string) *TaobaoTbkScVegasSendReportResult {
    s.MsgInfo = &v
    return s
}
func (s *TaobaoTbkScVegasSendReportResult) SetMsgCode(v string) *TaobaoTbkScVegasSendReportResult {
    s.MsgCode = &v
    return s
}
