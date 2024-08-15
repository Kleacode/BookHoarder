import {
	getFilenameWithoutExtension,
	buildIndexTs,
	isIndexFile,
	hasExportsFile,
	generateIndexFileName,
	isDir,
} from "./utils";
import Path from "node:path";
import fs from "node:fs";

export default function generateFeatureIndex(
	target: string,
	depth: number | undefined,
) {
	// featureディレクトリ配下のフォルダに指定されたパス配下のjsファイルからexport defaultの記述があるファイルを指定された深さで再帰的に集めてリストで返す
	const gatherFilesHasExport = (targetPath: string, depth = 1) => {
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
		outputFiles.filter((f: string) => {
			// 型定義ファイルは除外する
			if (Path.basename(f).includes(".d.ts")) return false;

			// 指定したpath直下のものと親ディレクトリと名前が同じファイル以外は除外
			if (Path.dirname(f) === targetPath) return true;
			if (getFilenameWithoutExtension(f) === Path.basename(Path.dirname(f)))
				return true;
			return false;
		});

	// targetファイルの順番をソートしてindexファイル生成の準備をする
	const sortFiles = (outputFiles: string[]) =>
		outputFiles
			.sort((pre: string, next: string) => {
				// ファイル名で並び替え
				if (Path.basename(pre) > Path.basename(next)) return 1;
				return -1;
			})
			.map((f) => {
				return { path: f };
			});

	const featureDirs = fs
		.readdirSync(target)
		.map((f) => Path.join(target, f))
		.filter((f) => !isIndexFile(f));

	const targets = featureDirs.reduce(
		// NOTE: 可読性重視でスプレッド演算子を使っています。
		// パフォーマンス上の問題が出た場合はconcatなどに切り替えてください
		// cf.) https://qiita.com/KtheS/items/08103c6ec7231abb50f9
		(prev: string[], curr) => [
			// biome-ignore lint/performance/noAccumulatingSpread: <explanation>
			...prev,
			...fs
				.readdirSync(curr)
				.map((f) => Path.join(curr, f))
				.filter((f) => !isIndexFile(f) && isDir(f)),
		],
		[],
	);

	// biome-ignore lint/complexity/noForEach: <explanation>
	targets.forEach((targetPath: string) => {
		const filesHasExport = gatherFilesHasExport(targetPath, depth);
		const filteredFiles = filterConfiguredFiles(filesHasExport, targetPath);
		const sortedFiles = sortFiles(filteredFiles);

		const indexText = buildIndexTs(sortedFiles, targetPath);
		const outputFilename = generateIndexFileName(true);
		fs.writeFile(Path.join(targetPath, outputFilename), indexText, (err) => {
			if (err) throw err;
			console.log(`output ${outputFilename} @${targetPath}`);
		});
	});
}
