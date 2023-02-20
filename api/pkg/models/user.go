package models

type User struct {
	Id          int    `json:"Id"`
	Username    string `json:"Username"`
	Email       string `json:"Email"`
	IsActivated bool   `json:"IsActivated"`
	Password    []byte `json:"-"`
}

type Token struct {
	Id           int    `json:"Id"`
	RefreshToken string `json:"RefreshToken"`
}

type SignInInput struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type SignUpInput struct {
	Email    string `json:"Email"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type AuthResponseModel struct {
	Status       int    `json:"Status"`
	Message      string `json:"Error"`
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
	User         User   `json:"User"`
}
