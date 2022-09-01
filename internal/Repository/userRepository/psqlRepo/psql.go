package psqlRepo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"sls/internal/Repository/userRepository"
	"sls/internal/entity/userEntity"
	"time"
)

type psqlRepository struct {
	conn *pgx.Conn
}

func (p *psqlRepository) GetById(id uuid.UUID) (*userEntity.UserAccess, error) {
	var userAccess userEntity.UserAccess

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	getStmt := fmt.Sprintf("SELECT firstname, lastname, email, profile_picture, phone, "+
		"password  FROM user_table WHERE user_id = '%v' AND user_status= 'active' ;", id)
	queryRes := p.conn.QueryRow(ctx, getStmt)

	err := queryRes.Scan(&userAccess.FirstName, &userAccess.LastName, &userAccess.Email, &userAccess.ProfilePicture,
		&userAccess.Phone, &userAccess.Password)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &userAccess, nil
}

func (p *psqlRepository) Persist(user *userEntity.CreateUser) (*userEntity.CreateUser, error) {
	ctx := context.Background()
	insertStmt := fmt.Sprintf("INSERT INTO user_table (user_id, firstname, lastname, email, profile_picture, phone, "+
		"password, default_location, created_at, user_status) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v',"+
		" '%v', '%v');",
		user.UserId, user.FirstName, user.LastName, user.Email, user.ProfilePicture, user.Phone,
		user.Password, user.DefaultLocation, user.CreatedAt, user.UserStatus)

	_, err := p.conn.Exec(ctx, insertStmt)
	if err != nil {
		fmt.Println("insider error place")
		return nil, err
	}
	return user, nil

}

func (p *psqlRepository) GetByEmail(email string) (*userEntity.UserAccess, error) {

	var userAccess userEntity.UserAccess

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	getStmt := fmt.Sprintf("SELECT firstname, lastname, email, profile_picture, phone, "+
		"password, user_id  FROM user_table WHERE email = '%v' AND user_status= 'active' ;", email)
	queryRes := p.conn.QueryRow(ctx, getStmt)

	err := queryRes.Scan(&userAccess.FirstName, &userAccess.LastName, &userAccess.Email, &userAccess.ProfilePicture,
		&userAccess.Phone, &userAccess.Password, &userAccess.UserId)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &userAccess, nil
}

func (p *psqlRepository) GetByPhone(phone string) (*userEntity.UserAccess, error) {
	var userAccess userEntity.UserAccess

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	getStmt := fmt.Sprintf("SELECT firstname, lastname, password, email, phone, "+
		"profile_picture, user_id FROM user_table where phone='%v' AND user_status ='active'",
		phone)
	queryRes := p.conn.QueryRow(ctx, getStmt)

	err := queryRes.Scan(&userAccess.FirstName, &userAccess.LastName, &userAccess.Password, &userAccess.Email,
		&userAccess.Phone, &userAccess.ProfilePicture, &userAccess.UserId)
	if err != nil {
		return nil, err
	}

	return &userAccess, nil
}

func (p *psqlRepository) ChangePassword(id uuid.UUID, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	updtStmt := fmt.Sprintf("UPDATE user_table SET password='%v', updated_at ='%v' WHERE user_id='%v';", password,
		time.Now(),
		id)
	_, err := p.conn.Exec(ctx, updtStmt)
	if err != nil {
		return err
	}
	return nil
}

func (p *psqlRepository) UpdateProfile(req *userEntity.UpdateUserReq) (*userEntity.UpdateUserReq, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()
	updateStmt := fmt.Sprintf("UPDATE user_table SET email='%v', phone='%v', firstname='%v', lastname='%v', "+
		"profile_picture='%v',updated_at='%v' WHERE user_id = '%v'",
		req.Email, req.Phone, req.FirstName, req.LastName, req.ProfilePicture, time.Now(), req.UserId)

	_, err := p.conn.Exec(ctx, updateStmt)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (p *psqlRepository) DeactivateAccount(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	deleteStmt := fmt.Sprintf("DELETE FROM user_table WHERE user_id='%v' ;", id)
	_, err := p.conn.Exec(ctx, deleteStmt)
	if err != nil {
		return err
	}
	return nil
}

func NewPsqlRepository(conn *pgx.Conn) userRepository.UserRepo {
	return &psqlRepository{conn: conn}
}
