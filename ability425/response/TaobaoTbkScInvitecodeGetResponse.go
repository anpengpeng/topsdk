package response

import (
	"github.com/liu8534584/topsdk/ability425/domain"
)

type TaobaoTbkScInvitecodeGetResponse struct {

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
	Data domain.TaobaoTbkScInvitecodeGetData `json:"data,omitempty" `
}
