package ability2475

import (
	"errors"
	"github.com/Anpengpeng/topsdk"
	"github.com/Anpengpeng/topsdk/ability2475/request"
	"github.com/Anpengpeng/topsdk/ability2475/response"
	"github.com/Anpengpeng/topsdk/util"
	"log"
)

type Ability2475 struct {
	Client *topsdk.TopClient
}

func NewAbility2475(client *topsdk.TopClient) *Ability2475 {
	return &Ability2475{client}
}

/*
   淘宝客-服务商-查询红包发放个数
*/
func (ability *Ability2475) TaobaoTbkScVegasSendReport(req *request.TaobaoTbkScVegasSendReportRequest, session string) (*response.TaobaoTbkScVegasSendReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2475 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.vegas.send.report", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScVegasSendReportResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScVegasSendReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
