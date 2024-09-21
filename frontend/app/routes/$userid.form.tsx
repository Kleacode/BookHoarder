import { Button, Header } from "@/components";
import { DropDownForm } from "@/components/form/DropDownForm";
import { InputTextForm } from "@/components/form/InputTextForm";
import { type SuggestItem, TagForm } from "@/components/form/tagform/TagForm";
import { postHoarder } from "@/hooks/postHoarder";
import { useGetTags } from "@/hooks/useGetTags";
import type { StatusType } from "@/libs/schemas/schemaHealper";
import { ConvertStatusTypeToLabel } from "@/utils/ConvertStatusType";
import { useParams } from "@remix-run/react";
import { useState } from "react";

export default function Index() {
	const [title, setTitle] = useState<string>("");
	const [suggests, setSuggests] = useState<SuggestItem[]>([]);
	const params = useParams();
	const userID = Number.parseInt(params.userid ?? "-1", 10);

	const postNewHoarder = () => {
		postHoarder(
			{ userId: userID },
			{
				book: { title: title },
				tags: [],
				status: "wip",
			},
		);
	};

	// TODO nestしたHook呼び出しは不可能
	const searchSuggest = (w: string) => {
		const tags = useGetTags({ userId: userID }, { params: { name: w } });
		const result: SuggestItem[] =
			tags?.map((e) => ({ key: e.tagId, label: e.name }));
		setSuggests(result);
	};

	const statusTypes: StatusType[] = ["todo", "wip", "done"];
	const statusOptions = statusTypes.map((e) => ConvertStatusTypeToLabel(e));

	return (
		<div className="font-sans p-4">
			<Header
				active="register"
				links={[
					{ id: "booklist", path: "booklist", name: "本の一覧" },
					{ id: "hoarder", path: "hoarder", name: "積読リスト" },
					{ id: "taglist", path: "taglist", name: "タグ一覧" },
					{ id: "register", path: "form", name: "本の登録" },
				]}
			/>
			<div className="flex flex-col gap-16">
				<InputTextForm
					label="タイトル"
					onChange={(e) => {
						setTitle(e.target.value);
					}}
				/>
				<DropDownForm<string> label="状態" options={statusOptions} />
				<TagForm
					SearchSuggest={(w) => {
						searchSuggest(w);
						return suggests;
					}}
				/>
				<Button label="Submit" onClick={postNewHoarder} />
			</div>
		</div>
	);
}
