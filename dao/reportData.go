package dao

// ReportData struct defination
type ReportData struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	Timestamp uint64 `json:"timestamp"`
	Type      int    `json:"type"`
	Value     int    `json:"value"`
	UUID      string `json:"uuid"`
}

// TodoList struct defination
type TodoList struct {
	ID          string `gorm:"primary_key" json:"id"`
	Text        string `json:"text"`
	IsCompleted int    `json:"isCompleted"`
}

// TableName return TableName
func (ReportData) TableName() string {
	return "report_data_tab"
}

// TableName Gorm 默认调用,取表名
func (TodoList) TableName() string {
	return "todos_database"
}

// ListReportData return result
func (d *DAO) ListReportData() []*ReportData {
	result := make([]*ReportData, 0)
	d.SelectDB("database").Limit(10).Find(&result)
	return result
}

// CreateReportData create a record
func (d *DAO) CreateReportData(timestamp uint64, dataType, value int, uuid string) error {
	reportData := &ReportData{
		Timestamp: timestamp,
		Type:      dataType,
		Value:     value,
		UUID:      uuid,
	}
	if err := d.SelectDB("database").Create(&reportData).Error; err != nil {
		return err
	}
	return nil
}

// CreateTodoRecord create a record
func (d *DAO) CreateTodoRecord(id string, text string, isCompleted int) error {
	data := &TodoList{
		ID:          id,
		Text:        text,
		IsCompleted: isCompleted,
	}
	if err := d.SelectDB("todo").Create(&data).Error; err != nil {
		return err
	}
	return nil
}

// QueryTodoList return result
func (d *DAO) QueryTodoList() []*TodoList {
	result := make([]*TodoList, 0)
	d.SelectDB("todo").Limit(100).Find(&result)
	return result
}

// DeleteTodo todo by id
func (d *DAO) DeleteTodo(id string) error {
	if err := d.SelectDB("todo").Table("todos_database").Where("id=?", id).Delete(&TodoList{}).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTodo a record
func (d *DAO) UpdateTodo(id string, isCompleted int) error {
	if err := d.SelectDB("todo").Table("todos_database").Where("id=?", id).Update(map[string]interface{}{
		"is_completed": isCompleted,
	}).Error; err != nil {
		return err
	}
	return nil
}
