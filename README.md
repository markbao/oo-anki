# oo-anki

### Converts OmniOutliner-exported OPML documents in flashcard format to CSV files compatible with Anki.

## Why?

I make my flashcards in OmniOutliner from my notes, and then I import them into Anki. OmniOutliner doesn't export CSV (for some bizarre reason), and even if it did, it would screw it up as badly as it does TXT exports (incorrectly formatted newlines, etc.). So my procedure for exporting was previously copy-and-paste into Numbers, remove the headers, remove the `- ` new line marker, fixing the damn newlines by merging cells, add a tag to each cell, export to CSV, and pray.

No longer.

## Usage

1. Use the included `example/example.oo3` OmniOutliner document, which is in question–answer flashcard format.
2. Add some content into it.
3. In OmniOutliner, go to File > Export and export an OPML file.
4. Run `./oo-anki PATH_TO_OPML TAG` where `PATH_TO_OPML` is the path of the exported OPML document and `TAG` is the tag you want to add to your Anki cards. (I use tags like `PSYC2360_LEC15` to keep cards organized.)
5. A CSV file will be placed in the same directory as your OPML file. Go to Anki, File > Import, find the CSV file, and import. Done.

## Limitations

- Doesn't support HTML/formatting other than newlines because OmniOutliner's OPML export does not include any formatting.
- Future revisions should make it more seamless, automatically hook into OmniOutliner with some AppleScript JS or something, exporting, running the script, and opening in Anki automatically. (Maybe next semester).