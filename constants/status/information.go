package status

type ScriptInformationStatus int

const (
	SCRIPT_INFORMATION_SUCCESS           ScriptInformationStatus = 0
	SCRIPT_INFORMATION_DB_ERROR          ScriptInformationStatus = -1
	SCRIPT_INFORMATION_TITLE_EXIST_ERROR ScriptInformationStatus = -2
	SCRIPT_INFORMATION_PARAM_ERROR       ScriptInformationStatus = -3
)
