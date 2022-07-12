package ability375

import (
	"errors"
	"github.com/Anpengpeng/topsdk"
	"github.com/Anpengpeng/topsdk/ability375/request"
	"github.com/Anpengpeng/topsdk/ability375/response"
	"github.com/Anpengpeng/topsdk/util"
	"log"
)

type Ability375 struct {
	Client *topsdk.TopClient
}

func NewAbility375(client *topsdk.TopClient) *Ability375 {
	return &Ability375{client}
}

/*
   淘宝客-公用-淘口令生成
*/
func (ability *Ability375) TaobaoTbkTpwdCreate(req *request.TaobaoTbkTpwdCreateRequest) (*response.TaobaoTbkTpwdCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability375 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.tpwd.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkTpwdCreateResponse{}
	if err != nil {
		log.Fatal("taobaoTbkTpwdCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponseNew(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
