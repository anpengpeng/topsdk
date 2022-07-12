package domain


type TaobaoTbkScMaterialOptionalMaifanPromotionDTO struct {
    /*
        猫超买返卡活动结束时间     */
    MaifanPromotionEndTime  *string `json:"maifan_promotion_end_time,omitempty" `

    /*
        猫超买返卡活动开始时间     */
    MaifanPromotionStartTime  *string `json:"maifan_promotion_start_time,omitempty" `

    /*
        猫超买返卡面额     */
    MaifanPromotionDiscount  *string `json:"maifan_promotion_discount,omitempty" `

    /*
        猫超买返卡总数，-1代表不限量，其他大于等于0的值为总数     */
    MaifanPromotionCondition  *string `json:"maifan_promotion_condition,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalMaifanPromotionDTO) SetMaifanPromotionEndTime(v string) *TaobaoTbkScMaterialOptionalMaifanPromotionDTO {
    s.MaifanPromotionEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMaifanPromotionDTO) SetMaifanPromotionStartTime(v string) *TaobaoTbkScMaterialOptionalMaifanPromotionDTO {
    s.MaifanPromotionStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMaifanPromotionDTO) SetMaifanPromotionDiscount(v string) *TaobaoTbkScMaterialOptionalMaifanPromotionDTO {
    s.MaifanPromotionDiscount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMaifanPromotionDTO) SetMaifanPromotionCondition(v string) *TaobaoTbkScMaterialOptionalMaifanPromotionDTO {
    s.MaifanPromotionCondition = &v
    return s
}
