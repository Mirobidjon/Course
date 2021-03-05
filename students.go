package course

type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Group    int    `json:"group"`
}

type SignUpStudent struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Group    int    `json:"group"`
}

type GetStudents struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Groups   int    `json:"group"`
}

type UpdateStudent struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
