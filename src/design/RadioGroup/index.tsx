import * as RadioGroupPrimitive from "@radix-ui/react-radio-group";
import * as React from "react";

import { Circle } from "lucide-react";
import clsx from "clsx";

const RadioGroup = React.forwardRef<
  React.ElementRef<typeof RadioGroupPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof RadioGroupPrimitive.Root>
>(({ className, ...props }, ref) => (
  <RadioGroupPrimitive.Root
    className={clsx("grid gap-2", className)}
    {...props}
    ref={ref}
  />
));
RadioGroup.displayName = RadioGroupPrimitive.Root.displayName;

const RadioGroupItem = React.forwardRef<
  React.ElementRef<typeof RadioGroupPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof RadioGroupPrimitive.Item>
>(({ className, ...props }, ref) => (
  <RadioGroupPrimitive.Item
    ref={ref}
    className={clsx(
      "text:fill-slate-50 h-4 w-4 rounded-full border  focus:outline-none focus:ring-2  focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
      "border-gray-6 text-gray-12 hover:border-gray-7 hover:text-gray-1 focus:ring-gray-7 focus:ring-offset-gray-1",
      "dark:border-gray-dark-6 dark:text-gray-dark-12  dark:hover:border-gray-dark-7 dark:hover:text-gray-dark-1 dark:focus:ring-gray-dark-7 dark:focus:ring-offset-gray-dark-1",
      className
    )}
    {...props}
  >
    <RadioGroupPrimitive.Indicator className="flex items-center justify-center">
      <Circle
        className={clsx(
          "h-2.5 w-2.5",
          " fill-gray-12",
          " dark:fill-gray-dark-12"
        )}
      />
    </RadioGroupPrimitive.Indicator>
  </RadioGroupPrimitive.Item>
));
RadioGroupItem.displayName = RadioGroupPrimitive.Item.displayName;

export { RadioGroup, RadioGroupItem };
