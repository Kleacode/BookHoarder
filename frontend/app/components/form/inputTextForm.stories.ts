import type { Meta, StoryObj } from "@storybook/react";
import { InputTextForm } from "./InputTextForm";

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories#default-export
const meta = {
	title: "Example/InputTextForm",
	component: InputTextForm,
	parameters: {
		// Optional parameter to center the component in the Canvas. More info: https://storybook.js.org/docs/configure/story-layout
		layout: "centered",
	},
	// This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
	tags: ["autodocs"],
} satisfies Meta<typeof InputTextForm>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Example1: Story = {
	args: {
		label: "title",
	},
};

export const Example2: Story = {
	args: {
		label: "placeholder",
		placeholder: "placeholder",
	},
};
