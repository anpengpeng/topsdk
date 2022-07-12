package response

import (
	"github.com/Anpengpeng/topsdk/ability2153/domain"
)

type TaobaoTbkScOptimusPromotionResponse struct {

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
	ResultList []domain.TaobaoTbkScOptimusPromotionMapData `json:"result_list,omitempty" `
}
