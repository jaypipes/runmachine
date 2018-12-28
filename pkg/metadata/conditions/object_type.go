package conditions

import pb "github.com/runmachine-io/runmachine/proto"

type HasObjectType interface {
	GetObjectType() string
}

type ObjectTypeCondition struct {
	Op         Op
	ObjectType *pb.ObjectType
}

func (c *ObjectTypeCondition) Matches(obj HasObjectType) bool {
	if c == nil || c.ObjectType == nil {
		return true
	}
	cmp := obj.GetObjectType()
	switch c.Op {
	case OP_EQUAL:
		return c.ObjectType.Code == cmp
	case OP_NOT_EQUAL:
		return c.ObjectType.Code != cmp
	default:
		return false
	}
}
