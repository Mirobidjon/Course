package course

type Course struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Student_group int    `json:"groups"`
}

type InputCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Group       int    `json:"groups"`
}

type UpdateCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
