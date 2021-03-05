package course

type Master struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type MasterInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type OutputMaster struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
