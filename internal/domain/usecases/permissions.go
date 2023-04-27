package usecases

import (
	"fmt"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/densmart/users-manager/internal/domain/utils"
	"time"
)

func CreatePermissions(s services.Service, data []dto.CreatePermissionDTO) ([]dto.PermissionDTO, error) {
	for idx, item := range data {
		binaryMask, err := utils.BinaryStringToDecimal(item.ActionsMask)
		if err != nil {
			return nil, err
		}
		data[idx].MethodMask = uint8(binaryMask)
	}

	permissions, err := s.Permissions.Create(data)
	if err != nil {
		return nil, &dto.APIError{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}
	var response []dto.PermissionDTO
	for _, perm := range permissions {
		permDTO := dto.PermissionDTO{
			ID:          perm.Id,
			CreatedAt:   perm.CreatedAt.Format(time.RFC3339),
			RoleID:      perm.RoleID,
			ResourceID:  perm.ResourceID,
			ActionsMask: fmt.Sprintf("%04b", perm.MethodMask),
		}
		response = append(response, permDTO)
	}
	return response, nil
}

func UpdatePermissions(s services.Service, data []dto.UpdatePermissionDTO) ([]dto.PermissionDTO, error) {
	for idx, item := range data {
		if item.ActionsMask != nil {
			binaryMask, err := utils.BinaryStringToDecimal(*item.ActionsMask)
			if err != nil {
				return nil, err
			}
			data[idx].MethodMask = uint8(binaryMask)
		}
	}
	permissions, err := s.Permissions.Update(data)
	if err != nil {
		return nil, err
	}
	var response []dto.PermissionDTO
	for _, perm := range permissions {
		permDTO := dto.PermissionDTO{
			ID:          perm.Id,
			CreatedAt:   perm.CreatedAt.Format(time.RFC3339),
			RoleID:      perm.RoleID,
			ResourceID:  perm.ResourceID,
			ActionsMask: fmt.Sprintf("%04b", perm.MethodMask),
		}
		response = append(response, permDTO)
	}
	return response, nil
}

func SearchPermissions(s services.Service, data dto.SearchPermissionDTO) (*dto.PermissionsDTO, error) {
	paginator := utils.NewPaginator(data.BaseSearchRequestDto)
	resources, err := s.Permissions.Search(data)
	if err != nil {
		return nil, err
	}

	response := dto.PermissionsDTO{
		Pagination: paginator.ToLinkHeader(),
	}
	for _, perm := range resources {
		permDTO := dto.PermissionDTO{
			ID:          perm.Id,
			CreatedAt:   perm.CreatedAt.Format(time.RFC3339),
			RoleID:      perm.RoleID,
			ResourceID:  perm.ResourceID,
			ActionsMask: fmt.Sprintf("%04b", perm.MethodMask),
		}
		response.Items = append(response.Items, permDTO)
	}

	return &response, nil
}

func DeletePermissions(s services.Service, id []uint64) error {
	return s.Permissions.Delete(id)
}
