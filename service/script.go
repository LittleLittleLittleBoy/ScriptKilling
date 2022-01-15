package service

import (
	"log"
	"time"

	"github.com/littlelittlelittleboy/scriptkilling/common/vo"
	"github.com/littlelittlelittleboy/scriptkilling/constants/status"
	"github.com/littlelittlelittleboy/scriptkilling/db"
	"github.com/littlelittlelittleboy/scriptkilling/models"
)

type ScriptInfoService struct{}

func (service *ScriptInfoService) Add(scriptInfo vo.CreateScriptRequest) (int64, status.ScriptInfoStatus) {
	scriptModel := models.ScriptInfo{
		Title:      scriptInfo.Title,
		RoleNum:    scriptInfo.RoleNum,
		Cover:      "",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	count, err := db.Engine.Insert(&scriptModel)

	if err != nil {
		log.Fatal(err)
		return -1, status.SCRPIT_INFO_DB_ERROR
	}

	if count == 0 {
		return 0, status.SCRPIT_INFO_TITLE_EXIST_ERROR
	}

	return count, status.SCRIPT_INFO_SUCCESS
}

func (service *ScriptInfoService) List() ([]models.ScriptInfo, status.ScriptInfoStatus) {
	scriptInfoList := make([]models.ScriptInfo, 0)
	err := db.Engine.SQL("select * from `script_info`").Find(&scriptInfoList)

	if err != nil {
		log.Fatalf("select script info from db error, %s", err)
		return scriptInfoList, status.SCRPIT_INFO_DB_ERROR
	}

	return scriptInfoList, status.SCRIPT_INFO_SUCCESS
}
