package qrybldr

import (
	"errors"
	"gorm.io/gorm"
)

type Qrybldr struct {
	db    *gorm.DB
	model interface{}
}

func Instance(db *gorm.DB) *Qrybldr {
	return &Qrybldr{db: db}
}

func (q *Qrybldr) Table(table string) *Qrybldr {
	q.db = q.db.Table(table)
	return q
}

func (q *Qrybldr) Select(columns ...string) *Qrybldr {
	q.db.Select(columns)

	return q
}

func (q *Qrybldr) Where(query string, args ...interface{}) *Qrybldr {
	q.db = q.db.Where(query, args...)
	return q
}

func (q *Qrybldr) OrderBy(order string) *Qrybldr {
	q.db = q.db.Order(order)
	return q
}

func (q *Qrybldr) Limit(limit int) *Qrybldr {
	q.db = q.db.Limit(limit)
	return q
}

func (q *Qrybldr) Offset(offset int) *Qrybldr {
	q.db = q.db.Offset(offset)
	return q
}

func (q *Qrybldr) GroupBy(columns ...string) *Qrybldr {
	q.db = q.db.Group(columns[0])
	for _, col := range columns[1:] {
		q.db = q.db.Group(col)
	}
	return q
}

func (q *Qrybldr) Having(cond string, args ...interface{}) *Qrybldr {
	q.db = q.db.Having(cond, args...)
	return q
}

func (q *Qrybldr) Distinct() *Qrybldr {
	q.db = q.db.Distinct()
	return q
}

func (q *Qrybldr) Join(table string, condition string) *Qrybldr {
	q.db = q.db.Joins("JOIN " + table + " ON " + condition)
	return q
}

func (q *Qrybldr) With(relations ...string) *Qrybldr {
	for _, r := range relations {
		q.db.Preload(r)
	}
	return q
}

func (q *Qrybldr) When(cond bool, fn func(*Qrybldr) *Qrybldr) *Qrybldr {
	if cond {
		return fn(q)
	}
	return q
}

func (q *Qrybldr) First(dest interface{}) error {
	return q.db.First(dest).Error
}

func (q *Qrybldr) FirstOrFail(dest interface{}) error {
	err := q.db.First(dest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}
	return err
}

func (q *Qrybldr) FirstOrCreate(dest interface{}, defaults map[string]interface{}) error {
	return q.db.FirstOrCreate(dest, defaults).Error
}

func (q *Qrybldr) FindByID(dest interface{}, id interface{}) error {
	return q.db.First(dest, id).Error
}

func (q *Qrybldr) Get(dest interface{}) error {
	return q.db.Find(dest).Error
}

func (q *Qrybldr) Pluck(dest interface{}, column string) error {
	return q.db.Pluck(column, dest).Error
}

func (q *Qrybldr) Count() (int64, error) {
	var count int64
	err := q.db.Count(&count).Error
	return count, err
}

func (q *Qrybldr) Exists() (bool, error) {
	c, err := q.Count()
	return c > 0, err
}

func (q *Qrybldr) Update(values map[string]interface{}) error {
	return q.db.Updates(values).Error
}

func (q *Qrybldr) Delete(model interface{}) error {
	return q.db.Delete(model).Error
}

func (q *Qrybldr) Begin() *Qrybldr {
	q.db = q.db.Begin()
	return q
}

func (q *Qrybldr) Commit() *Qrybldr {
	q.db.Commit()
	return q
}

func (q *Qrybldr) Rollback() *Qrybldr {
	q.db.Rollback()
	return q
}

type PaginationResult struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	LastPage int         `json:"last_page"`
}

func (q *Qrybldr) Paginate(dest interface{}, page int, perPage int) (PaginationResult, error) {
	var result PaginationResult
	var count int64

	q.db.Count(&count)
	result.Total = count
	result.Page = page
	result.PerPage = perPage
	if perPage > 0 {
		result.LastPage = int((count + int64(perPage) - 1) / int64(perPage))
	} else {
		result.LastPage = 1
	}

	offset := (page - 1) * perPage
	err := q.db.Offset(offset).Limit(perPage).Find(dest).Error
	result.Data = dest

	return result, err
}
