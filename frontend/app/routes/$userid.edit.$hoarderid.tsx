import { Button, Header } from "@/components";
import { DropDownForm } from "@/components/form/DropDownForm";
import { InputTextForm } from "@/components/form/InputTextForm";
import { TagForm } from "@/components/form/tagform/TagForm";

export default function Index() {
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
				<InputTextForm label="タイトル" />
				<DropDownForm<string> label="状態" options={[]} />
				<TagForm
					SearchSuggest={(s) => {
						return [];
					}}
				/>
				<Button label="Submit" />
			</div>
		</div>
	);
}
