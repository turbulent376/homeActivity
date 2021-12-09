package auth

const (
	ErrCodeAuthInvalidEmail     = "AUTH-001"
	ErrCodeAuthInvalidPassword  = "AUTH-002"
	ErrCodeAuthUserIdEmpty      = "AUTH-003"
	ErrCodeAuthUserEmptyField   = "AUTH-004"
	ErrCodeAuthObjectNotFound   = "AUTH-005"
	ErrCodeAuthObjectDelete     = "AUTH-006"
	ErrCodeAuthObjectCreate     = "AUTH-007"
	ErrCodeAuthObjectUpdate     = "AUTH-008"
	ErrCodeAuthToken            = "AUTH-009"
	ErrCodeAuthStorageGetDb     = "AUTH-010"
	ErrCodeAuthStorageGetCache  = "AUTH-011"
	ErrCodeAuthStorageGetIds    = "AUTH-012"
	ErrCodeAuthStorageSetCache  = "AUTH-013"
	ErrCodeAuthStorageUpdate    = "AUTH-014"
	ErrCodeAuthSendNotification = "AUTH-015"
)
