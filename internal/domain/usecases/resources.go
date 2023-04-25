package usecases

import (
	"fmt"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/densmart/users-manager/internal/domain/utils"
	"time"
)

func CreateResource(s services.Service, data dto.CreateResourceDTO) (*dto.ResourceDTO, error) {
	binaryMask, err := utils.BinaryStringToDecimal(data.ActionsMask)
	if err != nil {
		return nil, err
	}
	data.MethodMask = uint8(binaryMask)
	resource, err := s.Resources.Create(data)
	if err != nil {
		return nil, &dto.APIError{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}
	response := dto.ResourceDTO{
		ID:          resource.Id,
		CreatedAt:   resource.CreatedAt.Format(time.RFC3339),
		Name:        resource.Name,
		UriMask:     resource.UriMask,
		ActionsMask: fmt.Sprintf("%04b", resource.MethodMask),
	}
	return &response, nil
}

func UpdateResource(s services.Service, data dto.UpdateResourceDTO) (*dto.ResourceDTO, error) {
	if data.ActionsMask != nil {
		binaryMask, err := utils.BinaryStringToDecimal(*data.ActionsMask)
		if err != nil {
			return nil, err
		}
		bm8 := uint8(binaryMask)
		data.MethodMask = &bm8
	}
	resource, err := s.Resources.Update(data)
	if err != nil {
		return nil, err
	}
	response := dto.ResourceDTO{
		ID:          resource.Id,
		CreatedAt:   resource.CreatedAt.Format(time.RFC3339),
		Name:        resource.Name,
		UriMask:     resource.UriMask,
		ActionsMask: fmt.Sprintf("%04b", resource.MethodMask),
	}
	return &response, nil
}

func RetrieveResource(s services.Service, id uint64) (*dto.ResourceDTO, error) {
	resource, err := s.Resources.Retrieve(id)
	if err != nil {
		return nil, err
	}
	response := dto.ResourceDTO{
		ID:          resource.Id,
		CreatedAt:   resource.CreatedAt.Format(time.RFC3339),
		Name:        resource.Name,
		UriMask:     resource.UriMask,
		ActionsMask: fmt.Sprintf("%04b", resource.MethodMask),
	}
	return &response, nil
}

func SearchResources(s services.Service, data dto.SearchResourceDTO) (*dto.ResourcesDTO, error) {
	defaultOrderField := "res_group"
	data.OrderBy = &defaultOrderField

	resources, err := s.Resources.Search(data)
	if err != nil {
		return nil, err
	}

	response := dto.ResourcesDTO{}
	groups := make(map[string][]dto.ResourceDTO)
	for _, resource := range resources {
		resourceDTO := dto.ResourceDTO{
			ID:          resource.Id,
			CreatedAt:   resource.CreatedAt.Format(time.RFC3339),
			Name:        resource.Name,
			UriMask:     resource.UriMask,
			ActionsMask: fmt.Sprintf("%04b", resource.MethodMask),
		}
		groups[resource.ResGroup.String] = append(groups[resource.ResGroup.String], resourceDTO)
	}
	for k, v := range groups {
		m := dto.ResourceGroupDTO{}
		m.Name = k
		m.Resources = v
		response.Items = append(response.Items, m)
	}

	return &response, nil
}

func DeleteResource(s services.Service, id uint64) error {
	return s.Resources.Delete(id)
}
