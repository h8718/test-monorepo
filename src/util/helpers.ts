import clsx, { ClassValue } from "clsx";

import { FC } from "react";
import moment from "moment";
import { twMerge } from "tailwind-merge";

/**
 * this method combines the usage of two utility libraries:
 * - clsx https://github.com/lukeed/clsx#readme for more info
 *    clsx is a tiny utility for constructing className strings
 *    conditionally.
 * - tailwind-merge https://github.com/dcastil/tailwind-merge
 *    twMerge is a function to efficiently merge Tailwind CSS
 *    classes in JS without style conflicts.
 *
 *    input: twMergeClsx("bg-red", "bg-[#B91C1C]")
 *    output: bg-[#B91C1C]
 */
export const twMergeClsx = (...inputs: ClassValue[]) =>
  twMerge(clsx(...inputs));

type ConditionalWrapperProps = {
  condition: boolean;
  wrapper: (children: JSX.Element) => JSX.Element;
  children: JSX.Element;
};

export const ConditionalWrapper: FC<ConditionalWrapperProps> = ({
  condition,
  wrapper,
  children,
}) => (condition ? wrapper(children) : children);

export const formatLogTime = (time: string) =>
  moment(time).format("HH:mm:ss.mm");
