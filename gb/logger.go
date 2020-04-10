package gb

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var Logger *zap.SugaredLogger

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func LogConf() {
	now := time.Now()
	hook := &lumberjack.Logger{
		Filename:   fmt.Sprintf("log/%04d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), //filePath
		MaxSize:    500,                                                                                                                     // megabytes
		MaxBackups: 10,
		MaxAge:     180,   //days
		Compress:   false, // disabled by default
	}
	defer hook.Close()
	/*zap 的 Config 非常的繁琐也非常强大，可以控制打印 log 的所有细节，因此对于我们开发者是友好的，有利于二次封装。
	  但是对于初学者则是噩梦。因此 zap 提供了一整套的易用配置，大部分的姿势都可以通过一句代码生成需要的配置。
	*/
	enConfig := zap.NewProductionEncoderConfig() //生成配置

	// 时间格式
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 指定日期格式
	//enConfig.EncodeCaller = zapcore.NewJSONEncoder()

	level := zap.DebugLevel // 只有大于等于这个级别的日志才会被记录
	w := zapcore.AddSync(hook)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(enConfig), //编码器配置
		w,                                //打印到控制台和文件
		level,                            //日志等级
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	_log := log.New(hook, "", log.LstdFlags)
	Logger = logger.Sugar() // 初始化全局的Logger
	_log.Println("Start...")
}
