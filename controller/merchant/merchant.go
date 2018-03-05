package merchant

import (
	//"github.com/dalezhang/merchant_info_iris/controller/common"
	"fmt"

	"github.com/dalezhang/merchant_info_iris/config"
	"github.com/dalezhang/merchant_info_iris/controller/common"
	"github.com/dalezhang/merchant_info_iris/model"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

func Show(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON
	id := bson.ObjectIdHex(ctx.Params().Get("id"))
	fmt.Println(id)

	var merchant model.Merchant
	session := model.CloneSession() //调用这个获得session
	defer session.Close()           //一定要记得释放
	collection := session.DB(config.DBConfig.Database).C("application_records")

	err := collection.FindId(id).One(&merchant)
	if err != nil {
		SendErrJSON(err.Error(), ctx)
		return
	}
	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data": iris.Map{
			"merchant": merchant,
		},
	})

}

func Index(ctx iris.Context) {
	//SendErrJSON := common.SendErrJSON
	session := model.CloneSession() //调用这个获得session
	defer session.Close()           //一定要记得释放

	var merchant model.Merchant
	var merchants []model.Merchant

	iter := session.DB(config.DBConfig.Database).C("application_records").Find(bson.M{"_type": "Merchant"}).Iter()
	for iter.Next(&merchant) {
		merchants = append(merchants, merchant)
	}

	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data": iris.Map{
			"merchants": merchants,
		},
	})

}
