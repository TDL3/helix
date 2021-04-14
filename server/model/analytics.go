package model

type AnalyticsNewItems []struct {
	ItemsCount            []uint8 `gorm:"column:new_items;"`
	Months   	          []uint8 `gorm:"column:months;"`
	Year				  []uint8 `gorm:"column:year;"`
}

type AnalyticsRecoveredItems []struct {
	ItemsCount            []uint8 `gorm:"column:recovered_items;"`
	Months   	          []uint8 `gorm:"column:months;"`
	Year				  []uint8 `gorm:"column:year;"`
}

type AnalyticsNewItemsResp []struct {
	ItemsCount            string
	Months   	          string
	Year				  string
}

type AnalyticsRecoveredItemsResp []struct {
	ItemsCount            string
	Months   	          string
	Year				  string
}

type AnalyticsResp struct {
	NewItems struct {
		ItemsCount            []string `json:"itemsCount"`
		MonthsCount	          []string `json:"monthsCount"`
		Year				  []string `json:"year"`
	} `json:"newItems"`
	RecoveredItems struct {
		ItemsCount            []string `json:"itemsCount"`
		MonthsCount	          []string `json:"monthsCount"`
		Year				  []string `json:"year"`
	} `json:"recoveredItems"`
}