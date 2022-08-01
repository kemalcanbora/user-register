package models

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CreatedTime int64  `json:"created_time"`
	UpdatedTime int64  `json:"updated_time"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required,max=20,min=3"`
}

type UserResponse struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name,omitempty" bson:"firstname"`
	LastName  string `json:"last_name,omitempty" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
}

type UserUpdate struct {
	FirstName   *string `json:"first_name,omitempty" bson:"firstname,omitempty"`
	LastName    *string `json:"last_name,omitempty" bson:"lastname,omitempty"`
	Email       *string `json:"email,omitempty" bson:"email,omitempty"`
	Password    *string `json:"password,omitempty" bson:"password,omitempty"`
	UpdatedTime int64   `json:"updated_time"`
}
