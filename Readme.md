# pathutils

Ein kleines, plattformübergreifendes Hilfspaket für die Arbeit mit Dateisystempfaden in Go.

Es erweitert die Standardbibliothek (`path/filepath`) um konsistente Behandlung von:

- Windows-Langpfaden
- UNC-Pfaden
- Pfadnormalisierung
- grundlegenden Pfad-Hilfsfunktionen

Das Paket ist bewusst minimal und zustandslos gehalten.

---

## Installation

```bash
go get github.com/yourname/pathutils@v1.0.0
```

---

## Kernfunktionen

### Absoluter Pfad

```go
abs, err := pathutils.AbsPath("relative/pfad/datei.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(abs) // z.B. /home/user/projekt/relative/pfad/datei.txt
```

Löst einen relativen Pfad anhand des aktuellen Arbeitsverzeichnisses auf. Auf Nicht-Windows-Systemen wird geprüft, ob ein Windows-Pfad übergeben wurde — das wird als Fehler behandelt.

---

### Pfad normalisieren

```go
n := pathutils.Normalize("C:\\temp\\..\\temp\\datei.txt")
fmt.Println(n) // C:\temp\datei.txt

// Auf Windows mit sehr langen Pfaden:
n2 := pathutils.Normalize("C:\\sehr\\langer\\pfad\\datei.txt")
fmt.Println(n2) // \\?\C:\sehr\langer\pfad\datei.txt  (wenn nötig)
```

Bereinigt und normalisiert einen Pfad. Unter Windows wird bei Bedarf automatisch das Langpfad-Präfix gesetzt.

---

### Pfad analysieren

```go
info := pathutils.Analyze("C:\\temp\\datei.txt")

fmt.Println(info.Type)    // "absolute"
fmt.Println(info.Volume)  // "C:"
fmt.Println(info.IsLong)  // false
fmt.Println(info.Clean)   // C:\temp\datei.txt

// UNC-Pfad:
info2 := pathutils.Analyze(`\\server\freigabe\datei.txt`)
fmt.Println(info2.Type)   // "unc"
fmt.Println(info2.Volume) // \\server\freigabe
```

Gibt strukturierte Informationen über einen Pfad zurück, ohne das Dateisystem zu berühren.

---

## Langpfad-Behandlung (nur Windows)

Windows unterstützt Pfade über 260 Zeichen nur mit dem `\\?\`-Präfix.

```go
// Normalen Pfad in Langpfad umwandeln:
lp := pathutils.LongPath("C:\\sehr\\langer\\pfad\\datei.txt")
fmt.Println(lp) // \\?\C:\sehr\langer\pfad\datei.txt

// Langpfad zurück in normalen Pfad:
p := pathutils.FromLongPath(lp)
fmt.Println(p) // C:\sehr\langer\pfad\datei.txt

// Prüfen ob ein Pfad das Langpfad-Präfix trägt:
fmt.Println(pathutils.IsLongPath(lp)) // true
fmt.Println(pathutils.IsLongPath(p))  // false
```

---

## Dateiname & Erweiterung

### NameWithoutExt

```go
pathutils.NameWithoutExt("/home/user/bericht.pdf") // bericht
pathutils.NameWithoutExt("archiv.tar.gz")          // archiv.tar  ← nur letzte Ext
pathutils.NameWithoutExt(".gitignore")             // .gitignore  ← kein Ext, nur Name
pathutils.NameWithoutExt("datei")                  // datei
```

Gibt den Dateinamen ohne die letzte Dateiendung zurück.

---

### ReplaceExt

```go
pathutils.ReplaceExt("/home/user/bericht.txt", ".pdf") // /home/user/bericht.pdf
pathutils.ReplaceExt("/home/user/bericht.txt", "pdf")  // /home/user/bericht.pdf  ← Punkt optional
pathutils.ReplaceExt("/home/user/bericht.txt", "")     // /home/user/bericht      ← Ext entfernen
pathutils.ReplaceExt("archiv.tar.gz", ".bz2")         // archiv.tar.bz2          ← nur letzte Ext
```

Ersetzt die Dateiendung eines Pfades. Der führende Punkt im zweiten Parameter ist optional.

---

### EnsureExt

```go
pathutils.EnsureExt("/home/user/bericht", ".pdf")      // /home/user/bericht.pdf  ← Ext fehlt
pathutils.EnsureExt("/home/user/bericht.pdf", ".pdf")  // /home/user/bericht.pdf  ← schon korrekt
pathutils.EnsureExt("/home/user/bericht.txt", ".pdf")  // /home/user/bericht.pdf  ← ersetzt
pathutils.EnsureExt("/home/user/bericht.pdf", "pdf")   // /home/user/bericht.pdf  ← Punkt optional
```

Stellt sicher, dass ein Pfad mit der angegebenen Dateiendung endet. Falls eine andere Endung vorhanden ist, wird sie ersetzt.

---

## Pfadvergleich

### IsSubPath

```go
// Linux:
pathutils.IsSubPath("/home/user", "/home/user/docs/datei.txt") // true
pathutils.IsSubPath("/home/user", "/home/user")                // true  ← gleicher Pfad
pathutils.IsSubPath("/home/user", "/home/other")               // false
pathutils.IsSubPath("/home/user", "/home/userx")               // false ← kein echter Subpfad

