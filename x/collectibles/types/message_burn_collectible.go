package types

func NewMsgBurnCollectible(creator string, classId string, id string) *MsgBurnCollectible {
	return &MsgBurnCollectible{
		Creator: creator,
		ClassId: classId,
		Id:      id,
	}
}
