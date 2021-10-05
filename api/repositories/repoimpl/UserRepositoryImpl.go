package repoimpl

import (
	"go-architecture-mysql/api/entities"
	"go-architecture-mysql/api/payloads"
	"go-architecture-mysql/api/repositories"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func CreateUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) View() ([]entities.User, error) {
	row, err := r.DB.Raw("SELECT * FROM m_users").Rows()

	if err != nil {
		return nil, err
	}

	var users []entities.User
	for row.Next() {
		var user entities.User

		err := row.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (entities.User, error) {
	var user entities.User

	r.DB.Model(user).Where("username=?", username).Scan(&user)

	return user, nil
}

func (r *UserRepositoryImpl) FindById(userId uint) (entities.User, error) {
	var user entities.User

	r.DB.Model(user).Where("user_id=?", userId).Scan(&user)

	return user, nil
}

func (r *UserRepositoryImpl) Create(userReq payloads.CreateRequest) (bool, error) {
	err := r.DB.Exec("INSERT INTO m_users (first_name, last_name, username, password) VALUES(?, ?, ?, ?)",
			userReq.FirstName,
			userReq.LastName,
			userReq.Username,
			userReq.Password,
		).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepositoryImpl) Update(userReq payloads.CreateRequest, userId uint) (bool, error) {
	err := r.DB.Exec("UPDATE m_users SET first_name=?, last_name=?, username=?, password=? WHERE user_id=?",
		userReq.FirstName,
		userReq.LastName,
		userReq.Username,
		userReq.Password,
		userId,
	).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepositoryImpl) Delete(userId uint) (bool, error) {
	err := r.DB.Where("user_id=?", userId).Delete(&entities.User{}).Error
	
	if err != nil {
		return false, err
	}
	
	return true, nil
}