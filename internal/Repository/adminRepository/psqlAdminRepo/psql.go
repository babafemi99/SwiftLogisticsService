package psqlAdminRepo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"sls/internal/Repository/adminRepository"
	"sls/internal/entity/adminEntity"
	"sls/internal/entity/bikesEntity"
	"sls/internal/entity/riderEntity"
)

type psqlAdminRepository struct {
	conn *pgx.Conn
}

func (p *psqlAdminRepository) FindApplicationByName(ctx context.Context, name string) (*adminEntity.AdminViewRiderApplicationById, error) {
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
	var pdr adminEntity.AdminViewRiderApplicationById

	stmt1 := fmt.Sprintf(`
		SELECT  riders_table.rider_id as id, concat(riders_table.first_name, ' ', riders_table.last_name) as name, riders_table.email,
				riders_table.phone_number, riders_table.dob,riders_table.gender, riders_table.marital_status,
				riders_table.educational_level, riders_table.residential_address, riders_table.driver_license,
				riders_table.passport
		FROM riders_table
		WHERE concat(riders_table.first_name, ' ', riders_table.last_name) ='%v';`, name,
	)

	row := tx.QueryRow(ctx, stmt1)
	err = row.Scan(&pdr.Id, &pdr.Name, &pdr.Email, &pdr.PhoneNumber, &pdr.DateOfBirth, &pdr.Gender, &pdr.MaritalStatus,
		&pdr.EducationalLevel, &pdr.ResidentialAddress, &pdr.DriversLicense, &pdr.Passport)
	if err != nil {
		log.Println(2020)
		return nil, err
	}

	stmt2 := fmt.Sprintf(`
		SELECT concat(gt.guarantor_first_name, gt.guarantor_last_name) as guarantor_name,
			   gt.guarantor_email, gt.guarantor_phone, gt.guarantor_residential_address,
			   gt.guarantor_official_address, gt.guarantor_jobs, gt.guarantors_id_card
		FROM guarantors_table gt
		WHERE rider_id = '%v'`, &pdr.Id,
	)
	rows, err := tx.Query(ctx, stmt2)
	if err != nil {
		log.Printf("Error 201: %v", err)
		pdr.GuarantorDetails = nil
	}

	for rows.Next() {
		var gpr = &adminEntity.Guarantor{}
		err := rows.Scan(&gpr.GuarantorName, &gpr.GuarantorEmail, &gpr.GuarantorPhone, &gpr.GuarantorResidentialAddress,
			&gpr.GuarantorOfficeAddress, &gpr.GuarantorJob, &gpr.GuarantorIdCard)
		if err != nil {
			return nil, err
		}
		pdr.GuarantorDetails = append(pdr.GuarantorDetails, gpr)
	}

	return &pdr, nil
}

func (p *psqlAdminRepository) AcceptApplication(ctx context.Context, id string, application *adminEntity.AcceptRiderApplication) error {
	insertStmt := fmt.Sprintf(`
		UPDATE riders_table SET account_status='%v', salary='%v', bank_name='%v', bonus='%v', bank_account_no='%v', date_joined='%v'
		WHERE rider_id = '%v'`,
		application.AccountStatus, application.Salary, application.BankName, application.Bonus, application.BankAccountNo, application.DateJoined, id)

	_, err := p.conn.Exec(ctx, insertStmt)
	if err != nil {
		return err
	}
	return nil
}

