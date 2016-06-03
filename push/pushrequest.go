package push

// PushRequest struct represents a message to be pushed to mobile.
type PushRequest struct {
	// push target
	AppKey      string
	Target      string
	TargetValue string
	DeviceType  int

	// push configuration
	Type    int
	Title   string
	Body    string
	Summary string

	// IOS specific
	IOSBadge         string
	IOSMusic         string
	IOSExtParameters string
	ApnsEnv          string
	Remind           bool

	// Android specific
	AndroidOpenType      int
	AndroidOpenURL       string
	AndroidExtParameters string
	AndroidMusic         string
	AndroidActivity      string

	// Message control
	PushTime     string
	StoreOffline bool
	ExpireTime   string
	BatchNumber  string
}
