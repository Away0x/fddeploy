package upload

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/Away0x/fddeploy/config"

	"github.com/pkg/sftp"
)

func errMsg(msg string) {
	log.Fatal("[upload " + targetName + " error] " + msg)
}

func configString(name string) string {
	return config.String(targetName + "." + name)
}

func configInt(name string) int {
	return config.Int(targetName + "." + name)
}

// 上传文件
func uploadFile(client *sftp.Client, localPath, remotePath string) (err error) {
	srcFile, err := os.Open(localPath)
	if err != nil {
		fmt.Println("os.Open error : ", localPath)
		return err
	}
	defer srcFile.Close()

	// 创建远程文件
	remoteFileName := path.Base(localPath) // 获取要上传文件的文件名
	dstFile, err := client.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("client.Create error : ", path.Join(remotePath, remoteFileName))
		return err
	}
	defer dstFile.Close()

	// copy
	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll error : ", localPath)
		return err
	}

	dstFile.Write(ff)
	fmt.Println(localPath + " copy file to remote server finished!")
	return nil
}

// 上传文件夹
func uploadDir(client *sftp.Client, localPath, remotePath string) (err error) {
	localFiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		fmt.Println("read dir list fail ", err)
		return err
	}

	// 递归上传
	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())
		remoteFilePath := path.Join(remotePath, backupDir.Name())
		if backupDir.IsDir() {
			client.Mkdir(remoteFilePath) // 创建文件夹
			uploadDir(client, localFilePath, remoteFilePath)
		} else {
			uploadFile(client, path.Join(localPath, backupDir.Name()), remotePath)
		}
	}

	fmt.Println(localPath + " copy directory to remote server finished!")
	return nil
}
