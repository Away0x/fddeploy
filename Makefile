APP_NAME = "fddeploy"
PROXY = "https://goproxy.cn"
# PROXY = "https://goproxy.io"
# PROXY = "https://mirrors.aliyun.com/goproxy/"
# PROXY = "https://athens.azurefd.net/"
# PROXY = "https://proxy.golang.org/"

default:
	go build -o ${APP_NAME}
	# env GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

install:
	env GOPROXY=${PROXY} go mod download

clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi

help:
	@echo "make - compile the source code"
	@echo "make install - install dep"
	@echo "make clean - remove binary file"
