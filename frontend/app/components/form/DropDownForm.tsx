import _ from "lodash";

export interface DropDownFormType<T> {
	label: string;
	options: T[];
	onChange?: (e: React.ChangeEvent<HTMLSelectElement>) => void;
}

export const DropDownForm = <T,>({
	label,
	options,
	...props
}: DropDownFormType<T>) => {
	return (
		<div className="w-full mb-6 md:mb-0">
			<div className="text-gray-700 text-xs font-bold">{label}</div>
			<div className="relative">
				<select
					className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
					{...props}
				>
					{options.map((e) => {
						return <option key={_.toString(e)}>{_.toString(e)}</option>;
					})}
				</select>
				<div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
					<svg
						className="fill-current h-4 w-4"
						xmlns="http://www.w3.org/2000/svg"
						viewBox="0 0 20 20"
					>
						<title>arrow</title>
						<path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
					</svg>
				</div>
			</div>
		</div>
	);
};
