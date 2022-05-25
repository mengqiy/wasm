# KRM Function WebAssembly PoC

In this PoC, we build the set-namespace (`gcr.io/kpt-fn/set-namespace`) as [WebAssembly](https://webassembly.org/) and run the KRM function in the browser.

This PoC has been deployed using GitHub Pages. You can access it [here](https://mengqiy.github.io/wasm/).

`main.go` is a thin wrapper of the set-namespace logic. The business logic is imported from package `github.com/GoogleContainerTools/kpt-functions-catalog/functions/go/set-namespace/transformer`.

Build the WebAssembly:

```shell
$ GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
```

Copy the JavaScript support file. This file is available for all golang releases after 1.11. The one I'm using here is from golang 1.18.2.

```shell
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

Add an `index.html` to stitch things together.

Start a server to serve the contents.

```shell
# install goexec: go get -u github.com/shurcooL/goexec
$ goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

Open your browser with `localhost:8080` if you are running it locally or `https://mengqiy.github.io/wasm/`.

There is an input textbox with a sample resource list. You can change it and click the `evaluate` button. The output will be shown. You can change the input and evaluate it again.