// Windows (case-insensitive):
pathutils.IsSubPath(`C:\Projekte`, `C:\projekte\main.go`)      // true
pathutils.IsSubPath(`C:\Projekte`, `D:\projekte\main.go`)      // false ← anderes Volume
```

Prüft ob ein Pfad innerhalb eines anderen liegt. Auf Windows wird case-insensitiv verglichen.

---

## Versteckte Dateien

### IsHidden

```go
// Linux / macOS — Punkt-Prefix:
pathutils.IsHidden("/home/user/.gitignore") // true
pathutils.IsHidden("/home/user/datei.txt")  // false

// Windows — prüft das Hidden-Attribut:
pathutils.IsHidden(`C:\temp\datei.txt`)     // true oder false je nach Attribut
```

Prüft ob eine Datei oder ein Verzeichnis versteckt ist. Auf Linux und macOS anhand des Punkt-Prefix im Dateinamen, auf Windows anhand des Hidden-Attributs.

---

## Hilfsfunktionen

### Exists

```go
if pathutils.Exists("config.yaml") {
    // Datei oder Verzeichnis ist vorhanden
}

// Typischer Einsatz vor dem Öffnen:
if !pathutils.Exists(pfad) {
    return fmt.Errorf("Pfad nicht gefunden: %s", pfad)
}
```

Prüft ob eine Datei oder ein Verzeichnis existiert.

---

### IsDir

```go
if pathutils.IsDir("ausgabe") {
    // Verzeichnis vorhanden, kann beschrieben werden
} else {
    // Pfad existiert nicht oder ist eine Datei
}
```

Prüft ob ein Pfad auf ein Verzeichnis zeigt.

---

## Vollständiges Beispiel

```go
package main

import (
    "fmt"
    "log"

    "github.com/yourname/pathutils"
)

func main() {
    eingabe := `C:\projekte\meinprojekt\src\bericht.txt`

    // Pfad auflösen
    abs, err := pathutils.AbsPath(eingabe)
    if err != nil {
        log.Fatal(err)
    }

    // Informationen abrufen
    info := pathutils.Analyze(abs)
    fmt.Println("Typ:     ", info.Type)
    fmt.Println("Volume:  ", info.Volume)
    fmt.Println("Langpfad:", info.IsLong)

    // Dateiname ohne Endung
    fmt.Println("Name:", pathutils.NameWithoutExt(abs)) // bericht

    // Endung ersetzen
    pdf := pathutils.ReplaceExt(abs, ".pdf")
    fmt.Println("Als PDF:", pdf) // C:\projekte\meinprojekt\src\bericht.pdf

    // Subpfad prüfen
    root := `C:\projekte`
    fmt.Println("Ist Subpfad:", pathutils.IsSubPath(root, abs)) // true

    // Versteckt?
    fmt.Println("Versteckt:", pathutils.IsHidden(abs)) // false

    // Existenz prüfen
    if !pathutils.Exists(abs) {
        fmt.Println("Datei nicht gefunden")
        return
    }
}
```

---

## Designprinzipien

- Kein globaler Zustand
- Keine Konfiguration notwendig
- Kein verstecktes Verhalten
- Funktionen sind unabhängig und kombinierbar
- Sicher für plattformübergreifenden Einsatz