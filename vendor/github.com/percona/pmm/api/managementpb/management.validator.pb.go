// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/management.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/percona/pmm/api/inventory"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddMySQLRequest) Validate() error {
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	if !(this.Port > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '0'`, this.Port))
	}
	if this.PmmAgentId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PmmAgentId", fmt.Errorf(`value '%v' must not be an empty string`, this.PmmAgentId))
	}
	return nil
}
func (this *AddMySQLResponse) Validate() error {
	if this.Service != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Service); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Service", err)
		}
	}
	if this.MysqldExporter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MysqldExporter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MysqldExporter", err)
		}
	}
	if this.QanMysqlPerfschema != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.QanMysqlPerfschema); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("QanMysqlPerfschema", err)
		}
	}
	return nil
}
func (this *RegisterNodeRequest) Validate() error {
	if this.NodeName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeName", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeName))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *RegisterNodeResponse) Validate() error {
	if this.GenericNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GenericNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GenericNode", err)
		}
	}
	if this.ContainerNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContainerNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContainerNode", err)
		}
	}
	if this.PmmAgent != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PmmAgent); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PmmAgent", err)
		}
	}
	return nil
}