export type TargetDir = { targetPath: string; depth: number; tsx?: boolean };

export type Targets = {
	dirs: { targetPath: string; depth: number; tsx: boolean }[];
	excludes: string[];
	exportAll: string[];
};
