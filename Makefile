APP?=main
USERSPACE?=
RELEASE?=0.0.0
PROJECT?=github.com/${USERSPACE}/${APP}
GOOS?=linux
GOARCH=amd64
PORT?=8000

FUNCTION_NAME?=sdn-streaming-data-analysis
CODEPIPELINE_BUCKET?=sdn-streaming-data-analysis-us-east-1-codepipeline-631787010520
ZIP_FILE?=function.zip
AWS_PROFILE?=sdn

build:
	go build -o ${APP} main.go

clean:
	rm ${APP}
	rm ${ZIP_FILE}

run: build
	./${APP}

path:
	echo ${path}

prepare:
	GOOS=${GOOS} go build main.go
	zip ${ZIP_FILE} main

update: prepare
	aws s3 cp ${ZIP_FILE} s3://${CODEPIPELINE_BUCKET} --profile ${AWS_PROFILE}

deploy: update
	aws lambda update-function-code --function-name ${FUNCTION_NAME} \
	--s3-bucket ${CODEPIPELINE_BUCKET} --s3-key ${ZIP_FILE} \
	--profile ${AWS_PROFILE}
