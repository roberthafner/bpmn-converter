package handler

import (
        "github.com/roberthafner/bpmn-model"
)

type ParseHandler interface {
	Handle (element model.BaseElement, bpmnModel model.BpmnModel)
}

type SequenceFlowParseHandler struct {
	ParseHandler
}

func (handler SequenceFlowParseHandler) Handle (element model.BaseElement, bpmnModel model.BpmnModel) {
	sequenceFlow := element.(model.SequenceFlow)
	sourceRefElement := bpmnModel.GetFlowElement(sequenceFlow.SourceRef)
	targetRefElement := bpmnModel.GetFlowElement(sequenceFlow.TargetRef)
	sequenceFlow.SourceFlowElement = sourceRefElement.(model.FlowNode)
	sequenceFlow.TargetFlowElement = targetRefElement.(model.FlowNode)
}
