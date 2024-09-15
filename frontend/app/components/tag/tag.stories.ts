import type { Meta, StoryObj } from "@storybook/react";
import { Tag } from "./Tag";

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories#default-export
const meta = {
	title: "Example/Tag",
	component: Tag,
	parameters: {
		// Optional parameter to center the component in the Canvas. More info: https://storybook.js.org/docs/configure/story-layout
		layout: "centered",
	},
	// This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
	tags: ["autodocs"],
	// More on argTypes: https://storybook.js.org/docs/api/argtypes
	argTypes: {
		backgroundColor: { control: "color" },
	},
} satisfies Meta<typeof Tag>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Example1: Story = {
	args: {
		label: "AAAAA",
	},
};

export const Example2: Story = {
	args: {
		label: "あいうえお",
		backgroundColor: "#800080",
	},
};

export const Example3: Story = {
	args: {
		label: "あいうえおかきくけこさしすせそ",
		backgroundColor: "#000080",
	},
};

export const Example4: Story = {
	args: {
		label: "AAAAA",
		onClick: (e) => {},
	},
};

export const Example5: Story = {
	args: {
		label: "あいうえおかきくけこさしすせそ",
		onClick: (e) => {},
	},
};
