// pkg/notifications/notifications.go
package notifications

import "fmt"

func SendEmail(to, subject, body string) error {
	fmt.Printf("Email to %s: %s - %s\n", to, subject, body)
	return nil
}
