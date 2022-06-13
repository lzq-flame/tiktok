package logger

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/6 22:08
 **/

import (
	log "github.com/sirupsen/logrus"
)

var l *log.Logger

//Init
//@Description:初始化函数
func Init() {
	l = log.New()

	formatter := &log.TextFormatter{
		FullTimestamp:   true,                  //日志打印时间
		TimestampFormat: "2006-01-02 15:04:05", // 定义时间戳格式
		ForceColors:     true,
	}

	l.SetFormatter(formatter)
	l.SetLevel(log.TraceLevel)

	l.Info("init logger success")
}

func L() *log.Logger {
	return l
}
