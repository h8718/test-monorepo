import * as React from "react";

import { Loader2 } from "lucide-react";
import clsx from "clsx";

export type ButtonProps = {
  variant?: "destructive" | "outline" | "primary" | "ghost" | "link";
  size?: "sm" | "lg";
  loading?: boolean;
  circle?: boolean;
  block?: boolean;
  icon?: boolean;
};

const Button = React.forwardRef<
  HTMLButtonElement,
  React.ButtonHTMLAttributes<HTMLButtonElement> & ButtonProps
>(
  (
    {
      className,
      variant,
      size,
      circle,
      children,
      disabled,
      block,
      loading,
      icon,
      ...props
    },
    ref
  ) => (
    <button
      className={clsx(
        className,
        "inline-flex items-center justify-center text-sm font-medium transition-colors",
        "focus:outline-none focus:ring-2 focus:ring-gray-4 focus:ring-offset-2 focus:ring-offset-gray-1",
        "active:scale-95",
        "disabled:pointer-events-none disabled:opacity-50",
        "dark:focus:ring-gray-dark-4 dark:focus:ring-offset-gray-dark-1",
        !variant && [
          "bg-gray-12 text-gray-1",
          "dark:bg-gray-dark-12 dark:text-gray-dark-1",
        ],
        variant === "destructive" && [
          "bg-danger-10 text-gray-1 hover:bg-danger-11",
          "dark:bg-danger-dark-10 dark:text-gray-dark-1 dark:hover:bg-danger-dark-11",
        ],
        variant === "outline" && [
          "border border-gray-4 bg-transparent hover:bg-gray-2",
          "dark:border-gray-dark-4 dark:hover:bg-gray-dark-2",
        ],
        variant === "primary" && [
          "bg-primary-500  text-gray-1 hover:bg-primary-600",
          "dark:text-gray-dark-1",
        ],
        variant === "ghost" && [
          "bg-transparent hover:bg-gray-3 data-[state=open]:bg-transparent",
          "text-gray-12 hover:bg-gray-3",
          "dark:text-gray-dark-12 dark:hover:bg-gray-dark-3",
        ],
        variant === "link" && [
          "bg-transparent text-gray-12 underline-offset-4 hover:bg-transparent hover:underline",
          "dark:text-gray-dark-12",
        ],
        size === "sm" && "h-6 gap-1 [&>svg]:h-4",
        !size && "h-9 gap-2 py-2 [&>svg]:h-5",
        size === "lg" && "h-11 gap-3 [&>svg]:h-6",
        icon && "px-2",
        !icon && size === "sm" && "px-3",
        !icon && !size && "px-4",
        !icon && size === "lg" && "px-6",
        circle && "rounded-full",
        !circle && "rounded-md",
        block && "w-full"
      )}
      disabled={disabled || loading}
      ref={ref}
      {...props}
    >
      {loading && <Loader2 className="animate-spin" />}
      {children}
    </button>
  )
);
Button.displayName = "Button";

export default Button;
