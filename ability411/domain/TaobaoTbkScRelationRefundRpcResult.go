package domain


type TaobaoTbkScRelationRefundRpcResult struct {
    /*
        返回信息     */
    ResultMsg  *string `json:"result_msg,omitempty" `

    /*
        真正的业务数据结构     */
    Data  *TaobaoTbkScRelationRefundPageResult `json:"data,omitempty" `

    /*
        接口返回值信息，跟rpc架构保持一致     */
    ResultCode  *int64 `json:"result_code,omitempty" `

    /*
        业务错误信息     */
    BizErrorDesc  *string `json:"biz_error_desc,omitempty" `

    /*
        业务错误码 101, 102,103     */
    BizErrorCode  *int64 `json:"biz_error_code,omitempty" `

}

func (s *TaobaoTbkScRelationRefundRpcResult) SetResultMsg(v string) *TaobaoTbkScRelationRefundRpcResult {
    s.ResultMsg = &v
    return s
}
func (s *TaobaoTbkScRelationRefundRpcResult) SetData(v TaobaoTbkScRelationRefundPageResult) *TaobaoTbkScRelationRefundRpcResult {
    s.Data = &v
    return s
}
func (s *TaobaoTbkScRelationRefundRpcResult) SetResultCode(v int64) *TaobaoTbkScRelationRefundRpcResult {
    s.ResultCode = &v
    return s
}
func (s *TaobaoTbkScRelationRefundRpcResult) SetBizErrorDesc(v string) *TaobaoTbkScRelationRefundRpcResult {
    s.BizErrorDesc = &v
    return s
}
func (s *TaobaoTbkScRelationRefundRpcResult) SetBizErrorCode(v int64) *TaobaoTbkScRelationRefundRpcResult {
    s.BizErrorCode = &v
    return s
}
