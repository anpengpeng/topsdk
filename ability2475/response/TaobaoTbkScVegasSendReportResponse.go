package response

import (
	"github.com/liu8534584/topsdk/ability2475/domain"
)

type TaobaoTbkScVegasSendReportResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回model
	*/
	Result domain.TaobaoTbkScVegasSendReportResult `json:"result,omitempty" `
}
