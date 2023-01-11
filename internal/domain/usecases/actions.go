package usecases

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/densmart/users-manager/pkg"
	"time"
)

func CreateAction(s services.Service, data dto.CreateActionDTO) (*dto.ActionDTO, *dto.APIError) {
	action, err := s.Actions.Create(data)
	if err != nil {
		return nil, &dto.APIError{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}
	response := dto.ActionDTO{
		ID:        action.Id,
		CreatedAt: action.CreatedAt.Format(time.RFC3339),
		Name:      action.Name,
		Method:    action.Method,
	}
	return &response, nil
}

func UpdateAction(s services.Service, data dto.UpdateActionDTO) (*dto.ActionDTO, error) {
	action, err := s.Actions.Update(data)
	if err != nil {
		return nil, err
	}
	response := dto.ActionDTO{
		ID:        action.Id,
		CreatedAt: action.CreatedAt.Format(time.RFC3339),
		Name:      action.Name,
		Method:    action.Method,
	}
	return &response, nil
}

func RetrieveAction(s services.Service, id uint64) (*dto.ActionDTO, error) {
	action, err := s.Actions.Retrieve(id)
	if err != nil {
		return nil, err
	}
	response := dto.ActionDTO{
		ID:        action.Id,
		CreatedAt: action.CreatedAt.Format(time.RFC3339),
		Name:      action.Name,
		Method:    action.Method,
	}
	return &response, nil
}

func SearchActions(s services.Service, data dto.SearchActionDTO) (*dto.ActionsDTO, error) {
	paginator := pkg.NewPaginator(data.BaseSearchRequestDto)
	offset := paginator.GetOffset()
	limit := paginator.GetLimit()
	data.Offset = &offset
	data.Limit = &limit

	actions, err := s.Actions.Search(data)
	if err != nil {
		return nil, err
	}

	response := dto.ActionsDTO{
		Pagination: paginator.ToLinkHeader(),
	}
	for _, action := range actions {
		actionDTO := dto.ActionDTO{
			ID:        action.Id,
			CreatedAt: action.CreatedAt.Format(time.RFC3339),
			Name:      action.Name,
			Method:    action.Method,
		}
		response.Items = append(response.Items, actionDTO)
	}

	return &response, nil
}

func DeleteAction(s services.Service, id uint64) error {
	return s.Actions.Delete(id)
}
