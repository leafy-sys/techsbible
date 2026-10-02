# Doom Emacs Cheetsheet

## 1. Navigation & Movement (Normal Mode)

| Binding | Action |
| --- | --- |
| `h` / `j` / `k` / `l` | Move left / down / up / right |
| `w` / `b` | Jump forward / backward by word |
| `e` / `ge` | Jump to end / previous end of word |
| `0` / `^` / `$` | Move to line start / first non-blank char / line end |
| `gg` / `G` | Jump to beginning / end of buffer |
| `Ctrl+u` / `Ctrl+d` | Page up / page down |
| `%` | Jump between matching brackets `()`, `[]`, `{}` |
| `f{char}` / `F{char}` | Move forward / backward to `{char}` on current line |
| `;` / `,` | Repeat last `f` or `F` search forward / backward |

---

## 2. Editing & Operators

| Binding | Action |
| --- | --- |
| `i` / `a` | Insert before cursor / append after cursor |
| `I` / `A` | Insert at start of line / append at end of line |
| `o` / `O` | Open new line below / above |
| `x` / `X` | Delete character under / before cursor |
| `dd` / `D` | Delete entire line / delete to end of line |
| `dw` / `diw` | Delete word / delete inside word |
| `cc` / `C` | Change line / change to end of line |
| `ciw` / `ci"` | Change inside word / change inside double quotes |
| `yy` / `y$` | Yank (copy) line / yank to end of line |
| `p` / `P` | Paste after cursor / paste before cursor |
| `u` / `Ctrl+r` | Undo / redo |
| `.` | Repeat last change |
| `~` | Toggle character case |

---

## 3. Visual Mode & Text Objects

| Binding | Action |
| --- | --- |
| `v` / `V` / `Ctrl+v` | Character / Line / Block visual selection mode |
| `>` / `<` | Indent selection right / left |
| `~` | Toggle case of selected text |
| `iw` / `aw` | Inner word / A word (includes surrounding space) |
| `i"` / `a"` | Inside double quotes / Around double quotes |
| `i)` / `a)` | Inside parentheses / Around parentheses |
| `i]` / `a]` | Inside brackets / Around brackets |
| `ip` / `ap` | Inside paragraph / Around paragraph |

---

## 4. Doom Leader Navigation (`SPC`)

### Files & Projects (`SPC f` / `SPC p`)

* `SPC f f` – Find file (navigate file tree)
* `SPC f s` – Save buffer
* `SPC f r` – Open recent files
* `SPC f P` – Open private Doom configuration directory (`config.el`, `init.el`, etc.)
* `SPC p f` – Find file across current project

### Buffers & Windows (`SPC b` / `SPC w`)

* `SPC b b` – Switch active buffer
* `SPC b k` – Kill (close) active buffer
* `SPC w s` / `SPC w v` – Split window horizontally / vertically
* `SPC w c` – Close current split window
* `SPC w h/j/k/l` – Move focus to left / bottom / top / right split
* `SPC w =` – Balance all window split sizes equal

### Toggles & Utilities (`SPC t` / `SPC h`)

* `SPC t z` – Toggle Zen Mode (centered distraction-free reading column)
* `SPC t m` – Toggle Mixed Pitch mode (proportional reading font)
* `SPC t t` – Toggle file tree sidebar
* `SPC h r r` – Reload Doom configuration without restarting
* `SPC h t` – Load / switch color theme with live preview
* `SPC :` or `M-x` – Run any Emacs command prompt

---

## 5. Markdown Local Leader (`SPC m`)

* `SPC m m` – Toggle hiding raw Markdown markup (`#`, `**`, `[[wikilinks]]`)
* `SPC m i` – Toggle inline image display
* `SPC m l` – Follow link under cursor
* `SPC m t` – Insert Markdown table

