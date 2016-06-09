package controllers

import (
	"github.com/ljcastro/taskmanager/models"
)

// Models for JSON Users resources
type (
	// UserResource struct for POST - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	// LoginResource struct for POST - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	// AuthUserResource - Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// LoginModel for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// AuthUserModel for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)

// Models for JSON Tasks resources
type (
	// TaskResource for GET - /tasks/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	// TasksResource for GET - /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}
)

// Models for JSON Notes resources
type (
	// NoteResource for POST/PUT - /notes
	NoteResource struct {
		Data NoteModel `json:"data"`
	}
	// NotesResource for GET - /notes AND /notes/tasks/id
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}
	// NoteModel for a TaskNote
	NoteModel struct {
		TaskId      string `json:"taskid"`
		Description string `json:"description"`
	}
)
