# flags-poc

Proof of concept demonstrating how `jessevdk/go-flags` handles `[]string` values.

**tldr:** `jessevdk/go-flags` does not automatically split a comma seperated
string into an array. It simply puts the provided string in the first item of
the array, regardless of the value. The array always contains one item and it
needs to be split manually using `strings.Split()`.

```bash
$ go run .
(struct { Empty []string "long:\"empty\""; Single []string "long:\"single\""; Multiple []string "long:\"multiple\"" }) {
 Empty: ([]string) (len=1 cap=1) {
  (string) ""
 },
 Single: ([]string) (len=1 cap=1) {
  (string) (len=6) "string"
 },
 Multiple: ([]string) (len=1 cap=1) {
  (string) (len=23) "stringA,stringB,stringC"
 }
}
```
