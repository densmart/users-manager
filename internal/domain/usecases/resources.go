package usecases

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/densmart/users-manager/pkg"
	"time"
)

func CreateResource(s services.Service, data dto.CreateResourceDTO) (*dto.ResourceDTO, *APIError) {
	resource, err := s.Resources.Create(data)
	if err != nil {
		return nil, &APIError{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}
	response := dto.ResourceDTO{
		ID:        resource.Id,
		CreatedAt: resource.CreatedAt.Format(time.RFC3339),
		Name:      resource.Name,
		UriMask:   resource.UriMask,
	}
	return &response, nil
}

func UpdateResource(s services.Service, data dto.UpdateResourceDTO) (*dto.ResourceDTO, error) {
	resource, err := s.Resources.Update(data)
	if err != nil {
		return nil, err
	}
	response := dto.ResourceDTO{
		ID:        resource.Id,
		CreatedAt: resource.CreatedAt.Format(time.RFC3339),
		Name:      resource.Name,
		UriMask:   resource.UriMask,
	}
	return &response, nil
}

func RetrieveResource(s services.Service, id uint64) (*dto.ResourceDTO, error) {
	resource, err := s.Resources.Retrieve(id)
	if err != nil {
		return nil, err
	}
	response := dto.ResourceDTO{
		ID:        resource.Id,
		CreatedAt: resource.CreatedAt.Format(time.RFC3339),
		Name:      resource.Name,
		UriMask:   resource.UriMask,
	}
	return &response, nil
}

func SearchResources(s services.Service, data dto.SearchResourceDTO) (*dto.ResourcesDTO, error) {
	paginator := pkg.NewPaginator(data.BaseSearchRequestDto)
	offset := paginator.GetOffset()
	limit := paginator.GetLimit()
	data.Offset = &offset
	data.Limit = &limit

	resources, err := s.Resources.Search(data)
	if err != nil {
		return nil, err
	}

	response := dto.ResourcesDTO{
		Pagination: paginator.ToLinkHeader(),
	}
	for _, resource := range resources {
		resourceDTO := dto.ResourceDTO{
			ID:        resource.Id,
			CreatedAt: resource.CreatedAt.Format(time.RFC3339),
			Name:      resource.Name,
			UriMask:   resource.UriMask,
		}
		response.Items = append(response.Items, resourceDTO)
	}

	return &response, nil
}

func DeleteResource(s services.Service, id uint64) error {
	return s.Resources.Delete(id)
}
