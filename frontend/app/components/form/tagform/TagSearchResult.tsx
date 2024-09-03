import type { SuggestItem } from "./TagForm";
import { TagFormItem } from "./TagFormItem";

export interface TagSearchResultProps {
	isDisplay: boolean;
	suggests?: SuggestItem[];
	searchWords: string;
	createNewTag: (s: string) => void;
}

export const TagSearchResult = ({
	isDisplay,
	suggests,
	searchWords,
	createNewTag,
}: TagSearchResultProps) => {
	return (
		<div
			style={{ display: isDisplay ? "block" : "none" }}
			className="z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700"
		>
			<ul
				className="py-2 text-sm text-gray-700 dark:text-gray-200"
				aria-labelledby="dropdown-button"
			>
				{suggests?.map((e) => {
					return (
						<TagFormItem
							key={e.key}
							label={e.label}
							onMouseDown={e.onMouseDown}
						/>
					);
				})}
				<TagFormItem
					label={`Create·${searchWords}`}
					onMouseDown={(e) => createNewTag(searchWords)}
				/>
			</ul>
		</div>
	);
};
