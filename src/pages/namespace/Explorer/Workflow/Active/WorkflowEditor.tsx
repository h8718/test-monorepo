import { Bug, GitBranchPlus, GitMerge, Play, Save, Undo } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../../../../../design/Dropdown";
import { FC, useEffect, useState } from "react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "../../../../../design/Popover";

import Button from "../../../../../design/Button";
import { Card } from "../../../../../design/Card";
import Editor from "../../../../../design/Editor";
import { RxChevronDown } from "react-icons/rx";
import moment from "moment";
import { useCreateRevision } from "../../../../../api/tree/mutate/createRevision";
import { useNodeContent } from "../../../../../api/tree/query/get";
import { useTheme } from "../../../../../util/store/theme";
import { useUpdateWorkflow } from "../../../../../api/tree/mutate/updateWorkflow";

// get type of useNodeContent return value
type NodeContentType = ReturnType<typeof useNodeContent>["data"];

const WorkflowEditor: FC<{
  data: NonNullable<NodeContentType>;
  path: string;
}> = ({ data, path }) => {
  const [error, setError] = useState<string | undefined>();
  const [hasUnsavedChanged, setHasUnsavedChanged] = useState(false);
  const workflowData = data.revision?.source && atob(data?.revision?.source);

  const { mutate: updateWorkflow, isLoading } = useUpdateWorkflow({
    onError: (error) => {
      error && setError(error);
    },
  });

  const [value, setValue] = useState<string | undefined>(workflowData);
  const theme = useTheme();

  const { mutate: createRevision } = useCreateRevision();

  useEffect(() => {
    setHasUnsavedChanged(workflowData !== value);
  }, [value, workflowData]);

  const onSave = (toSave: string | undefined) => {
    if (toSave) {
      setError(undefined);
      updateWorkflow({
        path,
        fileContent: toSave,
      });
    }
  };

  return (
    <div className="relative flex grow flex-col space-y-4 p-4">
      <Card className="grow p-4">
        <Editor
          value={workflowData}
          onChange={(newData) => {
            setValue(newData);
          }}
          theme={theme ?? undefined}
          onSave={onSave}
        />
      </Card>
      <div className="flex flex-col justify-end gap-4 sm:flex-row sm:items-center">
        <div className="flex grow items-center justify-between gap-2 text-sm text-gray-8 dark:text-gray-dark-8">
          {/* must use fromNow(true) because otherwise after saving, it sometimes shows Updated in a few seconds */}
          {data.revision?.createdAt && (
            <span className="text-center">
              Updated {moment(data.revision?.createdAt).fromNow(true)} ago
            </span>
          )}
          {hasUnsavedChanged && (
            <span className="text-center">you have unsaved changes</span>
          )}
        </div>
        {error && (
          <Popover defaultOpen>
            <PopoverTrigger asChild>
              <Button variant="destructive">
                <Bug />
                There is one issue
              </Button>
            </PopoverTrigger>
            <PopoverContent asChild>
              <div className="flex">
                <div className="grow">{error}</div>
              </div>
            </PopoverContent>
          </Popover>
        )}
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" disabled={hasUnsavedChanged}>
              <GitMerge />
              Revisions <RxChevronDown />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-60">
            <DropdownMenuItem
              onClick={() => {
                createRevision({
                  path,
                });
              }}
            >
              <GitBranchPlus className="mr-2 h-4 w-4" /> Make Revision
            </DropdownMenuItem>
            <DropdownMenuItem>
              <Undo className="mr-2 h-4 w-4" /> Revert to previous revision
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
        <Button variant="outline">
          <Play />
          Run
        </Button>
        <Button
          variant="outline"
          disabled={isLoading}
          onClick={() => {
            onSave(value);
          }}
        >
          <Save />
          Save
        </Button>
      </div>
    </div>
  );
};

export default WorkflowEditor;
