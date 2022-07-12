package request

import (
	"github.com/liu8534584/topsdk/ability411/domain"
	"github.com/liu8534584/topsdk/util"
)

type TaobaoTbkScRelationRefundRequest struct {
	/*
	   参数option     */
	SearchOption *domain.TaobaoTbkScRelationRefundTopApiRefundRptOption `json:"search_option" required:"true" `
}

func (s *TaobaoTbkScRelationRefundRequest) SetSearchOption(v domain.TaobaoTbkScRelationRefundTopApiRefundRptOption) *TaobaoTbkScRelationRefundRequest {
	s.SearchOption = &v
	return s
}

func (req *TaobaoTbkScRelationRefundRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SearchOption != nil {
		paramMap["search_option"] = util.ConvertStruct(*req.SearchOption)
	}
	return paramMap
}

func (req *TaobaoTbkScRelationRefundRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
