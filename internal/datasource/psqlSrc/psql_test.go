package psqlSrc

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPsqlSrc(t *testing.T) {
	log := logrus.New()
	psq, _ := NewPsqlSrc(log, "postgres://postgres:mysecretpassword@localhost:5432/slsstore")
	err := psq.LoadDB("./../create.sql")
	require.NoError(t, err)

	type args struct {
		log     *logrus.Logger
		connStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *psqlSrc
		wantErr bool
	}{
		{
			name: "correct cred",
			args: args{
				log:     log,
				connStr: "postgres://postgres:mysecretpassword@localhost:5432/slsstore",
			},
			want:    psq,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPsqlSrc(tt.args.log, tt.args.connStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPsqlSrc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Nil(t, err)
			assert.NotNil(t, got)
			assert.IsType(t, got, psq)
		})
	}
}

func Test_psqlSrc_LoadDB(t *testing.T) {
	log := logrus.New()
	psq, _ := NewPsqlSrc(log, "postgres://postgres:mysecretpassword@localhost:5432/slsstore")

	type fields struct {
		log  *logrus.Logger
		conn *pgx.Conn
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "correct cred",
			fields: fields{
				log:  log,
				conn: psq.GetConn(),
			},
			args: args{
				"./../create.sql",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &psqlSrc{
				log:  tt.fields.log,
				conn: tt.fields.conn,
			}
			fmt.Println(1)
			err := p.LoadDB(tt.args.path)
			require.Nil(t, err)

		})
	}
}
