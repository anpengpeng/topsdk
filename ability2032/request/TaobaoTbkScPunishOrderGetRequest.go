package request

import (
	"github.com/liu8534584/topsdk/ability2032/domain"
	"github.com/liu8534584/topsdk/util"
)

type TaobaoTbkScPunishOrderGetRequest struct {
	/*
	   入参的对象     */
	AfOrderOption *domain.TaobaoTbkScPunishOrderGetTopApiAfOrderOption `json:"af_order_option,omitempty" required:"false" `
}

func (s *TaobaoTbkScPunishOrderGetRequest) SetAfOrderOption(v domain.TaobaoTbkScPunishOrderGetTopApiAfOrderOption) *TaobaoTbkScPunishOrderGetRequest {
	s.AfOrderOption = &v
	return s
}

func (req *TaobaoTbkScPunishOrderGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AfOrderOption != nil {
		paramMap["af_order_option"] = util.ConvertStruct(*req.AfOrderOption)
	}
	return paramMap
}

func (req *TaobaoTbkScPunishOrderGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
