package gzip

import (
	"code-in-go/projectExercise/common/log/zap"
	"compress/gzip"
	"fmt"
	"os"
)

func UnGzipFile() {
	r, err := os.Open("./resource/tt.gzip")
	if err != nil {
		fmt.Println(err)
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		fmt.Println(err)
	}
	defer gr.Close()

	buf := make([]byte, 1024*10) // 如果单独使用，需自己决定要读多少内容，根据官方文档的说法，你读出的内容可能超出你的所需（当你压缩gzip文件中有多个文件时，强烈建议直接和tar组合使用）
	n, err := gr.Read(buf)

	fmt.Println("xxxx", string(buf), n)
	zap.Logger().Sugar().Info(string(buf))
}
