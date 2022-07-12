package domain


type TaobaoTbkScPunishOrderGetResult struct {
    /*
        淘宝联盟unionid（该字段不再支持）     */
    UnionId  *string `json:"union_id,omitempty" `

    /*
        会员运营id（该字段不再支持）     */
    SpecialId  *int64 `json:"special_id,omitempty" `

    /*
        渠道关系id     */
    RelationId  *int64 `json:"relation_id,omitempty" `

    /*
        结算月份     */
    SettleMonth  *int64 `json:"settle_month,omitempty" `

    /*
        处罚状态。0 冻结，1 解冻     */
    PunishStatus  *string `json:"punish_status,omitempty" `

    /*
        处罚类型，目前包括 1:店铺淘宝客 2:订单虚假交易     */
    ViolationType  *string `json:"violation_type,omitempty" `

    /*
        淘客订单创建时间     */
    TkTradeCreateTime  *string `json:"tk_trade_create_time,omitempty" `

    /*
        子订单号     */
    TbTradeId  *int64 `json:"tb_trade_id,omitempty" `

    /*
        父订单号（该字段不再支持）     */
    TbTradeParentId  *int64 `json:"tb_trade_parent_id,omitempty" `

    /*
        pid里的adzoneid     */
    TkAdzoneId  *int64 `json:"tk_adzone_id,omitempty" `

    /*
        pid里的siteid     */
    TkSiteId  *int64 `json:"tk_site_id,omitempty" `

    /*
        pid里的pubid     */
    TkPubId  *int64 `json:"tk_pub_id,omitempty" `

}

func (s *TaobaoTbkScPunishOrderGetResult) SetUnionId(v string) *TaobaoTbkScPunishOrderGetResult {
    s.UnionId = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetSpecialId(v int64) *TaobaoTbkScPunishOrderGetResult {
    s.SpecialId = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetRelationId(v int64) *TaobaoTbkScPunishOrderGetResult {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetSettleMonth(v int64) *TaobaoTbkScPunishOrderGetResult {
    s.SettleMonth = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetPunishStatus(v string) *TaobaoTbkScPunishOrderGetResult {
    s.PunishStatus = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetViolationType(v string) *TaobaoTbkScPunishOrderGetResult {
    s.ViolationType = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetTkTradeCreateTime(v string) *TaobaoTbkScPunishOrderGetResult {
    s.TkTradeCreateTime = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetTbTradeId(v int64) *TaobaoTbkScPunishOrderGetResult {
    s.TbTradeId = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetTbTradeParentId(v int64) *TaobaoTbkScPunishOrderGetResult {
    s.TbTradeParentId = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetTkAdzoneId(v int64) *TaobaoTbkScPunishOrderGetResult {
    s.TkAdzoneId = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetTkSiteId(v int64) *TaobaoTbkScPunishOrderGetResult {
    s.TkSiteId = &v
    return s
}
func (s *TaobaoTbkScPunishOrderGetResult) SetTkPubId(v int64) *TaobaoTbkScPunishOrderGetResult {
    s.TkPubId = &v
    return s
}
