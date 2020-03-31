# flags-poc

Proof of concept demonstrating how `jessevdk/go-flags` handles `[]string` values.

If no value is provided at all, the array will be empty.
If any value is provided, including empty string, the array is of length one.
`jessevdk/go-flags` does not automatically split a comma seperated string into
an array.
It simply puts the provided string in the first item of the array, regardless of
the value.

```bash
(struct { NotSet []string "long:\"notset\""; Empty []string "long:\"empty\""; Single []string "long:\"single\""; Multiple []string "long:\"multiple\"" }) {
 NotSet: ([]string) <nil>,
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
