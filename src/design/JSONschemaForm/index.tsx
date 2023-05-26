import {
  RJSFSchema,
  WidgetProps,
} from "@rjsf/utils";

import Button from "../Button";
import Form from "@rjsf/core";
import Input from "../Input";
import React from "react";

const FormInput: React.FunctionComponent<WidgetProps> = (props) => {
  const [val, setVal] = React.useState<string>("");
  return (
    <Input
      className="form-control mb-2 mt-1 w-full"
      type={props.options.inputType}
      required={props.required}
      id={props.id}
      value={val}
      onChange={(e) => {
        props.onChange(e.target.value);
        setVal(e.target.value);
      }}
    />
  );
};

export const JSONschemaForm = () => {
  const uiSchema = {
    password: {
      "ui:widget": FormInput,
      "ui:options": {
        inputType: "password",
      },
    },
    firstName: {
      "ui:widget": FormInput,
      "ui:options": {
        inputType: "text",
      },
    },
    lastName: {
      "ui:widget": FormInput,
      "ui:options": {
        inputType: "text",
      },
    },
    bio: {
      "ui:widget": FormInput,
      "ui:options": {
        inputType: "text",
      },
    },
    age: {
      "ui:widget": FormInput,
      "ui:options": {
        inputType: "number",
      },
    },
    "ui:submitButtonOptions": {
      norender: true,
    },
  };
  const schema: RJSFSchema = {
    title: "A registration form",
    type: "object",
    required: ["firstName", "lastName"],
    properties: {
      password: {
        type: "string",
        title: "Password",
      },
      lastName: {
        type: "string",
        title: "Last name",
      },
      bio: {
        type: "string",
        title: "Bio",
      },
      firstName: {
        type: "string",
        title: "First name",
      },
      age: {
        type: "integer",
        title: "Age",
      },
    },
  };

  return (
    // <div className='[&>form>div>fieldset>legend]:text-primary-500'>
    <Form uiSchema={uiSchema} schema={schema}>
      <Button type="submit">Submit</Button>
    </Form>
    // </div>
  );
};
