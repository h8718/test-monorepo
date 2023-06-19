import {
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "~/design/Dialog";
import Editor, { EditorLanguagesType } from "~/design/Editor";
import MimeTypeSelect, {
  MimeTypeType,
  mimeTypeToLanguageDict,
} from "./MimeTypeSelect";
import { SubmitHandler, useForm } from "react-hook-form";
import { VarFormSchema, VarFormSchemaType } from "~/api/variables/schema";
import { useEffect, useState } from "react";

import Button from "~/design/Button";
import { Card } from "~/design/Card";
import FormErrors from "~/componentsNext/FormErrors";
import Input from "~/design/Input";
import { PlusCircle } from "lucide-react";
import { useTheme } from "~/util/store/theme";
import { useTranslation } from "react-i18next";
import { useUpdateVar } from "~/api/variables/mutate/updateVariable";
import { zodResolver } from "@hookform/resolvers/zod";

type CreateProps = { onSuccess: () => void };

const defaultMimeType: MimeTypeType = "application/json";

const Create = ({ onSuccess }: CreateProps) => {
  const { t } = useTranslation();
  const theme = useTheme();

  const [name, setName] = useState<string | undefined>();
  const [body, setBody] = useState<string | undefined>();
  const [mimeType, setMimeType] = useState<MimeTypeType>(defaultMimeType);
  const [editorLanguage, setEditorLanguage] =
    useState<EditorLanguagesType>("plaintext");

  const {
    handleSubmit,
    formState: { errors },
  } = useForm<VarFormSchemaType>({
    resolver: zodResolver(VarFormSchema),
    // mimeType should always be initialized to avoid backend defaulting to
    // "text/plain, charset=utf-8", which does not fit the options in
    // MimeTypeSelect
    values: {
      name: name ?? "",
      content: body ?? "",
      mimeType: mimeType ?? "application/json",
    },
  });

  useEffect(() => {
    setEditorLanguage(mimeTypeToLanguageDict[mimeType]);
  }, [mimeType]);

  const { mutate: createVarMutation } = useUpdateVar({
    onSuccess,
  });

  const onSubmit: SubmitHandler<VarFormSchemaType> = (data) => {
    createVarMutation(data);
  };

  return (
    <DialogContent>
      <form
        id="create-variable"
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col space-y-5"
      >
        <DialogHeader>
          <DialogHeader>
            <DialogTitle>
              <PlusCircle />
              {t("pages.settings.variables.create.title")}
            </DialogTitle>
          </DialogHeader>
        </DialogHeader>

        <FormErrors errors={errors} className="mb-5" />

        <fieldset className="flex items-center gap-5">
          <label className="w-[150px] text-right" htmlFor="name">
            {t("pages.settings.variables.create.name")}
          </label>
          <Input
            data-testid="new-variable-name"
            placeholder="variable-name"
            onChange={(event) => setName(event.target.value)}
          />
        </fieldset>

        <Card className="grow p-4 pl-0" background="weight-1">
          <div className="h-[500px]">
            <Editor
              value={body}
              onChange={(newData) => {
                setBody(newData);
              }}
              theme={theme ?? undefined}
              data-testid="variable-editor"
              language={editorLanguage}
            />
          </div>
        </Card>

        <fieldset className="flex items-center gap-5">
          <label className="w-[150px] text-right" htmlFor="template">
            {t("pages.settings.variables.edit.mimeType")}
          </label>
          <MimeTypeSelect mimeType={mimeType} onChange={setMimeType} />
        </fieldset>

        <DialogFooter>
          <DialogClose asChild>
            <Button variant="ghost">
              {t("components.button.label.cancel")}
            </Button>
          </DialogClose>
          <Button
            data-testid="variable-create-submit"
            type="submit"
            variant="primary"
          >
            {t("components.button.label.create")}
          </Button>
        </DialogFooter>
      </form>
    </DialogContent>
  );
};

export default Create;
