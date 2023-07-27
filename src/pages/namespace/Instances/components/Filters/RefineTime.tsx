import { Command, CommandGroup, CommandList } from "~/design/Command";

import { FiltersObj } from "~/api/instances/query/get";
import Input from "~/design/Input";
import moment from "moment";
import { useState } from "react";
import { useTranslation } from "react-i18next";

const RefineTime = ({
  field,
  date,
  setFilter,
}: {
  field: "AFTER" | "BEFORE";
  date: Date;
  setFilter: (filter: FiltersObj) => void;
}) => {
  const { t } = useTranslation();
  const [time, setTime] = useState<string>(moment(date).format("HH:mm:ss"));

  const setTimeOnDate = () => {
    const [hr, min, sec] = time.split(":").map((item) => Number(item));

    if (!hr || !min || !sec) {
      console.error("Invalid time string in setTimeOnDate");
      return;
    }

    date.setHours(hr, min, sec);
    setFilter({
      [field]: { type: "MATCH", value: date },
    });
  };

  const handleKeyDown = (event: { key: string }) => {
    event.key === "Enter" && setTimeOnDate();
  };

  return (
    <Command>
      <CommandList className="max-h-[460px]">
        <CommandGroup
          heading={t("pages.instances.list.filter.menuHeading.time")}
        >
          <Input
            type="time"
            step={1}
            value={time}
            onChange={(event) => {
              setTime(event.target.value);
            }}
            onKeyDown={handleKeyDown}
          />
        </CommandGroup>
      </CommandList>
    </Command>
  );
};

export default RefineTime;
