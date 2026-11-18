package services

// NotificationService: stub for push notifications (e.g. OneSignal).
type NotificationService struct{ AppID, APIKey string }

func NewNotificationService(appID, apiKey string) *NotificationService { return &NotificationService{AppID: appID, APIKey: apiKey} }

func (s *NotificationService) Send(target, title, message string) (map[string]any, error) {
	// TODO: integrate OneSignal API.
	return map[string]any{"sent": true, "target": target, "title": title}, nil
}
