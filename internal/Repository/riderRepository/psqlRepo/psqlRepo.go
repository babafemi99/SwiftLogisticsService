package psqlRepo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"sls/internal/Repository/riderRepository"
	"sls/internal/entity/riderEntity"
	"time"
)

type psql struct {
	conn *pgx.Conn
}

func (p *psql) GetById(id string) (*riderEntity.CreateRiderRes, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	var res riderEntity.CreateRiderRes

	queryStmt := fmt.Sprintf(`SELECT rider_id, first_name, last_name, email, phone_number, password, profile_picture, account_status FROM riders_table WHERE rider_id='%v';`, id)

	row := p.conn.QueryRow(ctx, queryStmt)
	err := row.Scan(&res.RiderId, &res.FirstName, &res.LastName, &res.Email, &res.Phone, &res.Password, &res.ProfilePicture, &res.AccountStatus)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *psql) UpdateProfile(id string, req *riderEntity.UpdateRiderReq) (*riderEntity.UpdateRiderReq, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancelFunc()

	updtStmt := fmt.Sprintf(`UPDATE riders_table SET first_name = '%v', last_name ='%v', email='%v', phone_number='%v', profile_picture='%v', date_updated='%v' WHERE rider_id = '%v';`, req.FirstName, req.LastName, req.Email, req.Phone,
		req.ProfilePicture, req.DateUpdated, id)
	_, err := p.conn.Exec(ctx, updtStmt)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (p *psql) ChangePassword(id string, password string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancelFunc()

	delStmt := fmt.Sprintf("UPDATE riders_table SET password = '%v' WHERE rider_id='%v' ;", password, id)
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

	deleteStmt2 := fmt.Sprintf(`DELETE FROM riders_table WHERE rider_id='%v';`, req.RiderId)
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

	queryStmt := fmt.Sprintf(`SELECT rider_id, first_name, last_name, email, phone_number, password, profile_picture, account_status FROM riders_table WHERE email='%v';`, email)

	row := p.conn.QueryRow(ctx, queryStmt)
	err := row.Scan(&res.RiderId, &res.FirstName, &res.LastName, &res.Email, &res.Phone, &res.Password, &res.ProfilePicture, &res.AccountStatus)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *psql) GetByPhone(phone string) (*riderEntity.CreateRiderRes, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	var res riderEntity.CreateRiderRes

	queryStmt := fmt.Sprintf(`SELECT rider_id, first_name, last_name, email, phone_number, password, profile_picture, account_status FROM riders_table WHERE phone_number='%v';`, phone)

	row := p.conn.QueryRow(ctx, queryStmt)
	err := row.Scan(&res.RiderId, &res.FirstName, &res.LastName, &res.Email, &res.Phone, &res.Password, &res.ProfilePicture, &res.AccountStatus)
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

	for i, guarantor := range req.Guarantor {
		insertStmt := fmt.Sprintf("INSERT INTO guarantors_table (guarantor_id, guarantor_first_name, guarantor_last_name, rider_id, guarantor_email, guarantor_phone, guarantor_residential_address, guarantor_official_address, guarantor_jobs) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')",
			guarantor.GuarantorId, guarantor.GuarantorFirstName, guarantor.GuarantorLastName, guarantor.RiderId, guarantor.GuarantorEmail, guarantor.GuarantorPhone, guarantor.GuarantorResidentialAddress, guarantor.GuarantorOfficeAddress, guarantor.GuarantorJob,
		)
		_, err = tx.Exec(ctx, insertStmt)
		if err != nil {
			log.Printf("inserted guarantor number %v succefully", i)
			return nil, err
		}
	}

	insertStmt2 := fmt.Sprintf(` INSERT INTO riders_table (rider_id, first_name, last_name, email, password,phone_number, DOB, gender, marital_status,educational_level, residential_address,driver_license,passport, profile_picture, account_status, date_created) 
 		VALUES('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')`,
		req.RiderId, req.FirstName, req.LastName, req.Email, req.Password, req.PhoneNumber, req.DOB, req.Gender, req.MaritalStatus, req.EducationLevel,
		req.ResidentialAddress, req.DriverLicense, req.Passport, req.ProfilePicture, req.AccountStatus, req.DateCreated)

	_, err = tx.Exec(ctx, insertStmt2)
	if err != nil {
		log.Printf("errror user: %v", err)
		return nil, err
	}

	return req, nil
}

func NewPsqlRiderRepo(conn *pgx.Conn) riderRepository.RiderRepo {
	return &psql{conn: conn}
}
