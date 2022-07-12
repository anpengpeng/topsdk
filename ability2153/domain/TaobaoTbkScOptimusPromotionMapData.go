package domain


type TaobaoTbkScOptimusPromotionMapData struct {
    /*
        权益类型。1 有价券（需要购买的店铺券或单品券） 2 公开券（直接领取的店铺券或单品券） 3 妈妈券（妈妈渠道领取的店铺券或单品券） 4.品类券 （跨店可用券，可与1，2，3叠加）     */
    PromotionType  *string `json:"promotion_type,omitempty" `

    /*
        优惠门槛类型： 1 满元 2 满件     */
    ConditionType  *string `json:"condition_type,omitempty" `

    /*
        优惠类型： 1 减钱 2 打折     */
    DiscountType  *string `json:"discount_type,omitempty" `

    /*
        权益信息-总量（权益初始库存量）     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        权益信息-剩余库存（权益剩余库存量）     */
    RemainCount  *int64 `json:"remain_count,omitempty" `

    /*
        权益信息展示开始时间，精确到毫秒时间戳     */
    DisplayStartTime  *string `json:"display_start_time,omitempty" `

    /*
        权益信息展示结束时间，精确到毫秒时间戳     */
    DisplayEndTime  *string `json:"display_end_time,omitempty" `

    /*
        权益信息     */
    PromotionList  *[]TaobaoTbkScOptimusPromotionPromotionList `json:"promotion_list,omitempty" `

    /*
        权益扩展信息     */
    PromotionExtend  *TaobaoTbkScOptimusPromotionPromotionExtend `json:"promotion_extend,omitempty" `

    /*
        店铺信息-卖家ID     */
    SellerId  *string `json:"seller_id,omitempty" `

    /*
        店铺信息-卖家昵称     */
    Nick  *string `json:"nick,omitempty" `

    /*
        店铺信息-店铺名称     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        店铺信息-店铺logo     */
    ShopPictureUrl  *string `json:"shop_picture_url,omitempty" `

}

func (s *TaobaoTbkScOptimusPromotionMapData) SetPromotionType(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.PromotionType = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetConditionType(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.ConditionType = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetDiscountType(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.DiscountType = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetTotalCount(v int64) *TaobaoTbkScOptimusPromotionMapData {
    s.TotalCount = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetRemainCount(v int64) *TaobaoTbkScOptimusPromotionMapData {
    s.RemainCount = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetDisplayStartTime(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.DisplayStartTime = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetDisplayEndTime(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.DisplayEndTime = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetPromotionList(v []TaobaoTbkScOptimusPromotionPromotionList) *TaobaoTbkScOptimusPromotionMapData {
    s.PromotionList = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetPromotionExtend(v TaobaoTbkScOptimusPromotionPromotionExtend) *TaobaoTbkScOptimusPromotionMapData {
    s.PromotionExtend = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetSellerId(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetNick(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.Nick = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetShopTitle(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionMapData) SetShopPictureUrl(v string) *TaobaoTbkScOptimusPromotionMapData {
    s.ShopPictureUrl = &v
    return s
}
