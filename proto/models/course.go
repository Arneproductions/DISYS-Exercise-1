/*
 * Mandatory exercise 1
 *
 * Mandatory exercse 1
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

import (
	pb "github.com/ArneProductions/DISYS-exercise-1/proto/course"
)

type Course struct {
	Id uint64 `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Teacher uint64 `json:"teacher,omitempty"`

	Students []User `json:"students,omitempty" gorm:"many2many:course_students;" `
}

func ProtoToModel(proto *pb.Course) Course {
	return Course{
		Id:          proto.Id,
		Name:        proto.Name,
		Description: proto.Description,
		Teacher:     proto.Teacher,
		Students:    IntArrayToUser(proto.Students),
	}
}

func ModelToProto(model Course) *pb.Course {
	val := pb.Course{
		Id:          model.Id,
		Name:        model.Name,
		Description: model.Description,
		Teacher:     model.Teacher,
		Students:    UserArrayToInt(model.Students),
	}

	return &val
}

func ProtoArrayToModel(protos []*pb.Course) []Course {
	models := make([]Course, len(protos))

	for i, value := range protos {
		models[i] = ProtoToModel(value)
	}

	return models
}

func ModelArrayToProto(models []Course) []*pb.Course {
	protos := make([]*pb.Course, len(models))

	for i, value := range models {
		protos[i] = ModelToProto(value)
	}

	return protos
}
