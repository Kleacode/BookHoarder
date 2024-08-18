# openapiドキュメントの結合

`$ref`で分割されているopenapiドキュメントを一つの.yamlに結合する手順です。

1. このreadmeと同じディレクトリ`(./openapi_schema)`に移動

2. イメージを作成する
```
docker build -f Dockerfile.Redocly -t <repo>:<tag> .
```
3. イメージを指定してコンテナ作成、起動
```
docker run --rm -v $PWD:/app -it <repo>:<tag>
```
4. 以下のコマンドを入力
```
npx @redocly/cli bundle ./openapi.yaml -o ./gen/openapi.yaml
```

以上の手順で、`./openapi_schema/gen`以下に結合されたopenapi.yamlが作成されます。