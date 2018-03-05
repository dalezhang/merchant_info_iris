package main

import (
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/dalezhang/merchant_info_iris/config"
	"github.com/dalezhang/merchant_info_iris/model"
	"github.com/dalezhang/merchant_info_iris/route"
	//	"github.com/dalezhang/merchant_info_iris/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//	"fmt"
	"time"
)

func init() {
	globalMgoSession, err := mgo.DialWithTimeout(config.DBConfig.URL, 10*time.Second) //连接数据库
	if err != nil {
		panic(err)
	}
	// defer dbSession.Close()
	globalMgoSession.SetMode(mgo.Monotonic, true)
	globalMgoSession.SetPoolLimit(300)
	//	collection := globalMgoSession.DB(config.DBConfig.Database).C("application_records") //数据库名称

	model.GlobalMgoSession = globalMgoSession

	govalidator.SetFieldsRequiredByDefault(false)
}

func main() {
	app := iris.New()

	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	app.Use(logger.New())

	route.Route(app)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errNo": model.ErrorCode.NotFound,
			"msg":   "Not Found",
			"data":  iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errNo": model.ErrorCode.ERROR,
			"msg":   "error",
			"data":  iris.Map{},
		})
	})

	addr := iris.Addr(":" + strconv.Itoa(config.ServerConfig.Port))
	if config.ServerConfig.Env == model.DevelopmentMode {
		app.Run(addr)
	} else {
		app.Run(addr, iris.WithoutVersionChecker)
	}
}
