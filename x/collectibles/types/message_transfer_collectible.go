package types

func NewMsgTransferCollectible(creator string, classId string, id string, receiver string) *MsgTransferCollectible {
	return &MsgTransferCollectible{
		Creator:  creator,
		ClassId:  classId,
		Id:       id,
		Receiver: receiver,
	}
}
