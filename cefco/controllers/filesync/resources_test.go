package filesync

import (
	"reflect"
	"testing"

	filesyncv1alpha1 "github.com/inspursoft/cefco/pkg/apis/filesync/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var pvcName = "example"

var srcPath = "/example-mini/src"
var rcvPath = "/example-mini/rcv"
var logPath = "/example-mini/log"
var trashPath = "/example-mini/trash"

var testFileSyncList = []filesyncv1alpha1.FileSync{
	// mini: without pvc
	{ObjectMeta: metav1.ObjectMeta{
		Name:      "example-mini",
		Namespace: "default",
	}, Spec: filesyncv1alpha1.FileSyncSpec{
		Reader: &filesyncv1alpha1.ReaderSpec{
			DataSliceSize: "1m",
			DataTopicAndThread: filesyncv1alpha1.DataTopicAndThread{
				DataTopic: "example.mini",
			},
			FmodeRead: 0,
			SyncMetadata: filesyncv1alpha1.SyncMetadata{
				Image:        "filereader:7.8.3-20210507-centos",
				NodeSelector: map[string]string{"kubernetes.io/hostname": "centos78-0"},
			},
			LogsyncMode: 0,
			SrcPath:     srcPath,
			UserRead:    0,
		},
		Writer: &filesyncv1alpha1.WriterSpec{
			DataTopicAndThread: filesyncv1alpha1.DataTopicAndThread{
				DataTopic: "example.mini",
			},
			FmodeWrite: 0,
			SyncMetadata: filesyncv1alpha1.SyncMetadata{
				Image:        "filewriter:7.8.3-20210507-centos",
				NodeSelector: map[string]string{"kubernetes.io/hostname": "centos78-1"},
			},
			RcvPath:   rcvPath,
			UserWrite: 0,
		},
	}},
	// mini: with pvc
	{ObjectMeta: metav1.ObjectMeta{
		Name:      "example-mini",
		Namespace: "default",
	}, Spec: filesyncv1alpha1.FileSyncSpec{
		Reader: &filesyncv1alpha1.ReaderSpec{
			DataSliceSize: "1m",
			DataTopicAndThread: filesyncv1alpha1.DataTopicAndThread{
				DataTopic: "example.mini",
			},
			FmodeRead: 0,
			SyncMetadata: filesyncv1alpha1.SyncMetadata{
				Image:        "filereader:7.8.3-20210507-centos",
				NodeSelector: map[string]string{"kubernetes.io/hostname": "centos78-0"},
			},
			LogsyncMode: 0,
			SrcPath:     srcPath,
			SrcPathPVC:  &pvcName,
			UserRead:    0,
		},
		Writer: &filesyncv1alpha1.WriterSpec{
			DataTopicAndThread: filesyncv1alpha1.DataTopicAndThread{
				DataTopic: "example.mini",
			},
			FmodeWrite: 0,
			SyncMetadata: filesyncv1alpha1.SyncMetadata{
				Image:        "filewriter:7.8.3-20210507-centos",
				NodeSelector: map[string]string{"kubernetes.io/hostname": "centos78-1"},
			},
			RcvPath:    rcvPath,
			RcvPathPVC: &pvcName,
			UserWrite:  0,
		},
	}},
	// mini: with a common pvc
	{ObjectMeta: metav1.ObjectMeta{
		Name:      "example-mini",
		Namespace: "default",
	}, Spec: filesyncv1alpha1.FileSyncSpec{
		Reader: &filesyncv1alpha1.ReaderSpec{
			DataSliceSize: "1m",
			DataTopicAndThread: filesyncv1alpha1.DataTopicAndThread{
				DataTopic: "example.mini",
			},
			FmodeRead: 0,
			SyncMetadata: filesyncv1alpha1.SyncMetadata{
				Image:        "filereader:7.8.3-20210507-centos",
				NodeSelector: map[string]string{"kubernetes.io/hostname": "centos78-0"},
			},
			LogsyncMode: 0,
			SrcPath:     srcPath,
			SrcPathPVC:  &pvcName,
			LogSpec: filesyncv1alpha1.LogSpec{
				LogPath:    &logPath,
				LogPathPVC: &pvcName,
			},
			UserRead: 0,
		},
		Writer: &filesyncv1alpha1.WriterSpec{
			DataTopicAndThread: filesyncv1alpha1.DataTopicAndThread{
				DataTopic: "example.mini",
			},
			FmodeWrite: 0,
			SyncMetadata: filesyncv1alpha1.SyncMetadata{
				Image:        "filewriter:7.8.3-20210507-centos",
				NodeSelector: map[string]string{"kubernetes.io/hostname": "centos78-1"},
			},
			RcvPath:      rcvPath,
			RcvPathPVC:   &pvcName,
			TrashPath:    &trashPath,
			TrashPathPVC: &pvcName,
			UserWrite:    0,
		},
	}},
}

func TestCreateVolInfoMap(t *testing.T) {
	t.Parallel()

	srcSubpath := "src"
	rcvSubpath := "rcv"
	logSubpath := "log"
	trashSubpath := "trash"
	wantVMs := []map[string]*volInfo{
		// without pvc
		{"src": {name: "src", path: &srcPath},
			"shadow": {name: "shadow"},
			"log":    {name: "log"},
			"trash":  {name: "trash"}},
		{"rcv": {name: "rcv", path: &rcvPath},
			"log":   {name: "log"},
			"trash": {name: "trash"}},
		// with pvc
		{"src": {name: "src", path: &srcPath, pvc: &pvcName},
			"shadow": {name: "shadow"},
			"log":    {name: "log"},
			"trash":  {name: "trash"}},
		{"rcv": {name: "rcv", path: &rcvPath, pvc: &pvcName},
			"log":   {name: "log"},
			"trash": {name: "trash"}},
		// with a common pvc
		{"src": {name: "src", path: &srcPath, pvc: &pvcName, subPath: &srcSubpath},
			"shadow": {name: "shadow"},
			"log":    {name: "src", path: &logPath, pvc: &pvcName, subPath: &logSubpath},
			"trash":  {name: "trash"}},
		{"rcv": {name: "rcv", path: &rcvPath, pvc: &pvcName, subPath: &rcvSubpath},
			"log":   {name: "log"},
			"trash": {name: "rcv", path: &trashPath, pvc: &pvcName, subPath: &trashSubpath}},
	}
	gotVMs := []map[string]*volInfo{}
	for _, fs := range testFileSyncList {
		if fs.Spec.Reader != nil {
			gotVMs = append(gotVMs, createVolInfoMap(&fs, true))
		}
		if fs.Spec.Writer != nil {
			gotVMs = append(gotVMs, createVolInfoMap(&fs, false))
		}
	}
	if len(wantVMs) != len(gotVMs) {
		t.Fatalf("want lenth: %d, got length: %d", len(wantVMs), len(gotVMs))
	} else {
		for index := range wantVMs {
			if !reflect.DeepEqual(wantVMs[index], gotVMs[index]) {
				t.Errorf("index:  %d\nwant:\n%v\ngot:\n%v\n", index, wantVMs[index], gotVMs[index])
				t.Fatal("failed")
			}
		}
	}
}
