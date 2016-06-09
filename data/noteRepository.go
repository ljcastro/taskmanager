package data

import (
	"time"

	"github.com/ljcastro/taskmanager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NoteRepository struct
type NoteRepository struct {
	C *mgo.Collection
}

// Create function
func (r *NoteRepository) Create(note *models.TaskNote) error {
	objID := bson.NewObjectId()
	note.Id = objID
	note.CreatedOn = time.Now()
	err := r.C.Insert(&note)
	return err
}

// Update function
func (r *NoteRepository) Update(note *models.TaskNote) error {
	// partial update on MongoDB
	err := r.C.Update(bson.M{"_id": note.Id},
		bson.M{"$set": bson.M{
			"description": note.Description,
		}})
	return err
}

// Delete function
func (r *NoteRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// GetByTask function
func (r *NoteRepository) GetByTask(id string) []models.TaskNote {
	var notes []models.TaskNote
	taskid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

// GetAll function
func (r *NoteRepository) GetAll() []models.TaskNote {
	var notes []models.TaskNote
	iter := r.C.Find(nil).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

// GetById function
func (r *NoteRepository) GetById(id string) (note models.TaskNote, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&note)
	return
}
