package server

import (
	"context"
	"errors"
	"student/domain/group"
	"student/domain/student"
	"student/service"

	"github.com/google/uuid"
	"github.com/shahTeam/crmconnect/errs"
	"github.com/shahTeam/crmprotos/studentpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func New(svc service.Service, studentFactory student.Factory, groupFactory group.Factory) Server {
	return Server{
		service:        svc,
		studentFactory: studentFactory,
		groupFactory:   groupFactory,
	}
}

type Server struct {
	studentpb.UnimplementedStudentServiceServer
	service        service.Service
	studentFactory student.Factory
	groupFactory   group.Factory
}

func (s Server) GetGroupStudents(ctx context.Context, req *studentpb.GetGroupStudentsRequest) (*studentpb.StudentList, error) {
	groupID, err := uuid.Parse(req.GetGroupId())
	if err != nil {
		return &studentpb.StudentList{}, status.Error(codes.InvalidArgument, "group id is not uuid")
	}
	students, err := s.service.GetStudentsByGroup(ctx, groupID)
	if err != nil {
		return &studentpb.StudentList{}, status.Error(codes.Internal, err.Error())
	}
	return toProtoStudents(students), nil
}

//ListGroups fetches list of groups
func (s Server) ListGroups(ctx context.Context, req *studentpb.ListGroupsRequest) (*studentpb.GroupList, error) {
	groups, _, err := s.service.ListGroup(ctx, req.Page, req.Limit)
	if err != nil {
		return &studentpb.GroupList{}, status.Error(codes.Internal, err.Error())
	}
	return toProtoGroups(groups), nil
}

//Delete Group deletes group by ID or Name
func (s Server) DeleteGroup(ctx context.Context, req *studentpb.DeleteGroupRequest) (*emptypb.Empty, error) {
	var groupBY group.GroupBy
	switch by := req.By.(type) {
	case *studentpb.DeleteGroupRequest_Name:
		groupBY = group.ByName{Name: by.Name}
	case *studentpb.DeleteGroupRequest_Id:
		id, err := uuid.Parse(req.GetId())
		if err != nil {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		groupBY = group.ByID{ID: id}
	default:
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "by is not provided")
	}
	if err := s.service.DeleteGroup(ctx, groupBY); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "group is not found")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil

}

func (s Server) UpdateGroup(ctx context.Context, req *studentpb.Group) (*studentpb.Group, error) {
	group, err := s.convertUpdateGroupRequestToDomainGroup(req)
	if err != nil {
		return &studentpb.Group{}, err
	}

	if err := s.service.UpdateGroup(ctx, group); err != nil {
		return &studentpb.Group{}, status.Error(codes.Internal, err.Error())
	}
	return toProtoGroup(group), nil
}

// ListStudents fetches a list of students from database
func (s Server) ListStudents(ctx context.Context, req *studentpb.ListStudentRequest) (*studentpb.StudentList, error) {
	students, _, err := s.service.ListStudent(ctx, req.Limit, req.Page)
	if err != nil {
		return &studentpb.StudentList{}, status.Error(codes.Internal, err.Error())
	}
	return toProtoStudents(students), nil
}

//DeleteStudent deltes by id, email, or phone_number
func (s Server) DeleteStudent(ctx context.Context, req *studentpb.DeleteStudentRequest) (*emptypb.Empty, error) {
	var std student.StudentBy
	switch by := req.By.(type) {
	case *studentpb.DeleteStudentRequest_Email:
		std = student.ByEamil{Email: by.Email}
	case *studentpb.DeleteStudentRequest_PhoneNumber:
		std = student.ByPhoneNumber{ByPhoneNumber: by.PhoneNumber}
	case *studentpb.DeleteStudentRequest_Id:
		id, err := uuid.Parse(req.GetId())
		if err != nil {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		std = student.ByID{ID: id}
	default:
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "by is not provided")
	}
	if err := s.service.DeleteStudent(ctx, std); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "students is not found")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

//GetGroup fetches Group data from database by id or name
func (s Server) GetGroup(ctx context.Context, req *studentpb.GetGroupRequest) (*studentpb.Group, error) {
	var gr group.GroupBy
	switch by := req.By.(type) {
	case *studentpb.GetGroupRequest_Name:
		gr = group.ByName{Name: by.Name}
	case *studentpb.GetGroupRequest_Id:
		id, err := uuid.Parse(req.GetId())
		if err != nil {
			return &studentpb.Group{}, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		gr = group.ByID{ID: id}
	default:
		return &studentpb.Group{}, status.Error(codes.Internal, "by is not provided")
	}
	group, err := s.service.GetGroup(ctx, gr)
	if err != nil {
		return &studentpb.Group{}, status.Error(codes.Internal, err.Error())
	}
	return toProtoGroup(group), nil
}

// GetStudent fetches student data from database by id, email or phone_number
func (s Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.StudentResponse, error) {
	var st student.StudentBy
	switch by := req.By.(type) {
	case *studentpb.GetStudentRequest_Email:
		st = student.ByEamil{Email: by.Email}
	case *studentpb.GetStudentRequest_PhoneNumber:
		st = student.ByPhoneNumber{ByPhoneNumber: by.PhoneNumber}
	case *studentpb.GetStudentRequest_Id:
		id, err := uuid.Parse(req.GetId())
		if err != nil {
			return &studentpb.StudentResponse{}, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		st = student.ByID{ID: id}
	default:
		return &studentpb.StudentResponse{}, status.Error(codes.Internal, "by is not provided")
	}
	student, err := s.service.GetStudent(ctx, st)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return &studentpb.StudentResponse{}, status.Error(codes.NotFound, "studentd is not found")
		}
		return &studentpb.StudentResponse{}, status.Error(codes.Internal, err.Error())
	}
	return toProtoStudent(student), nil
}

