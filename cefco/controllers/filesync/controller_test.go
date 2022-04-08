package filesync

import (
	"bytes"
	"github.com/inspursoft/cefco/pkg/apis/filesync/v1alpha1"
	"reflect"
	"testing"
	"text/template"
)

func TestTpl(t *testing.T) {
	t.Parallel()

	logPath := "/log"
	spec := v1alpha1.FileSyncSpec{
		Global: v1alpha1.GlobalSpec{MQ: "NATS"},
		NATS: &v1alpha1.NATSSpec{
			IP:       "127.0.0.1",
			Port:     4222,
			User:     "",
			Password: "",
		},
		Reader: &v1alpha1.ReaderSpec{
			SrcPath:            "/tmp",
			DataTopicAndThread: v1alpha1.DataTopicAndThread{DataTopic: "test"},
			DataSliceSize:      "1m",
			FmodeRead:          0,
			UserRead:           0,
			LogsyncMode:        0,
			LogSpec: v1alpha1.LogSpec{
				LogPath: &logPath,
			},
		},
	}
	tmpl, err := template.New("test").Parse(fileReaderConfig)
	if err != nil {
		t.Fatalf("read file failed! err: %v", err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, spec)
	if err != nil {
		t.Fatalf("prase failed! err: %v", err)
	}
	want := `[GLOBAL]
MQ=NATS
[NATS]
NATS_IP=127.0.0.1
NATS_PORT=4222
NATS_USER=
NATS_PASSWD=
[READER]
SRC_PATH=/tmp
LOG_PATH=/log
DATA_TOPIC=test
DATA_SLICE_SIZE=1m
FMODE_READ=0
USER_READ=0
LOGSYNC_MODE=0
`

	if !reflect.DeepEqual(buf.String(), want) {
		t.Error("unexpected finalizers")
		t.Fatalf("\n####### got: #######\n%v####### want: #######\n%v", buf.String(), want)
	}
}
