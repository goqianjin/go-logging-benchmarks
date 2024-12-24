module github.com/betterstack-community/go-logging-benchmarks

go 1.22

toolchain go1.22.2

require (
	github.com/apex/log v1.9.0
	github.com/goqianjin/common-libs/xlog v0.0.0-20240903075630-90a6286687dc
	github.com/inconshreveable/log15 v2.16.0+incompatible
	github.com/phuslu/log v1.0.99
	github.com/rs/zerolog v1.30.0
	github.com/sirupsen/logrus v1.9.3
	github.com/zerodha/logf v0.5.5
	go.uber.org/multierr v1.10.0
	go.uber.org/zap v1.25.0
	go.uber.org/zap/exp v0.2.0
	qiniu.com/kodo/libs v0.0.0-20240902112219-4a5404e2de20

)

require (
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/qiniu/errors v0.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.12.0 // indirect

)

replace github.com/goqianjin/common-libs/xlog => ../common-libs/xlog

// go list 报错

replace github.com/qiniu/errors => ../../qbox/kodo/src/github.com/qiniu/errors

replace github.com/qiniupd/qiniu-go-sdk => ../kodo/src/github.com/qiniupd/qiniu-go-sdk

replace github.com/globalsign/mgo v0.0.0 => ../../qbox/kodo/src/github.com/globalsign/mgo

replace qiniu.com/kodo/thirdparty/gopkg.in/redis.v5 => ../../qbox/kodo/thirdparty/gopkg.in/redis.v5

replace qbox.us/account v0.0.0 => ../../qbox/kodo/libs/base-replacement/account

replace qbox.us/account-api v0.0.0 => ../../qbox/kodo/libs/base-replacement/account-api

replace qbox.us/servend => ../../qbox/kodo/kodo/libs/base-replacement/servend

replace qiniupkg.com/api.v7 => ../../qbox/kodo/kodo/src/qiniupkg.com/api.v7

replace qiniupkg.com/x => ../../qbox/kodo/kodo/src/qiniupkg.com/x

replace github.com/koofr/go-cryptoutils => ../../qbox/kodo/kodo/src/github.com/koofr/go-cryptoutils

replace github.com/dolab/gogo v0.0.0 => ../../qbox/kodo/kodo/src/github.com/dolab/gogo

replace github.com/koofr/go-ioutils => ../../qbox/kodo/kodo/src/github.com/koofr/go-ioutils

replace github.com/koofr/pb => ../../qbox/kodo/kodo/src/github.com/koofr/pb

replace github.com/nsqio/go-nsq => ../../qbox/kodo/kodo/src/github.com/nsqio/go-nsq

replace github.com/sbunce/bson v0.0.0 => ../../qbox/kodo/kodo/src/github.com/sbunce/bson
