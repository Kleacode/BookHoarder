export interface TableHeaderCellProps {
	children: React.ReactNode;
}
export const TableHeaderCell = ({ children }: TableHeaderCellProps) => {
	return <div className="flex">{children}</div>;
};
