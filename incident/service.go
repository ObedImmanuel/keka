package incident

import (
	"context"
	"fmt"
	"keka/errors"
	"keka/notification"
	"log"
	"sync"
)

func NewService(notifService notification.Service) *Service {
	return &Service{
		mu:                   sync.Mutex{},
		globalCounter:        0,
		IncidentsStore:       make(map[string]*Incident),
		DedupKeyToIncidentId: make(map[string]string),
		NotificationService:  notifService,
		serviceIdPhoneNumMap: map[string]string{
			"svc-payments": "+919900112233",
			"svc-checkout": "+919900445566",
			"svc-search":   "+919900778899",
		},
	}
}

type Service struct {
	mu                   sync.Mutex
	globalCounter        int                  // counter to generate incident id
	IncidentsStore       map[string]*Incident // stores the list of incidents, key -> incident id
	DedupKeyToIncidentId map[string]string
	NotificationService  notification.Service
	serviceIdPhoneNumMap map[string]string
}

func (s *Service) Alert(ctx context.Context, alert *Alert) (*Incident, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	//check if incident is already created for the alert
	if id, ok := s.DedupKeyToIncidentId[alert.DedupKey]; ok {
		log.Printf("incident already created for %s with %s id", alert.DedupKey, id)
		return nil, errors.ErrAlreadyExists
	}

	// store incident
	id := "inc-" + fmt.Sprintf("%s", s.globalCounter)
	incident := &Incident{
		IncidentId:         id,
		ServiceId:          alert.ServiceId,
		Severity:           alert.Severity,
		Status:             TRIGGERED,
		NotificationStatus: "",
	}
	s.globalCounter++

	s.DedupKeyToIncidentId[alert.DedupKey] = id
	s.IncidentsStore[id] = incident

	// send notification if applicable and update the notification status
	switch s.IncidentsStore[id].Severity {
	case CRITICAL:
		err := s.placePhoneCall(ctx, id, alert.Title)
		if err != nil {
			s.IncidentsStore[id].NotificationStatus = FAILED
			return nil, err
		}

		s.IncidentsStore[id].NotificationStatus = SENT
		s.IncidentsStore[id].Status = NOTIFIED
	case INFO, WARNING:
		s.IncidentsStore[id].NotificationStatus = SKIPPED
	}

	return s.IncidentsStore[id], nil
}

func (s *Service) placePhoneCall(ctx context.Context, id string, title string) error {
	incident := s.IncidentsStore[id]
	callResp, callErr := s.NotificationService.Call(ctx, &notification.CallRequest{
		To:      s.serviceIdPhoneNumMap[incident.ServiceId],
		Message: fmt.Sprintf("%s: %s", incident.Severity, title),
	})
	if callErr != nil {
		log.Printf("call failed %v", callErr)
	}

	log.Printf("call performed, sid: %s, status: %s", callResp.CallSid, callResp.Status)
	return nil
}
