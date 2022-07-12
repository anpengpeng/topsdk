package ability2032

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability2032/request"
	"github.com/liu8534584/topsdk/ability2032/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability2032 struct {
	Client *topsdk.TopClient
}

func NewAbility2032(client *topsdk.TopClient) *Ability2032 {
	return &Ability2032{client}
}

/*
   淘宝客-服务商-处罚订单查询
*/
func (ability *Ability2032) TaobaoTbkScPunishOrderGet(req *request.TaobaoTbkScPunishOrderGetRequest, session string) (*response.TaobaoTbkScPunishOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2032 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.punish.order.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScPunishOrderGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScPunishOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
