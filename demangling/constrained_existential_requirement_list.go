package demangling

import "fmt"

func (ctx *Context) constrainedExistentialRequirementList() (*Node, error) {
	reqList := createNode(ConstrainedExistentialRequirementListKind)
	firstElem := false
	for {
		firstElem = ctx.popNodeKind(FirstElementMarkerKind) != nil
		req := ctx.popNodePred(isRequirement)
		if req == nil {
			return nil, fmt.Errorf("constrainedExistentialRequirementList: req is nil")
		}
		reqList.addChild(req)
		if firstElem {
			break
		}
	}
	reqList.reverseChildren(0)
	return reqList, nil
}
