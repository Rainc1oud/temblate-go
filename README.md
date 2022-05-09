# temblate-go

Minimalistic embedded multi-language text template generation for go programs

# Usage

1. Add the lib to your project:
   ```
   go get github.com/Rainc1oud/temblate-go
   ```
2. Make a file `temblate/generate.go` with the following content:
   ```go
   package temblate
   //go:generate go run -mod=mod github.com/Rainc1oud/temblate-go/cmd/gen ./templates ./messages.go
   ```
   and define your templates under `./temblate/templates/` with the file name formatted as `<keyname>.<lang>.gotmpl` (e.g. `mytopic.en.gotmpl`). \
3. Do
   ```
   go generate ./temblate/
   ```
   This will generate `./temblate/messages.go` (the name can be changed in the `//go:generate` line) and you can now use
   ```go
   mymsg := temblate.GetMessage(lang, key, &data)
   ```
   where `data` is a struct to the members of which you refer in the template as `{{ .MyField }}` (standard *go template* format)

## What's with the name

It's not a typo, *`temblate`* is a play on *embedded template*; the main feature compared to other `go` templating libs is that all template content is generated into the `temblate` package of your project and initialised at program start. That means you don't have to worry about files not found and multi-file resource bundling/distribution.
