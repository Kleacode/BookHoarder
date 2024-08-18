/* create-indexの代替
 * usage: npm run generate:index
 *
 * 指定したpath直下のもの全てと親ディレクトリと名前が同じファイルのindexを作成する
 * */

// argsの説明

// rootDir: string
//   ex) Path.resolve(".");

// targetSrcDirs: {targetPath: string, depth: number, tsx: boolean}[]
//   index.jsを作成するdir
//   frontend/から先を指定する
//   とにかく末端の深さまでを指定したいときはdepth = Infinityとか指定してください
// ex)
// [
//   { targetPath: "atoms", depth: 1 },
//   { targetPath: "components/features", depth: 1 },
//   { targetPath: "components/mols", depth: 2 },
//   { targetPath: "decorators", depth: 1 },
//   { targetPath: "defs/constants", depth: 1 },
//   { targetPath: "hooks", depth: 1 },
//   { targetPath: "layouts", depth: 1 },
//   { targetPath: "models", depth: 1 },
//   { targetPath: "types", depth: 1 },
//   { targetPath: "testing/factories", depth: 1 },
// ];

// excludeSrcDirs: string[]
//   除外したいディレクトリを指定

// exportAllSrcDirs: string[]
//   親ディレクトリと名前が異なる場合でもファイルを出力したいものを指定

// featuresDir: string
//   featuresディレクトリを指定

// featuresDepth: number
//   featuresディレクトリ向けのindex.tsを作成する際の深さの指定

import Path from "node:path";
import generateIndex from "./generate-index";
import type { TargetDir, Targets } from "./types";

type PropTypes = {
	rootDir: string;
	targetSrcDirs: TargetDir[];
	excludeSrcDirs: string[];
	exportAllSrcDirs: string[];
	featuresDir: string;
	featuresDepth: number;
};

const generator = ({
	rootDir,
	targetSrcDirs,
	excludeSrcDirs,
	exportAllSrcDirs,
}: PropTypes) => {
	const targets: Targets = {
		dirs: targetSrcDirs.map((t) => ({
			targetPath: Path.join(rootDir, t.targetPath),
			depth: t.depth,
			tsx: t.tsx ?? false,
		})),
		excludes: excludeSrcDirs.map((dir) => Path.join(rootDir, dir)),
		exportAll: exportAllSrcDirs.map((dir) => Path.join(rootDir, dir)),
	};
	// index.tsを生成/編集する
	generateIndex(targets);
};

export { generator };
