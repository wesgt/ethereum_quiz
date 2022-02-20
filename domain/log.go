package domain

type Log struct {
	Index  uint   `json:"index" gorm:"column:log_index;primary_key"`
	Data   string `json:"data"`
	TxHash string `gorm:"size:256;primary_key"`
}
