package dao

type ReportData struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	Timestamp uint64 `json:"timestamp"`
	Type      int    `json:"type"`
	Value     int    `json:"value"`
	UUID      string `json:"uuid"`
}

func (ReportData) TableName() string {
	return "report_data_tab"
}

func (d *DAO) ListReportData() []*ReportData {
	result := make([]*ReportData, 0)
	d.db.Limit(10).Find(&result)
	return result
}

func (d *DAO) CreateReportData(timestamp uint64, dataType, value int, uuid string) error {
	reportData := &ReportData{
		Timestamp: timestamp,
		Type:      dataType,
		Value:     value,
		UUID:      uuid,
	}
	if err := d.db.Create(&reportData).Error; err != nil {
		return err
	}
	return nil
}
