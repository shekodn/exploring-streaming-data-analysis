APP?=app
USERSPACE?=
RELEASE?=0.0.0
PROJECT?=github.com/${USERSPACE}/${APP}
GOOS?=linux
GOARCH=amd64
PORT?=8000


build:
	go build -o ${APP}

clean:
	rm app

run: build
	./${APP}
