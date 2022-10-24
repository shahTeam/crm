CREATE TABLE schedules (
    id UUID NOT NULL PRIMARY KEY,
    group_id UUID NOT NULL,
    subject_id UUID NOT NULL,
    teacher_id UUID NOT NULL,
    weekday INT NOT NULL,
    lesson_number INT NOT NULL  
);

