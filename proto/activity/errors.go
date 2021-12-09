package timesheet

const (
	ErrCodeTimesheetIdEmpty            = "TIMESHEET-001"
	ErrCodeTimesheetOwnerEmpty         = "TIMESHEET-002"
	ErrCodeTimesheetTimesheetIdEmpty   = "TIMESHEET-004"
	ErrCodeTimesheetSubjectEmpty       = "TIMESHEET-005"
	ErrCodeTimesheetWeekDayEmpty       = "TIMESHEET-006"
	ErrCodeTimesheetDescriptionEmpty   = "TIMESHEET-007"
	ErrCodeTimesheetNotFound           = "TIMESHEET-008"
	ErrCodeTimesheetDeleted            = "TIMESHEET-009"
	ErrCodeTimesheetStorageCreate      = "TIMESHEET-010"
	ErrCodeTimesheetStorageGetDb       = "TIMESHEET-011"
	ErrCodeTimesheetStorageGetCache    = "TIMESHEET-012"
	ErrCodeTimesheetStorageGetIds      = "TIMESHEET-013"
	ErrCodeTimesheetStorageSetCache    = "TIMESHEET-014"
	ErrCodeTimesheetStorageUpdate      = "TIMESHEET-015"
	ErrCodeTimesheetIndexSearch        = "TIMESHEET-016"
	ErrCodeTimesheetTimeEmpty          = "TIMESHEET-017"
	ErrCodeTimesheetStorageSearchCache = "TIMESHEET-018"
	ErrCodeTimesheetStorageSearchDb    = "TIMESHEET-019"
)
