package filesync

var credsBasePath = "/nats-accounts/creds/"

var fileReaderConfig = `[GLOBAL]
MQ={{.Global.MQ}}
{{- if .NATS}}
[NATS]
NATS_IP={{.NATS.IP}}
NATS_PORT={{.NATS.Port}}
NATS_USER={{.NATS.User}}
NATS_PASSWD={{.NATS.Password}}
{{- if .NATS.CredsSecret}}
NATS_CREDS=` + credsBasePath + `{{.NATS.SecretKey}}
{{- end}}
{{- end}}
{{- if .Monit}}
[MONIT]
MONIT_TOPIC={{.Monit.Topic}}
MONIT_ENCODING={{.Monit.Encoding}}
MONIT_HASH={{.Monit.Hash}}
END_MSG_TOPIC={{.Monit.EndMsgTopic}}
{{- end}}
[READER]
SRC_PATH={{.Reader.SrcPath}}
{{- if .Reader.ShadowPath}}
SHADOW_PATH={{.Reader.ShadowPath}}
{{- end}}
{{- if .Reader.DirShadowPrefix}}
DIR_SHADOW_PREFIX={{.Reader.DirShadowPrefix}}
{{- end}}
{{- if .Reader.LogPath}}
LOG_PATH={{.Reader.LogPath}}
{{- end}}
{{- if .Reader.LogLevel}}
LOG_LEVEL={{.Reader.LogLevel}}
{{- end}}
{{- if .Reader.HandleModeAfterRead}}
HANDLE_MODE_AFTER_READ={{.Reader.HandleModeAfterRead}}
{{- end}}
{{- if .Reader.TrashPath}}
TRASH_PATH={{.Reader.TrashPath}}
{{- end}}
{{- if .Reader.IncScanMode}}
INC_SCAN_MODE={{.Reader.IncScanMode}}
{{- end}}
{{- if .Reader.IncScanInterval}}
INC_SCAN_INTERVAL={{.Reader.IncScanInterval}}
{{- end}}
{{- if .Reader.IncQnotifyMode}}
INC_QNOTIFY_MODE={{.Reader.IncQnotifyMode}}
{{- end}}
DATA_TOPIC={{.Reader.DataTopic}}
{{- if .Reader.ThreadNum}}
THREAD_NUM={{.Reader.ThreadNum}}
{{- end}}
{{- if .Reader.DirLevelUseThread}}
DIR_LEVEL_USE_THREAD={{.Reader.DirLevelUseThread}}
{{- end}}
{{- if .Reader.DataSliceSize}}
DATA_SLICE_SIZE={{.Reader.DataSliceSize}}
{{- end}}
{{- if .Reader.IncSkipDays}}
INC_SKIP_DAYS={{.Reader.IncSkipDays}}
{{- end}}
{{- if .Reader.IncludedFiletype}}
INCLUDED_FILETYPE={{.Reader.IncludedFiletype}}
{{- end}}
{{- if .Reader.ExcludedFiletype}}
EXCLUDED_FILETYPE={{.Reader.ExcludedFiletype}}
{{- end}}
{{- if .Reader.FilepathRegex}}
FILEPATH_REGEX={{.Reader.FilepathRegex}}
{{- end}}
FMODE_READ={{.Reader.FmodeRead}}
USER_READ={{.Reader.UserRead}}
LOGSYNC_MODE={{.Reader.LogsyncMode}}
{{- if .Reader.ReaderLabel}}
READER_LABEL={{.Reader.ReaderLabel}}
{{- end}}
`

var fileWriterConfig = `[GLOBAL]
MQ={{.Global.MQ}}
{{- if .NATS}}
[NATS]
NATS_IP={{.NATS.IP}}
NATS_PORT={{.NATS.Port}}
NATS_USER={{.NATS.User}}
NATS_PASSWD={{.NATS.Password}}
{{- if .NATS.CredsSecret}}
NATS_CREDS=` + credsBasePath + `{{.NATS.SecretKey}}
{{- end}}
{{- end}}
{{- if .Monit}}
[MONIT]
MONIT_TOPIC={{.Monit.Topic}}
MONIT_ENCODING={{.Monit.Encoding}}
MONIT_HASH={{.Monit.Hash}}
{{- end}}
[WRITER]
DATA_TOPIC={{.Writer.DataTopic}}
{{- if .Writer.ThreadNum}}
THREAD_NUM={{.Writer.ThreadNum}}
{{- end}}
{{- if .Writer.RcvPath}}
RCV_PATH={{.Writer.RcvPath}}
{{- end}}
{{- if .Writer.TrashPath}}
TRASH_PATH={{.Writer.TrashPath}}
{{- end}}
{{- if .Writer.LogPath}}
LOG_PATH={{.Writer.LogPath}}
{{- end}}
{{- if .Writer.LogLevel}}
LOG_LEVEL={{.Writer.LogLevel}}
{{- end}}
FMODE_WRITE={{.Writer.FmodeWrite}}
USER_WRITE={{.Writer.UserWrite}}
{{- if .Writer.BisyncShadowPath}}
BISYNC_SHADOW_PATH={{.Writer.BisyncShadowPath}}
{{- end}}
{{- if .Writer.BisyncMode}}
BISYNC_MODE={{.Writer.BisyncMode}}
{{- end}}
{{- if .Writer.BigfileSliceCount}}
BIGFILE_SLICE_COUNT={{.Writer.BigfileSliceCount}}
{{- end}}
{{- if .Writer.WriterLabel}}
WRITER_LABEL={{.Writer.WriterLabel}}
{{- end}}
`