func (p *psqlAdminRepository) ViewAllPendingRiders(ctx context.Context) ([]*adminEntity.AdminViewRiderApplication, error) {
	findStmt := fmt.Sprintf(`
		SELECT rider_id, concat(first_name,' ', last_name) as name , phone_number, email
		FROM riders_table
		WHERE account_status = 'PENDING';`,
	)

	var results []*adminEntity.AdminViewRiderApplication

	query, err := p.conn.Query(ctx, findStmt)
	if err != nil {
		return nil, err
	}

	for query.Next() {
		var result adminEntity.AdminViewRiderApplication
		err := query.Scan(&result.RiderId, &result.Name, &result.PhoneNumber, &result.Email)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil

}

func (p *psqlAdminRepository) ViewPendingRiderById(ctx context.Context, id string) (*adminEntity.AdminViewRiderApplicationById, error) {

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
	var pdr adminEntity.AdminViewRiderApplicationById

	stmt1 := fmt.Sprintf(`
		SELECT  riders_table.rider_id as id, concat(riders_table.first_name,' ', riders_table.last_name) as name, riders_table.email,
				riders_table.phone_number, riders_table.dob,riders_table.gender, riders_table.marital_status,
				riders_table.educational_level, riders_table.residential_address, riders_table.driver_license,
				riders_table.passport
		FROM riders_table
		WHERE rider_id ='%v';`, id,
	)

	row := tx.QueryRow(ctx, stmt1)
	err = row.Scan(&pdr.Id, &pdr.Name, &pdr.Email, &pdr.PhoneNumber, &pdr.DateOfBirth, &pdr.Gender, &pdr.MaritalStatus,
		&pdr.EducationalLevel, &pdr.ResidentialAddress, &pdr.DriversLicense, &pdr.Passport)
	if err != nil {
		log.Println(2020)
		return nil, err
	}

	stmt2 := fmt.Sprintf(`
		SELECT concat(gt.guarantor_first_name, gt.guarantor_last_name) as guarantor_name,
			   gt.guarantor_email, gt.guarantor_phone, gt.guarantor_residential_address,
			   gt.guarantor_official_address, gt.guarantor_jobs, gt.guarantors_id_card
		FROM guarantors_table gt
		WHERE rider_id = '%v'`, id,
	)
	rows, err := tx.Query(ctx, stmt2)
	if err != nil {
		log.Printf("Error 201: %v", err)
		pdr.GuarantorDetails = nil
	}

	for rows.Next() {
		var gpr = &adminEntity.Guarantor{}
		err := rows.Scan(&gpr.GuarantorName, &gpr.GuarantorEmail, &gpr.GuarantorPhone, &gpr.GuarantorResidentialAddress,
			&gpr.GuarantorOfficeAddress, &gpr.GuarantorJob, &gpr.GuarantorIdCard)
		if err != nil {
			return nil, err
		}
		pdr.GuarantorDetails = append(pdr.GuarantorDetails, gpr)
	}

	return &pdr, nil
}

func (p *psqlAdminRepository) ViewAllRiders(ctx context.Context) ([]*adminEntity.AdminViewAllRider, error) {
	findStmt := fmt.Sprintf(`
		SELECT riders_id, salary, bonus,rating, concat(first_name, ' ',last_name) as Name
		FROM riders_table full join bikes_table bt on riders_table.rider_id = bt.riders_id
		WHERE riders_id is not null;`,
	)

	var riders []*adminEntity.AdminViewAllRider

	rows, err := p.conn.Query(ctx, findStmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rider adminEntity.AdminViewAllRider
		err := rows.Scan(&rider.RiderId, &rider.Salary, &rider.Bonus, &rider.Ratings, &rider.Name)
		if err != nil {
			return nil, err
		}
		oo := int32(12)
		rider.NumberOfOrders = &oo
		riders = append(riders, &rider)
	}

	return riders, nil
}

func (p *psqlAdminRepository) ViewRiderById(ctx context.Context, id string) (*adminEntity.AdminViewRider, error) {
	var bike adminEntity.AdminViewRider

	findStmt := fmt.Sprintf(`
		SELECT rider_id, bank_account_no, bank_name,driver_license, passport, salary, bonus, next_servicing_date,
			   rating, date_joined, plate_number, colour, engine_number, work_status
		FROM riders_table 
		    full join bikes_table bt on riders_table.rider_id = bt.riders_id
		WHERE riders_id='%v';`, id,
	)
	row := p.conn.QueryRow(ctx, findStmt)
	err := row.Scan(&bike.RiderId, &bike.BankAccountNo, &bike.BankName, &bike.DriversLicense, &bike.Passport,
		&bike.Salary, &bike.Bonus, &bike.NextServiceDate, &bike.Rating, &bike.DateJoined, &bike.PlateNumber,
		&bike.Color, &bike.EngineNumber, &bike.WorkStatus,
	)
	if err != nil {
		return nil, err
	}

	return &bike, nil
}

func (p *psqlAdminRepository) EditRider(ctx context.Context, id string, admin *riderEntity.UpdateRiderReqAdmin) (*riderEntity.UpdateRiderReqAdmin, error) {
	updateStmt := fmt.Sprintf(`UPDATE riders_table SET first_name='%v', last_name='%v', phone_number='%v', 
                        email='%v', profile_picture='%v', marital_status='%v', educational_level='%v', residential_address='%v',salary='%v',bank_name='%v',
                        bank_account_no='%v', bonus='%v', rating='%v', date_updated='%v' WHERE rider_id ='%v'`, admin.FirstName, admin.LastName, admin.PhoneNumber, admin.Email,
		admin.ProfilePicture, admin.MaritalStatus, admin.EducationLevel, admin.ResidentialAddress, admin.Salary, admin.BankName, admin.BankAccountNo,
		admin.Bonus, admin.Rating, admin.DateUpdated, id)
	_, err := p.conn.Exec(ctx, updateStmt)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (p *psqlAdminRepository) ViewAllBikes(ctx context.Context) ([]*adminEntity.AdminViewBike, error) {

	var bikes []*adminEntity.AdminViewBike

	findStmt := fmt.Sprintf(`
	select
		bike_id, bike_name as "name_of_bike",engine_number, plate_number, colour, first_name as "name"
	from
	    bikes_table b
		full outer join riders_table rt
		    on rt.rider_id = b.riders_id
	WHERE bike_id IS NOT NULL;
`)
	rows, err := p.conn.Query(ctx, findStmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var bike adminEntity.AdminViewBike
		err := rows.Scan(&bike.BikeId, &bike.NameOfBike, &bike.EngineNumber, &bike.PlateNumber, &bike.Colour, &bike.Rider)
		if err != nil {
			return nil, err
		}

		bikes = append(bikes, &bike)
	}
	return bikes, nil
}

func (p *psqlAdminRepository) CreateBike(ctx context.Context, bike *bikesEntity.CreateBike) (*bikesEntity.CreateBike, error) {
	insertStmt := fmt.Sprintf(`INSERT INTO bikes_table(bike_id, bike_name, colour, picture, engine_number,
                        plate_number, date_purchased) VALUES ('%v','%v','%v','%v','%v','%v','%v')`,
		bike.BikeId, bike.BikeName, bike.Colour, bike.Picture, bike.EngineNumber, bike.PlateNumber, bike.DatePurchased,
	)
	_, err := p.conn.Exec(ctx, insertStmt)
	if err != nil {
		return nil, err
	}
	return bike, err

}

func (p *psqlAdminRepository) EditBike(ctx context.Context, id string, bike *bikesEntity.UpdateBike) (*bikesEntity.UpdateBike, error) {
	updateStmt := fmt.Sprintf(`
		UPDATE bikes_table
		SET bike_name='%v', colour='%v', picture='%v', engine_number='%v',plate_number='%v' 
		WHERE bike_id='%v'`, bike.BikeName, bike.Colour, bike.Picture, bike.EngineNumber, bike.PlateNumber, id,
	)
	_, err := p.conn.Exec(ctx, updateStmt)
	if err != nil {
		return nil, err
	}
	return bike, nil
}

func (p *psqlAdminRepository) DeleteBike(ctx context.Context, id string) error {
	updateStmt := fmt.Sprintf(`DELETE FROM bikes_table WHERE bike_id='%v'`, id)
	_, err := p.conn.Exec(ctx, updateStmt)
	if err != nil {
		return err
	}
	return nil
}

func (p *psqlAdminRepository) AssignBikeToRider(ctx context.Context, riderId, bikesId string) error {
	assignStmt := fmt.Sprintf("UPDATE bikes_table SET riders_id ='%v' WHERE bike_id ='%v';", riderId, bikesId)

	_, err := p.conn.Exec(ctx, assignStmt)
	if err != nil {
		log.Println(err, 1)
		return err
	}

	return nil

}

func (p *psqlAdminRepository) UpdateBikeHistory(ctx context.Context, bike *bikesEntity.UpdateBikeHistory) error {
	inserStmt := fmt.Sprintf(`INSERT INTO bikes_history_table(bike_id, riders_name, duration) 
		VALUES ('%v', '%v', '%v')`, bike.BikeId, bike.RidersName, bike.Duration)

	_, err := p.conn.Exec(ctx, inserStmt)
	if err != nil {
		return err
	}

	return nil
}

func (p *psqlAdminRepository) FindBikeById(ctx context.Context, id string) (*adminEntity.AdminViewBikeDetails, error) {

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

	var bike adminEntity.AdminViewBikeDetails

	findStmt := fmt.Sprintf(`
		SELECT bike_id, bike_name, colour, picture, engine_number, plate_number,
			   next_servicing_date, date_purchased, first_name as "allocated_rider"
		FROM bikes_table b 
			full outer join riders_table rt on rt.rider_id = b.riders_id
		WHERE bike_id='%v';`, id,
	)

	row := p.conn.QueryRow(ctx, findStmt)
	err = row.Scan(&bike.BikeId, &bike.BikeName, &bike.Colour, &bike.Picture, &bike.EngineNumber, &bike.PlateNumber,
		&bike.NextServicingDate, &bike.AllocatedDriver, &bike.DatePurchased)
	if err != nil {
		return nil, err
	}
	fmt.Println("id is, ", id)
	secondStmt := fmt.Sprintf(`
		SELECT bike_id,riders_name,duration 
		from bikes_history_table 
		WHERE bike_id = '%v'`, id,
	)
	var histories []*adminEntity.BikeHistory
	rows, err := p.conn.Query(ctx, secondStmt)
	if err != nil {
		log.Printf("error 101:%v", err)
		histories = nil
	}

	for rows.Next() {
		var history = &adminEntity.BikeHistory{}
		err = rows.Scan(&history.BikeId, &history.RidersName, &history.Duration)
		if err != nil {
			if err == pgx.ErrNoRows {
				history = nil
			}
			return nil, err
		}
		fmt.Println("history-", history)

		histories = append(histories, history)
	}
	bike.History = histories

	// find for bike history
	return &bike, err
}

func NewPsqlAdminBikeRepository(conn *pgx.Conn) adminRepository.AdminBikers {
	return &psqlAdminRepository{conn: conn}
}
