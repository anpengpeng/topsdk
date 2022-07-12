package ability1556

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability1556/request"
	"github.com/liu8534584/topsdk/ability1556/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability1556 struct {
	Client *topsdk.TopClient
}

func NewAbility1556(client *topsdk.TopClient) *Ability1556 {
	return &Ability1556{client}
}

/*
   淘宝客-服务商-物料搜索
*/
func (ability *Ability1556) TaobaoTbkScMaterialOptional(req *request.TaobaoTbkScMaterialOptionalRequest, session string) (*response.TaobaoTbkScMaterialOptionalResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1556 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.material.optional", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScMaterialOptionalResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScMaterialOptional error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
