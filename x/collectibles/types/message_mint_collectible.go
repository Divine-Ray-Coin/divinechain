package types

func NewMsgMintCollectible(creator string, classId string, uri string, receiver string) *MsgMintCollectible {
	return &MsgMintCollectible{
		Creator:  creator,
		ClassId:  classId,
		Uri:      uri,
		Receiver: receiver,
	}
}
