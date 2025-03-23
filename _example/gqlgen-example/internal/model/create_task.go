package model

type CreateTasksInputItem struct {
	Name        string `json:"Name,omitempty" bson:"name"`
	Description string
	Status      Status
}

type Status string

const (
	StatusCreated    Status = "Created"
	StatusInProgress Status = "InProgress"
	StatusDone       Status = "Done"
	StatusRejected   Status = "Rejected"
)
