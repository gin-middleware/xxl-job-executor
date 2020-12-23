package xxl_job_executor_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/xxl-job/xxl-job-executor-go"
)

func XxlJobMux(e *gin.Engine, exec xxl.Executor) {
	//注册的gin的路由
	e.POST("run", gin.WrapF(exec.RunTask))
	e.POST("kill", gin.WrapF(exec.KillTask))
	e.POST("log", gin.WrapF(exec.TaskLog))
}
