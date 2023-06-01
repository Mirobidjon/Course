package course

type Course struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Student_group int            `json:"groups"`
	File_url      map[int]string `json:"file_url"`
}

type InputCourse struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Group       int            `json:"groups"`
	FileUrl     map[int]string `json:"file_url"`
}

type UpdateCourse struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	FileUrl     map[int]string `json:"file_url"`
}
