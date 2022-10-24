package server

import (
	"context"
	"teacher/domain/subject"
	"teacher/domain/teacher"
	"teacher/service"

	"github.com/google/uuid"
	"github.com/shahTeam/crmprotos/teacherpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func New(svc service.Service, subjectFactory subject.Factory, teacherFactory teacher.Factory) Server {
	return Server{
		UnimplementedTeacherServiceServer: teacherpb.UnimplementedTeacherServiceServer{},
		service: service.Service{},
		subjectFactory: subjectFactory,
		teacherFactory: teacherFactory,
	}
} 

type Server struct {
	teacherpb.UnimplementedTeacherServiceServer
	service        service.Service
	subjectFactory subject.Factory
	teacherFactory teacher.Factory
}

func (s Server) RegisterTeacher(ctx context.Context, request *teacherpb.RegisterTeacherRequest) (*teacherpb.Teacher, error) {
	teacher, err := s.convertRegisterTeacherRequestToDomainTeacher(request)
	if err != nil {
		return nil, err
	}
	creatTeacher, err := s.service.RegisterTeacher(ctx, teacher)
	if err != nil {
		return nil, err
	}
	return toProtoTeacher(creatTeacher), nil
} 

func (s Server) CreateSubject(ctx context.Context, req *teacherpb.CreateSubjectRequest) (*teacherpb.Subject, error) {
	sub, err := s.subjectFactory.NewSubject(req.Name, req.Description)
	if err != nil {
		return &teacherpb.Subject{}, err
	}
	sub, err = s.service.CreateSubject(ctx, sub)
	if err != nil {
		return &teacherpb.Subject{}, err
	} 
	return toProtoSubject(sub), nil
}

func (s Server) GetTeacher(ctx context.Context, req *teacherpb.GetTeacherRequest) (*teacherpb.Teacher, error) {
	var teacherBy teacher.By
	switch by := req.By.(type) {
	case *teacherpb.GetTeacherRequest_Email:
		teacherBy = teacher.ByEmail{Email: by.Email}
	case *teacherpb.GetTeacherRequest_PhoneNumber:
		teacherBy = teacher.ByPhoneNumber{PhoneNumber: by.PhoneNumber}
	case *teacherpb.GetTeacherRequest_TeacherId:
		id, err := uuid.Parse(by.TeacherId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		teacherBy = teacher.ByID{ID: id}
	default:
		return nil, status.Error(codes.InvalidArgument, "by is not provided")
	}

	t, err := s.service.GetTeacher(ctx, teacherBy)
	if err != nil {
		return &teacherpb.Teacher{}, err
	}
	return toProtoTeacher(t), nil
}

func (s Server) GetSubject(ctx context.Context, req *teacherpb.GetSubjectRequest) (*teacherpb.Subject, error) {
	var subjectBy subject.By
	switch sub := req.By.(type) {
	case *teacherpb.GetSubjectRequest_SubjectId:
		id, err := uuid.Parse(sub.SubjectId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "id is not uuid")			
		}
		subjectBy = subject.ByID{ID: id}
	default:
		return nil, status.Error(codes.InvalidArgument, "by is not provided")
	case *teacherpb.GetSubjectRequest_Name:
		subjectBy = subject.ByName{Name: sub.Name}
	}

	sub, err := s.service.GetSubject(ctx, subjectBy)
	if err != nil {
		return nil, err
	}
	return toProtoSubject(sub), nil
} 

func (s Server) DeleteTeacher(ctx context.Context, req *teacherpb.DeleteTeacherRequest) (*emptypb.Empty, error) {
	var teacherBy teacher.By
	switch t := req.By.(type) {
	case *teacherpb.DeleteTeacherRequest_Email:
		teacherBy = teacher.ByEmail{Email: t.Email}
	case *teacherpb.DeleteTeacherRequest_PhoneNumber:
		teacherBy = teacher.ByPhoneNumber{PhoneNumber: t.PhoneNumber}
	case *teacherpb.DeleteTeacherRequest_TeacherId:
		id, err := uuid.Parse(t.TeacherId)
		if err != nil {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		teacherBy = teacher.ByID{ID: id}
	default:
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "by is not provided")
	}

	if err := s.service.DeleteTeacher(ctx, teacherBy); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s Server) DeleteSubject(ctx context.Context, req *teacherpb.DeleteSubjectRequest) (*emptypb.Empty, error) {
	var subjectBy subject.By
	switch sub := req.By.(type) {
	case *teacherpb.DeleteSubjectRequest_Name:
		subjectBy = subject.ByName{Name: sub.Name}
	case *teacherpb.DeleteSubjectRequest_SubjectId:
		id, err := uuid.Parse(sub.SubjectId)
		if err != nil {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		subjectBy = subject.ByID{ID: id}
	default:
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "by is not provided")
	}

	if err := s.service.DeleteSubject(ctx, subjectBy); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s Server) ListTeachers(ctx context.Context, req *teacherpb.ListTeachersRequest) (*teacherpb.ListTeachersResponse, error) {
	list, _, err := s.service.ListTeacher(ctx, req.Page, req.Limit)
	if err != nil {
		return &teacherpb.ListTeachersResponse{}, status.Error(codes.Internal, err.Error())
	}

	protoTech := make([]*teacherpb.Teacher, 0, len(list))
	for _, item := range list {
		prt := toProtoTeacher(item)
		protoTech = append(protoTech, prt)
	}
	return &teacherpb.ListTeachersResponse{
		Teachers: protoTech,
	}, nil
}

func (s Server) ListSubjects(ctx context.Context, req *teacherpb.ListSubjectsRequest) (*teacherpb.ListSubjectsResponse, error) {
	list, _, err := s.service.ListSubject(ctx, req.Page, req.Limit)
	if err != nil {
		return &teacherpb.ListSubjectsResponse{}, status.Error(codes.Internal, err.Error())
	}

	prtoSub := make([]*teacherpb.Subject, 0, len(list))
	for _, item := range list {
		sub := toProtoSubject(item)
		prtoSub = append(prtoSub, sub)
	}
	return &teacherpb.ListSubjectsResponse{
		Subject: prtoSub,
	}, nil
}

func (s Server) convertRegisterTeacherRequestToDomainTeacher(toProtoTeacher *teacherpb.RegisterTeacherRequest) (teacher.Teacher, error) {
	subjectID, err := uuid.Parse(toProtoTeacher.SubjectId)
	if err != nil {
		return teacher.Teacher{}, status.Error(codes.InvalidArgument, "provided subject id is not uuid")
	}
	t, err := s.teacherFactory.NewFactory(
		toProtoTeacher.FirstName,
		toProtoTeacher.LastName,
		toProtoTeacher.Email,
		toProtoTeacher.PhoneNumber,
		toProtoTeacher.Password,
		subjectID,
	)
	if err != nil {
		return teacher.Teacher{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return t, nil
}
