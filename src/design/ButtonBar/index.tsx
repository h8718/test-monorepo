import React, { HTMLAttributes } from "react";

import { twMergeClsx } from "~/util/helpers";

export const ButtonBar: React.FC<HTMLAttributes<HTMLDivElement>> = ({
  children,
  className,
  ...props
}) => (
  <div
    {...props}
    className={twMergeClsx(
      "[&_button]:rounded-none",
      "[&_button]:mr-[-1px]",
      "[&_button]:active:outline-none",
      "[&_button]:border [&_button]:focus:ring-0 [&_button]:focus:ring-offset-0",
      "[&_button]:active:ring-0 [&_button]:active:ring-offset-0",

      "[&>*]:rounded-none",
      "[&>*]:mr-[-1px]",
      "[&>*]:active:outline-none",
      "[&>*]:border [&_button]:focus:ring-0 [&_button]:focus:ring-offset-0",
      "[&>*]:active:ring-0 [&_button]:active:ring-offset-0",
      //in case of label

      //
      "[&>*:first-child]:rounded-l-md",
      "[&>*:last-child]:rounded-r-md",
      // if button is not the direct child of the button bar (required e.g. for the tooltip)
      "[&>:first-child_button:first-of-type]:rounded-l-md",
      "[&>:last-child_button:first-of-type]:rounded-r-md",
      "flex items-end",
      className
    )}
  >
    {children}
  </div>
);
ButtonBar.displayName = "ButtonBar";
