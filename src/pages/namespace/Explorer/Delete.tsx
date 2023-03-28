import {
  DialogClose,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "../../../design/Dialog";

import Button from "../../../design/Button";
import { NodeSchemaType } from "../../../api/tree/schema";
import { Trash } from "lucide-react";
import { useDeleteNode } from "../../../api/tree/mutate/deleteNode";

const Delete = ({
  node,
  close,
}: {
  node: NodeSchemaType;
  close: () => void;
}) => {
  const { mutate, isLoading } = useDeleteNode({
    onSuccess: () => {
      close();
    },
  });

  return (
    <>
      <DialogHeader>
        <DialogTitle>
          <Trash /> Delete
        </DialogTitle>
      </DialogHeader>
      <div className="my-6All content of this directoy will be deleted as well.">
        Are you sure you want to delete <b>{node.name}</b>? This can not be
        undone.&nbsp;
        {node.type === "directory" &&
          "All content of this directoy will be deleted as well."}
      </div>
      <DialogFooter>
        <DialogClose asChild>
          <Button variant="ghost">Cancel</Button>
        </DialogClose>
        <Button
          onClick={() => {
            mutate({ node });
          }}
          variant="destructive"
          loading={isLoading}
        >
          {!isLoading && <Trash />}
          Delete
        </Button>
      </DialogFooter>
    </>
  );
};

export default Delete;
