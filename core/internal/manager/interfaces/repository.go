package interfaces

type IRepoClient interface {
}

type IRepoChat interface {
	//Create(member1, member2 int) ChatData
	//RecordMessage(msg ChatMessage) error
}

type Repository struct {
	IRepoChat
	IRepoClient
}
