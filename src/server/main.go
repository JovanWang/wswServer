package main

import (
	"server/gate"
	"server/login"

	"github.com/liangdas/mqant"
	"github.com/liangdas/mqant/module/modules"
	//"server/tracing"
	"github.com/opentracing/opentracing-go"
	"sourcegraph.com/sourcegraph/appdash"
	appdashtracer "sourcegraph.com/sourcegraph/appdash/opentracing"
)

var (
	collector *appdash.RemoteCollector = nil

	// Here we use the local collector to create a new opentracing.Tracer
	tracer opentracing.Tracer = nil
)

func DefaultTracer() opentracing.Tracer {
	return tracer
}

func main() {
	app := mqant.CreateApp()
	app.DefaultTracer(func() opentracing.Tracer {
		if collector == nil {
			collector = appdash.NewRemoteCollector("127.0.0.1:7701")
			tracer = appdashtracer.NewTracer(collector)
		}
		return tracer
	})
	//app.Route("Chat",ChatRoute)
	app.Run(true, //只有是在调试模式下才会在控制台打印日志, 非调试模式下只在日志文件中输出日志
		modules.MasterModule(),
		gate.Module(),  //这是默认网关模块,是必须的支持 TCP,websocket,MQTT协议
		login.Module(), //这是用户登录验证模块
	)

}
