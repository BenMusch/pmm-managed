// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
)

func serviceNewID(q *reform.Querier) (string, error) {
	id := "/service_id/" + uuid.New().String()
	service := &Service{ServiceID: id}
	switch err := q.Reload(service); err {
	case nil:
		return "", status.Errorf(codes.AlreadyExists, "Service with ID %q already exists.", id)
	case reform.ErrNoRows:
		return id, nil
	default:
		return "", errors.WithStack(err)
	}
}

func checkUniqueServiceName(q *reform.Querier, name string) error {
	if name == "" {
		return status.Error(codes.InvalidArgument, "Empty Service name.")
	}

	_, err := q.FindOneFrom(ServiceTable, "service_name", name)
	switch err {
	case nil:
		return status.Errorf(codes.AlreadyExists, "Service with name %q already exists.", name)
	case reform.ErrNoRows:
		return nil
	default:
		return errors.WithStack(err)
	}
}

// FindAllServices returns all Services.
func FindAllServices(q *reform.Querier) ([]*Service, error) {
	structs, err := q.SelectAllFrom(ServiceTable, "ORDER BY service_id")
	if err != nil {
		return nil, err
	}

	services := make([]*Service, len(structs))
	for i, s := range structs {
		services[i] = s.(*Service)
	}

	return services, nil
}

// FindServiceByID finds Service by ID.
func FindServiceByID(q *reform.Querier, id string) (*Service, error) {
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty Service ID.")
	}

	service := &Service{ServiceID: id}
	switch err := q.Reload(service); err {
	case nil:
		return service, nil
	case reform.ErrNoRows:
		return nil, status.Errorf(codes.NotFound, "Service with ID %q not found.", id)
	default:
		return nil, errors.WithStack(err)
	}
}

// FindServicesForAgentID returns all Services for which Agent with given ID provides insights.
func FindServicesForAgentID(q *reform.Querier, agentID string) ([]*Service, error) {
	structs, err := q.FindAllFrom(AgentServiceView, "agent_id", agentID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select Service IDs")
	}

	serviceIDs := make([]interface{}, len(structs))
	for i, s := range structs {
		serviceIDs[i] = s.(*AgentService).ServiceID
	}
	if len(serviceIDs) == 0 {
		return []*Service{}, nil
	}

	p := strings.Join(q.Placeholders(1, len(serviceIDs)), ", ")
	tail := fmt.Sprintf("WHERE service_id IN (%s) ORDER BY service_id", p) //nolint:gosec
	structs, err = q.SelectAllFrom(ServiceTable, tail, serviceIDs...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select Services")
	}

	res := make([]*Service, len(structs))
	for i, s := range structs {
		res[i] = s.(*Service)
	}
	return res, nil
}

// CreateServiceParams contains parameters for creating Services.
type CreateServiceParams struct {
	ServiceName  string
	NodeID       string
	CustomLabels map[string]string
	Address      *string
	Port         *uint16
}

// CreateService creates a Service.
func CreateService(q *reform.Querier, serviceType ServiceType, params *CreateServiceParams) (*Service, error) {
	id, err := serviceNewID(q)
	if err != nil {
		return nil, err
	}
	if err = checkUniqueServiceName(q, params.ServiceName); err != nil {
		return nil, err
	}

	if _, err = FindNodeByID(q, params.NodeID); err != nil {
		return nil, err
	}

	service := &Service{
		ServiceID:   id,
		ServiceType: serviceType,
		ServiceName: params.ServiceName,
		NodeID:      params.NodeID,
		Address:     params.Address,
		Port:        params.Port,
	}
	if err = service.SetCustomLabels(params.CustomLabels); err != nil {
		return nil, err
	}
	if err = q.Insert(service); err != nil {
		return nil, errors.WithStack(err)
	}

	return service, nil
}

// RemoveService removes a Service.
func RemoveService(q *reform.Querier, id string) error {
	err := q.Delete(&Service{ServiceID: id})
	if err == reform.ErrNoRows {
		return status.Errorf(codes.NotFound, "Service with ID %q not found.", id)
	}
	return nil
}