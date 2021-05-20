package model

import (
	"errors"
	"fmt"
)

type SetterState struct {
	StateCommon `yaml:",inline"`
	Variables   []SetterDefinition `yaml:"variables"`
	Transform   string             `yaml:"transform,omitempty"`
	Transition  string             `yaml:"transition,omitempty"`
}

type SetterDefinition struct {
	Scope string `yaml:"scope"`
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

func (o *SetterDefinition) Validate() error {

	if o.Scope == "" {
		return errors.New(`scope required ("instance", "workflow", or "namespace")`)
	}

	if o.Key == "" {
		return errors.New(`key required`)
	}

	if o.Value == "" {
		return errors.New(`value required`)
	}

	if err := validateTransformJQ(o.Value); err != nil {
		return err
	}

	return nil

}

func (o *SetterState) GetID() string {
	return o.ID
}

func (o *SetterState) getTransitions() map[string]string {
	transitions := make(map[string]string)
	if o.Transition != "" {
		transitions["transition"] = o.Transition
	}

	for i, errDef := range o.ErrorDefinitions() {
		if errDef.Transition != "" {
			transitions[fmt.Sprintf("errors[%v]", i)] = errDef.Transition
		}
	}

	return transitions
}

func (o *SetterState) GetTransitions() []string {
	transitions := make([]string, 0)
	if o.Transition != "" {
		transitions = append(transitions, o.Transition)
	}

	for _, errDef := range o.ErrorDefinitions() {
		if errDef.Transition != "" {
			transitions = append(transitions, errDef.Transition)
		}
	}

	return transitions
}

func (o *SetterState) Validate() error {
	if err := o.commonValidate(); err != nil {
		return err
	}

	if len(o.Variables) == 0 {
		return errors.New("variables required")
	}

	for i, varDef := range o.Variables {
		if err := varDef.Validate(); err != nil {
			return fmt.Errorf("variables[%d] is invalid: %v", i, err)
		}
	}

	if err := validateTransformJQ(o.Transform); err != nil {
		return err
	}

	for i, errDef := range o.ErrorDefinitions() {
		if err := errDef.Validate(); err != nil {
			return fmt.Errorf("catch[%v] is invalid: %v", i, err)
		}
	}

	return nil
}