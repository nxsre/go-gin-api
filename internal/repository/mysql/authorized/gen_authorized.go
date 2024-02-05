///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package authorized

import (
	"fmt"
	"time"

	"github.com/nxsre/go-gin-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Authorized {
	return new(Authorized)
}

func NewQueryBuilder() *authorizedQueryBuilder {
	return new(authorizedQueryBuilder)
}

func (t *Authorized) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type authorizedQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *authorizedQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *authorizedQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Authorized{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *authorizedQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Authorized{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *authorizedQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Authorized{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *authorizedQueryBuilder) First(db *gorm.DB) (*Authorized, error) {
	ret := &Authorized{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *authorizedQueryBuilder) QueryOne(db *gorm.DB) (*Authorized, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *authorizedQueryBuilder) QueryAll(db *gorm.DB) ([]*Authorized, error) {
	var ret []*Authorized
	err := qb.buildQuery(db).Limit(-1).Find(&ret).Error
	return ret, err
}

func (qb *authorizedQueryBuilder) Limit(limit int) *authorizedQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *authorizedQueryBuilder) Offset(offset int) *authorizedQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *authorizedQueryBuilder) WhereId(p mysql.Predicate, value int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereIdIn(value []int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereIdNotIn(value []int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderById(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessKey(p mysql.Predicate, value string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessKeyIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessKeyNotIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByBusinessKey(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_key "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessSecret(p mysql.Predicate, value string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessSecretIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessSecretNotIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByBusinessSecret(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_secret "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessDeveloper(p mysql.Predicate, value string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessDeveloperIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereBusinessDeveloperNotIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByBusinessDeveloper(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_developer "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereRemark(p mysql.Predicate, value string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereRemarkIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereRemarkNotIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByRemark(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "remark "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereIsUsed(p mysql.Predicate, value int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereIsUsedIn(value []int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereIsUsedNotIn(value []int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByIsUsed(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_used "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereIsDeletedIn(value []int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereIsDeletedNotIn(value []int32) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByIsDeleted(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereCreatedAtIn(value []time.Time) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByCreatedAt(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereCreatedUserIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereCreatedUserNotIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByCreatedUser(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereUpdatedAtIn(value []time.Time) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByUpdatedAt(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *authorizedQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereUpdatedUserIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) WhereUpdatedUserNotIn(value []string) *authorizedQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedQueryBuilder) OrderByUpdatedUser(asc bool) *authorizedQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
