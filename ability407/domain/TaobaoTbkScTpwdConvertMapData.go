package domain


type TaobaoTbkScTpwdConvertMapData struct {
    /*
        商品Id     */
    NumIid  *string `json:"num_iid,omitempty" `

    /*
        商品淘客转链链接     */
    ClickUrl  *string `json:"click_url,omitempty" `

    /*
        店铺卖家ID     */
    SellerId  *string `json:"seller_id,omitempty" `

    /*
        入参淘口令对应原始链接     */
    OriginUrl  *string `json:"origin_url,omitempty" `

    /*
        入参淘口令推广链接中的pid，如果不属于当前调用的推广者则展示“0”     */
    OriginPid  *string `json:"origin_pid,omitempty" `

}

func (s *TaobaoTbkScTpwdConvertMapData) SetNumIid(v string) *TaobaoTbkScTpwdConvertMapData {
    s.NumIid = &v
    return s
}
func (s *TaobaoTbkScTpwdConvertMapData) SetClickUrl(v string) *TaobaoTbkScTpwdConvertMapData {
    s.ClickUrl = &v
    return s
}
func (s *TaobaoTbkScTpwdConvertMapData) SetSellerId(v string) *TaobaoTbkScTpwdConvertMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkScTpwdConvertMapData) SetOriginUrl(v string) *TaobaoTbkScTpwdConvertMapData {
    s.OriginUrl = &v
    return s
}
func (s *TaobaoTbkScTpwdConvertMapData) SetOriginPid(v string) *TaobaoTbkScTpwdConvertMapData {
    s.OriginPid = &v
    return s
}
