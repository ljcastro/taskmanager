package data

import (
	"time"

	"github.com/ljcastro/taskmanager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TaskRepository struct definition
type TaskRepository struct {
	C *mgo.Collection
}

// Create Task function
func (r *TaskRepository) Create(task *models.Task) error {
	objID := bson.NewObjectId()
	task.Id = objID
	task.CreatedOn = time.Now()
	task.Status = "Created"
	err := r.C.Insert(&task)
	return err
}

// Update Task function
func (r *TaskRepository) Update(task *models.Task) error {
	// partial update on MongoDB
	err := r.C.Update(bson.M{"_id": task.Id},
		bson.M{"$set": bson.M{
			"name":        task.Name,
			"description": task.Description,
			"due":         task.Due,
			"status":      task.Status,
			"tags":        task.Tags,
		}})
	return err
}

// Delete Task function
func (r *TaskRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// GetAll tasks function
func (r *TaskRepository) GetAll() []models.Task {
	var tasks []models.Task
	iter := r.C.Find(nil).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}

// GetById tasks function
func (r *TaskRepository) GetById(id string) (task models.Task, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&task)
	return
}

// GetByUser tasks function
func (r *TaskRepository) GetByUser(user string) []models.Task {
	var tasks []models.Task
	iter := r.C.Find(bson.M{"createdby": user}).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}
