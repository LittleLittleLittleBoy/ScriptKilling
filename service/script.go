package service

import (
	"log"
	"time"

	"github.com/littlelittlelittleboy/scriptkilling/common/vo"
	"github.com/littlelittlelittleboy/scriptkilling/constants"
	"github.com/littlelittlelittleboy/scriptkilling/db"
	"github.com/littlelittlelittleboy/scriptkilling/models"
)

type ScriptInfoService struct{}

func (service *ScriptInfoService) Add(scriptInfo vo.CreateScriptRequest) (int64, constants.ScriptInfoStatus) {
	scriptModel := &models.ScriptInfo{
		Title:      scriptInfo.Title,
		RoleNum:    scriptInfo.RoleNum,
		Cover:      "",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}

	count, err := db.Engine.Insert(&scriptModel)

	if err != nil {
		log.Fatal(err)
		return -1, constants.SCRPIT_INFO_DB_ERROR
	}

	if count == 0 {
		return 0, constants.SCRPIT_INFO_TITLE_EXIST_ERROR
	}

	return count, constants.SCRIPT_INFO_SUCCESS
}

func (service *ScriptInfoService) List() ([]models.ScriptInfo, constants.ScriptInfoStatus) {
	scriptInfoLst := []models.ScriptInfo{}
	err := db.Engine.SQL("select * from script_info").Find(&scriptInfoLst)

	if err != nil {
		log.Fatal(err)
		return scriptInfoLst, constants.SCRPIT_INFO_DB_ERROR
	}

	return scriptInfoLst, constants.SCRIPT_INFO_SUCCESS
}
