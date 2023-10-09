import type { Meta, StoryObj } from "@storybook/react";
import Notification from "./index";

const meta = {
  title: "Components/Notification",
  component: Notification,
} satisfies Meta<typeof Notification>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  render: ({ ...args }) => <Notification {...args} />,
  argTypes: {
    hasMessage: {
      description: "Notification Icon signals existence of Messages",
      control: "boolean",
      type: { name: "boolean", required: false },
    },
  },
};

export const NotificationHasMessage = () => (
  <div className="flex space-x-2">
    <Notification hasMessage={true} />
  </div>
);
