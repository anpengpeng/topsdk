package request


type TaobaoTbkScTpwdConvertRequest struct {
    /*
        需要解析的淘口令     */
    PasswordContent  *string `json:"password_content" required:"true" `
    /*
        广告位ID，mm_xx_xx_xx pid三段式中的第三段     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接     */
    Dx  *string `json:"dx,omitempty" required:"false" `
    /*
        备案的网站id, mm_xx_xx_xx pid三段式中的第二段     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        会员人群ID，用于统计人群推广效果     */
    UcrowdId  *int64 `json:"ucrowd_id,omitempty" required:"false" `
}

func (s *TaobaoTbkScTpwdConvertRequest) SetPasswordContent(v string) *TaobaoTbkScTpwdConvertRequest {
    s.PasswordContent = &v
    return s
}
func (s *TaobaoTbkScTpwdConvertRequest) SetAdzoneId(v int64) *TaobaoTbkScTpwdConvertRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScTpwdConvertRequest) SetDx(v string) *TaobaoTbkScTpwdConvertRequest {
    s.Dx = &v
    return s
}
func (s *TaobaoTbkScTpwdConvertRequest) SetSiteId(v int64) *TaobaoTbkScTpwdConvertRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScTpwdConvertRequest) SetUcrowdId(v int64) *TaobaoTbkScTpwdConvertRequest {
    s.UcrowdId = &v
    return s
}

func (req *TaobaoTbkScTpwdConvertRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PasswordContent != nil) {
        paramMap["password_content"] = *req.PasswordContent
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.Dx != nil) {
        paramMap["dx"] = *req.Dx
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.UcrowdId != nil) {
        paramMap["ucrowd_id"] = *req.UcrowdId
    }
    return paramMap
}

func (req *TaobaoTbkScTpwdConvertRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}