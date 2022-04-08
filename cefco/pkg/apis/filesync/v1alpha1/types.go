package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FileSync is a specification for a FileSync resource
type FileSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FileSyncSpec   `json:"spec"`
	Status FileSyncStatus `json:"status"`
}

type GlobalSpec struct {
	MQ string `json:"mq"`
}

type NATSSpec struct {
	IP          string  `json:"ip"`
	Port        int     `json:"port"`
	User        string  `json:"user"`
	Password    string  `json:"password"`
	CredsSecret *string `json:"credsSecret,omitempty"`
	SecretKey   string  `json:"secretKey,omitempty"`
}

type MonitSpec struct {
	Topic       string  `json:"topic"`
	Encoding    string  `json:"encoding"`
	Hash        int     `json:"hash"`
	EndMsgTopic *string `json:"endMsgTopic,omitempty"`
}

type LogSpec struct {
	LogPath    *string `json:"logPath,omitempty"`
	LogPathPVC *string `json:"logPathPVC,omitempty"`
	LogLevel   *string `json:"logLevel,omitempty"`
}

type DataTopicAndThread struct {
	DataTopic          string `json:"dataTopic"`
	DontCheckDataTopic bool   `json:"dontCheckDataTopic,omitempty"`
	ThreadNum          *int   `json:"threadNum,omitempty"`
}

type SyncMetadata struct {
	NodeSelector map[string]string `json:"nodeSelector"`
	Image        string            `json:"image"`

	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Tolerations      []corev1.Toleration           `json:"tolerations,omitempty"`
	Labels           map[string]string             `json:"labels,omitempty"`
}

type ReaderSpec struct {
	SyncMetadata       `json:",inline"`
	LogSpec            `json:",inline"`
	DataTopicAndThread `json:",inline"`

	SrcPath         string  `json:"srcPath"`
	SrcPathPVC      *string `json:"srcPathPVC,omitempty"`
	ShadowPath      *string `json:"shadowPath,omitempty"`
	ShadowPathPVC   *string `json:"shadowPathPVC,omitempty"`
	DirShadowPrefix *string `json:"dirShadowPrefix,omitempty"`

	HandleModeAfterRead *string `json:"handleModeAfterRead,omitempty"`
	TrashPath           *string `json:"trashPath,omitempty"`
	TrashPathPVC        *string `json:"trashPathPVC,omitempty"`

	IncScanMode     *int `json:"incScanMode,omitempty"`
	IncScanInterval *int `json:"incScanInterval,omitempty"`
	IncQnotifyMode  *int `json:"incQnotifyMode,omitempty"`

	DirLevelUseThread *int   `json:"dirLevelUseThread,omitempty"`
	DataSliceSize     string `json:"dataSliceSize"`

	IncSkipDays *int `json:"incSkipDays,omitempty"`

	IncludedFiletype *string `json:"includedFiletype,omitempty"`
	ExcludedFiletype *string `json:"excludedFiletype,omitempty"`
	FilepathRegex    *string `json:"filepathRegex,omitempty"`

	FmodeRead int `json:"fmodeRead"`
	UserRead  int `json:"userRead"`

	LogsyncMode int `json:"logsyncMode"`

	ReaderLabel *string `json:"readerLabel,omitempty"`
}

type WriterSpec struct {
	SyncMetadata       `json:",inline"`
	LogSpec            `json:",inline"`
	DataTopicAndThread `json:",inline"`

	RcvPath      string  `json:"rcvPath"`
	RcvPathPVC   *string `json:"rcvPathPVC,omitempty"`
	TrashPath    *string `json:"trashPath,omitempty"`
	TrashPathPVC *string `json:"trashPathPVC,omitempty"`

	FmodeWrite int `json:"fmodeWrite"`
	UserWrite  int `json:"userWrite"`

	BisyncShadowPath *string `json:"bisyncShadowPath,omitempty"`
	BisyncMode       *int    `json:"bisyncMode,omitempty"`

	BigfileSliceCount *int    `json:"bigfileSliceCount,omitempty"`
	WriterLabel       *string `json:"writerLabel,omitempty"`
}

// FileSyncSpec is the spec for a FileSync resource
type FileSyncSpec struct {
	Global GlobalSpec  `json:"global"`
	NATS   *NATSSpec   `json:"nats,omitempty"`
	Monit  *MonitSpec  `json:"monit,omitempty"`
	Reader *ReaderSpec `json:"reader,omitempty"`
	Writer *WriterSpec `json:"writer,omitempty"`
}

// FileSyncStatus is the status for a FileSync resource
type FileSyncStatus struct {
	ObservedGeneration int64   `json:"observedGeneration"`
	ConfigMapStatus    string  `json:"configmapStatus"`
	ReaderStatus       *string `json:"readerStatus,omitempty"`
	WriterStatus       *string `json:"writerStatus,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FileSyncList is a list of FileSync resources
type FileSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []FileSync `json:"items"`
}
