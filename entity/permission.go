package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PermissionType string

const (
	CreateTask       PermissionType = "Create Task"
	ManageTask       PermissionType = "Manage Task"
	DeleteTask       PermissionType = "Deleted Task"
	StartDateEndDate PermissionType = "Start date - End date"
	DragAndDropTask  PermissionType = "Drag & Drop task"
	MoveTask         PermissionType = "Move task"
	Files            PermissionType = "Files"
	Comment          PermissionType = "Comment"
	CreateActivity   PermissionType = "Create Activity"
	EditActivity     PermissionType = "Edit Activity"
	DeleteActivity   PermissionType = "Delete Activity"
)

type Permission struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Desc      string             `json:"desc" bson:"desc"`
	Icon      string             `json:"icon" bson:"icon"`
	Action    string             `json:"action" bson:"action"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
