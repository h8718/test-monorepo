import { FC, PropsWithChildren } from "react";

import { DialogTrigger } from "~/design/Dialog";
import { Link } from "react-router-dom";
import { NodeSchemaType } from "~/api/tree/schema";
import { pages } from "~/util/router/pages";

type ConditionalLinkProps = PropsWithChildren & {
  node: NodeSchemaType;
  namespace: string;
  onPreviewClicked: (file: NodeSchemaType) => void;
};

export const ConditionalLink: FC<ConditionalLinkProps> = ({
  node,
  namespace,
  onPreviewClicked,
  children,
}) => {
  const isFile = node.expandedType === "file";

  if (isFile)
    return (
      <DialogTrigger
        className="flex-1 hover:underline"
        role="button"
        onClick={() => {
          onPreviewClicked(node);
        }}
        asChild
      >
        <a>{children}</a>
      </DialogTrigger>
    );

  const linkTarget = pages.explorer.createHref({
    namespace,
    path: node.path,
    subpage: node.expandedType === "workflow" ? "workflow" : undefined,
  });

  return (
    <Link
      data-testid={`explorer-item-link-${node.name}`}
      to={linkTarget}
      className="flex-1 hover:underline"
    >
      {children}
    </Link>
  );
};
