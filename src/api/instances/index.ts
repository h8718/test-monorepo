export const instanceKeys = {
  instancesList: (
    namespace: string,
    {
      apiKey,
      limit,
      offset,
    }: { apiKey?: string; limit: number; offset: number }
  ) =>
    [
      {
        scope: "instance-list",
        apiKey,
        namespace,
        limit,
        offset,
      },
    ] as const,
  instancesInput: (
    namespace: string,
    { apiKey, instanceId }: { apiKey?: string; instanceId: string }
  ) =>
    [
      {
        scope: "instance-input",
        apiKey,
        namespace,
        instanceId,
      },
    ] as const,
  instancesOutput: (
    namespace: string,
    { apiKey, instanceId }: { apiKey?: string; instanceId: string }
  ) =>
    [
      {
        scope: "instance-output",
        apiKey,
        namespace,
        instanceId,
      },
    ] as const,
};
