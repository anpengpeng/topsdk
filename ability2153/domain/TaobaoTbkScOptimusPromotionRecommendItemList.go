package domain


type TaobaoTbkScOptimusPromotionRecommendItemList struct {
    /*
        权益推荐商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品链接     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoTbkScOptimusPromotionRecommendItemList) SetItemId(v int64) *TaobaoTbkScOptimusPromotionRecommendItemList {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionRecommendItemList) SetUrl(v string) *TaobaoTbkScOptimusPromotionRecommendItemList {
    s.Url = &v
    return s
}
