package converter

import (
	"strings"
	"testing"
)

func TestConverter(t *testing.T) {
	definition :=
		`<definitions>
			<process id="1" name="test">
				<startEvent id="2" name="start"/>
				<sequenceFlow id="3" name="sequence flow 3" sourceRef="2" targetRef="4"/>
               	<userTask id="4" name="usertask"/>"
				<sequenceFlow id="5" name="sequence flow 5" sourceRef="4" targetRef="6"/>
				<endEvent id="6" name="end"/>
			</process>
		</definitions>`

	converter := NewBpmnXmlConverter()
	bpmnModel := converter.ConvertToBpmnModel(strings.NewReader(definition))

	processes := bpmnModel.Processes()
	if 1 != len(processes) {
		t.Fail()
	}

	process := processes[0]
	if 5 != len(process.FlowElements) {
		t.Fail()
	}

	startEvent := process.GetFlowElement("2")
	if startEvent.Name() != "start" {
		t.Fail()
	}

	sequenceFlow3 := process.GetFlowElement("3")
	if sequenceFlow3.Name() != "sequence flow 3" {
		t.Fail()
	}

	userTask := process.GetFlowElement("4")
	if userTask.Name() != "usertask" {
		t.Fail()
	}

	sequenceFlow5 := process.GetFlowElement("5")
	if sequenceFlow5.Name() != "sequence flow 5" {
		t.Fail()
	}

	endEvent := process.GetFlowElement("6")
	if endEvent.Name() != "end" {
		t.Fail()
	}
}
