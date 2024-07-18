package repository

import (
	"gorm.io/gorm"
)

type IRepository[T any] interface {
	First(where *T) error
	Find(where *T) error
	FindAlls(dest *[]T, where *T) error
	Create(entity *T) error
	Update(entity *T) error
	UpdateFields(entity *T, fields map[string]interface{}) error
	DeleteByID(id uint) error
	FindByID(id uint) (*T, error)
	FindAll(req PaginateRequest) (PaginatedResult[T], error)
	Count(where *T) (*int64, error)
}

type Repository[T any] struct {
	DB *gorm.DB
}

func NewRepository[T any](db *gorm.DB) IRepository[T] {
	return &Repository[T]{DB: db}
}

func (r *Repository[T]) First(where *T) error {
	return r.DB.First(where, where).Error
}

func (r *Repository[T]) Find(where *T) error {
	return r.DB.Find(where, where).Error
}

func (r *Repository[T]) FindAlls(dest *[]T, where *T) error {
	return r.DB.Find(dest, where).Error
}

func (r *Repository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.DB.Model(entity).Save(entity).Error
}

func (r *Repository[T]) UpdateFields(entity *T, fields map[string]any) error {
	ID, _ := fields["ID"]
	query := r.DB.Model(entity).Where("id = ?", ID)
	delete(fields, "CreatedAt") // Avoid updating CreatedAt
	delete(fields, "ID")

	return query.Updates(fields).Error
}

func (r *Repository[T]) DeleteByID(id uint) error {
	return r.DB.Delete(new(T), id).Error
}

func (r *Repository[T]) FindByID(id uint) (*T, error) {
	var entity T
	err := r.DB.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *Repository[T]) FindAll(req PaginateRequest) (PaginatedResult[T], error) {
	var entities []*T
	query := r.DB.Model(&entities)

	// Apply filters
	for _, filter := range req.Filters {
		query = query.Where(filter.Column+" = ?", filter.Value)
	}

	// Apply search
	for column, value := range req.Search {
		query = query.Where(column+" LIKE ?", "%"+value+"%")
	}

	// Get total count
	var total int64
	query.Count(&total)

	// Apply pagination
	limit := req.Pagination.PageSize
	offset := (req.Pagination.Page - 1) * req.Pagination.PageSize
	query = query.Limit(limit).Offset(offset)

	err := query.Find(&entities).Error
	if err != nil {
		return PaginatedResult[T]{}, err
	}

	pageCount := int((total + int64(req.Pagination.PageSize) - 1) / int64(req.Pagination.PageSize))

	return PaginatedResult[T]{
		Data:      entities,
		Total:     total,
		Page:      req.Pagination.Page,
		PageSize:  req.Pagination.PageSize,
		PageCount: pageCount,
	}, nil
}

func (r *Repository[T]) Count(where *T) (*int64, error) {
	var count int64

	err := r.DB.Model(where).Where(where).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return &count, nil
}
