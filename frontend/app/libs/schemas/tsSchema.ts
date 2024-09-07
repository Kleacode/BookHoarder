/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/books": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** 登録されている本を取得する */
        get: {
            parameters: {
                query?: {
                    title?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description 本の一覧 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["exist_book"][];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/tags": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** 登録されているタグを取得する */
        get: {
            parameters: {
                query?: {
                    name?: string;
                };
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description タグの一覧 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["exist_tag"][];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/books/{bookId}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** 特定の本1冊の情報を取得する */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    bookId: components["parameters"]["bookId"];
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description 本の情報 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["exist_book"];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/{userId}/books/{bookId}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /** ユーザーが登録した本を削除する */
        delete: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                    bookId: components["parameters"]["bookId"];
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description 削除成功 */
                204: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content?: never;
                };
                default: components["responses"]["default_response"];
            };
        };
        options?: never;
        head?: never;
        /** ユーザーが登録した本の情報を更新する */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                    bookId: components["parameters"]["bookId"];
                };
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["post_book"];
                };
            };
            responses: {
                /** @description 更新成功 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["exist_book"];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        trace?: never;
    };
    "/{userId}/hoarder": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** ユーザーの積読リストから、本の一覧を取得する。 */
        get: {
            parameters: {
                query?: {
                    status?: string;
                    tags?: number[];
                };
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description 本の一覧 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["hoarder_book"][];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        put?: never;
        /** 本を新しく登録する。そのままユーザーの積読リストにも登録する。 */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                };
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["post_hoarder"];
                };
            };
            responses: {
                /** @description 登録成功 */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["hoarder_book"];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/{userId}/hoarder/{bookId}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** ユーザーの積読リストに本を登録する */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                    bookId: components["parameters"]["bookId"];
                };
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["post_hoarder"];
                };
            };
            responses: {
                /** @description 登録成功 */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["hoarder_book"];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        /** ユーザーの積読リストにある本を削除する */
        delete: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                    bookId: components["parameters"]["bookId"];
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description 削除成功 */
                204: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content?: never;
                };
                default: components["responses"]["default_response"];
            };
        };
        options?: never;
        head?: never;
        /** ユーザーの積読リストにある本の状態を更新する */
        patch: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                    bookId: components["parameters"]["bookId"];
                };
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["post_hoarder"];
                };
            };
            responses: {
                /** @description 更新成功 */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["hoarder_book"];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        trace?: never;
    };
    "/{userId}/tags": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        /** タグを新しく登録する */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                };
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["tag_info"];
                };
            };
            responses: {
                /** @description 登録成功 */
                201: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["exist_tag"];
                    };
                };
                default: components["responses"]["default_response"];
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/{userId}/tags/{tagId}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get?: never;
        put?: never;
        post?: never;
        /** タグを削除する */
        delete: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    userId: components["parameters"]["userId"];
                    tagId: components["parameters"]["tagId"];
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description 削除成功 */
                204: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content?: never;
                };
                default: components["responses"]["default_response"];
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        error: {
            /** Format: int32 */
            code: number;
            message: string;
        };
        book_info: {
            title?: string;
        };
        new_book: components["schemas"]["book_info"] & {
            userId?: number;
        };
        exist_book: components["schemas"]["new_book"] & {
            bookId?: number;
        };
        tag_info: {
            name?: string;
        };
        exist_tag: components["schemas"]["tag_info"] & {
            tagId?: number;
            userId?: number;
        };
        post_book: components["schemas"]["book_info"];
        /** @enum {string} */
        status: "todo" | "wip" | "done";
        tag: {
            id?: number;
            name?: string;
        };
        tags: components["schemas"]["tag"][];
        hoarder_book: components["schemas"]["exist_book"] & {
            status?: components["schemas"]["status"];
            tags?: components["schemas"]["tags"];
        };
        post_hoarder: components["schemas"]["book_info"] & {
            status?: components["schemas"]["status"];
            tags?: components["schemas"]["tags"];
        };
    };
    responses: {
        /** @description unexpected error */
        default_response: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": components["schemas"]["error"];
            };
        };
    };
    parameters: {
        bookId: number;
        userId: number;
        tagId: number;
    };
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export type operations = Record<string, never>;
