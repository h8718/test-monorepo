import {
  Command,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "~/design/Command";
import {
  FiltersObj,
  statusValues,
  triggerValues,
} from "~/api/instances/query/get";

import { Datepicker } from "~/design/Datepicker";
import TextInput from "./TextInput";
import { useTranslation } from "react-i18next";

const optionMenus = {
  STATUS: statusValues,
  TRIGGER: triggerValues,
};

const FieldSubMenu = ({
  field,
  value,
  date,
  setFilter,
  clearFilter,
}: {
  field: keyof FiltersObj;
  value?: string;
  date?: Date;
  setFilter: (filter: FiltersObj) => void;
  clearFilter: (field: keyof FiltersObj) => void;
}) => {
  const { t } = useTranslation();

  const setDate = (type: "AFTER" | "BEFORE", value: Date) => {
    setFilter({
      [type]: { type, value },
    });
  };

  return (
    <>
      {field === "AS" && (
        <TextInput
          field={field}
          setFilter={setFilter}
          clearFilter={clearFilter}
          value={value}
        />
      )}
      {(field === "STATUS" || field === "TRIGGER") && (
        <Command value={value}>
          <CommandInput
            autoFocus
            placeholder={t("pages.instances.list.filter.placeholder.STATUS")}
          />
          <CommandList>
            <CommandGroup
              heading={t("pages.instances.list.filter.menuHeading.STATUS")}
            >
              {optionMenus[field].map((option) => (
                <CommandItem
                  key={option}
                  value={option}
                  onSelect={() =>
                    setFilter({
                      [field]: {
                        value: option,
                        // TODO: Move this decision to the API layer?
                        type: field === "TRIGGER" ? "CONTAINS" : "MATCH",
                      },
                    })
                  }
                >
                  {t(`pages.instances.list.filter.option.${option}`)}
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      )}
      {(field === "AFTER" || field === "BEFORE") && (
        <Command>
          <CommandList className="max-h-[460px]">
            <CommandGroup
              heading={t(`pages.instances.list.filter.menuHeading.${field}`)}
            >
              <Datepicker
                mode="single"
                selected={date}
                onSelect={(value) => value && setDate(field, value)}
              />
            </CommandGroup>
          </CommandList>
        </Command>
      )}
    </>
  );
};

export default FieldSubMenu;
