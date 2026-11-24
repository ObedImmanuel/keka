package incident

type Status string

const (
	TRIGGERED Status = "triggered"
	NOTIFIED  Status = "notified"
)

type NotificationStatus string

const (
	SENT    NotificationStatus = "sent"
	FAILED  NotificationStatus = "failed"
	SKIPPED NotificationStatus = "skipped"
)

type Incident struct {
	IncidentId         string
	ServiceId          string
	Severity           Severity
	Status             Status
	NotificationStatus NotificationStatus
}
