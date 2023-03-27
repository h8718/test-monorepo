import { TreeFolderCreatedSchema } from "../schema";
import { apiFactory } from "../../utils";
import { forceSlashIfPath } from "../utils";
import { useApiKey } from "../../../util/store/apiKey";
import { useMutation } from "@tanstack/react-query";
import { useNamespace } from "../../../util/store/namespace";
import { useToast } from "../../../componentsNext/Toast";

const createDirectory = apiFactory({
  pathFn: ({
    namespace,
    path,
    directory,
  }: {
    namespace: string;
    path?: string;
    directory: string;
  }) =>
    `/api/namespaces/${namespace}/tree${forceSlashIfPath(
      path
    )}/${directory}?op=create-directory`,
  method: "PUT",
  schema: TreeFolderCreatedSchema,
});

type ResolvedCreateDirectory = Awaited<ReturnType<typeof createDirectory>>;

export const useCreateDirectory = ({
  onSuccess,
}: { onSuccess?: (data: ResolvedCreateDirectory) => void } = {}) => {
  const apiKey = useApiKey();
  const namespace = useNamespace();
  const { toast } = useToast();

  if (!namespace) {
    throw new Error("namespace is undefined");
  }

  return useMutation({
    mutationFn: ({ path, directory }: { path?: string; directory: string }) =>
      createDirectory({
        apiKey: apiKey ?? undefined,
        params: undefined,
        pathParams: {
          directory,
          namespace: namespace,
          path,
        },
      }),

    onSuccess(data, variables) {
      toast({
        title: "Directory created",
        description: `Directory ${variables.directory} was created in ${variables.path}`,
        variant: "success",
      });
      onSuccess?.(data);
    },
    onError: () => {
      toast({
        title: "An error occurred",
        description: "could not create directory 😢",
        variant: "error",
      });
    },
  });
};