//UpdateStudent updates student's info
func (s Server) UpdateStudent(ctx context.Context, req *studentpb.StudentResponse) (*studentpb.StudentResponse, error) {
	student, err := s.convertUpdateStudentRequestToDomainStudent(req)
	if err != nil {
		return &studentpb.StudentResponse{}, err
	}

	if err := s.service.UpdateStudent(ctx, student); err != nil {
		return &studentpb.StudentResponse{}, status.Error(codes.Unavailable, "student data is not true")
	}
	return toProtoStudent(student), nil

}

//RegisterStudent creates a new student
func (s Server) RegisterStudent(ctx context.Context, req *studentpb.StudentRequest) (*studentpb.StudentResponse, error) {
	stt, err := s.convertRegisterStudentRequestToDomainStudent(req)
	if err != nil {
		return &studentpb.StudentResponse{}, err
	}
	if err := s.service.RegisterStudent(ctx, stt); err != nil {
		return &studentpb.StudentResponse{}, status.Error(codes.Internal, err.Error())
	}
	return toProtoStudent(stt), nil
}

//CreateGroup creates a new group
func (s Server) CreateGroup(ctx context.Context, req *studentpb.CreateGroupRequest) (*studentpb.Group, error) {
	grr, err := s.convertRegisterGroupRequestToDomainGroup(req)
	if err != nil {
		return &studentpb.Group{}, status.Error(codes.Internal, err.Error())
	}
	if err := s.service.CreateGroup(ctx, grr); err != nil {
		return &studentpb.Group{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return toProtoGroup(grr), nil
}


func (s Server) convertUpdateGroupRequestToDomainGroup(protoGroup *studentpb.Group) (group.Group, error) {
	id, err := uuid.Parse(protoGroup.GetId())
	if err != nil {
		return group.Group{}, status.Error(codes.InvalidArgument, "group id is not uuid")
	}

	mainTeacherID, err := uuid.Parse(protoGroup.MainTeacherId)
	if err != nil {
		return group.Group{}, status.Error(codes.InvalidArgument, "mainteacher_id is not uuid")
	}
	gr, err := group.UnmarshalGroup(group.UnmarshalGroupArgs{
		ID:            id,
		Name:          protoGroup.Name,
		MainTeacherID: mainTeacherID,
	})
	if err != nil {
		return group.Group{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return gr, nil
}

func (s Server) convertRegisterGroupRequestToDomainGroup(protoGroup *studentpb.CreateGroupRequest) (group.Group, error) {
	mainTeacherID, err := uuid.Parse(protoGroup.MainTeacherId)
	if err != nil {
		return group.Group{}, status.Error(codes.InvalidArgument, "mainteacher_id is not uuid")
	}
	gr, err := group.UnmarshalGroup(group.UnmarshalGroupArgs{
		Name:          protoGroup.Name,
		MainTeacherID: mainTeacherID,
	})
	if err != nil {
		return group.Group{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return gr, nil
}

func (s Server) convertUpdateStudentRequestToDomainStudent(protoStudent *studentpb.StudentResponse) (student.Student, error) {
	id, err := uuid.Parse(protoStudent.GetId())
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, "id is not uuid")
	}
	groupID, err := uuid.Parse(protoStudent.GetGroupId())
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, "groupID is not uuid")
	}
	st, err := student.UnmarshalStudent(student.UnmarshalStudentArgs{
		ID:          id,
		FirstName:   protoStudent.FirstName,
		LastName:    protoStudent.LastName,
		Email:       protoStudent.Email,
		PhoneNumber: protoStudent.PhoneNumber,
		Level:       protoStudent.Level,
		Password:    protoStudent.Password,
		GroupID:     groupID,
	})
	if err != nil {
		return student.Student{}, status.Error(codes.Internal, err.Error())
	}
	return st, nil
}

func (s Server) convertRegisterStudentRequestToDomainStudent(protoStudent *studentpb.StudentRequest) (student.Student, error) {
	groupID, err := uuid.Parse(protoStudent.GetGroupId())
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, "groupID is not uuid")
	}
	st, err := student.UnmarshalStudent(student.UnmarshalStudentArgs{
		FirstName:   protoStudent.FirstName,
		LastName:    protoStudent.LastName,
		Email:       protoStudent.Email,
		PhoneNumber: protoStudent.PhoneNumber,
		Level:       protoStudent.Level,
		Password:    protoStudent.Password,
		GroupID:     groupID,
	})
	if err != nil {
		return student.Student{}, status.Error(codes.Internal, err.Error())
	}
	return st, nil
}
