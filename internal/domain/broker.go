package domain

type BrokerPort interface {
	Subscribe(string, func(NotificationMsg))
	// Publish()
}
