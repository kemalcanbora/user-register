package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"rollic/internal/database"
	"rollic/internal/models"
	"rollic/pkg/services"
	"rollic/pkg/utils"
	"strings"
	"time"
)

var userService services.UserService

func init() {
	client := database.MongoConnection()
	userService = &client
}

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(request.Body).Decode(&user)
	validate := validator.New()
	errVal := validate.Struct(user)

	if errVal != nil {
		utils.HTTPErrorHandler(response, "Email or Password field is empty!", http.StatusBadRequest)
		return
	}

	result := userService.FindUserWithEmail(user)
	passErr := utils.CheckPasswordHash(user.Password, result.Password)

	if passErr != true {
		log.Println(passErr)
		utils.HTTPErrorHandler(response, "Wrong Password!", http.StatusUnauthorized)
		return
	}

	jwtToken, err := utils.GenerateJWT(result)
	if err != nil {
		utils.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))
}

func Add(response http.ResponseWriter, request *http.Request) {
	var user models.User
	response.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		utils.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	validate := validator.New()
	errVal := validate.Struct(user)

	if errVal != nil {
		var fields []string
		for _, fieldErr := range errVal.(validator.ValidationErrors) {
			fields = append(fields, fieldErr.Field())
		}
		fieldsStr := strings.Join(fields, ",")
		switch true {
		case len(fields) == 1:
			utils.HTTPErrorHandler(response, `Bad request. `+fieldsStr+` field can't be empty!'`, http.StatusBadRequest)
			return
		case len(fields) > 1:
			utils.HTTPErrorHandler(response, `Bad request. `+fieldsStr+` fields can't be empty!'`, http.StatusBadRequest)
			return
		}
	}

	userCheck := userService.FindUserWithEmail(user)
	if userCheck.Email == user.Email {
		utils.HTTPErrorHandler(response, "This email is already registered!", http.StatusForbidden)
		return
	}

	user.Password = utils.GetHash([]byte(user.Password))
	user.CreatedTime = time.Now().Unix()
	result, _ := userService.AddUser(user, "user")
	UserId := result.InsertedID.(primitive.ObjectID).Hex()
	userResponse := models.UserResponse{ID: UserId, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName}
	err = json.NewEncoder(response).Encode(userResponse)
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	user := request.Context().Value("data")
	tmp, _ := json.Marshal(user)
	var AuthUser models.JwtUserAuth

	err := json.Unmarshal(tmp, &AuthUser)
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
		return
	}

	var userModel models.UserUpdate
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewDecoder(request.Body).Decode(&userModel)
	validate := validator.New()
	errVal := validate.Struct(userModel)
	if errVal != nil {
		fmt.Println(errVal)
	}

	attributeCheck := *new(models.UserUpdate)
	if userModel == attributeCheck {
		utils.HTTPErrorHandler(response, "Probably you give me wrong attribute", http.StatusBadRequest)
		return
	}
	filter := bson.D{{"email", AuthUser.Email}}
	update := bson.D{{"$set", userModel}}
	userModel.UpdatedTime = time.Now().Unix()
	_, err = userService.UpdateUser(filter, update, "user")
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
		return
	}
	response.Write([]byte(`{"message":"User updated successfully!"}`))
}

func Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	id := vars["id"]
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.D{{"_id", objID}}
	_, err = userService.DeleteUser(filter, "user")
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
		return
	}
	response.Write([]byte(`{"message":"User deleted successfully!"}`))
}

func Get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
		return
	}
	userId := bson.D{{"_id", objID}}
	user, err := userService.FindUser(userId, "user")
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
		return
	}
	marshal, err := bson.Marshal(user)
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadGateway)
		return
	}
	var userResponse models.UserResponse
	err = bson.Unmarshal(marshal, &userResponse)
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadGateway)
		return
	}
	err = json.NewEncoder(response).Encode(userResponse)
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadGateway)
		return
	}
}

func All(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	users := userService.GetAllUser("user")
	err := json.NewEncoder(response).Encode(users)
	if err != nil {
		utils.HTTPErrorHandler(response, err.Error(), http.StatusBadGateway)
		return
	}
}

func Welcome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(`{"message":"Welcome to Rollic API"}`))
}
