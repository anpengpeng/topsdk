package domain


type TaobaoTbkScOrderDetailsGetPublisherOrderDto struct {
    /*
        订单在淘宝拍下付款的时间     */
    TbPaidTime  *string `json:"tb_paid_time,omitempty" `

    /*
        订单付款的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间     */
    TkPaidTime  *string `json:"tk_paid_time,omitempty" `

    /*
        买家确认收货的付款金额（不包含运费金额）     */
    PayPrice  *string `json:"pay_price,omitempty" `

    /*
        结算预估收入=结算金额*提成。以买家确认收货的付款金额为基数，预估您可能获得的收入。因买家退款、您违规推广等原因，可能与您最终收入不一致。最终收入以月结后您实际收到的为准     */
    PubShareFee  *string `json:"pub_share_fee,omitempty" `

    /*
        买家通过购物车购买的每个商品对应的订单编号，此订单编号并未在淘宝买家后台透出     */
    TradeId  *string `json:"trade_id,omitempty" `

    /*
        二方：佣金收益的第一归属者； 三方：从其他淘宝客佣金中进行分成的推广者     */
    TkOrderRole  *int64 `json:"tk_order_role,omitempty" `

    /*
        订单确认收货后且商家完成佣金支付的时间     */
    TkEarningTime  *string `json:"tk_earning_time,omitempty" `

    /*
        推广位管理下的推广位名称对应的ID，同时也是pid=mm_1_2_3中的“3”这段数字     */
    AdzoneId  *int64 `json:"adzone_id,omitempty" `

    /*
        从结算佣金中分得的收益比率     */
    PubShareRate  *string `json:"pub_share_rate,omitempty" `

    /*
        维权标签，0 含义为非维权 1 含义为维权订单     */
    RefundTag  *int64 `json:"refund_tag,omitempty" `

    /*
        平台给与的补贴比率，如天猫、淘宝、聚划算等     */
    SubsidyRate  *string `json:"subsidy_rate,omitempty" `

    /*
        提成=收入比率*分成比率。指实际获得收益的比率     */
    TkTotalRate  *string `json:"tk_total_rate,omitempty" `

    /*
        商品所属的根类目，即一级类目的名称     */
    ItemCategoryName  *string `json:"item_category_name,omitempty" `

    /*
        掌柜旺旺     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        推广者的会员id     */
    PubId  *int64 `json:"pub_id,omitempty" `

    /*
        推广者赚取佣金后支付给阿里妈妈的技术服务费用的比率     */
    AlimamaRate  *string `json:"alimama_rate,omitempty" `

    /*
        平台出资方，如天猫、淘宝、或聚划算等     */
    SubsidyType  *string `json:"subsidy_type,omitempty" `

    /*
        商品图片     */
    ItemImg  *string `json:"item_img,omitempty" `

    /*
        付款预估收入=付款金额*提成。指买家付款金额为基数，预估您可能获得的收入。因买家退款等原因，可能与结算预估收入不一致     */
    PubSharePreFee  *string `json:"pub_share_pre_fee,omitempty" `

    /*
        买家拍下付款的金额（不包含运费金额）     */
    AlipayTotalPrice  *string `json:"alipay_total_price,omitempty" `

    /*
        商品标题     */
    ItemTitle  *string `json:"item_title,omitempty" `

    /*
        媒体管理下的对应ID的自定义名称     */
    SiteName  *string `json:"site_name,omitempty" `

    /*
        商品数量     */
    ItemNum  *int64 `json:"item_num,omitempty" `

    /*
        补贴金额=结算金额*补贴比率     */
    SubsidyFee  *string `json:"subsidy_fee,omitempty" `

    /*
        技术服务费=结算金额*收入比率*技术服务费率。推广者赚取佣金后支付给阿里妈妈的技术服务费用     */
    AlimamaShareFee  *string `json:"alimama_share_fee,omitempty" `

    /*
        买家在淘宝后台显示的订单编号     */
    TradeParentId  *string `json:"trade_parent_id,omitempty" `

    /*
        订单所属平台类型，包括天猫、淘宝、聚划算等     */
    OrderType  *string `json:"order_type,omitempty" `

    /*
        订单创建的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间     */
    TkCreateTime  *string `json:"tk_create_time,omitempty" `

    /*
        产品类型     */
    FlowSource  *string `json:"flow_source,omitempty" `

    /*
        成交平台     */
    TerminalType  *string `json:"terminal_type,omitempty" `

    /*
        通过推广链接达到商品、店铺详情页的点击时间     */
    ClickTime  *string `json:"click_time,omitempty" `

    /*
        已付款：指订单已付款，但还未确认收货 已收货：指订单已确认收货，但商家佣金未支付 已结算：指订单已确认收货，且商家佣金已支付成功 已失效：指订单关闭/订单佣金小于0.01元，订单关闭主要有：1）买家超时未付款； 2）买家付款前，买家/卖家取消了订单；3）订单付款后发起售中退款成功；3：订单结算，12：订单付款， 13：订单失效，14：订单成功     */
    TkStatus  *int64 `json:"tk_status,omitempty" `

    /*
        商品单价     */
    ItemPrice  *string `json:"item_price,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        推广位管理下的自定义推广位名称     */
    AdzoneName  *string `json:"adzone_name,omitempty" `

    /*
        佣金比率     */
    TotalCommissionRate  *string `json:"total_commission_rate,omitempty" `

    /*
        商品链接     */
    ItemLink  *string `json:"item_link,omitempty" `

    /*
        媒体管理下的ID，同时也是pid=mm_1_2_3中的“2”这段数字     */
    SiteId  *int64 `json:"site_id,omitempty" `

    /*
        店铺名称     */
    SellerShopTitle  *string `json:"seller_shop_title,omitempty" `

    /*
        订单结算的佣金比率+平台的补贴比率     */
    IncomeRate  *string `json:"income_rate,omitempty" `

    /*
        佣金金额=结算金额＊佣金比率     */
    TotalCommissionFee  *string `json:"total_commission_fee,omitempty" `

    /*
        预估内容专项服务费：内容场景专项技术服务费，内容推广者在内容场景进行推广需要支付给阿里妈妈专项的技术服务费用。专项服务费＝付款金额＊专项服务费率。     */
    TkCommissionPreFeeForMediaPlatform  *string `json:"tk_commission_pre_fee_for_media_platform,omitempty" `

    /*
        结算内容专项服务费：内容场景专项技术服务费，内容推广者在内容场景进行推广需要支付给阿里妈妈专项的技术服务费用。专项服务费＝结算金额＊专项服务费率。     */
    TkCommissionFeeForMediaPlatform  *string `json:"tk_commission_fee_for_media_platform,omitempty" `

    /*
        内容专项服务费率：内容场景专项技术服务费率，内容推广者在内容场景进行推广需要按结算金额支付一定比例给阿里妈妈作为内容场景专项技术服务费，用于提供与内容平台实现产品技术对接等服务。     */
    TkCommissionRateForMediaPlatform  *string `json:"tk_commission_rate_for_media_platform,omitempty" `

    /*
        会员运营id     */
    SpecialId  *int64 `json:"special_id,omitempty" `

    /*
        渠道关系id     */
    RelationId  *int64 `json:"relation_id,omitempty" `

    /*
        预售时期，用户对预售商品支付定金的付款时间，可能略晚于在淘宝付定金时间     */
    TkDepositTime  *string `json:"tk_deposit_time,omitempty" `

    /*
        预售时期，用户对预售商品支付定金的付款时间     */
    TbDepositTime  *string `json:"tb_deposit_time,omitempty" `

    /*
        预售时期，用户对预售商品支付的定金金额     */
    DepositPrice  *string `json:"deposit_price,omitempty" `

    /*
        开发者调用api的appkey     */
    AppKey  *string `json:"app_key,omitempty" `

    /*
        口碑子订单号     */
    AlscId  *string `json:"alsc_id,omitempty" `

    /*
        口碑父订单号     */
    AlscPid  *string `json:"alsc_pid,omitempty" `

    /*
        服务费信息     */
    ServiceFeeDtoList  *[]TaobaoTbkScOrderDetailsGetServiceFeeDto `json:"service_fee_dto_list,omitempty" `

    /*
        激励池对应的rid     */
    LxRid  *string `json:"lx_rid,omitempty" `

    /*
        订单是否为激励池订单 1，表征是 0，表征否     */
    IsLx  *string `json:"is_lx,omitempty" `

    /*
        营销类型：该字段中视订单情况有单个或多个值。 例如：淘礼金（自助充值），特价版客户端染色，特价版客户端锁粉，特价版客户端推广。     */
    MarketingType  *string `json:"marketing_type,omitempty" `

    /*
        订单更新时间     */
    ModifiedTime  *string `json:"modified_time,omitempty" `

}

func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTbPaidTime(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TbPaidTime = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkPaidTime(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkPaidTime = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetPayPrice(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.PayPrice = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetPubShareFee(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.PubShareFee = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTradeId(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TradeId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkOrderRole(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkOrderRole = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkEarningTime(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkEarningTime = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetAdzoneId(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetPubShareRate(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.PubShareRate = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetRefundTag(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.RefundTag = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetSubsidyRate(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.SubsidyRate = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkTotalRate(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkTotalRate = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetItemCategoryName(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ItemCategoryName = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetSellerNick(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetPubId(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.PubId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetAlimamaRate(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.AlimamaRate = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetSubsidyType(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.SubsidyType = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetItemImg(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ItemImg = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetPubSharePreFee(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.PubSharePreFee = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetAlipayTotalPrice(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.AlipayTotalPrice = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetItemTitle(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ItemTitle = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetSiteName(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.SiteName = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetItemNum(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ItemNum = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetSubsidyFee(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.SubsidyFee = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetAlimamaShareFee(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.AlimamaShareFee = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTradeParentId(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TradeParentId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetOrderType(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.OrderType = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkCreateTime(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkCreateTime = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetFlowSource(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.FlowSource = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTerminalType(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TerminalType = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetClickTime(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ClickTime = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkStatus(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkStatus = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetItemPrice(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetItemId(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetAdzoneName(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.AdzoneName = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTotalCommissionRate(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TotalCommissionRate = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetItemLink(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ItemLink = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetSiteId(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetSellerShopTitle(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.SellerShopTitle = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetIncomeRate(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.IncomeRate = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTotalCommissionFee(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TotalCommissionFee = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkCommissionPreFeeForMediaPlatform(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkCommissionPreFeeForMediaPlatform = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkCommissionFeeForMediaPlatform(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkCommissionFeeForMediaPlatform = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkCommissionRateForMediaPlatform(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkCommissionRateForMediaPlatform = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetSpecialId(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.SpecialId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetRelationId(v int64) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTkDepositTime(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TkDepositTime = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetTbDepositTime(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.TbDepositTime = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetDepositPrice(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.DepositPrice = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetAppKey(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.AppKey = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetAlscId(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.AlscId = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetAlscPid(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.AlscPid = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetServiceFeeDtoList(v []TaobaoTbkScOrderDetailsGetServiceFeeDto) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ServiceFeeDtoList = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetLxRid(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.LxRid = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetIsLx(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.IsLx = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetMarketingType(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.MarketingType = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsGetPublisherOrderDto) SetModifiedTime(v string) *TaobaoTbkScOrderDetailsGetPublisherOrderDto {
    s.ModifiedTime = &v
    return s
}
