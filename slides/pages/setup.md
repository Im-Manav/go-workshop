---
layout: section
---

# Setting up a development environment

---

# Operating System

Most Go developers use Linux or macOS for development, but Windows and others work well too.

<img src="/os_dev.svg">

https://go.dev/blog/survey2024-h2-results#devenv

<Tip>
I recommend developing in WSL if you're using Windows, but do remember to install Go in the WSL environment, not in Windows, and tell VS Code to run in WSL too.
</Tip>

---
layout: two-cols-header
---

# Editor

::left::

The most popular editor for Go development is Visual Studio Code.

<img src="/editor.svg">

https://go.dev/blog/survey2024-h2-results#editor-awareness-and-preferences

::right::

<Tip>
Java developers often gravitate towards JetBrains's Goland IDE or IntelliJ with the Go plugin, but I like Neovim myself.

The Go team maintains the VS Code Go extension, and maintains `gopls`, the language server that provides IDE features like auto-completion and go-to-definition.
</Tip>

