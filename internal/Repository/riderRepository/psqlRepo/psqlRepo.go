package psqlRepo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"sls/internal/Repository/riderRepository"
	"sls/internal/entity/riderEntity"
	"time"
)

type psql struct {
	conn *pgx.Conn
}

func (p *psql) GetById(id uuid.UUID) (*riderEntity.CreateRiderRes, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	var res riderEntity.CreateRiderRes

	queryStmt := fmt.Sprintf("SELECT rider_id, first_name, last_name, email, phone, password, profile_picture, "+
		"verification_status, account_status FROM rider_table WHERE rider_id='%v';", id)

	row := p.conn.QueryRow(ctx, queryStmt)
	err := row.Scan(&res.RiderId, &res.FirstName, &res.LastName, &res.Email, &res.Phone, &res.Password, &res.ProfilePicture,
		&res.VerificationStatus, &res.AccountStatus)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *psql) UpdateProfile(req *riderEntity.UpdateRiderReq) (*riderEntity.UpdateRiderReq, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancelFunc()

	updtStmt := fmt.Sprintf("UPDATE rider_table SET first_name = '%v', last_name ='%v', email='%v', phone='%v', "+
		"profile_picture='%v', updated_at='%v';", req.FirstName, req.LastName, req.Email, req.Phone,
		req.ProfilePicture, req.UpdatedAt)
	_, err := p.conn.Exec(ctx, updtStmt)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (p *psql) ChangePassword(id uuid.UUID, password string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancelFunc()

	delStmt := fmt.Sprintf("UPDATE rider_table SET password = '%v' WHERE id='%v' ;", password, id)
	_, err := p.conn.Exec(ctx, delStmt)
	if err != nil {
		return err
	}

	return nil
}

func (p *psql) DeleteAccount(req *riderEntity.DeleteUser) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancelFunc()
	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	deleteStmt := fmt.Sprintf("DELETE FROM guarantors_table WHERE guarantor_id = '%v'; ", req.GuarantorId)
	_, err = tx.Exec(ctx, deleteStmt)
	if err != nil {
		return err
	}

	deleteStmt2 := fmt.Sprintf("DELETE FROM rider_table WHERE rider_id is ='%v'; ", req.RiderId)
	_, err = tx.Exec(ctx, deleteStmt2)
	if err != nil {
		return err
	}

	return nil
}

func (p *psql) GetByEmail(email string) (*riderEntity.CreateRiderRes, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	var res riderEntity.CreateRiderRes

	queryStmt := fmt.Sprintf("SELECT rider_id, first_name, last_name, email, phone, password, profile_picture, "+
		"verification_status, account_status FROM rider_table WHERE email='%v';", email)

	row := p.conn.QueryRow(ctx, queryStmt)
	err := row.Scan(&res.RiderId, &res.FirstName, &res.LastName, &res.Email, &res.Phone, &res.Password, &res.ProfilePicture,
		&res.VerificationStatus, &res.AccountStatus)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *psql) GetByPhone(phone string) (*riderEntity.CreateRiderRes, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	var res riderEntity.CreateRiderRes

	queryStmt := fmt.Sprintf("SELECT rider_id, first_name, last_name, email, phone, password, profile_picture, "+
		"verification_status, account_status FROM rider_table WHERE phone='%v';", phone)

	row := p.conn.QueryRow(ctx, queryStmt)
	err := row.Scan(&res.RiderId, &res.FirstName, &res.LastName, &res.Email, &res.Phone, &res.Password, &res.ProfilePicture,
		&res.VerificationStatus, &res.AccountStatus)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *psql) Persist(req *riderEntity.CreateRiderReq) (*riderEntity.CreateRiderReq, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	tx, err := p.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	insertStmt := fmt.Sprintf("INSERT INTO guarantors_table (guarantor_id, rider_id, first_name, last_name, email, "+
		"phone, guarantor_address, guarantor_identification) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')",
		req.Guarantor.GuarantorId, req.RiderId, req.Guarantor.FirstName, req.Guarantor.LastName,
		req.Guarantor.Email, req.Guarantor.Phone, req.Guarantor.GuarantorAddress, req.Guarantor.GuarantorJob)
	_, err = tx.Exec(ctx, insertStmt)
	if err != nil {
		fmt.Println("here")
		return nil, err
	}

	insertStmt2 := fmt.Sprintf("INSERT INTO rider_table (rider_id, guarantor_id, first_name, last_name, email, phone,"+
		" password, DOB, gender, marital_status, education_level, residential_address, driver_license, identity_card, "+
		"verification_status, account_status, created_at) VALUES('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', "+
		"'%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')", req.RiderId, req.Guarantor.GuarantorId,
		req.FirstName, req.LastName, req.Email, req.Phone, req.Password, req.DOB, req.Gender, req.MaritalStatus,
		req.EducationLevel,
		req.ResidentialAddress, req.DriverLicense, req.IdentityCard, req.VerificationStatus, req.AccountStatus, req.CreatedAt)

	_, err = tx.Exec(ctx, insertStmt2)
	if err != nil {
		fmt.Println("here 2")
		return nil, err
	}

	return req, nil
}

func NewPsql(conn *pgx.Conn) (riderRepository.RiderRepo, error) {
	return &psql{conn: conn}, nil
}
