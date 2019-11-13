package upload

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// ssh 密码连接服务器，创建 sftp client
func connect() (sftpClient *sftp.Client, session *ssh.Session, err error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(remoteServerInfo.Password)) // 使用密码连接

	// 配置
	clientConfig = &ssh.ClientConfig{
		User:            remoteServerInfo.UserName,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 如本地有 360 可能这里会报木马 -_- ...
	}

	// ssh 连接
	addr = fmt.Sprintf("%s:%d", remoteServerInfo.Host, remoteServerInfo.Port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, nil, err
	}

	// 创建 sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, nil, err
	}

	// 创建会话
	if session, err = sshClient.NewSession(); err != nil {
		return nil, nil, err
	}

	// 将远程的输出重定向为该 go 应用的输出
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	return sftpClient, session, nil
}
