package ability370

import (
	"errors"
	"github.com/Anpengpeng/topsdk"
	"github.com/Anpengpeng/topsdk/ability370/request"
	"github.com/Anpengpeng/topsdk/ability370/response"
	"github.com/Anpengpeng/topsdk/util"
	"log"
)

type Ability370 struct {
	Client *topsdk.TopClient
}

func NewAbility370(client *topsdk.TopClient) *Ability370 {
	return &Ability370{client}
}

/*
   淘宝客-推广者-物料搜索
*/
func (ability *Ability370) TaobaoTbkDgMaterialOptional(req *request.TaobaoTbkDgMaterialOptionalRequest) (*response.TaobaoTbkDgMaterialOptionalResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability370 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.material.optional", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgMaterialOptionalResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgMaterialOptional error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponseNew(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
