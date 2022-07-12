package domain


type TaobaoTbkScOrderDetailsGetServiceFeeDto struct {
    /*
        专项服务费率     */
    ShareRelativeRate  *string `json:"share_relative_rate,omitempty" `

    /*
        结算专项服务费     */
    ShareFee  *string `json:"share_fee,omitempty" `

    /*
        预估专项服务费     */
    SharePreFee  *string `json:"share_pre_fee,omitempty" `

    /*
        专项服务费来源，122-渠道     */
    TkShareRoleType  *int64 `json:"tk_share_role_type,omitempty" `

}

func (s *TaobaoTbkScOrderDetailsGetServiceFeeDto) SetShareRelativeRate(v string) *TaobaoTbkScOrderDetailsGetServiceFeeDto {
    s.ShareRelativeRate = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetServiceFeeDto) SetShareFee(v string) *TaobaoTbkScOrderDetailsGetServiceFeeDto {
    s.ShareFee = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetServiceFeeDto) SetSharePreFee(v string) *TaobaoTbkScOrderDetailsGetServiceFeeDto {
    s.SharePreFee = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetServiceFeeDto) SetTkShareRoleType(v int64) *TaobaoTbkScOrderDetailsGetServiceFeeDto {
    s.TkShareRoleType = &v
    return s
}
