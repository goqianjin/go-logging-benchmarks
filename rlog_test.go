package bench

import (
	"io"
	"os"
	"testing"

	"github.com/goqianjin/common-libs/xlog"
	"qiniu.com/kodo/libs/audit/largefile"
)

type rlogBenchmark interface {
	name() string
	new(w io.Writer) rlogBenchmark
	RLogEvent(msg string)
}

type xLogRLogBench struct {
	logger xlog.Logger
}

func (b *xLogRLogBench) new(w io.Writer) rlogBenchmark {
	w, err := os.OpenFile("run/xlog_rlog.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	l, _ := xlog.NewRLog(w, xlog.RawLogOption{AutoReqID: new(bool)})
	b.logger = l
	return b
}

func (b *xLogRLogBench) RLogEvent(msg string) {
	b.logger.Info(msg)
}

func (b *xLogRLogBench) name() string {
	return "XLogRLog"
}

type qiniuRLogBench struct {
	logger *largefile.Logger
}

func (b *qiniuRLogBench) new(w io.Writer) rlogBenchmark {
	l, err := largefile.Open("run/qiniu_rlog.log", 29)
	if err != nil {
		panic(err)
	}
	b.logger = l
	return b
}

func (b *qiniuRLogBench) RLogEvent(msg string) {
	b.logger.Log([]byte(msg))
}

func (b *qiniuRLogBench) name() string {
	return "QiniuRLog"
}

type fDirectRLogBench struct {
	w io.Writer
}

func (b *fDirectRLogBench) new(w io.Writer) rlogBenchmark {
	w, err := os.OpenFile("run/fDirect_rlog.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	b.w = w
	return b
}

func (b *fDirectRLogBench) RLogEvent(msg string) {
	b.w.Write([]byte(msg))
}

func (b *fDirectRLogBench) name() string {
	return "FDirectRLog"
}

var rLogMsg = `REQ	UP	%v	POST	/	{"Accept-Encoding":"gzip","Content-Length":"295418","Content-Type":"multipart/form-data; boundary=6c90c0bf5fd85c799c6fcd7595c4a627fec315807bac65f1229564054d38","Host":"up.qiniup.com","IP":"10.34.91.30","User-Agent":"Golang qiniu/client package","X-Forwarded-For":"43.143.171.244","X-Real-Ip":"43.143.171.244","X-Reqid":"dr0AAADDdbp64NcX","X-Scheme":"http"}		200	{"Content-Length":"131","Content-Type":"application/json","Tbl":"file-03","Token":{"appid":2420283688,"uid":1381724398,"utype":1067020},"X-Log":["body;lcy;msfa-lcy-360;put.at;msf.af0;ms.as0;disk.bp.lk0;disk.bp.a:2;disk.bp.ar:1;disk.bp.s0:90;fwd1-(disk.bp.lk0;disk.bp.a:15;disk.bp.ar:6;disk.bp.s0:14;disk.bp.bc;disk.bp.lk1;disk.bp.s1:8;disk.blk.p:38;disk.fpa:38;BLKSTG:38,fwd1:41);fwd0-(disk.bp.lk0;disk.bp.a:9;disk.bp.ar:1;disk.bp.s0:97;disk.bp.bc;disk.bp.lk1;disk.bp.s1:3;disk.blk.p:111;disk.fpa:111;BLKSTG:111,fwd0:111);disk.bp.bc:19;disk.bp.lk1;disk.bp.s1:6;disk.blk.p:118;disk.ppa:118;pa.body:1;BLKSTG:118;msf.bp1:119;msf.sp0:119;msf.np:119;msf.putv2:119;body;lc.set;line.set;dar.set;RsDbClusterID:default_v9_merged;rs40_shard.ins:5;v9.ins:5;mcl.ins:5;RS:6;rs.put:7;rs-upload.putFile:126;UP:128"],"X-Reqid":"dr0AAADDdbp64NcX","bucketRegion":"z0","rs-info":{"bucket":"file-03","itbl":485991365,"key":"1250097478917756417_c778b8f6-27be-11ef-b710-bd3a70f32030.gif","mimeType":"image/png"}}	{"hash":"FqDE2UI7efz6eFLHC8vJIHkT3qC3","key":"1250097478917756417_c778b8f6-27be-11ef-b710-bd3a70f32030.gif","x:name":"github logo"}	131	1285518`

var rLoggers = []rlogBenchmark{
	&xLogRLogBench{},
	&qiniuRLogBench{},
	&fDirectRLogBench{},
}

// BenchmarkEventAccumulatedCtx tests the impact of creating a logger with
// accumulated context and using it to log events.
func BenchmarkRLogEvent(b *testing.B) {
	b.Logf("RLog an event")

	for _, v := range rLoggers {
		b.Run(v.name(), func(b *testing.B) {
			out := &blackhole{}
			l := v.new(out)

			b.ResetTimer()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.RLogEvent(rLogMsg)
				}
			})

			/*if out.WriteCount() != uint64(b.N) {
				b.Fatalf(
					"Mismatch in log write count. Expected: %d, Actual: %d",
					b.N,
					out.WriteCount(),
				)
			}*/
		})
	}
}

func TestRLogEvent(t *testing.T) {
	for _, v := range rLoggers {
		out := &blackhole{}
		l := v.new(out)
		l.RLogEvent(rLogMsg)
	}
}
