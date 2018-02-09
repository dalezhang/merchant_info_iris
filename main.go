package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/dalezhang/merchant_info_iris/config"
	"github.com/dalezhang/merchant_info_iris/model"
	"github.com/dalezhang/merchant_info_iris/route"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	dbSession, err := mgo.Dial(config.DBConfig.URL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer dbSession.Close()
	dbSession.SetMode(mgo.Monotonic, true)
	db := dbSession.DB(config.DBConfig.Database).C("application_records") //数据库名称

	if config.ServerConfig.Env == model.DevelopmentMode {
		db.LogMode(true)
	}
	model.DB = db

	govalidator.SetFieldsRequiredByDefault(true)
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
