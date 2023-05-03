package repository

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

//-----------------------------UserSignUp----------------------

func (c *userDatabase) UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error) {
	var userData res.UserData
	insertQuery := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4) 
					RETURNING id,name,email,mobile`
	err := c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userData).Error
	return userData, err
}

//-----------------------------UserLogin------------------------

func (c *userDatabase) UserLogin(ctx context.Context, email string) (domain.Users, error) {
	var userData domain.Users
	err := c.DB.Raw("SELECT * FROM users WHERE email=?", email).Scan(&userData).Error
	return userData, err
}
func (c *userDatabase) IsSignIn(phone string) (bool, error) {
	query := "select exists(select 1 from users where mobie=?)"
	var IsSignIn bool
	err := c.DB.Raw(query, phone).Scan(&IsSignIn).Error
	return IsSignIn, err
}

func (c *userDatabase) OtpLogin(phone string) (int, error) {
	var id int
	query := "SELECT id from users where mobile=?"
	err := c.DB.Raw(query, phone).Scan(&id).Error
	return id, err
}
func (c *userDatabase) AddAddress(id int, address req.Address) error {
	//isDefault
	if address.IsDefault {
		changeAddress := `UPDATE addresses SET is_default=$1 WHERE users_id=$2 AND is_Default=$3`
		err := c.DB.Exec(changeAddress, false, id, true)
		if err != nil {
			fmt.Println("SET1")
		}
	}
	query := `INSERT INTO addresses (users_id,house_number,street,city, district,landmark,pincode,is_default)
	VALUES($1,$2,$3,$4,$5,$6,&7,&8)`
	err := c.DB.Raw(query, address.Id, address.House_number, address.Street, address.City, address.District, address.Landmark, address.Pincode, address.IsDefault).Error
	return err
}

func (c *userDatabase) ViewProfile(id int) (res.UserData, error) {
	var profile res.UserData
	findProfile := `SELECT name,email,mobile FROM users WHERE id=$1`
	err := c.DB.Raw(findProfile, id).Scan(&profile).Error
	fmt.Println(profile)
	return profile, err
}
func (c *userDatabase) EditProfile(id int, updatingDetails req.UserReq) (res.UserData, error) {
	var profile res.UserData
	UpdatedQuery := `UPDATE users SET name=$1,email=$2,mobile=$3 WHERE id=$4 RETURNING name,email,mobile`
	err := c.DB.Raw(UpdatedQuery, updatingDetails.Name, updatingDetails.Email, updatingDetails.Mobile, id).Scan(&profile).Error
	return profile, err
}
