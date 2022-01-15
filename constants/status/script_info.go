package status

type ScriptInfoStatus int

const (
	SCRIPT_INFO_SUCCESS           ScriptInfoStatus = 0
	SCRPIT_INFO_DB_ERROR          ScriptInfoStatus = -1
	SCRPIT_INFO_TITLE_EXIST_ERROR ScriptInfoStatus = -2
	SCRPIT_INFO_PARAM_ERROR       ScriptInfoStatus = -3
)
