package psqlRepo

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"reflect"
	"sls/internal/datasource/psqlSrc"
	"sls/internal/entity/riderEntity"
	"testing"
)

func Test_psql_CreateRider(t *testing.T) {
	log := logrus.New()

	psqlData, err := psqlSrc.NewPsqlSrc(log, "postgres://postgres:mysecretpassword@localhost:5432/slsstore")
	if err != nil {
		log.Fatalf("Error Starting Database: %v", err)
	}
	conn := psqlData.GetConn()
	id := uuid.New()
	g := riderEntity.Guarantor{
		GuarantorId:      uuid.New(),
		RiderId:          id,
		FirstName:        "Abc",
		LastName:         "Def",
		Phone:            "Efg",
		Email:            "ijk",
		GuarantorAddress: "lmno",
		GuarantorJob:     "abcd",
	}

	correctReq := &riderEntity.CreateRiderReq{
		RiderId:            id,
		FirstName:          "Bayo",
		LastName:           "Ade",
		Phone:              "090987654",
		Email:              "a@a.com",
		DOB:                "aas",
		Gender:             "ss",
		MaritalStatus:      "ddd",
		EducationLevel:     "aund",
		ResidentialAddress: "79083",
		Guarantor:          &g,
		DriverLicense:      "wws",
		IdentityCard:       "sref",
		VerificationStatus: "urof",
		AccountStatus:      "oehdb",
		CreatedAt:          "ysip",
		UpdatedAt:          "uoils",
	}

	type fields struct {
		conn *pgx.Conn
	}
	type args struct {
		req *riderEntity.CreateRiderReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *riderEntity.CreateRiderReq
		wantErr bool
	}{
		{
			name: "correct credentials",
			fields: fields{
				conn: conn,
			},
			args: args{
				req: correctReq,
			},
			want:    correctReq,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &psql{
				conn: tt.fields.conn,
			}
			got, err := p.Persist(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRider() got = %v, want %v", got, tt.want)
			}
		})
	}
}
