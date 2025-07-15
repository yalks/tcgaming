package tcgaming

import "fmt"

// Exception types matching Java SDK

// ProcessException represents business logic errors from the API
type ProcessException struct {
	Message string
	Status  int
}

func (e *ProcessException) Error() string {
	return fmt.Sprintf("ProcessException: %s (status: %d)", e.Message, e.Status)
}

// TransportException represents network/transport layer errors
type TransportException struct {
	Message string
	Cause   error
}

func (e *TransportException) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("TransportException: %s (cause: %v)", e.Message, e.Cause)
	}
	return fmt.Sprintf("TransportException: %s", e.Message)
}

// RemoteException represents remote server errors
type RemoteException struct {
	Message string
	Status  int
}

func (e *RemoteException) Error() string {
	return fmt.Sprintf("RemoteException: %s (status: %d)", e.Message, e.Status)
}

// Helper functions to create exceptions

func NewProcessException(message string, status int) *ProcessException {
	return &ProcessException{
		Message: message,
		Status:  status,
	}
}

func NewTransportException(message string, cause error) *TransportException {
	return &TransportException{
		Message: message,
		Cause:   cause,
	}
}

func NewRemoteException(message string, status int) *RemoteException {
	return &RemoteException{
		Message: message,
		Status:  status,
	}
}