import { Dialog, DialogContent } from "~/design/Dialog";
import { FC, useEffect, useState } from "react";
import { Table, TableBody, TableCell, TableRow } from "~/design/Table";

import { Card } from "~/design/Card";
import Delete from "./Delete";
import ExplorerHeader from "./Header";
import FileRow from "./FileRow";
import FileViewer from "./FileViewer";
import { FolderUp } from "lucide-react";
import { Link } from "react-router-dom";
import NoResult from "./NoResult";
import { NodeSchemaType } from "~/api/tree/schema";
import Rename from "./Rename";
import { analyzePath } from "~/util/router/utils";
import { pages } from "~/util/router/pages";
import { useNamespace } from "~/util/store/namespace";
import { useNodeContent } from "~/api/tree/query/node";
import { useTranslation } from "react-i18next";

const ExplorerPage: FC = () => {
  const namespace = useNamespace();
  const { path } = pages.explorer.useParams();
  const { data, isSuccess } = useNodeContent({ path });
  const { parent, isRoot } = analyzePath(path);
  const [dialogOpen, setDialogOpen] = useState(false);

  // we only want to use one dialog component for the whole list,
  // so when the user clicks on the delete button in the list, we
  // set the pointer to that node for the dialog
  const [deleteNode, setDeleteNode] = useState<NodeSchemaType>();
  const [renameNode, setRenameNode] = useState<NodeSchemaType>();
  const [previewNode, setPreviewNode] = useState<NodeSchemaType>();
  const { t } = useTranslation();

  useEffect(() => {
    if (dialogOpen === false) {
      setDeleteNode(undefined);
      setRenameNode(undefined);
      setPreviewNode(undefined);
    }
  }, [dialogOpen]);

  if (!namespace) return null;

  const results = data?.children?.results ?? [];
  const showTable = !isRoot || results.length > 0;
  const noResults = isSuccess && results.length === 0;

  return (
    <>
      <ExplorerHeader />
      <div className="p-5">
        <Card>
          {showTable && (
            <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
              <Table>
                <TableBody>
                  {!isRoot && (
                    <TableRow>
                      <TableCell colSpan={2}>
                        <Link
                          to={pages.explorer.createHref({
                            namespace,
                            path: parent?.absolute,
                          })}
                          className="flex items-center space-x-3 hover:underline"
                        >
                          <FolderUp className="h-5" />
                          <span>
                            {t("pages.explorer.tree.list.oneLevelUp")}
                          </span>
                        </Link>
                      </TableCell>
                    </TableRow>
                  )}
                  {results.map((file) => (
                    <FileRow
                      key={file.name}
                      namespace={namespace}
                      node={file}
                      onDeleteClicked={setDeleteNode}
                      onRenameClicked={setRenameNode}
                      onPreviewClicked={setPreviewNode}
                    />
                  ))}
                </TableBody>
              </Table>
              <DialogContent className="sm:max-w-xl md:max-w-2xl lg:max-w-3xl">
                {previewNode && <FileViewer node={previewNode} />}
                {deleteNode && (
                  <Delete
                    node={deleteNode}
                    close={() => {
                      setDialogOpen(false);
                    }}
                  />
                )}
                {renameNode && (
                  <Rename
                    node={renameNode}
                    close={() => {
                      setDialogOpen(false);
                    }}
                    unallowedNames={
                      data?.children?.results.map((x) => x.name) || []
                    }
                  />
                )}
              </DialogContent>
            </Dialog>
          )}
          {noResults && <NoResult />}
        </Card>
      </div>
    </>
  );
};

export default ExplorerPage;
