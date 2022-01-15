package constants

type ScriptRoleStatus int

const (
	SCRIPT_ROLE_SUCCESS          ScriptRoleStatus = 0
	SCRPIT_ROLE_DB_ERROR         ScriptRoleStatus = -1
	SCRPIT_ROLE_NAME_EXIST_ERROR ScriptRoleStatus = -2
	SCRPIT_ROLE_PARAM_ERROR      ScriptRoleStatus = -3
)
