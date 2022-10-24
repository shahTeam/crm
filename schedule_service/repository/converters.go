package repository

import "schedule/domain/schedules"

func convertToDomainSchedules(schedule []Schedule) ([]schedules.Schedule, error) {
	schs := make([]schedules.Schedule, 0, len(schedule))
	for _, item := range schedule {
		sch, err := schedules.UnmarshalSchedule(schedules.UnmarshalArgs(item))
		if err != nil {
			return []schedules.Schedule{}, err
		}
		schs = append(schs, sch)
	}
	return schs, nil
}

func toRepositorySchedule(sch schedules.Schedule) (Schedule) {
	return Schedule{
          ID: sch.ID(),
		  GroupID: sch.GroupID(),
		  SubjectID: sch.SubjectID(),
		  TeacherID: sch.TeacherID(),
		  Weekday: sch.Weekday(),
		  LessonNumber: sch.LessonNumber(),
	}
}