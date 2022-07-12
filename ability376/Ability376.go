package ability376

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability376/request"
	"github.com/liu8534584/topsdk/ability376/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability376 struct {
	Client *topsdk.TopClient
}

func NewAbility376(client *topsdk.TopClient) *Ability376 {
	return &Ability376{client}
}

/*
   淘宝客-公用-长链转短链
*/
func (ability *Ability376) TaobaoTbkSpreadGet(req *request.TaobaoTbkSpreadGetRequest) (*response.TaobaoTbkSpreadGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability376 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.spread.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkSpreadGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkSpreadGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
