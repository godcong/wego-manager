package model

// Spread ...
type Spread struct {
	Model `xorm:"extends"`
	//ActivityID    string `json:"activity_id"`
	//Code          string `json:"code"`
	UserID        string `json:"user_id"`
	ParentUserID1 string `json:"parent_user_id_1 comment(上1级)"`
	ParentUserID2 string `json:"parent_user_id_2 comment(上2级)"`
	ParentUserID3 string `json:"parent_user_id_3 comment(上3级)"`
	ParentUserID4 string `json:"parent_user_id_4 comment(上4级)"`
	ParentUserID5 string `json:"parent_user_id_5 comment(上5级)"`
	ParentUserID6 string `json:"parent_user_id_6 comment(上6级)"`
	ParentUserID7 string `json:"parent_user_id_7 comment(上7级)"`
	ParentUserID8 string `json:"parent_user_id_8 comment(上8级)"`
	ParentUserID9 string `json:"parent_user_id_9 comment(上9级)"`
}
