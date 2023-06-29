import { FC, useEffect, useState } from "react";
import { Table, TableBody } from "~/design/Table";

import { Card } from "~/design/Card";
import { Container } from "lucide-react";
import Create from "./Create";
import CreateItemButton from "../components/CreateItemButton";
import Delete from "./Delete";
import { Dialog } from "~/design/Dialog";
import EmptyList from "../components/EmptyList";
import ItemRow from "../components/ItemRow";
import { RegistrySchemaType } from "~/api/registries/schema";
import { useDeleteRegistry } from "~/api/registries/mutate/deleteRegistry";
import { useRegistries } from "~/api/registries/query/get";
import { useTranslation } from "react-i18next";

const RegistriesList: FC = () => {
  const { t } = useTranslation();

  const [dialogOpen, setDialogOpen] = useState(false);
  const [deleteRegistry, setDeleteRegistry] = useState<RegistrySchemaType>();
  const [createRegistry, setCreateRegistry] = useState(false);

  const { data, isFetched } = useRegistries();

  const { mutate: deleteRegistryMutation } = useDeleteRegistry({
    onSuccess: () => {
      setDeleteRegistry(undefined);
      setDialogOpen(false);
    },
  });

  useEffect(() => {
    if (dialogOpen === false) {
      setDeleteRegistry(undefined);
      setCreateRegistry(false);
    }
  }, [dialogOpen]);

  if (!isFetched) return null;

  return (
    <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
      <div className="mb-3 flex flex-row justify-between">
        <h3 className="flex items-center gap-x-2 pb-2 pt-1 font-bold text-gray-10 dark:text-gray-dark-10">
          <Container className="h-5" />
          {t("pages.settings.registries.list.title")}
        </h3>

        <CreateItemButton
          onClick={() => setCreateRegistry(true)}
          data-testid="registry-create"
        />
      </div>

      <Card>
        {data?.registries.length ? (
          <Table>
            <TableBody>
              {data?.registries.map((item, i) => (
                <ItemRow key={i} item={item} onDelete={setDeleteRegistry} />
              ))}
            </TableBody>
          </Table>
        ) : (
          <EmptyList>{t("pages.settings.registries.list.empty")}</EmptyList>
        )}
        {deleteRegistry && (
          <Delete
            name={deleteRegistry.name}
            onConfirm={() =>
              deleteRegistryMutation({ registry: deleteRegistry })
            }
          />
        )}
        {createRegistry && <Create onSuccess={() => setDialogOpen(false)} />}
      </Card>
    </Dialog>
  );
};

export default RegistriesList;
