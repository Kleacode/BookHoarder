export interface TableRowProps {
	children: React.ReactNode;
	onClick?: (e: React.MouseEvent<HTMLDivElement>) => void;
}
export const TableRow = ({ children, ...props }: TableRowProps) => {
	return (
		<div
			className="grid grid-cols-subgrid col-span-3 border p-2 hover:bg-gray-100"
			{...props}
		>
			{children}
		</div>
	);
};
