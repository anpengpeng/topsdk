package response

import (
	"github.com/liu8534584/topsdk/ability1269/domain"
)

type TaobaoTbkPrivilegeGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   result
	*/
	Result domain.TaobaoTbkPrivilegeGetRpcResult `json:"result,omitempty" `
}
