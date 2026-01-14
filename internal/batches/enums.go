package batches

type Status string

const (
	StatusCompleted              Status = "COMPLETED"
	StatusCreated                Status = "CREATED"
	StatusDiscarded              Status = "DISCARDED"
	StatusPendingManagerApproval Status = "PENDING_MANAGER_APPROVAL"
)
