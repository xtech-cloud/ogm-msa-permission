BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )
scope := $(shell cat /tmp/permission-scope)

.PHONY: build
build:
	go build -ldflags \
		"\
		-X 'main.BuildVersion=${BUILD_VERSION}' \
		-X 'main.BuildTime=${BUILD_TIME}' \
		-X 'main.CommitID=${COMMIT_SHA1}' \
		"\
		-o ./bin/${APP_NAME}

.PHONY: run
run:
	./bin/${APP_NAME}

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm -rf /tmp/ogm-permission.db

.PHONY: call
call:
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Healthy.Echo '{"msg":"hello"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Create '{"key":"test.1", "name":"test-1"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Create '{"key":"test.1", "name":"test-2"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Create '{"key":"test.2", "name":"test-1"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Create '{"key":"test.2", "name":"test-2"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.List '{"offset":0, "count":50}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.List '{"offset":0, "count":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.List '{"offset":1, "count":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Search '{"offset":0, "count":50, "key":"test"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Search '{"offset":0, "count":50, "key":"st."}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Search '{"offset":0, "count":50, "key":".1"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Search '{"offset":0, "count":50, "name":"test"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Search '{"offset":0, "count":50, "name":"st-"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Search '{"offset":0, "count":50, "name":"-1"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Scope.Search '{"offset":0, "count":50, "key":"1", "name":"1"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Rule.Add '{"scope":"${scope}", "key":"111", "name":"aaaaa", "state":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Rule.Add '{"scope":"${scope}", "key":"123", "name":"bbb", "state":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Rule.List '{"scope":"${scope}", "offset":0, "count":50}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Rule.Search '{"scope":"${scope}", "offset":0, "count":50, "key":"1"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Rule.Search '{"scope":"${scope}", "offset":0, "count":50, "key":"2"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.permission Rule.Export '{"scope":"${scope}"}'

.PHONY: post
post:
	curl -X POST -d '{"msg":"hello"}' localhost/ogm/permission/Healthy/Echo                                                                                     1

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

.PHONY: docker
	docker:
	docker build -t xtechcloud/${APP_NAME}:${BUILD_VERSION} .
	docker rm -f ${APP_NAME}
	docker run --restart=always --name=${APP_NAME} --net=host -v /data/${APP_NAME}:/ogm -e MSA_REGISTRY_ADDRESS='localhost:2379' -e MSA_CONFIG_DEFINE='{"source":"file","prefix":"/ogm/config","key":"${APP_NAME}.yaml"}' -d xtechcloud/${APP_NAME}:${BUILD_VERSION}
	docker logs -f ${APP_NAME}
