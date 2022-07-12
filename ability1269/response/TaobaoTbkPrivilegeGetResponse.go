package response

import (
	"github.com/Anpengpeng/topsdk/ability1269/domain"
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
