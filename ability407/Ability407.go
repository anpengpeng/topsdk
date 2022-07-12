package ability407

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability407/request"
	"github.com/liu8534584/topsdk/ability407/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability407 struct {
	Client *topsdk.TopClient
}

func NewAbility407(client *topsdk.TopClient) *Ability407 {
	return &Ability407{client}
}

/*
   淘宝客-服务商-淘口令解析&转链
*/
func (ability *Ability407) TaobaoTbkScTpwdConvert(req *request.TaobaoTbkScTpwdConvertRequest, session string) (*response.TaobaoTbkScTpwdConvertResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability407 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.tpwd.convert", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScTpwdConvertResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScTpwdConvert error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
