import { GitMerge, MoreVertical, Tag } from "lucide-react";
import {
  Table,
  TableBody,
  TableCell,
  TableRow,
} from "../../../../../design/Table";

import Button from "../../../../../design/Button";
import { Card } from "../../../../../design/Card";
import CopyButton from "../../../../../design/CopyButton";
import { FC } from "react";
import { Link } from "react-router-dom";
import { pages } from "../../../../../util/router/pages";
// import { useNodeContent } from "../../../../../api/tree/query/get";
import { useNodeRevisions } from "../../../../../api/tree/query/revisions";
import { useNodeTags } from "../../../../../api/tree/query/tags";

const WorkflowRevisionsPage: FC = () => {
  const { path, namespace } = pages.explorer.useParams();

  // const { data } = useNodeContent({
  //   path,
  //   revision,
  // });

  const { data: revisions } = useNodeRevisions({ path });
  const { data: tags } = useNodeTags({ path });

  if (!namespace) return null;

  return (
    <div className="p-5">
      <Card>
        <Table>
          <TableBody>
            {revisions?.results?.map((rev, i) => {
              const isTag = tags?.results?.some((tag) => tag.name === rev.name);
              const Icon = isTag ? Tag : GitMerge;
              return (
                <TableRow key={i} className="group">
                  <TableCell>
                    <div className="flex space-x-3">
                      <Icon aria-hidden="true" className="h-5" />

                      <Link
                        to={pages.explorer.createHref({
                          namespace,
                          path,
                          subpage: "workflow-revisions",
                          revision: rev.name,
                        })}
                      >
                        {rev.name}
                      </Link>
                    </div>
                  </TableCell>
                  <TableCell className="group w-0 text-right">
                    <CopyButton
                      value={rev.name}
                      buttonProps={{
                        variant: "outline",
                        className: "w-24 hidden group-hover:inline-flex",
                        size: "sm",
                      }}
                    >
                      {(copied) => (copied ? "copied" : "copy")}
                    </CopyButton>
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={(e) => e.preventDefault()}
                      icon
                    >
                      <MoreVertical />
                    </Button>
                  </TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </Card>
    </div>
  );
};

export default WorkflowRevisionsPage;
