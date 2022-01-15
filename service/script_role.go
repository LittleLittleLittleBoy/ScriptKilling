package service

import (
	"log"
	"time"

	"github.com/littlelittlelittleboy/scriptkilling/common/vo"
	"github.com/littlelittlelittleboy/scriptkilling/constants"
	"github.com/littlelittlelittleboy/scriptkilling/db"
	"github.com/littlelittlelittleboy/scriptkilling/models"
)

type ScriptRoleService struct{}

func (service *ScriptRoleService) Add(roleInfo vo.CreateRoleRequest) (int64, constants.ScriptRoleStatus) {
	scriptModel := models.ScriptRole{
		Sid:        roleInfo.Sid,
		Name:       roleInfo.Name,
		Cover:      roleInfo.Cover,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	count, err := db.Engine.Insert(&scriptModel)

	if err != nil {
		log.Fatal(err)
		return -1, constants.SCRPIT_ROLE_DB_ERROR
	}

	if count == 0 {
		return 0, constants.SCRPIT_ROLE_DB_ERROR
	}

	return count, constants.SCRIPT_ROLE_SUCCESS
}

func (service *ScriptRoleService) List(sid int) ([]models.ScriptRole, constants.ScriptRoleStatus) {
	scriptRoleList := make([]models.ScriptRole, 0)
	err := db.Engine.SQL("select * from `script_role` where sid = ?", sid).Find(&scriptRoleList)

	if err != nil {
		log.Fatalf("select script role from db error, %s", err)
		return scriptRoleList, constants.SCRPIT_ROLE_DB_ERROR
	}

	return scriptRoleList, constants.SCRIPT_ROLE_SUCCESS
}
