package request


type TaobaoTbkScOptimusPromotionRequest struct {
    /*
        页大小，每次请求限制10以内 defalutValue��10    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        第几页，默认：1 defalutValue��1    */
    PageNum  *int64 `json:"page_num,omitempty" required:"false" `
    /*
        mm_xxx_xxx_xxx的第3段数字     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        官方提供的权益物料Id。有价券-37104、大额店铺券-37116，更多权益物料id敬请期待！     */
    PromotionId  *int64 `json:"promotion_id" required:"true" `
    /*
        mm_xxx_xxx_xxx的第2段数字     */
    SiteId  *int64 `json:"site_id" required:"true" `
}

func (s *TaobaoTbkScOptimusPromotionRequest) SetPageSize(v int64) *TaobaoTbkScOptimusPromotionRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionRequest) SetPageNum(v int64) *TaobaoTbkScOptimusPromotionRequest {
    s.PageNum = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionRequest) SetAdzoneId(v int64) *TaobaoTbkScOptimusPromotionRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionRequest) SetPromotionId(v int64) *TaobaoTbkScOptimusPromotionRequest {
    s.PromotionId = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionRequest) SetSiteId(v int64) *TaobaoTbkScOptimusPromotionRequest {
    s.SiteId = &v
    return s
}

func (req *TaobaoTbkScOptimusPromotionRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNum != nil) {
        paramMap["page_num"] = *req.PageNum
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.PromotionId != nil) {
        paramMap["promotion_id"] = *req.PromotionId
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    return paramMap
}

func (req *TaobaoTbkScOptimusPromotionRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}