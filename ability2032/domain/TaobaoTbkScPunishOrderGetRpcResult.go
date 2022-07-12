package domain


type TaobaoTbkScPunishOrderGetRpcResult struct {
    /*
        结果     */
    Data  *TaobaoTbkScPunishOrderGetPageResult `json:"data,omitempty" `

    /*
        业务出错的描述     */
    BizErrorDesc  *string `json:"biz_error_desc,omitempty" `

    /*
        业务出错的状态码     */
    BizErrorCode  *int64 `json:"biz_error_code,omitempty" `

    /*
        执行结果     */
    ResultMsg  *string `json:"result_msg,omitempty" `

    /*
        执行结果状态码     */
    ResultCode  *int64 `json:"result_code,omitempty" `

}

func (s *TaobaoTbkScPunishOrderGetRpcResult) SetData(v TaobaoTbkScPunishOrderGetPageResult) *TaobaoTbkScPunishOrderGetRpcResult {
    s.Data = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetRpcResult) SetBizErrorDesc(v string) *TaobaoTbkScPunishOrderGetRpcResult {
    s.BizErrorDesc = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetRpcResult) SetBizErrorCode(v int64) *TaobaoTbkScPunishOrderGetRpcResult {
    s.BizErrorCode = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetRpcResult) SetResultMsg(v string) *TaobaoTbkScPunishOrderGetRpcResult {
    s.ResultMsg = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetRpcResult) SetResultCode(v int64) *TaobaoTbkScPunishOrderGetRpcResult {
    s.ResultCode = &v
    return s
}
