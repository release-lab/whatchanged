# How it works?

I layered it in the following steps

### 1. Parser

parser is mainly to parse the input range

eg. `HEAD~v1.0.0`/`v2.0.0~v1.0.0`/`770ed02~585445d`/`HEAD~`/`770ed02~v2.0.0`

In this step, you need to determine the scope of the changelog to be generated

find their corresponding tags/commits.

The final interface output from parser

```go
type Range struct {
	Commit
	*client.Tag
}

type Scope struct {
	Start Range         // The range start
	End   Range         // The range end
	Tags  []*client.Tag // Tags included in the range
}
```

### 2. Extractor

After we specify the scope, we need, in the entire scope, which ones involve tags/commits

### 3. Transform

At this step, we already have relevant data

1. scope
2. tags/commit of each scope

With this data, we can generate a template. But before that, for better abstraction, We introduce an intermediate context. The entire context describes a completed version information. Including `version`/`date`/`commits`

From the above scope, we convert it to a context array

### 4. Generator

With the above context array, we can easily render the template.

Up to this point, the role of generator is to generate templates and which template to generate.

- [x] JSON
- [x] Markdown
- [ ] PlainText

### 5. formatter

In this step we will format the output. For example, format markdown/json