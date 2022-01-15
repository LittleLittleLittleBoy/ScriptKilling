package service

import (
	"log"
	"time"

	"github.com/littlelittlelittleboy/scriptkilling/common/vo"
	"github.com/littlelittlelittleboy/scriptkilling/constants/status"
	"github.com/littlelittlelittleboy/scriptkilling/db"
	"github.com/littlelittlelittleboy/scriptkilling/models"
)

type ScriptInformationService struct{}

func (service *ScriptInformationService) Add(information vo.CreateInformationRequest) (int64, status.ScriptInformationStatus) {
	scriptInformation := models.ScriptInformation{
		Type:       information.Type,
		Title:      information.Title,
		Stage:      information.Stage,
		Url:        information.URL,
		Rid:        information.Rid,
		Sid:        information.Sid,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	count, err := db.Engine.Insert(&scriptInformation)

	if err != nil {
		log.Fatal(err)
		return -1, status.SCRIPT_INFORMATION_DB_ERROR
	}

	if count == 0 {
		return 0, status.SCRIPT_INFORMATION_DB_ERROR
	}

	return count, status.SCRIPT_INFORMATION_SUCCESS
}

func (service *ScriptInformationService) List(sid int) ([]models.ScriptInformation, status.ScriptInformationStatus) {
	scriptInformationList := make([]models.ScriptInformation, 0)
	err := db.Engine.SQL("select * from `script_information` where sid = ?", sid).Find(&scriptInformationList)

	if err != nil {
		log.Fatalf("select script information from db error, %s", err)
		return scriptInformationList, status.SCRIPT_INFORMATION_DB_ERROR
	}

	return scriptInformationList, status.SCRIPT_INFORMATION_SUCCESS
}
