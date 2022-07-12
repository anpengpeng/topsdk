package domain


type TaobaoTbkScMaterialOptionalFoodMapData struct {
    /*
        本地化-商品图片     */
    FoodPic  *string `json:"food_pic,omitempty" `

    /*
        本地化-商品标题     */
    FoodTitle  *string `json:"food_title,omitempty" `

    /*
        本地化-商品促销价     */
    FoodPromotionPrice  *string `json:"food_promotion_price,omitempty" `

    /*
        本地化-商品原价     */
    FoodReservePrice  *string `json:"food_reserve_price,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalFoodMapData) SetFoodPic(v string) *TaobaoTbkScMaterialOptionalFoodMapData {
    s.FoodPic = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalFoodMapData) SetFoodTitle(v string) *TaobaoTbkScMaterialOptionalFoodMapData {
    s.FoodTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalFoodMapData) SetFoodPromotionPrice(v string) *TaobaoTbkScMaterialOptionalFoodMapData {
    s.FoodPromotionPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalFoodMapData) SetFoodReservePrice(v string) *TaobaoTbkScMaterialOptionalFoodMapData {
    s.FoodReservePrice = &v
    return s
}
