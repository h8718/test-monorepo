import React, { FC, HTMLAttributes } from "react";

import clsx from "clsx";

export type AvatarProps = HTMLAttributes<HTMLDivElement> & {
  className?: string;
  children?: React.ReactNode;
};

export const Avatar = React.forwardRef<HTMLDivElement, React.HTMLAttributes<HTMLDivElement>>(({ children, className, ...props }, ref) => (
  <div
    {...props}
    className={clsx(
      "flex h-7 w-7 items-center justify-center rounded-full text-xs",
      "bg-primary-500 text-gray-1 dark:text-gray-dark-1",
      className
    )}
    ref={ref}
  >
    {children ? children : ""}
  </div>
));
Avatar.displayName = "Avatar";

export default Avatar;
