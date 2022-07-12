package domain


type TaobaoTbkScOptimusPromotionPromotionExtend struct {
    /*
        权益推荐商品     */
    RecommendItemList  *[]TaobaoTbkScOptimusPromotionRecommendItemList `json:"recommend_item_list,omitempty" `

    /*
        有价券信息     */
    YoujiaCouponInfo  *TaobaoTbkScOptimusPromotionYoujiacouponinfo `json:"youjia_coupon_info,omitempty" `

    /*
        权益链接     */
    PromotionUrl  *string `json:"promotion_url,omitempty" `

}

func (s *TaobaoTbkScOptimusPromotionPromotionExtend) SetRecommendItemList(v []TaobaoTbkScOptimusPromotionRecommendItemList) *TaobaoTbkScOptimusPromotionPromotionExtend {
    s.RecommendItemList = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionPromotionExtend) SetYoujiaCouponInfo(v TaobaoTbkScOptimusPromotionYoujiacouponinfo) *TaobaoTbkScOptimusPromotionPromotionExtend {
    s.YoujiaCouponInfo = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionPromotionExtend) SetPromotionUrl(v string) *TaobaoTbkScOptimusPromotionPromotionExtend {
    s.PromotionUrl = &v
    return s
}
