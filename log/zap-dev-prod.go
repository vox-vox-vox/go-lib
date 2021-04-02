package main

import (
    "go.uber.org/zap"
    "time"
)

type User struct {
    Username string
    Age int 
}

func main() {
    // format NewProduction json序列化输出; NewDevelopment: 普通Info format
    logger, _ := zap.NewDevelopment()
    logger, _ = zap.NewProduction()
    logger = zap.NewExample()
    defer logger.Sync()
    logger.Info("无法获取网址",
        zap.String("url", "http://www.baidu.com"),
        zap.Int("attempt", 3),
        zap.Duration("backoff", time.Second),
    )

    loggerS := logger.Named("日志名").Sugar()
    loggerS.Debugw("Sugar", "k1", 123, "k2", "v2")
    loggerS.Debugw("Sugar", "k1", map[string]interface{}{"k":1}, User{})
}
