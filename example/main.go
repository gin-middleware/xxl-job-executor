package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-middleware/xxl-job-executor"
	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
	"log"
)

const Port = "9999"

func main() {
	//初始化执行器
	exec := xxl.NewExecutor(
		xxl.ServerAddr("http://127.0.0.1/xxl-job-admin"),
		xxl.AccessToken(""),            //请求令牌(默认为空)
		xxl.ExecutorIp("127.0.0.1"),    //可自动获取
		xxl.ExecutorPort(Port),         //默认9999（此处要与gin服务启动port必需一至）
		xxl.RegistryKey("golang-jobs"), //执行器名称
	)
	exec.Init()
	//添加到gin路由
	r := gin.Default()
	xxl_job_executor_gin.XxlJobMux(r, exec)

	//注册gin的handler
	r.GET("ping", func(cxt *gin.Context) {
		cxt.JSON(200, "pong")
	})

	//注册任务handler
	exec.RegTask("task.test", task.Test)
	exec.RegTask("task.test2", task.Test2)
	exec.RegTask("task.panic", task.Panic)

	log.Fatal(r.Run(":" + Port))
}
