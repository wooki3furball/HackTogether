package mtd

import "github.com/gin-gonic/contrib/sessions"

// Memento Function // Prevent SQL Injection/Brute Force?

// Caretaker holds the list of mementos
type Caretaker struct {
	mementoList []*SessionMemento
}

func (c *Caretaker) AddMemento(m *SessionMemento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *Caretaker) GetMemento(index int) *SessionMemento {
	if index < 0 || index >= len(c.mementoList) {
		return nil // Or handle the error as appropriate
	}
	return c.mementoList[index]
}

// SessionMemento struct holds the state of the session
type SessionMemento struct {
	data map[string]interface{}
}

func NewSessionMemento(data map[string]interface{}) *SessionMemento {
	return &SessionMemento{data: data}
}

func (s *SessionMemento) GetData() map[string]interface{} {
	return s.data
}

// SessionOriginator represents the session whose state we want to save and restore
type SessionOriginator struct {
	session sessions.Session
}

func (so *SessionOriginator) SaveToSessionMemento() *SessionMemento {
	data := make(map[string]interface{})
	for _, k := range so.session.Keys() {
		data[k] = so.session.Get(k)
	}
	return NewSessionMemento(data)
}

func (so *SessionOriginator) RestoreFromSessionMemento(m *SessionMemento) {
	for k, v := range m.GetData() {
		so.session.Set(k, v)
	}
	so.session.Save()
}
