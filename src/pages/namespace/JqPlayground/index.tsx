import { FC, useState } from "react";
import { FileInput, FileOutput, Play, PlaySquare } from "lucide-react";
import {
  useJqPlaygroundActions,
  useJqPlaygroundInput,
  useJqPlaygroundQuery,
} from "~/util/store/jqPlaygrpund";

import Button from "~/design/Button";
import { Card } from "~/design/Card";
import CopyButton from "~/design/CopyButton";
import Editor from "~/design/Editor";
import FormErrors from "~/componentsNext/FormErrors";
import Input from "~/design/Input";
import { JqQueryErrorSchema } from "~/api/jq/schema";
import { useExecuteJQuery } from "~/api/jq/mutate/executeQuery";
import { useTheme } from "~/util/store/theme";
import { useTranslation } from "react-i18next";

const data = {
  foo: [
    { name: "JSON", good: true },
    { name: "XML", good: false },
  ],
};

const JqPlaygroundPage: FC = () => {
  const { t } = useTranslation();
  const theme = useTheme();
  const defaultInput = JSON.parse(useJqPlaygroundInput() ?? "{}");
  const defaultQuery = useJqPlaygroundQuery() ?? ".";
  const {
    setInput: storeInputInLocalstorage,
    setQuery: storeQueryInLocalstorage,
  } = useJqPlaygroundActions();

  const [query, setQuery] = useState(defaultQuery ?? ".");
  const [input, setInput] = useState(
    JSON.stringify(defaultInput ?? "{}", null, 2)
  );
  const [result, setResult] = useState("");
  const [error, setError] = useState("");

  const { mutate: executeQuery, isLoading } = useExecuteJQuery({
    onSuccess: (data) => {
      const resultAsJson = JSON.parse(data.results?.[0] ?? "{}");
      setResult(JSON.stringify(resultAsJson, null, 2));
    },
    onError: (error) => {
      const errorParsed = JqQueryErrorSchema.safeParse(error);
      if (errorParsed.success) {
        setError(errorParsed.data.message);
      }
    },
  });

  const formId = "jq-playground-form";

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    executeQuery({ query, inputJsonString: input });
  };

  const onQueryChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setQuery(e.target.value);
    storeQueryInLocalstorage(e.target.value);
    setError("");
  };

  const onInputChange = (newData: string | undefined) => {
    if (newData) {
      setInput(newData);
      storeInputInLocalstorage(newData);
      setError("");
    }
  };

  return (
    <div className="flex grow flex-col gap-y-4 p-5">
      <h3 className="flex items-center gap-x-2 font-bold">
        <PlaySquare className="h-5" />
        {t("pages.jqPlayground.title")}
      </h3>
      <Card className="p-5">
        <form
          id={formId}
          onSubmit={handleSubmit}
          className="flex flex-col gap-5"
        >
          <div className="flex flex-col gap-5 sm:flex-row">
            <Input
              placeholder={t("pages.jqPlayground.queryPlaceholder")}
              value={query}
              onChange={onQueryChange}
            />
            <Button
              className="grow sm:w-64"
              type="submit"
              disabled={isLoading}
              loading={isLoading}
            >
              {!isLoading && <Play />}
              {t("pages.jqPlayground.submitBtn")}
            </Button>
          </div>
          {error && <FormErrors errors={{ error: { message: error } }} />}
          <div className="flex gap-5">
            <Card className="h-96 w-full p-4" noShadow background="weight-1">
              <div className="mb-5 flex">
                <h3 className="flex grow items-center gap-x-2 font-medium">
                  <FileInput className="h-5" />
                  {t("pages.jqPlayground.output")}
                </h3>
                <CopyButton
                  value={input}
                  buttonProps={{
                    variant: "outline",
                    size: "sm",
                    type: "button",
                    disabled: !input,
                  }}
                />
              </div>
              <Editor
                value={input}
                language="json"
                onChange={onInputChange}
                theme={theme ?? undefined}
              />
            </Card>
            <Card className="h-96 w-full p-4" noShadow background="weight-1">
              <div className="mb-5 flex">
                <h3 className="flex grow items-center gap-x-2 font-medium">
                  <FileOutput className="h-5" />
                  {t("pages.jqPlayground.output")}
                </h3>
                <CopyButton
                  value={result}
                  buttonProps={{
                    variant: "outline",
                    size: "sm",
                    type: "button",
                    disabled: !result,
                  }}
                />
              </div>
              <Editor
                language="json"
                value={result}
                options={{
                  readOnly: true,
                }}
                theme={theme ?? undefined}
              />
            </Card>
          </div>
        </form>
      </Card>
    </div>
  );
};

export default JqPlaygroundPage;
