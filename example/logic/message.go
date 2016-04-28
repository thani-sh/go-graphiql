package logic

import "sync"

var (
	message    = ""
	messageMtx = sync.RWMutex{}
)

// GetMessage ...
func GetMessage() string {
	messageMtx.RLock()
	defer messageMtx.RUnlock()
	return message
}

// SetMessage ...
func SetMessage(msg string) {
	messageMtx.Lock()
	defer messageMtx.Unlock()
	message = msg
}
