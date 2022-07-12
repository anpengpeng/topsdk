package ability2553

import (
	"errors"
	"github.com/Anpengpeng/topsdk"
	"github.com/Anpengpeng/topsdk/ability2553/request"
	"github.com/Anpengpeng/topsdk/ability2553/response"
	"github.com/Anpengpeng/topsdk/util"
	"log"
)

type Ability2553 struct {
	Client *topsdk.TopClient
}

func NewAbility2553(client *topsdk.TopClient) *Ability2553 {
	return &Ability2553{client}
}

/*
   淘宝客-服务商-官方活动转链
*/
func (ability *Ability2553) TaobaoTbkScActivityInfoGet(req *request.TaobaoTbkScActivityInfoGetRequest, session string) (*response.TaobaoTbkScActivityInfoGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2553 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.activity.info.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScActivityInfoGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScActivityInfoGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
