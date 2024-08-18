import type { Targets } from "./types";
import Path from "node:path";
import fs from "node:fs";
import _ from "lodash";
import {
	getFilenameWithoutExtension,
	buildIndexTs,
	hasExportsFile,
	generateIndexFileName,
	isDir,
	hasExportDefaultFile,
} from "./utils";

export default function generateIndex(targets: Targets) {
	// 指定されたパス配下のtsファイルからexport の記述があるファイルを指定された深さで再帰的に集めてリストで返す
	const gatherFilesHasExport = (targetPath: string, depth = 1) => {
		console.log(targetPath);
		let targetFileList: string[] = [];
		const targetFiles = fs
			.readdirSync(targetPath)
			.map((f) => Path.join(targetPath, f));
		// biome-ignore lint/complexity/noForEach: <explanation>
		targetFiles.forEach((f) => {
			if (hasExportsFile(f)) targetFileList.push(f);

			if (isDir(f) && depth > 0) {
				targetFileList = targetFileList.concat(
					gatherFilesHasExport(f, depth - 1),
				);
			}
		});
		return targetFileList;
	};

	// 設定されているファイルかどうかをフィルタリングする
	const filterConfiguredFiles = (outputFiles: string[], targetPath: string) =>
		outputFiles.filter((f) => {
			// excludesに定義されているファイルは除外する
			if (targets.excludes.some((dir) => f.includes(dir))) return false;
			// 型定義ファイルは除外する
			if (Path.basename(f).includes(".d.ts")) return false;

			if (targets.exportAll.some((dir) => f.includes(dir))) return true;
			// 指定したpath直下のものと親ディレクトリと名前が同じファイル以外は除外
			if (Path.dirname(f) === targetPath) return true;
			if (getFilenameWithoutExtension(f) === Path.basename(Path.dirname(f)))
				return true;
			return false;
		});

	// targetファイルの順番をソートしてindexファイル生成の準備をする
	const sortFiles = (
		outputFiles: string[],
	): { path: string; expDefault: boolean }[] =>
		outputFiles
			.sort((pre, next) => {
				// ファイル名で並び替え
				if (Path.basename(pre) > Path.basename(next)) return 1;
				return -1;
			})
			.map((f) => {
				return { path: f, expDefault: hasExportDefaultFile(f) };
			});

	for (const { targetPath, depth, tsx } of targets.dirs) {
		const filesHasExport = gatherFilesHasExport(targetPath, depth);
		const filteredFiles = filterConfiguredFiles(filesHasExport, targetPath);
		const sortedFiles = sortFiles(filteredFiles);

		const indexText = buildIndexTs(sortedFiles, targetPath);
		const outputFilename = generateIndexFileName(tsx);
		fs.writeFile(
			Path.join(targetPath, outputFilename),
			indexText,
			// biome-ignore lint/suspicious/noExplicitAny: <explanation>
			(err: any) => {
				if (err) throw err;
				console.log(`output ${outputFilename} @${targetPath}`);
			},
		);
	}
}
