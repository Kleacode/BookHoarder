export type Obj = { [x: string]: unknown };

export const isObject = (value: unknown): value is object => {
	const type = typeof value;
	return value !== null && (type === "object" || type === "function");
};
