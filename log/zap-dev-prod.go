package main

import (
    "go.uber.org/zap"
    "time"
)

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

    loggerS := logger.Named("service").Sugar()
    loggerS.Debugw("msg", "k1", 123, "k2", "v2")
    loggerS.Debugw("msg", "k1", map[string]interface{}{"k":1})
}
