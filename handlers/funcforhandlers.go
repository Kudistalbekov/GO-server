package handlers

import (
	"projects/server/models"
	"sync"
)

func change(d *models.User, new *models.User) {
	w := &sync.WaitGroup{}
	w.Add(7)
	go func() {
		if (*new).Name != "" {
			(*d).Name = (*new).Name
		}
		w.Done()
	}()
	go func() {
		if (*new).Password != "" {
			(*d).Password = (*new).Password
		}
		w.Done()
	}()
	go func() {
		if (*new).Surname.IsZero() != true {
			(*d).Surname = (*new).Surname
		}
		w.Done()
	}()
	go func() {
		if (*new).Username.IsZero() != true {
			(*d).Username = (*new).Username
		}
		w.Done()
	}()
	go func() {
		if (*new).Email != "" {
			(*d).Email = (*new).Email
		}
		w.Done()
	}()
	go func() {
		(*d).RegisterDate = (*d).RegisterDate
		w.Done()
	}()
	go func() {
		(*d).ID = (*d).ID
		w.Done()
	}()
	w.Wait()
}
