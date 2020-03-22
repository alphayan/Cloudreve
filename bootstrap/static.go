package bootstrap

import (
	"net/http"

	"github.com/HFO4/cloudreve/pkg/util"
	_ "github.com/HFO4/cloudreve/statik"
	"github.com/gin-contrib/static"
	"github.com/rakyll/statik/fs"
)

type GinFS struct {
	FS http.FileSystem
}

// StaticFS 内置静态文件资源
var StaticFS static.ServeFileSystem

// Open 打开文件
func (b *GinFS) Open(name string) (http.File, error) {
	return b.FS.Open(name)
}

// Exists 文件是否存在
func (b *GinFS) Exists(prefix string, filepath string) bool {

	if _, err := b.FS.Open(filepath); err != nil {
		return false
	}
	return true

}

// InitStatic 初始化静态资源文件
func InitStatic() {
	var err error

	if util.Exists(util.RelativePath("statics")) {
		util.Log().Info("检测到 statics 目录存在，将使用此目录下的静态资源文件")
		StaticFS = static.LocalFile(util.RelativePath("statics"), false)
	} else {
		StaticFS = &GinFS{}
		StaticFS.(*GinFS).FS, err = fs.New()
		if err != nil {
			util.Log().Panic("无法初始化静态资源, %s", err)
		}
	}

}
