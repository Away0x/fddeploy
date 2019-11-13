package upload

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pkg/sftp"
)

func ftpUpload(client *sftp.Client) (err error) {
	start := time.Now()

	// 本地路径和远程路径的监测
	_, err = client.Stat(uploadInfo.RemoteDir)
	if err != nil {
		return errors.New("ftpUpload error - " + uploadInfo.RemoteDir + " remote path not exists!")
	}
	_, err = ioutil.ReadDir(uploadInfo.LocalDir)
	if err != nil {
		return errors.New("ftpUpload error - " + uploadInfo.LocalDir + " local path not exists!")
	}

	// 上传
	if err = uploadDir(client, uploadInfo.LocalDir, uploadInfo.RemoteDir); err != nil {
		return errors.New("ftpUpload error - " + err.Error())
	}

	fmt.Println("elapsed time : ", time.Since(start))
	return nil
}
