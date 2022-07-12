package response

import (
	"github.com/liu8534584/topsdk/ability2032/domain"
)

type TaobaoTbkScPunishOrderGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   查询的对象
	*/
	Result domain.TaobaoTbkScPunishOrderGetRpcResult `json:"result,omitempty" `
}
