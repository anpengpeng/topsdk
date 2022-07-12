package domain


type TaobaoTbkScMaterialOptionalLocalizationMapData struct {
    /*
        本地化-配送时间     */
    OrderLeadTime  *string `json:"order_lead_time,omitempty" `

    /*
        本地化-用户评分     */
    UserRating  *string `json:"user_rating,omitempty" `

    /*
        本地化-起送价     */
    DeliveryMinPrice  *string `json:"delivery_min_price,omitempty" `

    /*
        本地化-配送费     */
    DeliveryFee  *string `json:"delivery_fee,omitempty" `

    /*
        本地化-配送费原价     */
    OriginalDeliveryFee  *string `json:"original_delivery_fee,omitempty" `

    /*
        本地化-配送类型；0：蜂鸟专送 1：蜂鸟快送 2：商家自配 3: 全城送     */
    DeliveryType  *string `json:"delivery_type,omitempty" `

    /*
        本地化-推荐理由     */
    RecommendReasons  *[]string `json:"recommend_reasons,omitempty" `

    /*
        本地化-营销标签     */
    SaleTags  *[]string `json:"sale_tags,omitempty" `

    /*
        本地化-单店商品列表     */
    FoodItemList  *[]TaobaoTbkScMaterialOptionalFoodMapData `json:"food_item_list,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetOrderLeadTime(v string) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.OrderLeadTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetUserRating(v string) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.UserRating = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetDeliveryMinPrice(v string) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.DeliveryMinPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetDeliveryFee(v string) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.DeliveryFee = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetOriginalDeliveryFee(v string) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.OriginalDeliveryFee = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetDeliveryType(v string) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.DeliveryType = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetRecommendReasons(v []string) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.RecommendReasons = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetSaleTags(v []string) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.SaleTags = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalLocalizationMapData) SetFoodItemList(v []TaobaoTbkScMaterialOptionalFoodMapData) *TaobaoTbkScMaterialOptionalLocalizationMapData {
    s.FoodItemList = &v
    return s
}
