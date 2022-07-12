package domain


type TaobaoTbkPrivilegeGetStepRateDto struct {
    /*
        前N件佣金结束时间， 当前N件佣金 失效，本字段置空     */
    TopnEndTime  *int64 `json:"topn_end_time,omitempty" `

    /*
        前N件佣金开始时间，当前N件佣金失效，本字段置空     */
    TopnStartTime  *int64 `json:"topn_start_time,omitempty" `

    /*
        前N件剩余库存，当前N件佣金失效，本字段置空     */
    TopnQuantity  *int64 `json:"topn_quantity,omitempty" `

    /*
        前N件初始总库存，当前N件佣金失效，本字段置空（失效：任务完成、时间结束、商品下架）     */
    TopnTotalCount  *int64 `json:"topn_total_count,omitempty" `

}

func (s *TaobaoTbkPrivilegeGetStepRateDto) SetTopnEndTime(v int64) *TaobaoTbkPrivilegeGetStepRateDto {
    s.TopnEndTime = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetStepRateDto) SetTopnStartTime(v int64) *TaobaoTbkPrivilegeGetStepRateDto {
    s.TopnStartTime = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetStepRateDto) SetTopnQuantity(v int64) *TaobaoTbkPrivilegeGetStepRateDto {
    s.TopnQuantity = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetStepRateDto) SetTopnTotalCount(v int64) *TaobaoTbkPrivilegeGetStepRateDto {
    s.TopnTotalCount = &v
    return s
}
