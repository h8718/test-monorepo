import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "~/design/Dialog";
import { FC, useState } from "react";
import { PlusCircle, SquareAsterisk, Trash } from "lucide-react";
import { SubmitHandler, useForm } from "react-hook-form";
import { Table, TableBody, TableCell, TableRow } from "~/design/Table";
import { Trans, useTranslation } from "react-i18next";

import Button from "~/design/Button";
import { Card } from "~/design/Card";
import Input from "~/design/Input";
import { RegistrySchemaType } from "~/api/registries/schema";
import { Textarea } from "~/design/TextArea";
// import { useCreateRegistry } from "~/api/registries/mutate/createRegistry";
import { useDeleteRegistry } from "~/api/registries/mutate/deleteRegistry";
import { useRegistries } from "~/api/registries/query/get";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

type RegistryFormInput = {
  name: string;
  value: string;
};

const RegistriesList: FC = () => {
  const { t } = useTranslation();

  const [dialogOpen, setDialogOpen] = useState(false);
  const [deleteRegistry, setDeleteRegistry] = useState<RegistrySchemaType>();
  const [createRegistry, setCreateRegistry] = useState(false);

  const registries = useRegistries();

  const { mutate: deleteRegistryMutation } = useDeleteRegistry({
    onSuccess: () => {
      setDeleteRegistry(undefined);
      setDialogOpen(false);
    },
  });

  // const { mutate: createRegistryMutation } = useCreateRegistry({
  //   onSuccess: () => {
  //     setCreateRegistry(false);
  //     setDialogOpen(false);
  //   },
  // });

  // const onSubmit: SubmitHandler<RegistryFormInput> = ({ name, value }) => {
  //   createRegistryMutation({
  //     name,
  //     value,
  //   });
  // };

  const { register, handleSubmit, reset } = useForm<RegistryFormInput>({
    resolver: zodResolver(
      z.object({
        name: z.string(),
        value: z.string(),
      })
    ),
  });

  const resetDialog = (isOpening: boolean) => {
    if (!isOpening) {
      setDeleteRegistry(undefined);
      setCreateRegistry(false);
      reset();
    }
    setDialogOpen(isOpening);
  };

  return (
    <Dialog
      open={dialogOpen}
      onOpenChange={(isOpening) => resetDialog(isOpening)}
    >
      <div className="mb-3 flex flex-row justify-between">
        <h3 className="flex items-center gap-x-2 font-bold text-gray-10 dark:text-gray-dark-10">
          <SquareAsterisk className="h-5" />
          {t("pages.settings.registries.list.title")}
        </h3>

        <DialogTrigger
          asChild
          data-testid="registry-create"
          onClick={() => setCreateRegistry(true)}
        >
          <Button variant="ghost">
            <PlusCircle />
          </Button>
        </DialogTrigger>
      </div>

      <Card>
        <Table>
          <TableBody>
            {registries.data?.registries.map((item, i) => (
              <TableRow key={i}>
                <TableCell>{item.name}</TableCell>
                <TableCell className="w-0">
                  <DialogTrigger
                    asChild
                    data-testid="registry-delete"
                    onClick={() => {
                      setDeleteRegistry(item);
                    }}
                  >
                    <Button variant="ghost">
                      <Trash />
                    </Button>
                  </DialogTrigger>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        {deleteRegistry && (
          <DialogContent>
            <DialogHeader>
              <DialogTitle>
                <Trash /> Delete
              </DialogTitle>
            </DialogHeader>
            <div className="my-3">
              <Trans
                i18nKey="pages.settings.registries.delete.description"
                values={{ name: deleteRegistry.name }}
              />
            </div>
            <DialogFooter>
              <DialogClose asChild>
                <Button variant="ghost">Cancel</Button>
              </DialogClose>
              <Button
                data-testid="registry-delete-confirm"
                onClick={() =>
                  deleteRegistryMutation({ registry: deleteRegistry })
                }
                variant="destructive"
              >
                Delete
              </Button>
            </DialogFooter>
          </DialogContent>
        )}
        {createRegistry && (
          <DialogContent>
            <form
              id="create-registry"
              // onSubmit={handleSubmit(onSubmit)}
              className="flex flex-col space-y-5"
            >
              <DialogHeader>
                <DialogTitle>
                  <Trash /> Create
                </DialogTitle>
              </DialogHeader>

              <fieldset className="flex items-center gap-5">
                <label
                  className="w-[150px] text-right text-[15px]"
                  htmlFor="name"
                >
                  Name
                </label>
                <Input
                  data-testid="new-registry-name"
                  id="name"
                  placeholder="registry-name"
                  // {...register("name")}
                />
              </fieldset>

              <fieldset className="flex items-start gap-5">
                <Textarea
                  className="h-96"
                  data-testid="new-workflow-editor"
                  // {...register("value")}
                />
              </fieldset>

              <DialogFooter>
                <DialogClose asChild>
                  <Button variant="ghost">Cancel</Button>
                </DialogClose>
                <Button
                  data-testid="registry-create-submit"
                  type="submit"
                  variant="primary"
                >
                  Create
                </Button>
              </DialogFooter>
            </form>
          </DialogContent>
        )}
      </Card>
    </Dialog>
  );
};

export default RegistriesList;
