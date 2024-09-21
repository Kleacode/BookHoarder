export interface TableProps {
	children: React.ReactNode;
}
export const Table = ({ children }: TableProps) => {
	return <div className="grid">{children}</div>;
};
