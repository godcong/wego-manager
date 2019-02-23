package model

// UserInformation ...
type UserInformation struct {
	Model         `xorm:"extends" json:",inline"`
	UserID        string `xorm:"notnull default('') unique user_id" json:"user_id"`
	Email         string `json:"email" xorm:"email"`                     //邮件
	Mobile1       string `json:"mobile" xorm:"mobile1"`                  //移动电话
	Mobile2       string `json:"mobile" xorm:"mobile2"`                  //移动电话
	Mobile3       string `json:"mobile" xorm:"mobile3"`                  //移动电话
	Mobile4       string `json:"mobile" xorm:"mobile4"`                  //移动电话
	Mobile5       string `json:"mobile" xorm:"mobile5"`                  //移动电话
	IDCardFacade  string `json:"id_card_facade" xorm:"id_card_facade"`   //身份证(正)
	IDCardObverse string `json:"id_card_obverse" xorm:"id_card_obverse"` //身份证(反)
}

// Get ...
func (obj *UserInformation) Get() (bool, error) {
	return Get(nil, obj)
}
