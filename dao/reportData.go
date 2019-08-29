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

// ListTodoList return result
func (d *DAO) ListTodoList() []*TodoList {
	result := make([]*TodoList, 0)
	d.SelectDB("todo").Limit(10).Find(&result)
	return result
}
