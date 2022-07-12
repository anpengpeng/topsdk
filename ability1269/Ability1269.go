package ability1269

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability1269/request"
	"github.com/liu8534584/topsdk/ability1269/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability1269 struct {
	Client *topsdk.TopClient
}

func NewAbility1269(client *topsdk.TopClient) *Ability1269 {
	return &Ability1269{client}
}

/*
   淘宝客-服务商-单品券高效转链
*/
func (ability *Ability1269) TaobaoTbkPrivilegeGet(req *request.TaobaoTbkPrivilegeGetRequest, session string) (*response.TaobaoTbkPrivilegeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1269 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.privilege.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkPrivilegeGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkPrivilegeGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
