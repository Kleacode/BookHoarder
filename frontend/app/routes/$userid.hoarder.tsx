import { Header } from "@/components";
import { StatusLabel } from "@/components/status/StatusLabel";
import { Table } from "@/components/table/Table";
import { TableBody } from "@/components/table/TableBody";
import { TableCell } from "@/components/table/TableCell";
import { TableHeader } from "@/components/table/TableHeader";
import { TableHeaderCell } from "@/components/table/TableHeaderCell";
import { TableRow } from "@/components/table/TableRow";
import { Tag, type TagProps } from "@/components/tag/Tag";
import { useGetHoarderBooks } from "@/hooks/useGetHoarderBooks";
import { useParams } from "@remix-run/react";

export default function Index() {
	// TODO useParams Hook化
	const params = useParams();
	const userID = Number.parseInt(params.userid ?? "-1", 10);
	const hoarderBooks = useGetHoarderBooks({ userId: userID }, {});
	// TODO header 共通化
	return (
		<div className="font-sans p-4">
			<Header
				active="hoarder"
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
						<TableHeaderCell>tags</TableHeaderCell>
						<TableHeaderCell>status</TableHeaderCell>
					</TableRow>
				</TableHeader>
				<TableBody>
					{hoarderBooks?.map((e) => {
						return (
							<TableRow key={e.hoarderId}>
								<TableCell>{e.book?.title}</TableCell>
								<TableCell>
									{e.tags?.map((t) => {
										return <Tag key={t.id} label={t.name} />;
									})}
								</TableCell>
								<TableCell>
									<StatusLabel status={e.status} />
								</TableCell>
							</TableRow>
						);
					})}
				</TableBody>
			</Table>
		</div>
	);
}
