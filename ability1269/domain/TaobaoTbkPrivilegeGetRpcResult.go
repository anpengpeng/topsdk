package domain


type TaobaoTbkPrivilegeGetRpcResult struct {
    /*
        data     */
    Data  *TaobaoTbkPrivilegeGetMaterialDto `json:"data,omitempty" `

}

func (s *TaobaoTbkPrivilegeGetRpcResult) SetData(v TaobaoTbkPrivilegeGetMaterialDto) *TaobaoTbkPrivilegeGetRpcResult {
    s.Data = &v
    return s
}
