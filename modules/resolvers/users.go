package resolvers

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/graphql-go/graphql"
	"github.com/Arkangel12/curso/graphql-gorm/mail"
	"github.com/Arkangel12/curso/graphql-gorm/db"
	"github.com/Arkangel12/curso/graphql-gorm/models"
)

// Me => Returns the current user logged
func Me(params graphql.ResolveParams) (interface{}, error) {
	// Bring currentUser from context
	currentUser, ok := params.Context.Value("currentUser").(models.User)
	fmt.Println("step 1")
	// If user is anonimous, return error
	if ok && currentUser.Username == "invited" {
		return nil, errors.New("User must be logged")
	}
	fmt.Println("step 2")
	return &currentUser, nil
}

// CreateUser => Creates new user account in database
func CreateUser(params graphql.ResolveParams) (interface{}, error) {
	email, _ := params.Args["email"].(string)
	fullname, _ := params.Args["full_name"].(string)
	username, _ := params.Args["username"].(string)
	password, _ := params.Args["password"].(string)

	user := models.User{}

	db := db.GetConnection()

	// Find and verify if username already exists
	db.Where(&models.User{Username:username}).First(&user)

	if user.Username == username {
		return nil, errors.New("Username is taken")
	}

	// hashes password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Println(err.Error())
	}

	password = string(hashedPassword)

	user = models.User{
		Username: username,
		Password: password,
		Email:    email,
		FullName: fullname,
	}

	// Create new user in database
	db.Create(&user)

	return &user, nil
}

// GetUserByID get user by ID
func GetUserByID(id int) (*models.User, error) {

	user := models.User{}

	db := db.GetConnection()

	db.Where(&models.User{ID: id}).First(&user)

	return &user, nil
}

// AuthenticateUser => Allow user to log in
func AuthenticateUser(params graphql.ResolveParams) (interface{}, error) {
	username, _ := params.Args["username"].(string)
	password, _ := params.Args["password"].(string)

	user := models.User{}

	db := db.GetConnection()
	db.Where(&models.User{Username: username}).First(&user)

	//Password verification
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New("Hash comparation failed!")
	}

	db.Save(&user)

	return &user, nil
}

func PasswordRecovery(params graphql.ResolveParams) (interface{}, error) {
	emailRecover, _ := params.Args["email"].(string)

	user := models.User{}

	db := db.GetConnection()
	err := db.Where(&models.User{Email: emailRecover}).First(&user)

	if err == nil {
		return nil, errors.New("Email not found!")
	}

	var email mail.Mail

	email.Name = user.Username
	email.Email = user.Email
	email.Subject = "Password Recovery"
	email.Body = "Email body!"

	mail.Send(email)

	return &user, nil
}
