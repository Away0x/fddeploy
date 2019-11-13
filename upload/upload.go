package upload

import (
	"fmt"
)

type (
	// UserInfo 用户信息
	UserInfo struct {
		Name string
	}
	// RemoteServerInfo 服务器信息
	RemoteServerInfo struct {
		Host     string
		Port     int
		UserName string
		Password string
	}
	// UploadInfo 上传信息
	UploadInfo struct {
		LocalDir  string // 需上传的本地文件夹路径
		RemoteDir string // 上传到的远程文件夹路径
	}
	// DingDingInfo 推送上传信息到 "钉钉"
	DingDingInfo struct {
		Tag     string // 推送信息的 tag
		URL     string // 点击推送信息跳转的 url
		Push    string // 钉钉的推送地址
		Message string // 推送信息的描述
	}
)

var (
	// 存储一些全局配置
	targetName       string
	uploadMessage    string
	userInfo         *UserInfo
	remoteServerInfo *RemoteServerInfo
	uploadInfo       *UploadInfo
	dingDingInfo     *DingDingInfo
)

// Run 执行上传
func Run(target, message string) {
	var err error
	targetName, uploadMessage = target, message
	fmt.Println(targetName + " upload command running...... ")

	// 初始化 upload config
	if err = initConfig(); err != nil {
		errMsg("config error: " + err.Error())
	}

	// 获取 client
	sftpClient, session, err := connect()
	if err != nil {
		errMsg("connect error: " + err.Error())
	}
	defer sftpClient.Close()
	defer session.Close()

	// 远程执行命令
	if err = sshSessionAction(session); err != nil {
		errMsg("ssh session action error: " + err.Error())
	}

	// 执行上传
	if err = ftpUpload(sftpClient); err != nil {
		errMsg(err.Error())
	}

	// 发送钉钉消息
	if dingDingInfo.Push != "" {
		if err = dingdingAction(); err != nil {
			errMsg(err.Error())
		}
	}
}
