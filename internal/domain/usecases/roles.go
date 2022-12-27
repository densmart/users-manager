package usecases

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/densmart/users-manager/pkg"
	"time"
)

func CreateRole(s services.Service, data dto.CreateRoleDTO) (*dto.RoleDTO, error) {
	role, err := s.Roles.Create(data)
	if err != nil {
		return nil, err
	}
	response := dto.RoleDTO{
		ID:          role.Id,
		CreatedAt:   role.CreatedAt.Format(time.RFC3339),
		Name:        role.Name,
		Slug:        role.Slug,
		IsPermitted: role.IsPermitted,
	}
	return &response, nil
}

func UpdateRole(s services.Service, data dto.UpdateRoleDTO) (*dto.RoleDTO, error) {
	role, err := s.Roles.Update(data)
	if err != nil {
		return nil, err
	}
	response := dto.RoleDTO{
		ID:          role.Id,
		CreatedAt:   role.CreatedAt.Format(time.RFC3339),
		Name:        role.Name,
		Slug:        role.Slug,
		IsPermitted: role.IsPermitted,
	}
	return &response, nil
}

func RetrieveRole(s services.Service, id uint64) (*dto.RoleDTO, error) {
	role, err := s.Roles.Retrieve(id)
	if err != nil {
		return nil, err
	}
	response := dto.RoleDTO{
		ID:          role.Id,
		CreatedAt:   role.CreatedAt.Format(time.RFC3339),
		Name:        role.Name,
		Slug:        role.Slug,
		IsPermitted: role.IsPermitted,
	}
	return &response, nil
}

func SearchRoles(s services.Service, data dto.SearchRoleDTO) (*dto.RolesDTO, error) {
	paginator := pkg.NewPaginator(data.BaseSearchRequestDto)
	offset := paginator.GetOffset()
	limit := paginator.GetLimit()
	data.Offset = &offset
	data.Limit = &limit

	roles, err := s.Roles.Search(data)
	if err != nil {
		return nil, err
	}

	response := dto.RolesDTO{
		Pagination: paginator.ToRepresentation(),
	}
	for _, role := range roles {
		roleDTO := dto.RoleDTO{
			ID:          role.Id,
			CreatedAt:   role.CreatedAt.Format(time.RFC3339),
			Name:        role.Name,
			Slug:        role.Slug,
			IsPermitted: role.IsPermitted,
		}
		response.Items = append(response.Items, roleDTO)
	}

	return &response, nil
}

func DeleteRole(s services.Service, id uint64) error {
	return s.Roles.Delete(id)
}
