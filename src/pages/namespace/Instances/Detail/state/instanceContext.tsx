import {
  FC,
  PropsWithChildren,
  createContext,
  useContext,
  useReducer,
} from "react";

import { FiltersObj } from "~/api/logs/query/get";
import { Actions as InstanceActions } from "./type";

type StateType = {
  instanceId: string;
  filters?: FiltersObj;
};

const InstanceStateContext = createContext<StateType | null>(null);

const InstanceDispatchContext =
  createContext<React.Dispatch<InstanceActions> | null>(null);

function stateReducer(state: StateType, action: InstanceActions) {
  switch (action.type) {
    case "UPDATE_FILTER_STATE_NAME": {
      const newState: StateType = {
        ...state,
        filters: {
          ...state.filters,
          QUERY: {
            type: "MATCH",
            workflowName: state.filters?.QUERY?.workflowName ?? undefined,
            stateName: action.payload.stateName,
          },
        },
      };
      return newState;
    }
    case "UPDATE_FILTER_WORKFLOW": {
      const newState: StateType = {
        ...state,
        filters: {
          ...state.filters,
          QUERY: {
            type: "MATCH",
            workflowName: action.payload.workflowName,
            stateName: state.filters?.QUERY?.stateName ?? undefined,
          },
        },
      };
      return newState;
    }

    default: {
      return state;
    }
  }
}

const Provider: FC<PropsWithChildren & { instanceId: string }> = ({
  children,
  instanceId,
}) => {
  const [state, dispatch] = useReducer(stateReducer, { instanceId });

  return (
    <InstanceStateContext.Provider value={state}>
      <InstanceDispatchContext.Provider value={dispatch}>
        {children}
      </InstanceDispatchContext.Provider>
    </InstanceStateContext.Provider>
  );
};

const useInstanceId = () => {
  const context = useContext(InstanceStateContext);
  if (!context) {
    throw new Error("useFilter must be used within a InstanceContext");
  }

  return context.instanceId;
};

const useFilters = () => {
  const context = useContext(InstanceStateContext);
  if (!context) {
    throw new Error("useFilter must be used within a InstanceContext");
  }

  return context.filters;
};

const useActions = () => {
  const dispatch = useContext(InstanceDispatchContext);
  if (!dispatch) {
    throw new Error("useActions must be used within a InstanceDispatchContext");
  }

  return {
    updateFilterStateName: (stateName: string) => {
      dispatch({
        type: "UPDATE_FILTER_STATE_NAME",
        payload: {
          stateName,
        },
      });
    },
    updateFilterWorkflow: (workflowName: string) => {
      dispatch({
        type: "UPDATE_FILTER_WORKFLOW",
        payload: {
          workflowName,
        },
      });
    },
    resetFilterStateName: () => {
      dispatch({
        type: "UPDATE_FILTER_STATE_NAME",
        payload: {
          stateName: undefined,
        },
      });
    },
    resetFilterWorkflow: () => {
      dispatch({
        type: "UPDATE_FILTER_WORKFLOW",
        payload: {
          workflowName: undefined,
        },
      });
    },
  };
};

export {
  Provider as InstanceStateProvider,
  useFilters,
  useInstanceId,
  useActions,
};
