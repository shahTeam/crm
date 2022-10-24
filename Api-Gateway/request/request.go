package request

type RegisterTeacherRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	SubjectID   string `json:"subject_id"`
}

type CreateSubjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetTeacherRequest struct {
	TeacherID   string `json:"teacher_id, omitempty"`
	Email       string `json:"email, omitempty"`
	PhoneNumber string `json:"phone_number, omitempty"`
}

type DeleteTeacherRequest struct {
	TeacherID   string `json:"teacher_id, omitempty"`
	Email       string `json:"email, omitempty"`
	PhoneNumber string `json:"phone_number, omitempty"`
}

type GetSubjectRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteSubjectRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
