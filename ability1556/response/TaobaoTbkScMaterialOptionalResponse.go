package response

import (
	"github.com/liu8534584/topsdk/ability1556/domain"
)

type TaobaoTbkScMaterialOptionalResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   搜索到符合条件的结果总数
	*/
	TotalResults int64 `json:"total_results,omitempty" `
	/*
	   resultList
	*/
	ResultList []domain.TaobaoTbkScMaterialOptionalMapData `json:"result_list,omitempty" `
	/*
	   本地化-lbs唯一分页标示，请在翻页时作为入参透传
	*/
	PageResultKey string `json:"page_result_key,omitempty" `
}
