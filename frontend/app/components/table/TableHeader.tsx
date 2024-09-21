export interface TableHeaderProps {
	children: React.ReactNode;
}

export const TableHeader = ({ children }: TableHeaderProps) => {
	return <div className="grid grid-cols-3 bg-gray-200">{children}</div>;
};
