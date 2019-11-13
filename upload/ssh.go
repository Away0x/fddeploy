package upload

import "golang.org/x/crypto/ssh"

func sshSessionAction(session *ssh.Session) error {
	// 远程删除
	return session.Run("rm -fr " + uploadInfo.RemoteDir + "; " + "mkdir " + uploadInfo.RemoteDir)
}
