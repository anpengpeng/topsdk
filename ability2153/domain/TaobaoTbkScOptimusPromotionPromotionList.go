package domain


type TaobaoTbkScOptimusPromotionPromotionList struct {
    /*
        权益开始时间，精确到毫秒时间戳     */
    EntryUsedStartTime  *string `json:"entry_used_start_time,omitempty" `

    /*
        权益结束时间，精确到毫秒时间戳     */
    EntryUsedEndTime  *string `json:"entry_used_end_time,omitempty" `

    /*
        权益起用门槛，满X元可用，券场景为满元，精确到分，如满100元可用     */
    EntryCondition  *string `json:"entry_condition,omitempty" `

    /*
        权益面额，券场景为减钱，精确到分     */
    EntryDiscount  *string `json:"entry_discount,omitempty" `

}

func (s *TaobaoTbkScOptimusPromotionPromotionList) SetEntryUsedStartTime(v string) *TaobaoTbkScOptimusPromotionPromotionList {
    s.EntryUsedStartTime = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionPromotionList) SetEntryUsedEndTime(v string) *TaobaoTbkScOptimusPromotionPromotionList {
    s.EntryUsedEndTime = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionPromotionList) SetEntryCondition(v string) *TaobaoTbkScOptimusPromotionPromotionList {
    s.EntryCondition = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionPromotionList) SetEntryDiscount(v string) *TaobaoTbkScOptimusPromotionPromotionList {
    s.EntryDiscount = &v
    return s
}
