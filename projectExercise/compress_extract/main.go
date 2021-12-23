package main

import (
	"code-in-go/projectExercise/common/log/zap"
	"code-in-go/projectExercise/compress_extract/gzip"
)

func main() {
	zap.Logger().Sugar().Info("start ..")
	gzip.UnGzipFile()
}

func init() {
	zap.InitLogger("./all.log", "debug")
}
