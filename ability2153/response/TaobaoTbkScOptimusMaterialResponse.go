package response

import (
	"github.com/liu8534584/topsdk/ability2153/domain"
)

type TaobaoTbkScOptimusMaterialResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   resultList
	*/
	ResultList []domain.TaobaoTbkScOptimusMaterialMapData `json:"result_list,omitempty" `
	/*
	   推荐信息-是否抄底
	*/
	IsDefault string `json:"is_default,omitempty" `
}
