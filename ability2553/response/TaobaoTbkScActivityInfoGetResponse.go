package response

import (
	"github.com/liu8534584/topsdk/ability2553/domain"
)

type TaobaoTbkScActivityInfoGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果对象
	*/
	Data domain.TaobaoTbkScActivityInfoGetData `json:"data,omitempty" `
}
