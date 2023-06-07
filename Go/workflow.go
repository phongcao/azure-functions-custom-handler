package main

import (
	"errors"
	"reflect"

	"github.com/serverlessworkflow/sdk-go/v2/model"
	"github.com/serverlessworkflow/sdk-go/v2/parser"
)

func ParseWorkflow(yaml []byte) (*model.Workflow, error) {
	workflow, err := parser.FromYAMLSource(yaml)
	if err != nil {
		return nil, err
	}
	return workflow, nil
}

func ExecuteStateAndReturnNext(states []model.State, currentState model.State) (model.State, error) {
	operationState := currentState.OperationState
	if operationState == nil {
		return currentState, errors.New("Not yet implemented")
	}

	action := operationState.Actions[0]
	methodName := action.FunctionRef.RefName
	method := reflect.ValueOf(TaskType(0)).MethodByName(methodName)
	if !method.IsValid() {
		return currentState, errors.New("Task not found")
	} else {
		method.Call(nil)
	}

	baseState := currentState.BaseState
	transition := baseState.Transition
	if transition == nil {
		return currentState, nil
	}

	for i := 1; i < len(states); i++ {
		if states[i].Name == transition.NextState {
			return ExecuteStateAndReturnNext(states, states[i])
		}
	}

	return currentState, nil
}
