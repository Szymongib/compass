package application

import (
	"github.com/kyma-incubator/compass/components/director/internal/labelfilter"
	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/pkg/pagination"
	"github.com/pkg/errors"
)

type inMemoryRepository struct {
	store map[string]*model.Application
}

func NewRepository() *inMemoryRepository {
	return &inMemoryRepository{store: make(map[string]*model.Application)}
}

func (r *inMemoryRepository) GetByID(tenant, id string) (*model.Application, error) {
	application := r.store[id]

	if application == nil || application.Tenant != tenant {
		return nil, errors.New("application not found")
	}

	return application, nil
}

// TODO: Make filtering and paging
func (r *inMemoryRepository) List(tenant string, filter []*labelfilter.LabelFilter, pageSize *int, cursor *string) (*model.ApplicationPage, error) {
	var items []*model.Application
	for _, item := range r.store {
		if item.Tenant == tenant {
			items = append(items, item)
		}
	}

	return &model.ApplicationPage{
		Data:       items,
		TotalCount: len(items),
		PageInfo: &pagination.Page{
			StartCursor: "",
			EndCursor:   "",
			HasNextPage: false,
		},
	}, nil
}

func (r *inMemoryRepository) Create(item *model.Application) error {
	if item == nil {
		return errors.New("item can not be empty")
	}

	r.store[item.ID] = item

	return nil
}

func (r *inMemoryRepository) Update(item *model.Application) error {
	if item == nil {
		return errors.New("item can not be empty")
	}

	if r.store[item.ID] == nil {
		return errors.New("application not found")
	}

	r.store[item.ID] = item

	return nil
}

func (r *inMemoryRepository) Delete(item *model.Application) error {
	if item == nil {
		return nil
	}

	delete(r.store, item.ID)

	return nil
}
