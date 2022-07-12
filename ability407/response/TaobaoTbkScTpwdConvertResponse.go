package response

import (
	"github.com/liu8534584/topsdk/ability407/domain"
)

type TaobaoTbkScTpwdConvertResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   data
	*/
	Data domain.TaobaoTbkScTpwdConvertMapData `json:"data,omitempty" `
}
