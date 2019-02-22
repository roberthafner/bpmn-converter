package converter

import (
        "github.com/roberthafner/bpmn-model"
	"github.com/roberthafner/bpmn-converter/parser"
	"github.com/roberthafner/bpmn-converter/parser/handler"
	"encoding/xml"
	"fmt"
	"io"
)

type BpmnXMLConverter struct {
	parsers  map[string]parser.Parser
	handlers map[string]handler.ParseHandler
}

func NewBpmnXmlConverter() BpmnXMLConverter {
	parsers := make(map[string]parser.Parser)
	parsers[model.ElementProcess] = parser.ProcessParser{}
	parsers[model.ElementStartEvent] = parser.StartEventParser{}
	parsers[model.ElementEndEvent] = parser.EndEventParser{}
	parsers[model.ElementUserTask] = parser.UserTaskParser{}
	parsers[model.ElementSequenceFlow] = parser.SequenceFlowParser{}

	handlers := make(map[string]handler.ParseHandler)
	handlers[model.ElementSequenceFlow] = handler.SequenceFlowParseHandler{}

	return BpmnXMLConverter{
		parsers:  parsers,
		handlers: handlers,
	}
}

func (converter BpmnXMLConverter) ConvertToBpmnModel(reader io.Reader) model.BpmnModel {
	decoder := xml.NewDecoder(reader)
	bpmnModel := model.BpmnModel{}
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {

		case xml.StartElement:
			parser := converter.parsers[v.Name.Local]
			if parser != nil {
				bpmnModel = parser.Parse(t, *decoder, bpmnModel)
			}
			break
		}
	}

	for i := 0; i < len(bpmnModel.Processes); i++ {
		process := bpmnModel.Processes[i]
		flowElements := process.FlowElements
		for j := 0; j < len(flowElements); j++ {
			flowElement := flowElements[j]
			switch v := flowElement.(type) {
			case *model.SequenceFlow:
				h := converter.handlers[model.ElementSequenceFlow]
				h.Handle(v, bpmnModel)
			}
		}
	}

	return bpmnModel
}

func (convert BpmnXMLConverter) ConvertToXML(model model.BpmnModel) []byte {
	return nil
}

//type BpmnXMLConverter interface {
//	convertToModel(decoder xml.Decoder, model model.BpmnModel, process model.Process)
//	convertToXML(encoder xml.Encoder, element model.BaseElement, model model.BpmnModel)
//
//	convertElementToXML(encoder xml.Encoder, bpmnModel model.BpmnModel)
//	convertXMLToElement(decoder xml.Decoder, model model.BpmnModel) model.BaseElement
//}

//type BaseBpmnXMLConverter struct {
//	BpmnXMLConverter
//}
