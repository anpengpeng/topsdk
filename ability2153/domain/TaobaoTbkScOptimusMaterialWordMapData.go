package domain


type TaobaoTbkScOptimusMaterialWordMapData struct {
    /*
        商品信息-商品相关的关联词     */
    Word  *string `json:"word,omitempty" `

    /*
        链接-商品相关关联词落地页地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoTbkScOptimusMaterialWordMapData) SetWord(v string) *TaobaoTbkScOptimusMaterialWordMapData {
    s.Word = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialWordMapData) SetUrl(v string) *TaobaoTbkScOptimusMaterialWordMapData {
    s.Url = &v
    return s
}
