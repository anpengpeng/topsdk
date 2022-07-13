package ability370

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Anpengpeng/topsdk"
	"github.com/Anpengpeng/topsdk/ability370/request"
	"github.com/Anpengpeng/topsdk/ability370/response"
	"github.com/Anpengpeng/topsdk/util"
	"log"
	"reflect"
	"strings"
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
	err = ability.HandleJsonResponseNew(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

func (ability *Ability370) HandleJsonResponseNew(jsonStr string, v interface{}) (err error) {

	if strings.Contains(jsonStr[0:20], "error_response") {
		err := &util.TopApiRequestError{}
		jsonStr = jsonStr[18 : len(jsonStr)-1]
		err2 := json.Unmarshal([]byte(jsonStr), err)
		if err2 != nil {
			return err2
		}
		return err
	}
	var res map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &res)
	for firKey, v := range res {
		if reflect.TypeOf(v).Kind().String() == "slice" {
			newSlice := []interface{}{}
			for _, value := range v.([]interface{}) {
				if reflect.TypeOf(value).Kind().String() == "map" {
					var newMap = make(map[string]interface{})
					for key, value := range value.(map[string]interface{}) {
						s := reflect.TypeOf(value).Kind().String()
						if (key == "item_id" || key == "num_iid") && (s == "int64" || s == "float64") {
							newMap[key] = fmt.Sprintf("%d", value)
						} else {
							newMap[key] = value
						}
					}
					newSlice = append(newSlice, newMap)
				} else {
					newSlice = append(newSlice, value)
				}
			}
			res[firKey] = newSlice
		} else if reflect.TypeOf(v).Kind().String() == "map" {
			var m = make(map[string]interface{})
			for key, value := range v.(map[string]interface{}) {
				s := reflect.TypeOf(value).Kind().String()
				if (key == "item_id" || key == "num_iid") && (s == "int64" || s == "float64") {
					m[key] = fmt.Sprintf("%d", value)
				} else {
					m[key] = value
				}
			}
			res[firKey] = m
		}
	}

	newJsonStr, _ := json.Marshal(res)

	return json.Unmarshal(newJsonStr, v)
}
