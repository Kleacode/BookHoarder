import { Header } from "@/components";
import { Table } from "@/components/table/Table";
import { TableBody } from "@/components/table/TableBody";
import { TableCell } from "@/components/table/TableCell";
import { TableHeader } from "@/components/table/TableHeader";
import { TableHeaderCell } from "@/components/table/TableHeaderCell";
import { TableRow } from "@/components/table/TableRow";
import { useGetBooks } from "@/hooks/useGetBooks";

export default function Index() {
	const books = useGetBooks(undefined, {});
	return (
		<div className="font-sans p-4">
			<Header
				active="booklist"
				links={[
					{ id: "booklist", path: "booklist", name: "本の一覧" },
					{ id: "hoarder", path: "hoarder", name: "積読リスト" },
					{ id: "taglist", path: "taglist", name: "タグ一覧" },
					{ id: "register", path: "form", name: "本の登録" },
				]}
			/>
			<Table>
				<TableHeader>
					<TableRow>
						<TableHeaderCell>title</TableHeaderCell>
					</TableRow>
				</TableHeader>
				<TableBody>
					{books.map((e) => {
						return (
							<TableRow key={e.bookId}>
								<TableCell>{e.title}</TableCell>
							</TableRow>
						);
					})}
				</TableBody>
			</Table>
		</div>
	);
}
