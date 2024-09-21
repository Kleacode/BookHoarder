export interface TableCellProps {
	children: React.ReactNode;
}
export const TableCell = ({ children }: TableCellProps) => {
	return <div className="flex">{children}</div>;
};
