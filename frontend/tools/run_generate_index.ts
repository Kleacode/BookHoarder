/* create-indexの代替
 * usage: npm run generate:index
 *
 * 指定したpath直下のもの全てと親ディレクトリと名前が同じファイルのindexを作成する
 * */

type TargetDirType = { targetPath: string; depth: number; tsx?: boolean };
type ExcludesType = string[];
type ExportAllType = string[];

import Path from "node:path";

import { generator } from "./generate-index/main";

const rootDir = Path.resolve("./app");

// index.jsを作成するdir
// frontend/から先を指定する
// とにかく末端の深さまでを指定したいときはdepth = Infinityとか指定してください
const targetSrcDirs: TargetDirType[] = [
	{ targetPath: "components", depth: 1, tsx: true },
	{ targetPath: "models", depth: 1 },
	{ targetPath: "utils", depth: 1 },
];

// 除外したいディレクトリを指定
const excludeSrcDirs: ExcludesType = [];

// 親ディレクトリと名前が異なる場合でもファイルを出力したいものを指定
const exportAllSrcDirs: ExportAllType = [];

// featuresディレクトリを指定
const featuresDir: string = "";

// featuresディレクトリ向けのindex.tsを作成する際の深さの指定
const featuresDepth: number = 1;

const params = {
	rootDir,
	targetSrcDirs,
	excludeSrcDirs,
	exportAllSrcDirs,
	featuresDir,
	featuresDepth,
};

generator(params);
