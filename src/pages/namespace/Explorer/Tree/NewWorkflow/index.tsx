import {
  DialogClose,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "~/design/Dialog";
import { Play, PlusCircle } from "lucide-react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "~/design/Select";
import { SubmitHandler, useForm } from "react-hook-form";

import Alert from "~/design/Alert";
import Button from "~/design/Button";
import { Card } from "~/design/Card";
import Editor from "~/design/Editor";
import Input from "~/design/Input";
import { Textarea } from "~/design/TextArea";
import { addYamlFileExtension } from "./utils";
import { fileNameSchema } from "~/api/tree/schema";
import { pages } from "~/util/router/pages";
import { useCreateWorkflow } from "~/api/tree/mutate/createWorkflow";
import { useNamespace } from "~/util/store/namespace";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import { useTheme } from "~/util/store/theme";
import { useTranslation } from "react-i18next";
import workflowTemplates from "./templates";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

type FormInput = {
  name: string;
  fileContent: string;
};

const defaultWorkflowTemplate = workflowTemplates[0];

const NewWorkflow = ({
  path,
  close,
  unallowedNames,
}: {
  path?: string;
  close: () => void;
  unallowedNames?: string[];
}) => {
  const { t } = useTranslation();
  const namespace = useNamespace();
  const navigate = useNavigate();

  const theme = useTheme();
  const [workflowData, setWorkflowData] = useState<string>(
    defaultWorkflowTemplate.data
  );
  const {
    register,
    handleSubmit,
    setValue,
    formState: { isDirty, errors, isValid, isSubmitted },
  } = useForm<FormInput>({
    resolver: zodResolver(
      z.object({
        name: fileNameSchema.and(
          z
            .string()
            .refine((name) => !(unallowedNames ?? []).some((n) => n === name), {
              message: t("pages.explorer.tree.newWorkflow.nameAlreadyExists"),
            })
        ),
        fileContent: z.string(),
      })
    ),
    defaultValues: {
      fileContent: defaultWorkflowTemplate.data,
    },
  });

  const { mutate: createWorkflow, isLoading } = useCreateWorkflow({
    onSuccess: (data) => {
      namespace &&
        navigate(
          pages.explorer.createHref({
            namespace,
            path: data.node.path,
            subpage: "workflow",
          })
        );
      close();
    },
  });

  const onSubmit: SubmitHandler<FormInput> = ({ name, fileContent }) => {
    createWorkflow({ path, name: addYamlFileExtension(name), fileContent });
  };

  // you can not submit if the form has not changed or if there are any errors and
  // you have already submitted the form (errors will first show up after submit)
  const disableSubmit = !isDirty || (isSubmitted && !isValid);

  const formId = `new-worfklow-${path}`;
  return (
    <>
      <DialogHeader>
        <DialogTitle>
          <Play />
          {t("pages.explorer.tree.newWorkflow.title")}
        </DialogTitle>
      </DialogHeader>
      <div className="my-3">
        {!!errors.name && (
          <Alert variant="error" className="mb-5">
            <p>{errors.name.message}</p>
          </Alert>
        )}
        <form
          id={formId}
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col gap-y-5"
        >
          <fieldset className="flex items-center gap-5">
            <label className="w-[100px] text-right text-[14px]" htmlFor="name">
              {t("pages.explorer.tree.newWorkflow.nameLabel")}
            </label>
            <Input
              data-testid="new-workflow-name"
              id="name"
              placeholder={t("pages.explorer.tree.newWorkflow.namePlaceholder")}
              {...register("name")}
            />
          </fieldset>
          <fieldset className="flex items-center gap-5">
            <label
              className="w-[100px] text-right text-[14px]"
              htmlFor="template"
            >
              {t("pages.explorer.tree.newWorkflow.templateLabel")}
            </label>
            <Select
              onValueChange={(value) => {
                const matchingWf = workflowTemplates.find(
                  (t) => t.name === value
                );
                if (matchingWf) {
                  setValue("fileContent", matchingWf.data);
                  setWorkflowData(matchingWf.data);
                }
              }}
            >
              <SelectTrigger id="template" variant="outline" block>
                <SelectValue
                  placeholder={defaultWorkflowTemplate.name}
                  defaultValue={defaultWorkflowTemplate.data}
                />
              </SelectTrigger>
              <SelectContent>
                {workflowTemplates.map((t) => (
                  <SelectItem value={t.name} key={t.name}>
                    {t.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </fieldset>
          <fieldset className="flex items-start gap-5">
            <Textarea className="hidden" {...register("fileContent")} />
            <Card className="h-96 w-full p-4" noShadow weight={1}>
              <Editor
                value={workflowData}
                onChange={(newData) => {
                  if (newData) {
                    setWorkflowData(newData);
                    setValue("fileContent", newData);
                  }
                }}
                theme={theme ?? undefined}
              />
            </Card>
          </fieldset>
        </form>
      </div>
      <DialogFooter>
        <DialogClose asChild>
          <Button variant="ghost">
            {t("pages.explorer.tree.newWorkflow.cancelBtn")}
          </Button>
        </DialogClose>
        <Button
          data-testid="new-workflow-submit"
          type="submit"
          disabled={disableSubmit}
          loading={isLoading}
          form={formId}
        >
          {!isLoading && <PlusCircle />}
          {t("pages.explorer.tree.newWorkflow.createBtn")}
        </Button>
      </DialogFooter>
    </>
  );
};

export default NewWorkflow;
