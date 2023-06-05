import "./styles.css";

import {
  BaseInputTemplateProps,
  IconButtonProps,
  RJSFSchema,
  RegistryWidgetsType,
  SubmitButtonProps,
  TitleFieldProps,
  WidgetProps,
  getSubmitButtonOptions,
} from "@rjsf/utils";
import {
  ChevronDownIcon,
  ChevronUpIcon,
  MinusIcon,
  PlusIcon,
} from "lucide-react";
import React, { useMemo } from "react";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../Select";

import Button from "../Button";
import { Checkbox } from "../Checkbox";
import Form from "@rjsf/core";
import Input from "../Input";
import validator from "@rjsf/validator-ajv8";

const CustomSelectWidget: React.FunctionComponent<WidgetProps> = (props) => (
  <div className="my-4">
    <Select onValueChange={props.onChange}>
      <SelectTrigger>
        <SelectValue placeholder={`Select ${props.label}`} />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {props.options.enumOptions?.map((op) => (
            <SelectItem key={`select-${op.value}`} value={op.value}>
              {op.label}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  </div>
);
function SubmitButton(props: SubmitButtonProps) {
  const { uiSchema } = props;
  const { norender } = getSubmitButtonOptions(uiSchema);
  if (norender) {
    return null;
  }
  return (
    <Button type="submit" className="float-right mt-4">
      Submit
    </Button>
  );
}
function AddButton(props: IconButtonProps) {
  const { ...btnProps } = props;
  return (
    <div className="inline-block w-full divide-y divide-solid">
      <Button
        {...btnProps}
        icon
        className="float-right mt-4 w-1/4 bg-danger-10"
      >
        <PlusIcon />
      </Button>
      <div className="mt-2 w-full" />
    </div>
  );
}

function RemoveButton(props: IconButtonProps) {
  const { ...btnProps } = props;
  return (
    <Button {...btnProps} icon>
      <MinusIcon />
    </Button>
  );
}

function MoveUpButton(props: IconButtonProps) {
  const { ...btnProps } = props;
  return (
    <Button {...btnProps} icon>
      <ChevronUpIcon />
    </Button>
  );
}

function MoveDownButton(props: IconButtonProps) {
  const { ...btnProps } = props;
  return (
    <Button {...btnProps} icon>
      <ChevronDownIcon />
    </Button>
  );
}

function BaseInputTemplate(props: BaseInputTemplateProps) {
  const type = useMemo(() => {
    if (props.schema.type === "integer") {
      return "number";
    } else if (props.type === "file") {
      return "file";
    } else {
      return undefined;
    }
  }, [props.type, props.schema.type]);

  return (
    <Input
      className="mb-2 mt-1 w-full"
      type={type}
      required={props.required}
      id={props.id}
      onChange={(e) => {
        if (
          props.type === "file" &&
          e.target.files &&
          e.target.files.length > 0
        ) {
          const reader = new FileReader();
          reader.onloadend = () => {
            props.onChange(reader.result);
          };
          reader.readAsDataURL(e.target.files[0] as Blob);
        } else {
          props.onChange(e.target.value);
        }
      }}
    />
  );
}

function TitleFieldTemplate(props: TitleFieldProps) {
  const { id, required, title } = props;
  return (
    <header id={id} className="mb-2 text-2xl">
      {title}
      {required && <mark>*</mark>}
    </header>
  );
}

const CustomCheckbox = function (props: WidgetProps) {
  return (
    <div className="flex space-x-2 p-2 ">
      <Checkbox
        onClick={() => props.onChange(!props.value)}
        id={`wgt-checkbox-${props.id}`}
      />
      <div className="grid gap-1.5 leading-none">
        <label
          htmlFor={`wgt-checkbox-${props.id}`}
          className="text-sm font-medium leading-none text-gray-10 peer-disabled:cursor-not-allowed peer-disabled:opacity-70 dark:text-gray-dark-10"
        >
          {props.label}
        </label>
      </div>
    </div>
  );
};

const widgets: RegistryWidgetsType = {
  CheckboxWidget: CustomCheckbox,
  SelectWidget: CustomSelectWidget,
};
export interface JSONSchemaFormProps {
  schema: RJSFSchema;
}
export const JSONSchemaForm: React.FC<JSONSchemaFormProps> = (props) => (
  <Form
    schema={props.schema}
    templates={{
      BaseInputTemplate,
      TitleFieldTemplate,
      ButtonTemplates: {
        SubmitButton,
        AddButton,
        RemoveButton,
        MoveUpButton,
        MoveDownButton,
      },
    }}
    validator={validator}
    widgets={widgets}
  />
);

JSONSchemaForm.displayName = "JSONSchemaForm";
