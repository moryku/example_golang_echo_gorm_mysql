package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"github.com/heroku/go-getting-started/app/config"
	"github.com/heroku/go-getting-started/app/db"
	"github.com/heroku/go-getting-started/app/model"
	"github.com/heroku/go-getting-started/app/util"
	"strconv"
)

func CreateUser(c echo.Context) error {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	anOddCondition := true
	if anOddCondition {
		panic("I am panicking")
	}

	var myDb = db.DbManager()
	config := config.GetConfig()

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm_password")

	if (password == "") {
		return util.PublishFailureMessage("Please enter password", c)
	}

	if (email == "") {
		return util.PublishFailureMessage("Please enter email", c)
	}

	if (name == "") {
		return util.PublishFailureMessage("Please enter name", c)
	}

	if (confirmPassword == "") {
		return echo.NewHTTPError(http.StatusBadRequest, "Please enter confirm password")
	}

	if (password != confirmPassword) {
		return util.PublishFailureMessage("Please enter password", c)
	}

	if (!util.ValidateEmail(email)) {
		return util.PublishFailureMessage("Email Not Valid", c)
	}

	if (CheckUserExists(email)) {
		return util.PublishFailureMessage("Email alraedy taken", c)
	}

	enc := util.EncryptString(password)
	user := model.User{Name: name, Email: email, Password: enc}
	myDb.Create(&user)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["email"] = user.Email

	t, err := token.SignedString([]byte(config.ENCRYCPTION_KEY)) // "secret" &gt;&gt; EncryptionKey
	if err != nil {
		return util.PublishFailureMessage("Error Encryption Key Token", c)
	}

	authentication := model.Authentication{}
	if myDb.First(&authentication, "user_id =?", user.ID).RecordNotFound() {
		myDb.Create(&model.Authentication{User: user, Token: t})
	} else {
		authentication.User = user
		authentication.Token = t
		myDb.Save(&authentication)
	}

	return util.PublishSuccessData(authentication, c)
}


func CheckUserExists(email string) bool {
	user1 := model.User{}
	db := db.DbManager()
	if db.Where(model.User{Email: email}).First(&user1).RecordNotFound() {
		return false
	}
	return true
}

func Login(c echo.Context) error {
	myDb := db.DbManager()
	email := c.FormValue("email")
	password := c.FormValue("password")

	if (password == "") {
		return echo.NewHTTPError(http.StatusBadRequest, "Please enter password")
	}

	if (email == "") {
		return echo.NewHTTPError(http.StatusBadRequest, "Please enter email")
	}
	var user model.User
	if myDb.First(&user, "email =?", email).RecordNotFound() {
		_error := util.CustomHTTPError{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "Email not found",
		}
		return c.JSONPretty(http.StatusBadGateway, _error, "  ")
	}

	configuration := config.GetConfig()
	hashPassword := util.EncryptString(password)

	if hashPassword == user.Password {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Name
		claims["email"] = user.Email
		claims["id"] = user.ID

		t, err := token.SignedString([]byte(configuration.ENCRYCPTION_KEY)) // "secret" &gt;&gt; EncryptionKey
		if err != nil {
			return err
		}

		authentication := model.Authentication{}
		if myDb.First(&authentication, "user_id =?", user.ID).RecordNotFound() {
			myDb.Create(&model.Authentication{User: user, Token: t})
		} else {
			authentication.User = user
			authentication.Token = t
			myDb.Save(&authentication)
		}
		authentication.User.Password = ""
		return util.PublishSuccessData(authentication, c)

	} else {
		_error := util.CustomHTTPError{
			Code:    http.StatusBadRequest,
			Message: "Invalid email & password",
			Status:  false,
		}
		return c.JSONPretty(http.StatusBadGateway, _error, "  ")
	}
}

func Logout(c echo.Context) error {
	var myDb = db.DbManager()
	tokenRequester := c.Get("user").(*jwt.Token)
	claims := tokenRequester.Claims.(jwt.MapClaims)
	fRequesterID := claims["id"].(float64)
	iRequesterID := int(fRequesterID)
	sRequesterID := strconv.Itoa(iRequesterID)

	requester := model.User{}
	if myDb.First(&requester, "id =?", sRequesterID).RecordNotFound() {
		return echo.ErrUnauthorized
	}

	authentication := model.Authentication{}
	if myDb.First(&authentication, "user_id =?", requester.ID).RecordNotFound() {
		return echo.ErrUnauthorized
	}
	myDb.Delete(&authentication)
	return c.String(http.StatusAccepted, "")
}
