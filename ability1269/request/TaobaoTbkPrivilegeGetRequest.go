package request


type TaobaoTbkPrivilegeGetRequest struct {
    /*
        推广位id，mm_xx_xx_xx pid三段式中的第三段     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        淘客商品id defalutValue��-1    */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        1：PC，2：无线，默认：１ defalutValue��1    */
    Platform  *int64 `json:"platform,omitempty" required:"false" `
    /*
        备案的网站id, mm_xx_xx_xx pid三段式中的第二段     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        渠道关系ID，仅适用于渠道推广场景     */
    RelationId  *string `json:"relation_id,omitempty" required:"false" `
    /*
        会员运营ID     */
    SpecialId  *string `json:"special_id,omitempty" required:"false" `
    /*
        淘宝客外部用户标记，如自身系统账户ID；微信ID等     */
    ExternalId  *string `json:"external_id,omitempty" required:"false" `
    /*
        团长与下游渠道合作的特殊标识，用于统计渠道推广效果     */
    Xid  *string `json:"xid,omitempty" required:"false" `
    /*
        会员人群ID，用于统计人群推广效果     */
    UcrowdId  *int64 `json:"ucrowd_id,omitempty" required:"false" `
    /*
        是否获取前N件佣金 ,0-否，1-是,其他值-否 defalutValue��0    */
    GetTopnRate  *int64 `json:"get_topn_rate,omitempty" required:"false" `
    /*
        是否需要获取小程序链接，需要设置1。(暂未对外开放) defalutValue��0    */
    MiniProgramLink  *int64 `json:"mini_program_link,omitempty" required:"false" `
}

func (s *TaobaoTbkPrivilegeGetRequest) SetAdzoneId(v int64) *TaobaoTbkPrivilegeGetRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetItemId(v int64) *TaobaoTbkPrivilegeGetRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetPlatform(v int64) *TaobaoTbkPrivilegeGetRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetSiteId(v int64) *TaobaoTbkPrivilegeGetRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetRelationId(v string) *TaobaoTbkPrivilegeGetRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetSpecialId(v string) *TaobaoTbkPrivilegeGetRequest {
    s.SpecialId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetExternalId(v string) *TaobaoTbkPrivilegeGetRequest {
    s.ExternalId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetXid(v string) *TaobaoTbkPrivilegeGetRequest {
    s.Xid = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetUcrowdId(v int64) *TaobaoTbkPrivilegeGetRequest {
    s.UcrowdId = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetGetTopnRate(v int64) *TaobaoTbkPrivilegeGetRequest {
    s.GetTopnRate = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetRequest) SetMiniProgramLink(v int64) *TaobaoTbkPrivilegeGetRequest {
    s.MiniProgramLink = &v
    return s
}

func (req *TaobaoTbkPrivilegeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.Platform != nil) {
        paramMap["platform"] = *req.Platform
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.SpecialId != nil) {
        paramMap["special_id"] = *req.SpecialId
    }
    if(req.ExternalId != nil) {
        paramMap["external_id"] = *req.ExternalId
    }
    if(req.Xid != nil) {
        paramMap["xid"] = *req.Xid
    }
    if(req.UcrowdId != nil) {
        paramMap["ucrowd_id"] = *req.UcrowdId
    }
    if(req.GetTopnRate != nil) {
        paramMap["get_topn_rate"] = *req.GetTopnRate
    }
    if(req.MiniProgramLink != nil) {
        paramMap["mini_program_link"] = *req.MiniProgramLink
    }
    return paramMap
}

func (req *TaobaoTbkPrivilegeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}