package request


type TaobaoTbkScOptimusMaterialRequest struct {
    /*
        页大小，默认20，1~100 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        第几页，默认：1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        mm_xxx_xxx_xxx的第3段数字     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        官方提供的物料Id（详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a）     */
    MaterialId  *int64 `json:"material_id" required:"true" `
    /*
        mm_xxx_xxx_xxx的第2段数字     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID     */
    DeviceType  *string `json:"device_type,omitempty" required:"false" `
    /*
        智能匹配-设备号加密类型：MD5，类型为OAID时不传     */
    DeviceEncrypt  *string `json:"device_encrypt,omitempty" required:"false" `
    /*
        智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值     */
    DeviceValue  *string `json:"device_value,omitempty" required:"false" `
    /*
        内容专用-内容详情ID     */
    ContentId  *int64 `json:"content_id,omitempty" required:"false" `
    /*
        内容专用-内容渠道信息     */
    ContentSource  *string `json:"content_source,omitempty" required:"false" `
    /*
        商品ID，用于相似商品推荐     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
}

func (s *TaobaoTbkScOptimusMaterialRequest) SetPageSize(v int64) *TaobaoTbkScOptimusMaterialRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetPageNo(v int64) *TaobaoTbkScOptimusMaterialRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetAdzoneId(v int64) *TaobaoTbkScOptimusMaterialRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetMaterialId(v int64) *TaobaoTbkScOptimusMaterialRequest {
    s.MaterialId = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetSiteId(v int64) *TaobaoTbkScOptimusMaterialRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetDeviceType(v string) *TaobaoTbkScOptimusMaterialRequest {
    s.DeviceType = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetDeviceEncrypt(v string) *TaobaoTbkScOptimusMaterialRequest {
    s.DeviceEncrypt = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetDeviceValue(v string) *TaobaoTbkScOptimusMaterialRequest {
    s.DeviceValue = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetContentId(v int64) *TaobaoTbkScOptimusMaterialRequest {
    s.ContentId = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetContentSource(v string) *TaobaoTbkScOptimusMaterialRequest {
    s.ContentSource = &v
    return s
}
func (s *TaobaoTbkScOptimusMaterialRequest) SetItemId(v int64) *TaobaoTbkScOptimusMaterialRequest {
    s.ItemId = &v
    return s
}

func (req *TaobaoTbkScOptimusMaterialRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.MaterialId != nil) {
        paramMap["material_id"] = *req.MaterialId
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.DeviceType != nil) {
        paramMap["device_type"] = *req.DeviceType
    }
    if(req.DeviceEncrypt != nil) {
        paramMap["device_encrypt"] = *req.DeviceEncrypt
    }
    if(req.DeviceValue != nil) {
        paramMap["device_value"] = *req.DeviceValue
    }
    if(req.ContentId != nil) {
        paramMap["content_id"] = *req.ContentId
    }
    if(req.ContentSource != nil) {
        paramMap["content_source"] = *req.ContentSource
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    return paramMap
}

func (req *TaobaoTbkScOptimusMaterialRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}