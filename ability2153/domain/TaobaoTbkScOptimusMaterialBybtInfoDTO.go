package domain


type TaobaoTbkScOptimusMaterialBybtInfoDTO struct {
    /*
        百亿补贴品牌logo     */
    BybtBrandLogo  *string `json:"bybt_brand_logo,omitempty" `

    /*
        百亿补贴白底图     */
    BybtPicUrl  *string `json:"bybt_pic_url,omitempty" `

    /*
        百亿补贴商品特征标签，eg.今日发货、晚发补偿、限购一件等     */
    BybtItemTags  *[]string `json:"bybt_item_tags,omitempty" `

    /*
        百亿补贴专属券面额，仅限百亿补贴场景透出     */
    BybtCouponAmount  *string `json:"bybt_coupon_amount,omitempty" `

    /*
        百亿补贴页面实时价     */
    BybtShowPrice  *string `json:"bybt_show_price,omitempty" `

    /*
        全网对比参考价格     */
    BybtLowestPrice  *string `json:"bybt_lowest_price,omitempty" `

    /*
        商品的百亿补贴开始时间     */
    BybtEndTime  *string `json:"bybt_end_time,omitempty" `

    /*
        商品的百亿补贴结束时间     */
    BybtStartTime  *string `json:"bybt_start_time,omitempty" `

}

func (s *TaobaoTbkScOptimusMaterialBybtInfoDTO) SetBybtBrandLogo(v string) *TaobaoTbkScOptimusMaterialBybtInfoDTO {
    s.BybtBrandLogo = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialBybtInfoDTO) SetBybtPicUrl(v string) *TaobaoTbkScOptimusMaterialBybtInfoDTO {
    s.BybtPicUrl = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialBybtInfoDTO) SetBybtItemTags(v []string) *TaobaoTbkScOptimusMaterialBybtInfoDTO {
    s.BybtItemTags = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialBybtInfoDTO) SetBybtCouponAmount(v string) *TaobaoTbkScOptimusMaterialBybtInfoDTO {
    s.BybtCouponAmount = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialBybtInfoDTO) SetBybtShowPrice(v string) *TaobaoTbkScOptimusMaterialBybtInfoDTO {
    s.BybtShowPrice = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialBybtInfoDTO) SetBybtLowestPrice(v string) *TaobaoTbkScOptimusMaterialBybtInfoDTO {
    s.BybtLowestPrice = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialBybtInfoDTO) SetBybtEndTime(v string) *TaobaoTbkScOptimusMaterialBybtInfoDTO {
    s.BybtEndTime = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialBybtInfoDTO) SetBybtStartTime(v string) *TaobaoTbkScOptimusMaterialBybtInfoDTO {
    s.BybtStartTime = &v
    return s
}
