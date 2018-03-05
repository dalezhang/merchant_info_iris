package sub_merchant

import (
	"fmt"

	"github.com/dalezhang/merchant_info_iris/controller/common"
	"github.com/dalezhang/merchant_info_iris/model"
	"github.com/dalezhang/merchant_info_iris/utils"
	"github.com/kataras/iris"
	//	"github.com/kataras/iris/context"

	"github.com/dalezhang/merchant_info_iris/config"
	"gopkg.in/mgo.v2/bson"
)

type SubMerchantParams struct {
	PlatformID string `json:"platform_id"`
	Sign       string `json:"sign"`
	Method     string `json:"method"`
	model.SubMerchant
}

var merchant model.Merchant
var subMerchant model.SubMerchant

func Create(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON
	var params SubMerchantParams
	if err := ctx.ReadJSON(&params); err != nil {
		fmt.Println(err.Error())
		SendErrJSON("参数无效", ctx)
		return
	}
	if err := md5Decode(params); err != "" {
		SendErrJSON(err, ctx)
		return
	}
	switch params.Method {
	case "sub_merchant.create":
		var merchant1 model.Merchant
		var subMerchant1 model.SubMerchant
		partnerMchID := params.PartnerMchID
		session := model.CloneSession() //调用这个获得session
		defer session.Close()           //一定要记得释放
		collection := session.DB(config.DBConfig.Database).C("application_records")

		collection.Find(bson.M{"partner_mch_id": partnerMchID}).One(&merchant1)
		collection.Find(bson.M{"partner_mch_id": partnerMchID}).One(&subMerchant1)
		if merchant1.ID != "" || subMerchant1.ID != "" {
			SendErrJSON("该partner_mch_id已经存在，请修改后再尝试进件", ctx)
			return
		}
		subMerchant = params.SubMerchant
		subMerchant.MerchantID = merchant.ID
		fmt.Println("merchant_id: ", subMerchant.MerchantID)
		subMerchant.ID = bson.NewObjectId()
		fmt.Println("id: ", subMerchant.ID)
		saveErr := collection.Insert(&subMerchant)
		if saveErr == nil {
			ctx.JSON(iris.Map{
				"errNo": model.ErrorCode.SUCCESS,
				"msg":   "success",
				"data": iris.Map{
					"subMerchant": subMerchant,
				},
			})
			return
		} else {
			SendErrJSON(saveErr.Error(), ctx)
			return
		}
	default:
		SendErrJSON(`invalid method, should be one of ["merchant.create","merchant.update","merchant.query"]`, ctx)
		return
	}

}

func prepareMerchant(platformID string) (model.Merchant, string) {
	if platformID == "" {
		return model.Merchant{}, "platform_id为空"
	}
	session := model.CloneSession() //调用这个获得session
	defer session.Close()           //一定要记得释放
	collection := session.DB(config.DBConfig.Database).C("application_records")

	collection.Find(bson.M{"merchant_id": platformID}).One(&merchant)
	if merchant.ID == "" {
		return model.Merchant{}, "找不到商户信息，platform_id无效。"
	}
	return merchant, ""
}

func md5Decode(params SubMerchantParams) string {
	merchant, err := prepareMerchant(params.PlatformID)
	if err != "" {
		return err
	}
	shareKey := merchant.ShareKey
	fmt.Println("shareKey: ", shareKey)
	js, err1 := utils.MapStruct(&params)
	if err1 != nil {
		return err1.Error()
	}
	sign := utils.GetMac(js, shareKey)

	if sign == params.Sign {
		return ""
	}

	return "签名不匹配"
}
