import { Tag } from "@/models/model";
import { useRef, useState } from "react";
import { Tag as TagComp } from "../../tag/Tag";
import { InputTextForm } from "../InputTextForm";
import type { TagFormItemProps } from "./TagFormItem";
import { TagSearchResult } from "./TagSearchResult";

export interface SuggestItem extends TagFormItemProps {
	key: React.Key;
}

export interface TagFormProps {
	SearchSuggest: (searchWord: string) => SuggestItem[];
}

export const TagForm = ({ SearchSuggest }: TagFormProps) => {
	const [tags, setTags] = useState<Tag[]>([]);
	const [suggestItems, setSuggestItems] = useState<SuggestItem[]>([]);
	const [searchWord, setSearchWord] = useState<string>("");
	const [isShowDropDown, SetIsShowDropDown] = useState<boolean>(false);

	const newIdCounter = useRef<number>(-1);

	const onChangeSearchWord = (word: string) => {
		setSearchWord(word);
		setSuggestItems(SearchSuggest(word));
	};

	const CreateNewTag = (name: string) => {
		setTags([...tags, Tag.fromJson({ id: newIdCounter.current, name })]);
		newIdCounter.current -= 1;
	};

	const OnClickTag =
		(target: Tag) => (event: React.MouseEvent<HTMLButtonElement>) => {
			setTags(
				tags.filter((tag, idx) => {
					return !tag.isEqual(target);
				}),
			);
		};

	const OnFocus = (event: React.FocusEvent<HTMLInputElement>) => {
		SetIsShowDropDown(true);
	};

	const OnBlur = (event: React.FocusEvent<HTMLInputElement>) => {
		SetIsShowDropDown(false);
	};

	const OnChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		onChangeSearchWord(event.target.value);
	};

	return (
		<div>
			<InputTextForm
				label="tag"
				placeholder="Search for tags"
				onChange={OnChange}
				onFocus={OnFocus}
				onBlur={OnBlur}
			/>

			<TagSearchResult
				isDisplay={isShowDropDown}
				searchWords={searchWord}
				createNewTag={CreateNewTag}
				suggests={suggestItems}
			/>

			<div>
				{tags.map((e) => {
					return <TagComp key={e.id} label={e.name} onClick={OnClickTag(e)} />;
				})}
			</div>
		</div>
	);
};
