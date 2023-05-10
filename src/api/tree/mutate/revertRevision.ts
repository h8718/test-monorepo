import { WorkflowCreatedSchema } from "../schema";
import { apiFactory } from "../../utils";
import { forceLeadingSlash } from "../utils";
import { useApiKey } from "../../../util/store/apiKey";
import { useMutation } from "@tanstack/react-query";
import { useNamespace } from "../../../util/store/namespace";
import { useToast } from "../../../design/Toast";

const revertRevision = apiFactory({
  pathFn: ({ namespace, path }: { namespace: string; path: string }) =>
    `/api/namespaces/${namespace}/tree${forceLeadingSlash(
      path
    )}?op=discard-workflow&ref=latest`,
  method: "POST",
  schema: WorkflowCreatedSchema,
});

export const useRevertRevision = () => {
  const apiKey = useApiKey();
  const namespace = useNamespace();
  const { toast } = useToast();

  if (!namespace) {
    throw new Error("namespace is undefined");
  }

  return useMutation({
    mutationFn: ({ path }: { path: string }) =>
      revertRevision({
        apiKey: apiKey ?? undefined,
        params: undefined,
        pathParams: {
          namespace: namespace,
          path,
        },
      }),
    onSuccess: () => {
      toast({
        title: "Restored workflow",
        description: `The latest revision was restored`,
        variant: "success",
      });
    },
    onError: () => {
      toast({
        title: "An error occurred",
        description: "could not revert workflow 😢",
        variant: "error",
      });
    },
  });
};
