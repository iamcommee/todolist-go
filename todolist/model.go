package todolist

type Todolists struct {
	Todolists []Todolist `json:"todolists"`
}

type Todolist struct {
	Id   string `json:"id,omitempty" bson:"_id,omitempty"`
	Task string `json:"task"`
}
