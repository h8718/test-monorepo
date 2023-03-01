import {
  Box,
  Bug,
  Calendar,
  ChevronsUpDown,
  FolderOpen,
  FolderTree,
  Github,
  Home,
  Layers,
  LogOut,
  Network,
  Play,
  PlusCircle,
  Settings,
  Settings2,
  Users,
} from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "../dropdown";

import Button from "../button";
import Logo from "./Logo";
import clsx from "clsx";

export default {
  title: "Design System/Compositions/App Shell",
};

const navigation = [
  { name: "Explorer", href: "#", icon: FolderTree, current: true },
  { name: "Monitoring", href: "#", icon: Bug, current: false },
  { name: "Instances", href: "#", icon: Box, current: false },
  { name: "Events", href: "#", icon: Calendar, current: false },
  { name: "Services", href: "#", icon: Layers, current: false },
  { name: "Settings", href: "#", icon: Settings, current: false },
];

const enterprise = [
  { name: "Gateway", href: "#", icon: Network, current: false },
  { name: "Permissions", href: "#", icon: Users, current: false },
];

export const Default = () => (
  <div className="min-h-full">
    <div className="lg:fixed lg:inset-y-0 lg:flex lg:w-64 lg:flex-col lg:border-r lg:border-gray-200 lg:pt-5">
      <div className="px-6">
        <Logo className="h-8 w-auto" />
      </div>
      <div className="mt-5 flex h-0 flex-1 flex-col overflow-y-auto pt-1">
        <nav className="mt-6 px-3 space-y-3">
          <div className="space-y-1">
            {navigation.map((item) => (
              <a
                key={item.name}
                href={item.href}
                className={clsx(
                  item.current
                    ? "bg-primary50 text-gray-900"
                    : "text-gray-700 hover:text-gray-900 hover:bg-gray-50",
                  "group flex items-center px-2 py-2 text-sm font-medium rounded-md"
                )}
                aria-current={item.current ? "page" : undefined}
              >
                <item.icon
                  className={clsx(
                    item.current
                      ? "text-gray-500"
                      : "text-gray-400 group-hover:text-gray-500",
                    "mr-3 flex-shrink-0 h-6 w-6"
                  )}
                  aria-hidden="true"
                />
                {item.name}
              </a>
            ))}
          </div>
          <div className="divider"></div>
          <div className="space-y-1">
            {enterprise.map((item) => (
              <a
                key={item.name}
                href={item.href}
                className={clsx(
                  item.current
                    ? "bg-gray-200 text-gray-900"
                    : "text-gray-700 hover:text-gray-900 hover:bg-gray-50",
                  "group flex items-center px-2 py-2 text-sm font-medium rounded-md"
                )}
                aria-current={item.current ? "page" : undefined}
              >
                <item.icon
                  className={clsx(
                    item.current
                      ? "text-gray-500"
                      : "text-gray-400 group-hover:text-gray-500",
                    "mr-3 flex-shrink-0 h-6 w-6"
                  )}
                  aria-hidden="true"
                />
                {item.name}
              </a>
            ))}
          </div>
        </nav>
      </div>
      <div className="flex flex-shrink-0 border-t border-gray-200 p-2 group w-full rounded-md py-5 text-left text-sm font-medium text-gray-700">
        <span className="flex w-full min-w-0 items-center justify-between space-x-3">
          <div className="avatar placeholder">
            <div className="bg-neutral-focus text-neutral-content rounded-full w-10">
              <span className="text-3xl">A</span>
            </div>
          </div>
          <span className="flex min-w-0 flex-1 flex-col">
            <span className="truncate text-sm font-medium text-gray-900">
              admin
            </span>
            <span className="truncate text-sm text-gray-400">
              Version: 78c688e
            </span>
          </span>
          <Button color="link" className="text-grayDark-gray9">
            <LogOut />
          </Button>
        </span>
      </div>
    </div>
    <div className="flex flex-col lg:pl-64">
      <main className="flex-1">
        <div className="border-b border-gray-200 px-4 py-4 sm:flex sm:items-center sm:justify-between">
          <div className="min-w-0 flex-1">
            <h1 className="text-lg font-medium leading-6 text-gray-900 sm:truncate">
              <div className="text-sm breadcrumbs">
                <ul>
                  <li>
                    <a className="gap-2">
                      <Home className="h-4 w-auto" />
                      My-namespace
                    </a>
                    &nbsp;
                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button size="xs" color="ghost" circle>
                          <ChevronsUpDown />
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent className="w-56">
                        <DropdownMenuLabel>Namespaces</DropdownMenuLabel>
                        <DropdownMenuSeparator />
                        <DropdownMenuCheckboxItem checked>
                          My-namespace
                        </DropdownMenuCheckboxItem>
                        <DropdownMenuCheckboxItem>
                          second-namespace
                        </DropdownMenuCheckboxItem>
                        <DropdownMenuCheckboxItem>
                          another-namespace
                        </DropdownMenuCheckboxItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem>
                          <PlusCircle className="mr-2 h-4 w-4" />
                          <span>Create new namespace</span>
                        </DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </li>
                  <li>
                    <a className="gap-2">
                      <Github className="h-4 w-auto" />
                      Example Mirror
                    </a>
                  </li>
                  <li>
                    <a className="gap-2">
                      <FolderOpen className="h-4 w-auto" />
                      Folder
                    </a>
                  </li>
                  <li>
                    <a className="gap-2">
                      <Play className="h-4 w-auto" />
                      workflow.yml
                    </a>
                  </li>
                </ul>
              </div>
            </h1>
          </div>
          <div className="mt-4 flex sm:mt-0 sm:ml-4">
            <Button color="ghost" className="px-1">
              <Settings2 className="w-5 h-auto" />
            </Button>
          </div>
        </div>
      </main>
    </div>
  </div>
);
