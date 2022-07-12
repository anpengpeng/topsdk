package domain


type TaobaoTbkPrivilegeGetMaterialDto struct {
    /*
        优惠券剩余量     */
    CouponRemainCount  *int64 `json:"coupon_remain_count,omitempty" `

    /*
        优惠券总量     */
    CouponTotalCount  *int64 `json:"coupon_total_count,omitempty" `

    /*
        优惠券面额     */
    CouponInfo  *string `json:"coupon_info,omitempty" `

    /*
        优惠券结束时间     */
    CouponEndTime  *string `json:"coupon_end_time,omitempty" `

    /*
        优惠券开始时间     */
    CouponStartTime  *string `json:"coupon_start_time,omitempty" `

    /*
        优惠券(商品优惠券推广链接中的券)类型，1 公开券，2 私有券，3 妈妈券     */
    CouponType  *int64 `json:"coupon_type,omitempty" `

    /*
        商品优惠券推广链接     */
    CouponClickUrl  *string `json:"coupon_click_url,omitempty" `

    /*
        后台一级类目     */
    CategoryId  *int64 `json:"category_id,omitempty" `

    /*
        当不入参special_id、relation_id、external_id时，展示常规佣金率(%)     */
    MaxCommissionRate  *string `json:"max_commission_rate,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品淘客链接     */
    ItemUrl  *string `json:"item_url,omitempty" `

    /*
        预售有礼-推广链接     */
    YsylClickUrl  *string `json:"ysyl_click_url,omitempty" `

    /*
        预售有礼-预估淘礼金（元）     */
    YsylTljFace  *string `json:"ysyl_tlj_face,omitempty" `

    /*
        预售有礼-淘礼金发放时间     */
    YsylTljSendTime  *string `json:"ysyl_tlj_send_time,omitempty" `

    /*
        预售有礼-淘礼金使用开始时间     */
    YsylTljUseStartTime  *string `json:"ysyl_tlj_use_start_time,omitempty" `

    /*
        预售有礼-佣金比例（%）（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）     */
    YsylCommissionRate  *string `json:"ysyl_commission_rate,omitempty" `

    /*
        预售有礼-淘礼金使用结束时间     */
    YsylTljUseEndTime  *string `json:"ysyl_tlj_use_end_time,omitempty" `

    /*
        当入参special_id、relation_id、external_id时，该字段展示预估最低佣金率(%)     */
    MinCommissionRate  *string `json:"min_commission_rate,omitempty" `

    /*
        比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0     */
    RewardInfo  *int64 `json:"reward_info,omitempty" `

    /*
        前N件佣金信息-当入参get_topn_rate=1，前N件佣金生效且最高，透出该组字段     */
    TopnInfo  *TaobaoTbkPrivilegeGetStepRateDto `json:"topn_info,omitempty" `

    /*
        小程序链接(暂未对外开放)     */
    MiniProgram  *TaobaoTbkPrivilegeGetMiniProgramDto `json:"mini_program,omitempty" `

}

func (s *TaobaoTbkPrivilegeGetMaterialDto) SetCouponRemainCount(v int64) *TaobaoTbkPrivilegeGetMaterialDto {
    s.CouponRemainCount = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetCouponTotalCount(v int64) *TaobaoTbkPrivilegeGetMaterialDto {
    s.CouponTotalCount = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetCouponInfo(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.CouponInfo = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetCouponEndTime(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.CouponEndTime = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetCouponStartTime(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.CouponStartTime = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetCouponType(v int64) *TaobaoTbkPrivilegeGetMaterialDto {
    s.CouponType = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetCouponClickUrl(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.CouponClickUrl = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetCategoryId(v int64) *TaobaoTbkPrivilegeGetMaterialDto {
    s.CategoryId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetMaxCommissionRate(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.MaxCommissionRate = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetItemId(v int64) *TaobaoTbkPrivilegeGetMaterialDto {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetItemUrl(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.ItemUrl = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetYsylClickUrl(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.YsylClickUrl = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetYsylTljFace(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.YsylTljFace = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetYsylTljSendTime(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.YsylTljSendTime = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetYsylTljUseStartTime(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.YsylTljUseStartTime = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetYsylCommissionRate(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.YsylCommissionRate = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetYsylTljUseEndTime(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.YsylTljUseEndTime = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetMinCommissionRate(v string) *TaobaoTbkPrivilegeGetMaterialDto {
    s.MinCommissionRate = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetRewardInfo(v int64) *TaobaoTbkPrivilegeGetMaterialDto {
    s.RewardInfo = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetTopnInfo(v TaobaoTbkPrivilegeGetStepRateDto) *TaobaoTbkPrivilegeGetMaterialDto {
    s.TopnInfo = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMaterialDto) SetMiniProgram(v TaobaoTbkPrivilegeGetMiniProgramDto) *TaobaoTbkPrivilegeGetMaterialDto {
    s.MiniProgram = &v
    return s
}
