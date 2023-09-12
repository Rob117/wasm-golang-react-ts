# What is this

Run Go code, using web assembly, in a React app that uses typescript. This is the minimum proof of concept that I could get working.

## How do I use this

`yarn build`
`yarn start`

## Positives

You can run go code. That's pretty cool, and lets you port your go applications really easily using the web.

You can call that Go code from Javascript.

You can do things to JS from Go Code. [Small example of an alert](https://stackoverflow.com/a/69974936/4982995)

## Negatives

When you are done using a function and it will never be called again, you have to call Func.Release or the memory won't be freed up

Funcs can deadlock really easily, apparently - [See this](https://pkg.go.dev/syscall/js#FuncOf)

The LoadWasm component and wasm_exec.js is just a huge bundle of boilerplate

Type safety in Go when crossing the JS bridge is _annihilated_. Go can't ensure anything about what is coming across that bridge, so the only args that can be used when defining bridge functions are `[]js.Value`. There are a lot of methods on js.Value (`.IsNumber`, `IsUndefined`, etc.) to help you with this, but there is going to be a lot of conversion on your part.

The best way to handle this seems to just be minimizing what you pass back-and-forth between go and js, for the most part. This should be entirely possible - if you've worked with React Native and native modules, it should be a similar feeling.

Threads don't seem to really exist in the go implementation? Based on https://github.com/golang/go/issues/28631 and https://news.ycombinator.com/item?id=30357078, it might be doable, but I really don't know.


## Acknowledgements

Created while reading the following:

https://dev.to/royhadad/how-to-create-a-react-app-with-go-support-using-webassembly-in-under-60-seconds-4oa3
https://xebia.com/blog/golang-webassembly/
