package domain


type TaobaoTbkScMaterialOptionalMapData struct {
    /*
        优惠券信息-优惠券开始时间     */
    CouponStartTime  *string `json:"coupon_start_time,omitempty" `

    /*
        优惠券信息-优惠券开始时间     */
    CouponEndTime  *string `json:"coupon_end_time,omitempty" `

    /*
        商品信息-是否包含定向计划     */
    InfoDxjh  *string `json:"info_dxjh,omitempty" `

    /*
        商品信息-淘客30天推广量     */
    TkTotalSales  *string `json:"tk_total_sales,omitempty" `

    /*
        商品信息-月支出佣金(该字段废弃，请勿再用)     */
    TkTotalCommi  *string `json:"tk_total_commi,omitempty" `

    /*
        优惠券信息-优惠券id     */
    CouponId  *string `json:"coupon_id,omitempty" `

    /*
        商品信息-宝贝id(该字段废弃，请勿再用)     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品信息-商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        商品信息-商品主图     */
    PictUrl  *string `json:"pict_url,omitempty" `

    /*
        商品信息-商品小图列表     */
    SmallImages  *[]string `json:"small_images,omitempty" `

    /*
        商品信息-一口价通常显示为划线价     */
    ReservePrice  *string `json:"reserve_price,omitempty" `

    /*
        商品信息-在线售卖价(元)。若属于预售商品，付定金时间内，在线售卖价=预售价     */
    ZkFinalPrice  *string `json:"zk_final_price,omitempty" `

    /*
        店铺信息-卖家类型。0表示集市，1表示天猫     */
    UserType  *int32 `json:"user_type,omitempty" `

    /*
        商品信息-宝贝所在地     */
    Provcity  *string `json:"provcity,omitempty" `

    /*
        商品信息-商品地址     */
    ItemUrl  *string `json:"item_url,omitempty" `

    /*
        商品信息-是否包含营销计划     */
    IncludeMkt  *string `json:"include_mkt,omitempty" `

    /*
        商品信息-是否包含定向计划     */
    IncludeDxjh  *string `json:"include_dxjh,omitempty" `

    /*
        商品信息-佣金比率。1550表示15.5%     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        商品信息-30天销量(饿了么卡券信息-总销量）     */
    Volume  *int32 `json:"volume,omitempty" `

    /*
        店铺信息-卖家id     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        店铺信息-店铺名称     */
    ShopTitle  *string `json:"shop_title,omitempty" `

    /*
        优惠券信息-优惠券总量     */
    CouponTotalCount  *int32 `json:"coupon_total_count,omitempty" `

    /*
        优惠券信息-优惠券剩余量     */
    CouponRemainCount  *int32 `json:"coupon_remain_count,omitempty" `

    /*
        优惠券信息-优惠券满减信息     */
    CouponInfo  *string `json:"coupon_info,omitempty" `

    /*
        商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划     */
    CommissionType  *string `json:"commission_type,omitempty" `

    /*
        链接-宝贝推广链接     */
    Url  *string `json:"url,omitempty" `

    /*
        链接-宝贝+券二合一页面链接     */
    CouponShareUrl  *string `json:"coupon_share_url,omitempty" `

    /*
        店铺信息-店铺dsr评分     */
    ShopDsr  *int64 `json:"shop_dsr,omitempty" `

    /*
        商品信息-商品白底图     */
    WhiteImage  *string `json:"white_image,omitempty" `

    /*
        商品信息-商品短标题     */
    ShortTitle  *string `json:"short_title,omitempty" `

    /*
        商品信息-叶子类目id     */
    CategoryId  *int64 `json:"category_id,omitempty" `

    /*
        商品信息-叶子类目名称     */
    CategoryName  *string `json:"category_name,omitempty" `

    /*
        商品信息-一级类目ID     */
    LevelOneCategoryId  *int64 `json:"level_one_category_id,omitempty" `

    /*
        商品信息-一级类目名称     */
    LevelOneCategoryName  *string `json:"level_one_category_name,omitempty" `

    /*
        拼团专用-拼团结束时间     */
    Oetime  *string `json:"oetime,omitempty" `

    /*
        拼团专用-拼团开始时间     */
    Ostime  *string `json:"ostime,omitempty" `

    /*
        拼团专用--拼团几人团     */
    JddNum  *int64 `json:"jdd_num,omitempty" `

    /*
        拼团专用-拼团拼成价，单位元     */
    JddPrice  *string `json:"jdd_price,omitempty" `

    /*
        预售专用-预售数量     */
    UvSumPreSale  *int64 `json:"uv_sum_pre_sale,omitempty" `

    /*
        优惠券信息-优惠券(元)。若属于预售商品，该优惠券付尾款可用，付定金不可用     */
    CouponAmount  *string `json:"coupon_amount,omitempty" `

    /*
        优惠券信息-优惠券起用门槛，满X元可用，如：满299元减20元     */
    CouponStartFee  *string `json:"coupon_start_fee,omitempty" `

    /*
        商品信息-宝贝描述(推荐理由,不一定有)     */
    ItemDescription  *string `json:"item_description,omitempty" `

    /*
        店铺信息-卖家昵称     */
    Nick  *string `json:"nick,omitempty" `

    /*
        链接-物料块id(测试中请勿使用)     */
    XId  *string `json:"x_id,omitempty" `

    /*
        拼团专用-拼团一人价(原价)，单位元     */
    OrigPrice  *string `json:"orig_price,omitempty" `

    /*
        拼团专用-拼团库存数量     */
    TotalStock  *int32 `json:"total_stock,omitempty" `

    /*
        拼团专用-拼团已售数量     */
    SellNum  *int32 `json:"sell_num,omitempty" `

    /*
        拼团专用-拼团剩余库存     */
    Stock  *int32 `json:"stock,omitempty" `

    /*
        营销-天猫营销玩法     */
    TmallPlayActivityInfo  *string `json:"tmall_play_activity_info,omitempty" `

    /*
        商品信息-宝贝id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品信息-商品邮费     */
    RealPostFee  *string `json:"real_post_fee,omitempty" `

    /*
        商品信息-锁佣开始时间     */
    LockRateStartTime  *int64 `json:"lock_rate_start_time,omitempty" `

    /*
        商品信息-锁佣结束时间     */
    LockRateEndTime  *int64 `json:"lock_rate_end_time,omitempty" `

    /*
        商品信息-锁住的佣金率     */
    LockRate  *string `json:"lock_rate,omitempty" `

    /*
        预售商品-优惠信息     */
    PresaleDiscountFeeText  *string `json:"presale_discount_fee_text,omitempty" `

    /*
        预售商品-付尾款结束时间(毫秒)     */
    PresaleTailEndTime  *int64 `json:"presale_tail_end_time,omitempty" `

    /*
        预售商品-付尾款开始时间(毫秒)     */
    PresaleTailStartTime  *int64 `json:"presale_tail_start_time,omitempty" `

    /*
        预售商品-付定金结束时间(毫秒)     */
    PresaleEndTime  *int64 `json:"presale_end_time,omitempty" `

    /*
        预售商品-付定金开始时间(毫秒)     */
    PresaleStartTime  *int64 `json:"presale_start_time,omitempty" `

    /*
        预售商品-定金(元)     */
    PresaleDeposit  *string `json:"presale_deposit,omitempty" `

    /*
        预售有礼-淘礼金发放时间     */
    YsylTljSendTime  *string `json:"ysyl_tlj_send_time,omitempty" `

    /*
        预售有礼-推广链接     */
    YsylClickUrl  *string `json:"ysyl_click_url,omitempty" `

    /*
        预售有礼-佣金比例(预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）     */
    YsylCommissionRate  *string `json:"ysyl_commission_rate,omitempty" `

    /*
        预售有礼-预估淘礼金(元)     */
    YsylTljFace  *string `json:"ysyl_tlj_face,omitempty" `

    /*
        预售有礼-淘礼金使用结束时间     */
    YsylTljUseEndTime  *string `json:"ysyl_tlj_use_end_time,omitempty" `

    /*
        预售有礼-淘礼金使用开始时间     */
    YsylTljUseStartTime  *string `json:"ysyl_tlj_use_start_time,omitempty" `

    /*
        本地化-可用店铺名称     */
    UsableShopName  *string `json:"usable_shop_name,omitempty" `

    /*
        本地化-可用店铺id     */
    UsableShopId  *string `json:"usable_shop_id,omitempty" `

    /*
        本地化-到门店距离(米)     */
    Distance  *string `json:"distance,omitempty" `

    /*
        本地化-销售开始时间     */
    SaleEndTime  *string `json:"sale_end_time,omitempty" `

    /*
        本地化-销售结束时间     */
    SaleBeginTime  *string `json:"sale_begin_time,omitempty" `

    /*
        商品信息-大促活动预热价     */
    SalePrice  *string `json:"sale_price,omitempty" `

    /*
        营销-跨店满减信息     */
    KuadianPromotionInfo  *string `json:"kuadian_promotion_info,omitempty" `

    /*
        是否品牌精选，0不是，1是     */
    SuperiorBrand  *string `json:"superior_brand,omitempty" `

    /*
        比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0     */
    RewardInfo  *int32 `json:"reward_info,omitempty" `

    /*
        是否品牌快抢，0不是，1是     */
    IsBrandFlashSale  *string `json:"is_brand_flash_sale,omitempty" `

    /*
        本地化-扩展信息     */
    LocalizationExtend  *TaobaoTbkScMaterialOptionalLocalizationMapData `json:"localization_extend,omitempty" `

    /*
        物料评估-收益分     */
    CommiScore  *string `json:"commi_score,omitempty" `

    /*
        物料评估-匹配分     */
    MatchScore  *string `json:"match_score,omitempty" `

    /*
        是否是热门商品，0不是，1是     */
    HotFlag  *string `json:"hot_flag,omitempty" `

    /*
        前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
    TopnInfo  *TaobaoTbkScMaterialOptionalTopNInfoDTO `json:"topn_info,omitempty" `

    /*
        百亿补贴信息     */
    BybtInfo  *TaobaoTbkScMaterialOptionalBybtInfoDTO `json:"bybt_info,omitempty" `

    /*
        商品入驻淘特后产生的所有销量量级，不特指某段具体时间     */
    TtSoldCount  *string `json:"tt_sold_count,omitempty" `

    /*
        猫超买返卡信息     */
    MaifanPromotion  *TaobaoTbkScMaterialOptionalMaifanPromotionDTO `json:"maifan_promotion,omitempty" `

    /*
        额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=单单奖励(已失效)，1=超级单单奖励(已失效)，2=年货节单单奖励     */
    CpaRewardType  *string `json:"cpa_reward_type,omitempty" `

    /*
        额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割     */
    CpaRewardAmount  *string `json:"cpa_reward_amount,omitempty" `

    /*
        合作伙伴单单补ID，用作“年货节超级单单补”活动合作伙伴奖励统计依据     */
    ActivityId  *string `json:"activity_id,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponStartTime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponEndTime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetInfoDxjh(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.InfoDxjh = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetTkTotalSales(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.TkTotalSales = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetTkTotalCommi(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.TkTotalCommi = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponId(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetNumIid(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.NumIid = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetTitle(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.Title = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetPictUrl(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.PictUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetSmallImages(v []string) *TaobaoTbkScMaterialOptionalMapData {
    s.SmallImages = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetReservePrice(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.ReservePrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetZkFinalPrice(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.ZkFinalPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetUserType(v int32) *TaobaoTbkScMaterialOptionalMapData {
    s.UserType = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetProvcity(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.Provcity = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetItemUrl(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.ItemUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetIncludeMkt(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.IncludeMkt = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetIncludeDxjh(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.IncludeDxjh = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCommissionRate(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CommissionRate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetVolume(v int32) *TaobaoTbkScMaterialOptionalMapData {
    s.Volume = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetSellerId(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetShopTitle(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.ShopTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponTotalCount(v int32) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponTotalCount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponRemainCount(v int32) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponRemainCount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponInfo(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCommissionType(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CommissionType = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetUrl(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.Url = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponShareUrl(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponShareUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetShopDsr(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.ShopDsr = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetWhiteImage(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.WhiteImage = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetShortTitle(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.ShortTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCategoryId(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.CategoryId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCategoryName(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CategoryName = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetLevelOneCategoryId(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.LevelOneCategoryId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetLevelOneCategoryName(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.LevelOneCategoryName = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetOetime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.Oetime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetOstime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.Ostime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetJddNum(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.JddNum = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetJddPrice(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.JddPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetUvSumPreSale(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.UvSumPreSale = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponAmount(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponAmount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCouponStartFee(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CouponStartFee = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetItemDescription(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.ItemDescription = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetNick(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.Nick = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetXId(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.XId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetOrigPrice(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.OrigPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetTotalStock(v int32) *TaobaoTbkScMaterialOptionalMapData {
    s.TotalStock = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetSellNum(v int32) *TaobaoTbkScMaterialOptionalMapData {
    s.SellNum = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetStock(v int32) *TaobaoTbkScMaterialOptionalMapData {
    s.Stock = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetTmallPlayActivityInfo(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.TmallPlayActivityInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetItemId(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetRealPostFee(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.RealPostFee = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetLockRateStartTime(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.LockRateStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetLockRateEndTime(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.LockRateEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetLockRate(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.LockRate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetPresaleDiscountFeeText(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.PresaleDiscountFeeText = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetPresaleTailEndTime(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.PresaleTailEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetPresaleTailStartTime(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.PresaleTailStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetPresaleEndTime(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.PresaleEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetPresaleStartTime(v int64) *TaobaoTbkScMaterialOptionalMapData {
    s.PresaleStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetPresaleDeposit(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.PresaleDeposit = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetYsylTljSendTime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.YsylTljSendTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetYsylClickUrl(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.YsylClickUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetYsylCommissionRate(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.YsylCommissionRate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetYsylTljFace(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.YsylTljFace = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetYsylTljUseEndTime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.YsylTljUseEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetYsylTljUseStartTime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.YsylTljUseStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetUsableShopName(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.UsableShopName = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetUsableShopId(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.UsableShopId = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetDistance(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.Distance = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetSaleEndTime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.SaleEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetSaleBeginTime(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.SaleBeginTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetSalePrice(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.SalePrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetKuadianPromotionInfo(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.KuadianPromotionInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetSuperiorBrand(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.SuperiorBrand = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetRewardInfo(v int32) *TaobaoTbkScMaterialOptionalMapData {
    s.RewardInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetIsBrandFlashSale(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.IsBrandFlashSale = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetLocalizationExtend(v TaobaoTbkScMaterialOptionalLocalizationMapData) *TaobaoTbkScMaterialOptionalMapData {
    s.LocalizationExtend = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCommiScore(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CommiScore = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetMatchScore(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.MatchScore = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetHotFlag(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.HotFlag = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetTopnInfo(v TaobaoTbkScMaterialOptionalTopNInfoDTO) *TaobaoTbkScMaterialOptionalMapData {
    s.TopnInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetBybtInfo(v TaobaoTbkScMaterialOptionalBybtInfoDTO) *TaobaoTbkScMaterialOptionalMapData {
    s.BybtInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetTtSoldCount(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.TtSoldCount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetMaifanPromotion(v TaobaoTbkScMaterialOptionalMaifanPromotionDTO) *TaobaoTbkScMaterialOptionalMapData {
    s.MaifanPromotion = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCpaRewardType(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CpaRewardType = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetCpaRewardAmount(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.CpaRewardAmount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalMapData) SetActivityId(v string) *TaobaoTbkScMaterialOptionalMapData {
    s.ActivityId = &v
    return s
}
