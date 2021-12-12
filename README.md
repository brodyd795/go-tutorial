# Go Tutorial

https://go.dev/doc/tutorial/getting-started

### Thoughts

Things I liked:
- nice compilation errors; used {} instead of (), and it directed me to the exact line of the error
- nice type errors; said what types needed to be returned as I was in the middle of adding a new error type
- nice intellisense out of the box with "logs" module
- interesting use of capital vs. lowercase function names to indicate exports
- nice that arrays can be of dynamic size (Slices), and it's easy to declare/instantiate
- neet use of init function that runs automatically
- compilation was fast!

Things I didn't like:
- takes a few seconds to adjust types; annoying
- strange that I have to return a string AND an error from my function; why not one or the other?
- formatOnSave removes unused imports, even if I'm about to use them; probably a VSCode issue and not a Go issue though
- docs aren't consistent about showing changed lines in tutorials
- slice initialization syntax is unexpected. Why {} instead of []?
- can I direct the executable to a /build directory or something? I need to gitignore it, but I'd have to do that manually for every package...

Testing!
- love that testing capability comes out of the box!
- interesting use of TestFunctionName syntax for declaring tests, and filename_test.go for test files
- was surprised to see that the debugging to "just worked". Maybe there was something in the recommended extensions that I installed? Cool anyway
- test output is a bit hard to read
- interesting to see that the regexp package had a built-in MustCompile method
