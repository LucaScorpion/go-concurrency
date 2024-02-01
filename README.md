# Golangang: Concurrency

Deze opdracht is een minimale opzet voor een TCP server, of hub.
De server opent een poort waarmee meerdere clients kunnen verbinden,
en stuurt berichten door naar alle verbonden clients.

## Prerequisites

- Go 1.21
- Netcat

[Netcat](https://netcat.sourceforge.net) is een tool waarmee je eenvoudig TCP verbindingen kan openen,
dit zullen we gebruiken als client om met de server te praten.

- Voor [Homebrew](https://formulae.brew.sh/formula/netcat#default): `brew install netcat`
- Voor Ubuntu: `apt install netcat`

## Quick Start

Start de server:

```shell
go run ./cmd/server
```

Verbind met de server:

```shell
nc localhost 7000
```

Input aan de server zou naar alle clients gestuurd moeten worden.
Input aan een client zou naar de server, en vervolgens naar alle clients gestuurd moeten worden.
