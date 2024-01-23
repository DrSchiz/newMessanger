package user

type RequestLoginUser struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type RequestRegisterUserStepOne struct {
	Email string `json:"email" validate:"required,email"`
}

type RequestRegisterUserStepTwo struct {
	RegisterId       string `json:"register_id" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	VerificationCode string `json:"verification_code" validate:"required"`
}

type RequestRegisterUserStepThree struct {
	RegisterId string `json:"register_id" validate:"required"`
	Username   string `json:"username" validate:"required,min=3,max=20"`
	Firstname  string `json:"firstname" validate:"min=0,max=20"`
	Lastname   string `json:"lastname" validate:"min=0,max=20"`
}

type RequestRegisterUserStepFour struct {
	RegisterId string `json:"register_id" validate:"required"`
	Password   string `json:"password" validate:"required,min=6,max=50"`
}

type RequestSetPassword struct {
	Password       string `json:"password" validate:"required,min=6,max=50"`
	RepeatPassword string `json:"repeat_password" validate:"required,min=6,max=50"`
	UserId         string `json:"user_id" validate:"required"`
}
