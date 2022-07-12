package domain


type TaobaoTbkPrivilegeGetMiniProgramDto struct {
    /*
        小程序APPID     */
    MiniProgramAppid  *string `json:"mini_program_appid,omitempty" `

    /*
        小程序路径     */
    MiniProgramPath  *string `json:"mini_program_path,omitempty" `

    /*
        小程序码url地址     */
    MiniProgramQrcodeUrl  *string `json:"mini_program_qrcode_url,omitempty" `

}

func (s *TaobaoTbkPrivilegeGetMiniProgramDto) SetMiniProgramAppid(v string) *TaobaoTbkPrivilegeGetMiniProgramDto {
    s.MiniProgramAppid = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMiniProgramDto) SetMiniProgramPath(v string) *TaobaoTbkPrivilegeGetMiniProgramDto {
    s.MiniProgramPath = &v
    return s
}
func (s *TaobaoTbkPrivilegeGetMiniProgramDto) SetMiniProgramQrcodeUrl(v string) *TaobaoTbkPrivilegeGetMiniProgramDto {
    s.MiniProgramQrcodeUrl = &v
    return s
}
