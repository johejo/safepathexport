# safepathexport

Generate `export` for shell.

No matter how many times you run, the path will not grow longer.

## Example Usage

```sh
eval "$(safepathexport -key PATH -value $HOME/go/bin -push)"
```
