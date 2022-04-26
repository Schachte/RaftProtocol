package fsm

type FSMOperation int

const (
	ADD_REMINDER FSMOperation = FSMOperation(iota)
	DELETE_REMINDER
	RETRIEVE_REMINDERS
	COMPLETE_REMINDER
)

type Command struct {
	Operation FSMOperation
	Value     interface{}
}
