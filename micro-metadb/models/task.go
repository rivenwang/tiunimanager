package models

import (
	"gorm.io/gorm"
)

type FlowDO struct {
	Data
	FlowName    string
	StatusAlias string
}

func (do FlowDO) TableName() string {
	return "flows"
}

type TaskDO struct {
	Data
	ParentType      int8		`gorm:"default:0"`
	ParentId 		string
	TaskName 		string
	TaskReturnType 	string
	Parameters 		string
	Result 			string
}

func (do TaskDO) TableName() string {
	return "tasks"
}

func (do *TaskDO) BeforeCreate(tx *gorm.DB) (err error) {
	do.Status = 0
	return nil
}

func (do *FlowDO) BeforeCreate(tx *gorm.DB) (err error) {
	do.Status = 0
	return nil
}

func CreateFlow(flowName string, statusAlias string, bizId string) (flow *FlowDO, err error) {
	flow = &FlowDO{
		FlowName: flowName,
		StatusAlias: statusAlias,
		Data: Data{
			BizId: bizId,
		},
	}
	err = MetaDB.Create(&flow).Error
	return
}

func CreateTask(parentType int8, parentId string, taskName, bizId string, taskReturnType string, parameters, result string) (task *TaskDO, err error) {
	task = &TaskDO{
		ParentType: parentType,
		ParentId: parentId,
		TaskName: taskName,
		TaskReturnType: taskReturnType,

		Parameters: parameters,
		Result: result,
		Data: Data{
			BizId: bizId,
		},
	}
	err = MetaDB.Create(&task).Error
	return
}

func FetchFlow(id uint) (flow FlowDO, err error) {
	err = MetaDB.Find(&flow, id).Error
	return
}

func FetchFlowDetail(id uint) (flow *FlowDO, tasks []*TaskDO, err error) {
	flow = &FlowDO{}
	err = MetaDB.Find(flow, id).Error

	if err != nil {
		return
	}
	err = MetaDB.Where("parent_type = 0 and parent_id = ?", id).Find(&tasks).Error
	return
}

func FetchTask(id uint) (task TaskDO, err error) {
	err = MetaDB.Find(&task, id).Error
	return
}

func QueryTask(bizId string, taskType string) (tasks []TaskDO, err error) {
	err = MetaDB.Find(&tasks, "biz_id = ?" ,bizId).Error
	return
}

func UpdateFlow(flow FlowDO) (FlowDO, error) {
	err := MetaDB.Save(&flow).Error
	if err != nil {
		return flow, err
	}
	return flow,nil
}

func BatchSaveTasks(tasks []*TaskDO) (returnTasks []*TaskDO, err error) {
	err = MetaDB.Save(tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks,nil
}

func UpdateFlowAndTasks(flow FlowDO, tasks []TaskDO) (FlowDO, []TaskDO, error) {
	err := MetaDB.Save(&flow).Error

	if err != nil {
		return flow, tasks, err
	}

	err = MetaDB.Save(&tasks).Error

	return flow, tasks, nil
}

func UpdateTask(task TaskDO)  (returnTask TaskDO, err error) {
	return task, MetaDB.Save(&task).Error
}
