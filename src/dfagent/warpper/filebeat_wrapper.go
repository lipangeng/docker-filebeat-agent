package warpper

import (
	"os"
	"os/exec"
	"log"
	"errors"
	"syscall"
)

type FilebeatConf struct {
	FilePath   string
	ConfigPath string
}

// filebeat的命令
var filebeatExec *exec.Cmd
// filebeat的配置文件
var filebeatConf FilebeatConf

// 启动Filebeat
func StartFilebeat() error {
	if filebeatExec != nil {
		return errors.New("进程已经存在")
	}
	log.Println("准备启动Filebeat")
	filebeatExec = exec.Command("/usr/local/filebeat", "-c", "/app/filebeat.yml")
	filebeatExec.Stderr = os.Stderr
	filebeatExec.Stdout = os.Stdout
	return filebeatExec.Start()
}

// 启动Filebeat
func StopFilebeat() error {
	if filebeatExec == nil {
		return errors.New("进程已经关闭，无需重复关闭")
	}
	return filebeatExec.Process.Signal(syscall.SIGQUIT)
}

// 启动Filebeat
func ReloadFilebeat() error {
	err := StopFilebeat()
	if err != nil {
		log.Println("关闭Filebeat时发生错误:" + err.Error())
		return err
	}
	return StartFilebeat()
}

// 初始化配置文件
func InitFilebeatConf(filePath string, configPath string) {
	filebeatConf = FilebeatConf{filePath, configPath}
}
