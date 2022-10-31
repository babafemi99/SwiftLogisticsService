package validationService

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"sls/internal/entity/adminEntity"
	"sls/internal/entity/userEntity"
	"testing"
)

func Test_validationSrv_Validate(t *testing.T) {

	//validData := userEntity.UpdateUserReq{
	//	UserId:         uuid.New(),
	//	Email:          "oo@o.com",
	//	Phone:          "+2349098140976",
	//	FirstName:      "Bayo",
	//	LastName:       "Banda",
	//	ProfilePicture: "https://hhasdiude.com",
	//}

	InvalidData := userEntity.UpdateUserReq{
		UserId:         uuid.New(),
		Email:          "ooo.com",
		Phone:          "+2349098140976",
		FirstName:      "Bayo",
		LastName:       "Banda",
		ProfilePicture: "hhasdicom",
	}

	validData := adminEntity.Availability{
		Status: "INACTIVE",
		Start:  "start",
		Stop:   "stop",
	}

	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Correct Credentials",
			args: args{
				data: validData,
			},
			wantErr: false,
		},
		{
			name: "Incorrect Credentials",
			args: args{
				data: InvalidData,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validationSrv{}
			if err := v.Validate(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDPR(t *testing.T) {
	vld := adminEntity.DPR{
		Status:     "NIL",
		Kilometers: 4,
		Seconds:    456,
	}
	v := validationSrv{}

	err := v.ValidateDPR(&vld)
	fmt.Println(vld)

	require.Nil(t, err)
}
