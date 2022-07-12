package ability2153

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability2153/request"
	"github.com/liu8534584/topsdk/ability2153/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability2153 struct {
	Client *topsdk.TopClient
}

func NewAbility2153(client *topsdk.TopClient) *Ability2153 {
	return &Ability2153{client}
}

/*
   淘宝客-服务商-权益物料精选
*/
func (ability *Ability2153) TaobaoTbkScOptimusPromotion(req *request.TaobaoTbkScOptimusPromotionRequest, session string) (*response.TaobaoTbkScOptimusPromotionResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2153 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.optimus.promotion", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScOptimusPromotionResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScOptimusPromotion error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-服务商-物料精选
*/
func (ability *Ability2153) TaobaoTbkScOptimusMaterial(req *request.TaobaoTbkScOptimusMaterialRequest, session string) (*response.TaobaoTbkScOptimusMaterialResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2153 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.optimus.material", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScOptimusMaterialResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScOptimusMaterial error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
