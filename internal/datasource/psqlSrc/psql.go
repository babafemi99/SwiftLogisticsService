package psqlSrc

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
	"time"
)

//var (
//	connectionStr = "postgres://postgres:mysecretpassword@localhost:5432/postgres"
//)

type psqlSrc struct {
	log  *logrus.Logger
	conn *pgx.Conn
}

func NewPsqlSrc(log *logrus.Logger, connStr string) (*psqlSrc, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancelFunc()
	log.Info("Database Starting.....")

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database :%v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("Error Pinging DB: %v", err)
	}

	log.Info("Database connected successfully")
	return &psqlSrc{conn: conn}, nil
}

func (p *psqlSrc) GetConn() *pgx.Conn {
	return p.conn
}

func (p *psqlSrc) CloseConn() {
	p.CloseConn()
}
func (p *psqlSrc) LoadDB(path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	file, queryErr := os.ReadFile(path)
	if queryErr != nil {
		log.Fatalf("unable to read sql file: %v", queryErr)
	}

	tx, err := p.conn.Begin(ctx)
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

	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(ctx, q); err != nil {
			return err
		}
	}

	return nil
}
