export interface TableBodyProps {
	children: React.ReactNode;
}
export const TableBody = ({ children }: TableBodyProps) => {
	return <div className="grid grid-cols-3">{children}</div>;
};
