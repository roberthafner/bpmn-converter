package parser

import (
        "github.com/roberthafner/bpmn-model"
	"encoding/xml"
)

type Parser interface {
	Parse (token xml.Token, decoder xml.Decoder, bpmnModel model.BpmnModel) model.BpmnModel
}

type ProcessParser struct {
	Parser
}

func (parser ProcessParser) Parse (token xml.Token, decoder xml.Decoder, bpmnModel model.BpmnModel) model.BpmnModel{
	processElement := token.(xml.StartElement)
	attributeMap := readAttributes(processElement.Attr)
	// TODO: parse process documentation.
	process := model.NewProcess(attributeMap[model.AttributeId], attributeMap[model.AttributeName], nil)
	bpmnModel = bpmnModel.Add(process)
	return bpmnModel
}

type SequenceFlowParser struct {
	Parser
}

func (parser SequenceFlowParser) Parse (token xml.Token, decoder xml.Decoder, bpmnModel model.BpmnModel) model.BpmnModel {
	sequenceFlowElement := token.(xml.StartElement)
	attributeMap := readAttributes(sequenceFlowElement.Attr)

	sequenceFlow := model.NewSequenceFlow(attributeMap[model.AttributeId], attributeMap[model.AttributeName],
		nil, attributeMap[model.AttributeSourceRef], attributeMap[model.AttributeTargetRef])

	process := bpmnModel.CurrentProcess()
	process = process.Add(sequenceFlow)
	bpmnModel.SetCurrentProcess(process)
	return bpmnModel
}

type StartEventParser struct {
	Parser
}

func (parser StartEventParser) Parse (token xml.Token, decoder xml.Decoder, bpmnModel model.BpmnModel) model.BpmnModel {
	startEventElement := token.(xml.StartElement)
	attributeMap := readAttributes(startEventElement.Attr)

	startEvent := model.NewStartEvent(attributeMap[model.AttributeId], attributeMap[model.AttributeName], nil)
	process := bpmnModel.CurrentProcess()
	process = process.Add(startEvent)
	bpmnModel.SetCurrentProcess(process)
	return bpmnModel
}

type UserTaskParser struct {
	Parser
}

func (parser UserTaskParser) Parse (token xml.Token, decoder xml.Decoder, bpmnModel model.BpmnModel) model.BpmnModel {
	userTaskElement := token.(xml.StartElement)
	attributeMap := readAttributes(userTaskElement.Attr)

	userTask := model.NewUserTask(attributeMap[model.AttributeId], attributeMap[model.AttributeName], nil)
	process := bpmnModel.CurrentProcess()
	process = process.Add(userTask)
	bpmnModel.SetCurrentProcess(process)
	return bpmnModel
}

type EndEventParser struct {
	Parser
}

func (parser EndEventParser) Parse (token xml.Token, decoder xml.Decoder, bpmnModel model.BpmnModel) model.BpmnModel {
	endEventElement := token.(xml.StartElement)
	attributeMap := readAttributes(endEventElement.Attr)

	endEvent := model.NewEndEvent(attributeMap[model.AttributeId], attributeMap[model.AttributeName], nil)
	process := bpmnModel.CurrentProcess()
	process = process.Add(endEvent)
	bpmnModel.SetCurrentProcess(process)
	return bpmnModel
}

func readAttributes(attributes []xml.Attr) map[string]string {
	attributeMap := make(map[string]string, len(attributes))
	for _, attr := range attributes {
		attributeMap[attr.Name.Local] = attr.Value
	}
	return attributeMap
}
