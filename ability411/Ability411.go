package ability411

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability411/request"
	"github.com/liu8534584/topsdk/ability411/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability411 struct {
	Client *topsdk.TopClient
}

func NewAbility411(client *topsdk.TopClient) *Ability411 {
	return &Ability411{client}
}

/*
   淘宝客-服务商-维权退款订单查询
*/
func (ability *Ability411) TaobaoTbkScRelationRefund(req *request.TaobaoTbkScRelationRefundRequest, session string) (*response.TaobaoTbkScRelationRefundResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability411 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.relation.refund", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScRelationRefundResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScRelationRefund error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-服务商-所有订单查询
*/
func (ability *Ability411) TaobaoTbkScOrderDetailsGet(req *request.TaobaoTbkScOrderDetailsGetRequest, session string) (*response.TaobaoTbkScOrderDetailsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability411 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.order.details.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScOrderDetailsGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScOrderDetailsGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
