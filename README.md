# Go Tutorial

https://go.dev/doc/tutorial/getting-started

## Thoughts

### Create a Module

#### Summary
Overall, pleasantly surprised with Go! Definitely interested in playing around with it more. 
- The syntax is relatively easy to latch onto thus far, coming from a JS background
- Nice type errors at time of writing
- Nice compilation errors
- Intellisense
- Testing functionality out of the box
- Debugging functionality out of the box (at least inside VS Code)
- Easy-to-use CLI
- Fast, at least for a small "hello, world" program. Will be interested to see how it fares with larger programs.
- Biggeset annoyance was the type checking latency. 

#### Things I liked:
- nice compilation errors; used {} instead of (), and it directed me to the exact line of the error
- nice type errors; said what types needed to be returned as I was in the middle of adding a new error type
- nice intellisense out of the box with "logs" module
- interesting use of capital vs. lowercase function names to indicate exports
- nice that arrays can be of dynamic size (Slices), and it's easy to declare/instantiate
- neet use of init function that runs automatically
- compilation was fast
- testing out of the box!
    - love that testing capability comes out of the box!
    - interesting use of TestFunctionName syntax for declaring tests, and filename_test.go for test files
    - was surprised to see that the debugging to "just worked". Maybe there was something in the recommended extensions that I installed? Cool anyway
    - test output is a bit hard to read
    - interesting to see that the regexp package had a built-in MustCompile method
#### Things I didn't like:
- takes a few seconds to adjust types – annoying. There were times where I'd look at the tutorial example, write the code, see a type error, double check that I didn't misread the tutorial code, and come back to the code to see that it resolved itself after a couple of seconds.
- strange that I have to return a string AND an error from my function; why not one or the other?
- formatOnSave removes unused imports, even if I'm about to use them; probably a VSCode issue and not a Go issue though
- docs aren't consistent about showing changed lines in tutorials
- slice initialization syntax is unexpected. Why {} instead of []?
- can I direct the executable to a /build directory or something? I need to gitignore it, but I'd have to do that manually for every package...

### Accessing a relational database

#### Summary
- Custom type definitions are nice and intuitive given TS background. Still not sure what a "struct" is.
- Nice how they have a wiki with recommended RDBMS drivers. Saw drivers for PG as well. 
- Connection to db was fairly intuitive given Node background as well. 
- Unsure of difference between log and fmt
- Hope they have something like JS's template strings. This string formatting is getting old quick.
- Also hope they have good ORM options
- Syntax of for loop was surprising, and simple
- Absolutely no idea what's going on in this conditional:
```
if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
```
- Overall pleased with how easy the db connection and querying worked!