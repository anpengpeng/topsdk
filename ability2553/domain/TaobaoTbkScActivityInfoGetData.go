package domain


type TaobaoTbkScActivityInfoGetData struct {
    /*
        【本地化】微信推广二维码地址     */
    WxQrcodeUrl  *string `json:"wx_qrcode_url,omitempty" `

    /*
        推广长链接     */
    ClickUrl  *string `json:"click_url,omitempty" `

    /*
        推广短链接     */
    ShortClickUrl  *string `json:"short_click_url,omitempty" `

    /*
        投放平台：1-PC，2-无线     */
    TerminalType  *string `json:"terminal_type,omitempty" `

    /*
        物料素材下载地址     */
    MaterialOssUrl  *string `json:"material_oss_url,omitempty" `

    /*
        会场名称     */
    PageName  *string `json:"page_name,omitempty" `

    /*
        活动开始时间     */
    PageStartTime  *string `json:"page_start_time,omitempty" `

    /*
        活动结束时间     */
    PageEndTime  *string `json:"page_end_time,omitempty" `

    /*
        【本地化】微信小程序推广地址     */
    WxMiniprogramPath  *string `json:"wx_miniprogram_path,omitempty" `

}

func (s *TaobaoTbkScActivityInfoGetData) SetWxQrcodeUrl(v string) *TaobaoTbkScActivityInfoGetData {
    s.WxQrcodeUrl = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetData) SetClickUrl(v string) *TaobaoTbkScActivityInfoGetData {
    s.ClickUrl = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetData) SetShortClickUrl(v string) *TaobaoTbkScActivityInfoGetData {
    s.ShortClickUrl = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetData) SetTerminalType(v string) *TaobaoTbkScActivityInfoGetData {
    s.TerminalType = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetData) SetMaterialOssUrl(v string) *TaobaoTbkScActivityInfoGetData {
    s.MaterialOssUrl = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetData) SetPageName(v string) *TaobaoTbkScActivityInfoGetData {
    s.PageName = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetData) SetPageStartTime(v string) *TaobaoTbkScActivityInfoGetData {
    s.PageStartTime = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetData) SetPageEndTime(v string) *TaobaoTbkScActivityInfoGetData {
    s.PageEndTime = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetData) SetWxMiniprogramPath(v string) *TaobaoTbkScActivityInfoGetData {
    s.WxMiniprogramPath = &v
    return s
}
