package userctrl

type RegisterReq struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ImageName string `json:"image_name"`

}

type LoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

