package incident

type Severity string

const (
	CRITICAL Severity = "critical"
	WARNING  Severity = "warning"
	INFO     Severity = "info"
)

type Alert struct {
	ServiceId string
	Severity  Severity
	Title     string
	DedupKey  string
}
