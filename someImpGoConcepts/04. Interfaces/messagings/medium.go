package medium

type Medium interface {
	SendMessage(userID int64) error
}
