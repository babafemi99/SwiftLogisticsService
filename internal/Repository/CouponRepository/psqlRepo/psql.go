package psqlRepo

//
//import (
//	"context"
//	"fmt"
//	"github.com/jackc/pgx/v4"
//	"sls/internal/entity/couponEntity"
//	"time"
//)
//
//type psql struct {
//	conn *pgx.Conn
//}
//
//func NewCouponPsql(conn *pgx.Conn) *psql {
//	return &psql{conn: conn}
//}
//
//func (p *psql) CreateCoupon(coupon *couponEntity.Coupon) (*couponEntity.Coupon, error) {
//	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancelFunc()
//
//	insertStmt := fmt.Sprintf("INSERT INTO"+
//		"coupons (coupon_id, coupon_code, coupon_name, coupon_description, coupon_discount, coupon_expiry_date) "+
//		"VALUES ('%v','%v','%v','%v','%v','%v');",
//		coupon.CouponId, coupon.CouponCode, coupon.CouponName, coupon.CouponDescription,
//		coupon.CouponDiscount, coupon.CouponExpiryDate,
//	)
//
//	_, err := p.conn.Exec(ctx, insertStmt)
//	if err != nil {
//		return nil, err
//	}
//
//	return coupon, nil
//}
//
//func (p *psql) DeleteCoupon(id string) error {
//	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancelFunc()
//
//	deleteStmt := fmt.Sprintf("DELETE FROM coupons WHERE id = '%v'; ", id)
//	_, err := p.conn.Exec(ctx, deleteStmt)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (p *psql) EditCoupon(id string, coupon *couponEntity.Coupon) (*couponEntity.Coupon, error) {
//	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancelFunc()
//
//	editStmt := fmt.Sprintf("UPDATE coupons SET"+
//		"coupon_name='%v', coupon_description='%v', coupon_code='%v',"+
//		"coupon_discount='%v', coupon_expiry_date='%v', coupon_maximum='%v' WHERE id='%v';",
//		coupon.CouponName, coupon.CouponDescription, coupon.CouponCode,
//		coupon.CouponDiscount, coupon.CouponExpiryDate, coupon.CouponMaximum, id,
//	)
//
//	_, err := p.conn.Exec(ctx, editStmt)
//	if err != nil {
//		return nil, err
//	}
//
//	return coupon, err
//}
//
//func (p *psql) FindByCode(code string) (*couponEntity.Coupon, error) {
//	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancelFunc()
//
//	var coupon couponEntity.Coupon
//	findStmt := fmt.Sprintf("SELECT * FROM coupons WHERE coupon_code='%v'", code)
//	err := p.conn.QueryRow(ctx, findStmt).Scan(&coupon)
//	if err != nil {
//		return nil, err
//	}
//	return &coupon, nil
//}
//
//func (p *psql) FindAllCoupon() ([]*couponEntity.Coupon, error) {
//	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
//	defer cancelFunc()
//
//	var coupons []*couponEntity.Coupon
//	findStmt := fmt.Sprintf("SELECT * FROM coupons WHERE coupon_code='%v'", code)
//	rows, err := p.conn.Query(ctx, findStmt)
//	if err != nil {
//		return nil, err
//	}
//
//	for rows.Next() {
//		var coupon couponEntity.Coupon
//		err := rows.Scan(&coupon)
//		if err != nil {
//			return nil, err
//		}
//		coupons = append(coupons, &coupon)
//	}
//	return coupons, nil
//}
//
//func (p *psql) UseCoupon(id, couponId string) error {
//	//- check if value already exists
//	// insert coupon into db if not exists
//	// value already exist increase count
//
//	//update coupon count value
//	panic("implement me")
//}
