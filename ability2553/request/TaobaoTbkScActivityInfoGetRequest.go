package request


type TaobaoTbkScActivityInfoGetRequest struct {
    /*
        官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取     */
    ActivityMaterialId  *string `json:"activity_material_id" required:"true" `
    /*
        mm_xxx_xxx_xxx的第3段数字     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        mm_xxx_xxx_xxx的第2段数字     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        渠道关系id     */
    RelationId  *int64 `json:"relation_id,omitempty" required:"false" `
    /*
        自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道     */
    UnionId  *string `json:"union_id,omitempty" required:"false" `
}

func (s *TaobaoTbkScActivityInfoGetRequest) SetActivityMaterialId(v string) *TaobaoTbkScActivityInfoGetRequest {
    s.ActivityMaterialId = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetRequest) SetAdzoneId(v int64) *TaobaoTbkScActivityInfoGetRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetRequest) SetSiteId(v int64) *TaobaoTbkScActivityInfoGetRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetRequest) SetRelationId(v int64) *TaobaoTbkScActivityInfoGetRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScActivityInfoGetRequest) SetUnionId(v string) *TaobaoTbkScActivityInfoGetRequest {
    s.UnionId = &v
    return s
}

func (req *TaobaoTbkScActivityInfoGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActivityMaterialId != nil) {
        paramMap["activity_material_id"] = *req.ActivityMaterialId
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.UnionId != nil) {
        paramMap["union_id"] = *req.UnionId
    }
    return paramMap
}

func (req *TaobaoTbkScActivityInfoGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}