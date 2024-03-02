// mtd/memento.go
package mtd

import "github.com/gin-gonic/contrib/sessions"

// https://refactoring.guru/design-patterns/memento

var RegisteredPaths = make(map[string]bool)

// Caretaker holds the list of mementos
type Caretaker struct {
	MementoList []*SessionMemento
}

func (c *Caretaker) AddMemento(m *SessionMemento) {
	c.MementoList = append(c.MementoList, m)
}

func (c *Caretaker) GetMemento(index int) *SessionMemento {
	if index < 0 || index >= len(c.MementoList) {
		return nil // Or handle the error as appropriate
	}
	return c.MementoList[index]
}

// SessionMemento struct holds the state of the session
type SessionMemento struct {
	Data map[string]interface{}
}

func NewSessionMemento(data map[string]interface{}) *SessionMemento {
	return &SessionMemento{Data: data}
}

func (s *SessionMemento) GetData() map[string]interface{} {
	return s.Data
}

// SessionOriginator represents the session whose state we want to save and restore
type SessionOriginator struct {
	Session sessions.Session
}

func (so *SessionOriginator) SaveToSessionMemento() *SessionMemento {
	// Assume you have a special key in your session that holds a list of all other keys
	keysKey := "session_keys" // A reserved key to track all other keys in the session
	var keys []string
	if k, exists := so.Session.Get(keysKey).([]string); exists {
		keys = k
	}

	data := make(map[string]interface{})
	for _, k := range keys {
		if k != keysKey { // Avoid including the keys list itself
			data[k] = so.Session.Get(k)
		}
	}
	return NewSessionMemento(data)
}

func (so *SessionOriginator) RestoreFromSessionMemento(m *SessionMemento) {
	for k, v := range m.GetData() {
		so.Session.Set(k, v)
	}
	so.Session.Save()
}
