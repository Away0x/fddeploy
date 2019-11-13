package upload

import (
	"errors"
	"fddeploy/config"
)

func initConfig() error {
	targetConfig := config.StringMap(targetName)
	if len(targetConfig) == 0 {
		errMsg("target " + targetName + " config not found")
	}

	var (
		remoteHost     = configString("REMOTE_HOST")
		remotePort     = configInt("REMOTE_PORT")
		remoteUserName = configString("REMOTE_USER_NAME")
		remotePassword = configString("REMOTE_PASSWORD")

		uploadLocalDir  = configString("LOCAL_DIR")
		uploadRemoteDir = configString("REMOTE_DIR")
	)

	if remoteHost == "" || remotePort == 0 || remoteUserName == "" || remotePassword == "" {
		return errors.New("remote config (REMOTE_HOST & REMOTE_PORT & REMOTE_USER_NAME & REMOTE_PASSWORD) not found")
	}
	if uploadLocalDir == "" || uploadRemoteDir == "" {
		return errors.New("upload config (LOCAL_DIR & REMOTE_DIR) not found")
	}

	userInfo = &UserInfo{
		Name: configString("USER_NAME"),
	}

	remoteServerInfo = &RemoteServerInfo{
		Host:     remoteHost,
		Port:     remotePort,
		UserName: remoteUserName,
		Password: remotePassword,
	}

	uploadInfo = &UploadInfo{
		LocalDir:  uploadLocalDir,
		RemoteDir: uploadRemoteDir,
	}

	dingDingInfo = &DingDingInfo{
		Tag:     configString("DINGDING_TAG"),
		URL:     configString("DINGDING_URL"),
		Push:    configString("DINGDING_PUSH"),
		Message: uploadMessage,
	}

	return nil
}
