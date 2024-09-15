export interface TagFormItemProps {
	label: string;
	onMouseDown?: (e: React.MouseEvent<HTMLButtonElement>) => void;
}

export const TagFormItem = ({ label, ...props }: TagFormItemProps) => {
	return (
		<li>
			<button
				type="button"
				className="truncate inline-flex w-full px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
				{...props}
			>
				{label}
			</button>
		</li>
	);
};
