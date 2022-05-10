package command

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"esp_webrtc/conf"
	"esp_webrtc/utils/filetil"

	"github.com/lifei6671/gocaptcha"
)

/***
解析命令行
***/
func ResolveCommand(args []string) {
	flagset := flag.NewFlagSet("leewiki command:", flag.ExitOnError)
	flagset.StringVar(&conf.ConfigurationFile, "config", "", "leewiki configuration file.")
	flagset.StringVar(&conf.WorkingDirectory, "dir", "", "MinDoc working directory.")
	flagset.StringVar(&conf.LogFile, "log", "", "MinDoc log file path.")
	if err := flagset.Parse(args); err != nil {
		log.Fatal("解析命令失败->", err)
	}
	//判断工作目录是否为空，为空则返回当前路径的绝对路径
	if conf.WorkingDirectory == "" {
		if p, err := filepath.Abs(os.Args[0]); err == nil {
			conf.WorkingDirectory = filepath.Dir(p)
		}
	}
	//判断配置文件选项是否为空，若为空且app.conf不存在，则复制app.conf.exmple
	if conf.ConfigurationFile == "" {
		conf.ConfigurationFile = conf.WorkingDir("conf", "app.conf")
		config := conf.WorkingDir("conf", "app.conf.example")
		if !filetil.FileExists(conf.ConfigurationFile) && filetil.FileExists(config) {
			_ = filetil.CopyFile(conf.ConfigurationFile, config)
		}
	}
	if err := gocaptcha.ReadFonts(conf.WorkingDir("static", "fonts"), ".ttf"); err != nil {
		log.Fatal("读取字体文件时出错 -> ", err)
	}
}
